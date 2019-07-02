package controllers

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"docker.io/go-docker/api/types/mount"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/google/uuid"
	"github.com/rescale-labs/scaleshift/api/src/auth"
	"github.com/rescale-labs/scaleshift/api/src/config"
	"github.com/rescale-labs/scaleshift/api/src/db"
	"github.com/rescale-labs/scaleshift/api/src/generated/models"
	"github.com/rescale-labs/scaleshift/api/src/generated/restapi/operations"
	"github.com/rescale-labs/scaleshift/api/src/generated/restapi/operations/job"
	"github.com/rescale-labs/scaleshift/api/src/kubernetes"
	"github.com/rescale-labs/scaleshift/api/src/lib"
	"github.com/rescale-labs/scaleshift/api/src/log"
	"github.com/rescale-labs/scaleshift/api/src/queue"
	"github.com/rescale-labs/scaleshift/api/src/rescale"
)

func jobRoute(api *operations.ScaleShiftAPI) {
	api.JobGetJobsHandler = job.GetJobsHandlerFunc(getJobs)
	api.JobPostNewJobHandler = job.PostNewJobHandlerFunc(postNewJob)
	api.JobModifyJobHandler = job.ModifyJobHandlerFunc(modifyJob)
	api.JobDeleteJobHandler = job.DeleteJobHandlerFunc(deleteJob)
}

func getJobs(params job.GetJobsParams, principal *auth.Principal) middleware.Responder {
	creds := auth.FindCredentials(principal.Username)
	ctx := params.HTTPRequest.Context()

	payload := []*models.Job{}
	if jobs, err := db.GetJobs(); err == nil {
		for _, j := range jobs {
			status := swag.String(j.Status)
			var externalLink string
			var ended time.Time

			switch j.Status {
			case db.K8sJobStart:
				status, err = kubernetes.PodStatus(creds.Base.K8sConfig, j.TargetID, "default")
				if err != nil {
					log.Debug("Kubernetes Status", err, nil)
					status = swag.String(db.StatusUnknown)
				}
			case db.RescaleStart:
				candidate, e := rescale.Status(ctx, creds.Base.RescaleKey, j.TargetID)
				if candidate == nil || e != nil {
					break
				}
				externalLink = fmt.Sprintf(
					"%s/jobs/%s/status/",
					config.Config.RescaleEndpoint,
					j.TargetID)
				switch candidate.Status {
				case rescale.JobStatusPending, rescale.JobStatusQueued,
					rescale.JobStatusWait4Cls, rescale.JobStatusWaitQueue:
					status = swag.String(db.RescaleStart)
				case rescale.JobStatusStarted, rescale.JobStatusValidated,
					rescale.JobStatusExecuting:
					status = swag.String(db.RescaleRunning)
				case rescale.JobStatusCompleted:
					status = swag.String(db.RescaleSucceed)
					if candidate.StatusDate != nil {
						ended = *candidate.StatusDate
					}
				case rescale.JobStatusStopping, rescale.JobStatusForceStop:
					status = swag.String(db.RescaleFailed)
					if candidate.StatusDate != nil {
						ended = *candidate.StatusDate
					}
				}
			}
			payload = append(payload, &models.Job{
				ID:           swag.String(j.ID),
				Platform:     j.Platform.String(),
				Status:       swag.StringValue(status),
				Image:        j.DockerImage,
				Mounts:       j.Workspaces,
				Commands:     j.Commands,
				ExternalLink: externalLink,
				Started:      strfmt.DateTime(j.Started),
				Ended:        strfmt.DateTime(ended),
			})
		}
	}
	return job.NewGetJobsOK().WithPayload(payload)
}

