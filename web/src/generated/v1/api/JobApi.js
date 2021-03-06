/**
 * ScaleShift
 * A platform for machine learning & high performance computing 
 *
 * OpenAPI spec version: 1.0.0
 *
 * NOTE: This class is auto generated by the swagger code generator program.
 * https://github.com/swagger-api/swagger-codegen.git
 *
 * Swagger Codegen version: 2.3.1
 *
 * Do not edit the class manually.
 *
 */

(function(root, factory) {
  if (typeof define === 'function' && define.amd) {
    // AMD. Register as an anonymous module.
    define(['ApiClient', 'model/Error', 'model/Job', 'model/JobAttrs', 'model/JobAttrs1', 'model/JobDetail', 'model/JobFiles', 'model/JobLogs', 'model/NewJobID'], factory);
  } else if (typeof module === 'object' && module.exports) {
    // CommonJS-like environments that support module.exports, like Node.
    module.exports = factory(require('../ApiClient'), require('../model/Error'), require('../model/Job'), require('../model/JobAttrs'), require('../model/JobAttrs1'), require('../model/JobDetail'), require('../model/JobFiles'), require('../model/JobLogs'), require('../model/NewJobID'));
  } else {
    // Browser globals (root is window)
    if (!root.ScaleShift) {
      root.ScaleShift = {};
    }
    root.ScaleShift.JobApi = factory(root.ScaleShift.ApiClient, root.ScaleShift.Error, root.ScaleShift.Job, root.ScaleShift.JobAttrs, root.ScaleShift.JobAttrs1, root.ScaleShift.JobDetail, root.ScaleShift.JobFiles, root.ScaleShift.JobLogs, root.ScaleShift.NewJobID);
  }
}(this, function(ApiClient, Error, Job, JobAttrs, JobAttrs1, JobDetail, JobFiles, JobLogs, NewJobID) {
  'use strict';

  /**
   * Job service.
   * @module api/JobApi
   * @version 1.0.0
   */

  /**
   * Constructs a new JobApi. 
   * @alias module:api/JobApi
   * @class
   * @param {module:ApiClient} [apiClient] Optional API client implementation to use,
   * default to {@link module:ApiClient#instance} if unspecified.
   */
  var exports = function(apiClient) {
    this.apiClient = apiClient || ApiClient.instance;


    /**
     * Callback function to receive the result of the deleteJob operation.
     * @callback module:api/JobApi~deleteJobCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * delete a job 
     * @param {String} id Job ID
     * @param {module:api/JobApi~deleteJobCallback} callback The callback function, accepting three arguments: error, data, response
     */
    this.deleteJob = function(id, callback) {
      var postBody = null;

      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling deleteJob");
      }


      var pathParams = {
        'id': id
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = ['api-authorizer'];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = null;

      return this.apiClient.callApi(
        '/jobs/{id}', 'DELETE',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the getJobDetail operation.
     * @callback module:api/JobApi~getJobDetailCallback
     * @param {String} error Error message, if any.
     * @param {module:model/JobDetail} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * returns the details of a job 
     * @param {String} id Job ID
     * @param {module:api/JobApi~getJobDetailCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/JobDetail}
     */
    this.getJobDetail = function(id, callback) {
      var postBody = null;

      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling getJobDetail");
      }


      var pathParams = {
        'id': id
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = ['api-authorizer'];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = JobDetail;

      return this.apiClient.callApi(
        '/jobs/{id}', 'GET',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the getJobFiles operation.
     * @callback module:api/JobApi~getJobFilesCallback
     * @param {String} error Error message, if any.
     * @param {module:model/JobFiles} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * returns the list of output files 
     * @param {String} id Job ID
     * @param {module:api/JobApi~getJobFilesCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/JobFiles}
     */
    this.getJobFiles = function(id, callback) {
      var postBody = null;

      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling getJobFiles");
      }


      var pathParams = {
        'id': id
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = ['api-authorizer'];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = JobFiles;

      return this.apiClient.callApi(
        '/jobs/{id}/files', 'GET',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the getJobLogs operation.
     * @callback module:api/JobApi~getJobLogsCallback
     * @param {String} error Error message, if any.
     * @param {module:model/JobLogs} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * returns the logs of a job 
     * @param {String} id Job ID
     * @param {module:api/JobApi~getJobLogsCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/JobLogs}
     */
    this.getJobLogs = function(id, callback) {
      var postBody = null;

      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling getJobLogs");
      }


      var pathParams = {
        'id': id
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = ['api-authorizer'];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = JobLogs;

      return this.apiClient.callApi(
        '/jobs/{id}/logs', 'GET',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the getJobs operation.
     * @callback module:api/JobApi~getJobsCallback
     * @param {String} error Error message, if any.
     * @param {Array.<module:model/Job>} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * returns training jobs on cloud 
     * @param {module:api/JobApi~getJobsCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link Array.<module:model/Job>}
     */
    this.getJobs = function(callback) {
      var postBody = null;


      var pathParams = {
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = ['api-authorizer'];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = [Job];

      return this.apiClient.callApi(
        '/jobs', 'GET',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the modifyJob operation.
     * @callback module:api/JobApi~modifyJobCallback
     * @param {String} error Error message, if any.
     * @param data This operation does not return a value.
     * @param {String} response The complete HTTP response.
     */

    /**
     * modify the job status 
     * @param {String} id Job ID
     * @param {module:model/JobAttrs1} body 
     * @param {module:api/JobApi~modifyJobCallback} callback The callback function, accepting three arguments: error, data, response
     */
    this.modifyJob = function(id, body, callback) {
      var postBody = body;

      // verify the required parameter 'id' is set
      if (id === undefined || id === null) {
        throw new Error("Missing the required parameter 'id' when calling modifyJob");
      }

      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling modifyJob");
      }


      var pathParams = {
        'id': id
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = ['api-authorizer'];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = null;

      return this.apiClient.callApi(
        '/jobs/{id}', 'PATCH',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }

    /**
     * Callback function to receive the result of the postNewJob operation.
     * @callback module:api/JobApi~postNewJobCallback
     * @param {String} error Error message, if any.
     * @param {module:model/NewJobID} data The data returned by the service call.
     * @param {String} response The complete HTTP response.
     */

    /**
     * Submit a job with the specified image 
     * @param {module:model/JobAttrs} body 
     * @param {module:api/JobApi~postNewJobCallback} callback The callback function, accepting three arguments: error, data, response
     * data is of type: {@link module:model/NewJobID}
     */
    this.postNewJob = function(body, callback) {
      var postBody = body;

      // verify the required parameter 'body' is set
      if (body === undefined || body === null) {
        throw new Error("Missing the required parameter 'body' when calling postNewJob");
      }


      var pathParams = {
      };
      var queryParams = {
      };
      var collectionQueryParams = {
      };
      var headerParams = {
      };
      var formParams = {
      };

      var authNames = ['api-authorizer'];
      var contentTypes = ['application/json'];
      var accepts = ['application/json'];
      var returnType = NewJobID;

      return this.apiClient.callApi(
        '/jobs', 'POST',
        pathParams, queryParams, collectionQueryParams, headerParams, formParams, postBody,
        authNames, contentTypes, accepts, returnType, callback
      );
    }
  };

  return exports;
}));