func postNewJob(params job.PostNewJobParams, principal *auth.Principal) middleware.Responder {
	creds := auth.FindCredentials(principal.Username)
	if swag.IsZero(creds.Base.K8sConfig) && swag.IsZero(creds.Base.RescaleKey) {
		code := http.StatusForbidden
		return job.NewPostNewJobDefault(code).WithPayload(newerror(code))
	}
	cli, _, code := dockerClient(nil)
	if code != 0 {
		return job.NewPostNewJobDefault(code).WithPayload(newerror(code))
	}
	defer cli.Close()

	// FIXME allow only emerald yet, dolomite will come soon
	if params.Body.Coretype != "emerald" {
		code := http.StatusNotAcceptable
		return job.NewPostNewJobDefault(code).WithPayload(newerror(code))
	}

	ctx := context.Background()
	container, err := cli.ContainerInspect(ctx, params.Body.NotebookID)
	if err != nil {
		log.Error("ContainerInspect@postNewJob", err, nil)
		code := http.StatusBadRequest
		return job.NewPostNewJobDefault(code).WithPayload(newerror(code))
	}
	mounts := []mount.Mount{}
	workspaces := []string{}
	src := ""
	for _, mnt := range container.Mounts {
		mounts = append(mounts, mount.Mount{
			Type:   mount.TypeBind,
			Source: mnt.Source,
			Target: mnt.Destination,
		})
		workspaces = append(workspaces, strings.TrimLeft(strings.Replace(
			mnt.Source, config.Config.WorkspaceHostDir, "", -1), "/"))
		src = strings.Replace(mnt.Source, config.Config.WorkspaceHostDir,
			config.Config.WorkspaceContainerDir, 1)
	}
	ipynb := params.Body.EntrypointFile
	if ipynb != "none" {
		nb := lib.ParseIPython(filepath.Join(src, ipynb))
		cmd := "python"
		if strings.EqualFold(nb.Meta.KernelSpec.Lang, "bash") {
			cmd = "script"
		}
		workdir := lib.DetectImageWorkDir(ctx, container.Image)
		if err := lib.ConvertNotebook(ctx, container.Image, ipynb, cmd, workdir, mounts); err != nil {
			log.Error("ConvertNotebook@postNewJob", err, nil)
			code := http.StatusBadRequest
			return job.NewPostNewJobDefault(code).WithPayload(newerror(code))
		}
		ipynb = strings.Replace(ipynb, ".ipynb", nb.Meta.Lang.FileExt, -1)
	}
	commands := []string{}
	for _, cmd := range params.Body.Commands {
		switch {
		case strings.HasSuffix(ipynb, ".py"):
			cmd = strings.Replace(cmd, "<converted-notebook.py>",
				filepath.Join("/workspace", ipynb), -1)
		case strings.HasSuffix(ipynb, ".sh"):
			if strings.EqualFold(cmd, "python") {
				cmd = "bash"
			}
			cmd = strings.Replace(cmd, "<converted-notebook.py>",
				filepath.Join("/workspace", ipynb), -1)
		}
		if command := strings.TrimSpace(cmd); command != "" {
			commands = append(commands, command)
		}
	}
	if len(commands) < 1 {
		code := http.StatusBadRequest
		return job.NewPostNewJobDefault(code).WithPayload(newerror(code))
	}
	platform := db.PlatformKubernetes
	if params.Body.PlatformID == models.PostNewJobParamsBodyPlatformIDRescale {
		platform = db.PlatformRescale
	}
	image, _, _ := lib.ContainerAttrs(container.Config.Labels)
	newjob := &db.Job{
		Platform:    platform,
		ID:          uuid.New().String(),
		Status:      db.BuildingJob,
		DockerImage: image,
		PythonFile:  ipynb,
		Workspaces:  workspaces,
		Commands:    commands,
		CPU:         params.Body.CPU,
		Memory:      params.Body.Mem,
		GPU:         params.Body.Gpu,
		CoreType:    params.Body.Coretype,
		Cores:       params.Body.Cores,
		Started:     time.Now(),
	}
	if err := db.SetJobMeta(newjob); err != nil {
		log.Error("SetJobMeta@postNewJob", err, nil)
		code := http.StatusInternalServerError
		return job.NewPostNewJobDefault(code).WithPayload(newerror(code))
	}
	credential := ""
	if strings.HasPrefix(image, config.Config.DockerRegistryHostName) {
		credential = creds.Base.DockerPassword
	}
	if strings.HasPrefix(image, config.Config.NgcRegistryHostName) {
		credential = creds.Base.NgcApikey
	}
	switch params.Body.PlatformID {
	case models.PostNewJobParamsBodyPlatformIDKubernetes:
		config.Config.DockerRegistryUserName = creds.Base.DockerUsername
		if err := queue.BuildJobDockerImage(
			newjob.ID,
			creds.Base.DockerPassword,
			principal.Username,
			creds.Base.K8sConfig,
		); err != nil {
			log.Error("BuildDockerJobImage@postNewJob", err, nil)
			code := http.StatusInternalServerError
			return job.NewPostNewJobDefault(code).WithPayload(newerror(code))
		}
	case models.PostNewJobParamsBodyPlatformIDRescale:
		if err := queue.BuildSingularityImageJob(
			newjob.ID,
			credential,
			creds.Base.RescaleKey,
			principal.Username,
		); err != nil {
			log.Error("BuildSingularityImageJob@postNewJob", err, nil)
			code := http.StatusInternalServerError
			return job.NewPostNewJobDefault(code).WithPayload(newerror(code))
		}
	}
	return job.NewPostNewJobCreated().WithPayload(&models.PostNewJobCreatedBody{
		ID: newjob.ID,
	})
}

func modifyJob(params job.ModifyJobParams, principal *auth.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	j, err := db.GetJob(params.ID)
	if err != nil {
		log.Error("GetJob@modifyJob", err, nil)
		code := http.StatusBadRequest
		return job.NewModifyJobDefault(code).WithPayload(newerror(code))
	}
	switch params.Body.Status { // nolint:gocritic
	case models.ModifyJobParamsBodyStatusStopped:
		switch j.Status {
		case db.K8sJobStart:
			// There is no proper method
		case db.RescaleStart:
			if e := rescale.Stop(ctx, config.Config.RescaleAPIToken, j.TargetID); e != nil {
				log.Error("StopRescaleJob@modifyJob", e, nil)
			}
			if e := db.UpdateJob(j.ID, db.RescaleJobStoreKey, db.RescaleJobStoreKey, db.RescaleFailed); e != nil {
				log.Error("UpdateJob@modifyJob", e, nil)
			}
		}
	}
	return job.NewModifyJobOK()
}

func deleteJob(params job.DeleteJobParams, principal *auth.Principal) middleware.Responder {
	ctx := params.HTTPRequest.Context()

	j, err := db.GetJob(params.ID)
	if err != nil {
		log.Error("GetJob@deleteJob", err, nil)
		code := http.StatusBadRequest
		return job.NewDeleteJobDefault(code).WithPayload(newerror(code))
	}
	switch j.Status {
	case db.K8sJobStart:
		// TODO
		// if e := kubernetes.DeleteJob(config.Config.KubernetesConfig, j.ID, "default"); e != nil {
		// 	log.Error("DeleteJob@deleteJob", e, nil)
		// 	code := http.StatusBadRequest
		// 	return job.NewModifyJobDefault(code).WithPayload(newerror(code))
		// }
	case db.RescaleStart, db.RescaleRunning, db.RescaleSucceed, db.RescaleFailed:
		if e := rescale.Delete(ctx, config.Config.RescaleAPIToken, j.TargetID); e != nil {
			log.Error("DeleteRescaleJob@deleteJob", e, nil)
		}
	}
	err = db.RemoveJob(j.ID)
	if err != nil {
		log.Error("RemoveJob@deleteJob", err, nil)
		code := http.StatusInternalServerError
		return job.NewDeleteJobDefault(code).WithPayload(newerror(code))
	}
	err = os.RemoveAll(filepath.Join(config.Config.SingImgContainerDir, j.ID))
	if err != nil {
		log.Error("RemoveAll@deleteJob", err, nil)
		code := http.StatusInternalServerError
		return job.NewDeleteJobDefault(code).WithPayload(newerror(code))
	}
	return job.NewDeleteJobNoContent()
}
