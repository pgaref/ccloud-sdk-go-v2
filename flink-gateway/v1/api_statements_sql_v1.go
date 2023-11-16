// Copyright 2021 Confluent Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

/*
SQL API v1

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 0.0.1
Contact: flink-control-plane@confluent.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package v1

import (
	"bytes"
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

type StatementsSqlV1Api interface {

	/*
		CreateSqlv1Statement Create a Statement

		[![Preview](https://img.shields.io/badge/Lifecycle%20Stage-Preview-%2300afba)](#section/Versioning/API-Lifecycle-Policy)
	Make a request to create a statement.

		 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param organizationId The unique identifier for the organization.
		 @param environmentId The unique identifier for the environment.
		 @return ApiCreateSqlv1StatementRequest
	*/
	CreateSqlv1Statement(ctx _context.Context, organizationId string, environmentId string) ApiCreateSqlv1StatementRequest

	// CreateSqlv1StatementExecute executes the request
	//  @return SqlV1Statement
	CreateSqlv1StatementExecute(r ApiCreateSqlv1StatementRequest) (SqlV1Statement, *_nethttp.Response, error)

	/*
		DeleteSqlv1Statement Delete a Statement

		[![Preview](https://img.shields.io/badge/Lifecycle%20Stage-Preview-%2300afba)](#section/Versioning/API-Lifecycle-Policy)
	Make a request to delete a statement.

		 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param organizationId The unique identifier for the organization.
		 @param environmentId The unique identifier for the environment.
		 @param statementName The unique identifier for the statement.
		 @return ApiDeleteSqlv1StatementRequest
	*/
	DeleteSqlv1Statement(ctx _context.Context, organizationId string, environmentId string, statementName string) ApiDeleteSqlv1StatementRequest

	// DeleteSqlv1StatementExecute executes the request
	DeleteSqlv1StatementExecute(r ApiDeleteSqlv1StatementRequest) (*_nethttp.Response, error)

	/*
		GetSqlv1Statement Read a Statement

		[![Preview](https://img.shields.io/badge/Lifecycle%20Stage-Preview-%2300afba)](#section/Versioning/API-Lifecycle-Policy)
	Make a request to read a statement.

		 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param organizationId The unique identifier for the organization.
		 @param environmentId The unique identifier for the environment.
		 @param statementName The unique identifier for the statement.
		 @return ApiGetSqlv1StatementRequest
	*/
	GetSqlv1Statement(ctx _context.Context, organizationId string, environmentId string, statementName string) ApiGetSqlv1StatementRequest

	// GetSqlv1StatementExecute executes the request
	//  @return SqlV1Statement
	GetSqlv1StatementExecute(r ApiGetSqlv1StatementRequest) (SqlV1Statement, *_nethttp.Response, error)

	/*
		ListSqlv1Statements List of Statements

		[![Preview](https://img.shields.io/badge/Lifecycle%20Stage-Preview-%2300afba)](#section/Versioning/API-Lifecycle-Policy)
	Retrieve a sorted, filtered, paginated list of all statements.

		 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param organizationId The unique identifier for the organization.
		 @param environmentId The unique identifier for the environment.
		 @return ApiListSqlv1StatementsRequest
	*/
	ListSqlv1Statements(ctx _context.Context, organizationId string, environmentId string) ApiListSqlv1StatementsRequest

	// ListSqlv1StatementsExecute executes the request
	//  @return SqlV1StatementList
	ListSqlv1StatementsExecute(r ApiListSqlv1StatementsRequest) (SqlV1StatementList, *_nethttp.Response, error)

	/*
		UpdateSqlv1Statement Update a Statement

		[![Preview](https://img.shields.io/badge/Lifecycle%20Stage-Preview-%2300afba)](#section/Versioning/API-Lifecycle-Policy)
	Make a request to update a statement.

		 @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		 @param organizationId The unique identifier for the organization.
		 @param environmentId The unique identifier for the environment.
		 @param statementName The unique identifier for the statement.
		 @return ApiUpdateSqlv1StatementRequest
	*/
	UpdateSqlv1Statement(ctx _context.Context, organizationId string, environmentId string, statementName string) ApiUpdateSqlv1StatementRequest

	// UpdateSqlv1StatementExecute executes the request
	UpdateSqlv1StatementExecute(r ApiUpdateSqlv1StatementRequest) (*_nethttp.Response, error)
}

// StatementsSqlV1ApiService StatementsSqlV1Api service
type StatementsSqlV1ApiService service

type ApiCreateSqlv1StatementRequest struct {
	ctx            _context.Context
	ApiService     StatementsSqlV1Api
	organizationId string
	environmentId  string
	sqlV1Statement *SqlV1Statement
}

func (r ApiCreateSqlv1StatementRequest) SqlV1Statement(sqlV1Statement SqlV1Statement) ApiCreateSqlv1StatementRequest {
	r.sqlV1Statement = &sqlV1Statement
	return r
}

func (r ApiCreateSqlv1StatementRequest) Execute() (SqlV1Statement, *_nethttp.Response, error) {
	return r.ApiService.CreateSqlv1StatementExecute(r)
}

/*
CreateSqlv1Statement Create a Statement

[![Preview](https://img.shields.io/badge/Lifecycle%20Stage-Preview-%2300afba)](#section/Versioning/API-Lifecycle-Policy)
Make a request to create a statement.

	@param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param organizationId The unique identifier for the organization.
	@param environmentId The unique identifier for the environment.
	@return ApiCreateSqlv1StatementRequest
*/
func (a *StatementsSqlV1ApiService) CreateSqlv1Statement(ctx _context.Context, organizationId string, environmentId string) ApiCreateSqlv1StatementRequest {
	return ApiCreateSqlv1StatementRequest{
		ApiService:     a,
		ctx:            ctx,
		organizationId: organizationId,
		environmentId:  environmentId,
	}
}

// Execute executes the request
//
//	@return SqlV1Statement
func (a *StatementsSqlV1ApiService) CreateSqlv1StatementExecute(r ApiCreateSqlv1StatementRequest) (SqlV1Statement, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SqlV1Statement
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StatementsSqlV1ApiService.CreateSqlv1Statement")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sql/v1/organizations/{organization_id}/environments/{environment_id}/statements"
	localVarPath = strings.Replace(localVarPath, "{"+"organization_id"+"}", _neturl.PathEscape(parameterToString(r.organizationId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"environment_id"+"}", _neturl.PathEscape(parameterToString(r.environmentId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.sqlV1Statement
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 409 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiDeleteSqlv1StatementRequest struct {
	ctx            _context.Context
	ApiService     StatementsSqlV1Api
	organizationId string
	environmentId  string
	statementName  string
}

func (r ApiDeleteSqlv1StatementRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.DeleteSqlv1StatementExecute(r)
}

/*
DeleteSqlv1Statement Delete a Statement

[![Preview](https://img.shields.io/badge/Lifecycle%20Stage-Preview-%2300afba)](#section/Versioning/API-Lifecycle-Policy)
Make a request to delete a statement.

	@param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param organizationId The unique identifier for the organization.
	@param environmentId The unique identifier for the environment.
	@param statementName The unique identifier for the statement.
	@return ApiDeleteSqlv1StatementRequest
*/
func (a *StatementsSqlV1ApiService) DeleteSqlv1Statement(ctx _context.Context, organizationId string, environmentId string, statementName string) ApiDeleteSqlv1StatementRequest {
	return ApiDeleteSqlv1StatementRequest{
		ApiService:     a,
		ctx:            ctx,
		organizationId: organizationId,
		environmentId:  environmentId,
		statementName:  statementName,
	}
}

// Execute executes the request
func (a *StatementsSqlV1ApiService) DeleteSqlv1StatementExecute(r ApiDeleteSqlv1StatementRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodDelete
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StatementsSqlV1ApiService.DeleteSqlv1Statement")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sql/v1/organizations/{organization_id}/environments/{environment_id}/statements/{statement_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"organization_id"+"}", _neturl.PathEscape(parameterToString(r.organizationId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"environment_id"+"}", _neturl.PathEscape(parameterToString(r.environmentId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"statement_name"+"}", _neturl.PathEscape(parameterToString(r.statementName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiGetSqlv1StatementRequest struct {
	ctx            _context.Context
	ApiService     StatementsSqlV1Api
	organizationId string
	environmentId  string
	statementName  string
}

func (r ApiGetSqlv1StatementRequest) Execute() (SqlV1Statement, *_nethttp.Response, error) {
	return r.ApiService.GetSqlv1StatementExecute(r)
}

/*
GetSqlv1Statement Read a Statement

[![Preview](https://img.shields.io/badge/Lifecycle%20Stage-Preview-%2300afba)](#section/Versioning/API-Lifecycle-Policy)
Make a request to read a statement.

	@param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param organizationId The unique identifier for the organization.
	@param environmentId The unique identifier for the environment.
	@param statementName The unique identifier for the statement.
	@return ApiGetSqlv1StatementRequest
*/
func (a *StatementsSqlV1ApiService) GetSqlv1Statement(ctx _context.Context, organizationId string, environmentId string, statementName string) ApiGetSqlv1StatementRequest {
	return ApiGetSqlv1StatementRequest{
		ApiService:     a,
		ctx:            ctx,
		organizationId: organizationId,
		environmentId:  environmentId,
		statementName:  statementName,
	}
}

// Execute executes the request
//
//	@return SqlV1Statement
func (a *StatementsSqlV1ApiService) GetSqlv1StatementExecute(r ApiGetSqlv1StatementRequest) (SqlV1Statement, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SqlV1Statement
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StatementsSqlV1ApiService.GetSqlv1Statement")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sql/v1/organizations/{organization_id}/environments/{environment_id}/statements/{statement_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"organization_id"+"}", _neturl.PathEscape(parameterToString(r.organizationId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"environment_id"+"}", _neturl.PathEscape(parameterToString(r.environmentId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"statement_name"+"}", _neturl.PathEscape(parameterToString(r.statementName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiListSqlv1StatementsRequest struct {
	ctx               _context.Context
	ApiService        StatementsSqlV1Api
	organizationId    string
	environmentId     string
	specComputePoolId *string
	pageSize          *int32
	pageToken         *string
}

// Filter the results by exact match for spec.compute_pool.
func (r ApiListSqlv1StatementsRequest) SpecComputePoolId(specComputePoolId string) ApiListSqlv1StatementsRequest {
	r.specComputePoolId = &specComputePoolId
	return r
}

// A pagination size for collection requests.
func (r ApiListSqlv1StatementsRequest) PageSize(pageSize int32) ApiListSqlv1StatementsRequest {
	r.pageSize = &pageSize
	return r
}

// An opaque pagination token for collection requests.
func (r ApiListSqlv1StatementsRequest) PageToken(pageToken string) ApiListSqlv1StatementsRequest {
	r.pageToken = &pageToken
	return r
}

func (r ApiListSqlv1StatementsRequest) Execute() (SqlV1StatementList, *_nethttp.Response, error) {
	return r.ApiService.ListSqlv1StatementsExecute(r)
}

/*
ListSqlv1Statements List of Statements

[![Preview](https://img.shields.io/badge/Lifecycle%20Stage-Preview-%2300afba)](#section/Versioning/API-Lifecycle-Policy)
Retrieve a sorted, filtered, paginated list of all statements.

	@param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param organizationId The unique identifier for the organization.
	@param environmentId The unique identifier for the environment.
	@return ApiListSqlv1StatementsRequest
*/
func (a *StatementsSqlV1ApiService) ListSqlv1Statements(ctx _context.Context, organizationId string, environmentId string) ApiListSqlv1StatementsRequest {
	return ApiListSqlv1StatementsRequest{
		ApiService:     a,
		ctx:            ctx,
		organizationId: organizationId,
		environmentId:  environmentId,
	}
}

// Execute executes the request
//
//	@return SqlV1StatementList
func (a *StatementsSqlV1ApiService) ListSqlv1StatementsExecute(r ApiListSqlv1StatementsRequest) (SqlV1StatementList, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  SqlV1StatementList
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StatementsSqlV1ApiService.ListSqlv1Statements")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sql/v1/organizations/{organization_id}/environments/{environment_id}/statements"
	localVarPath = strings.Replace(localVarPath, "{"+"organization_id"+"}", _neturl.PathEscape(parameterToString(r.organizationId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"environment_id"+"}", _neturl.PathEscape(parameterToString(r.environmentId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	if r.specComputePoolId != nil {
		localVarQueryParams.Add("spec.compute_pool_id", parameterToString(*r.specComputePoolId, ""))
	}
	if r.pageSize != nil {
		localVarQueryParams.Add("page_size", parameterToString(*r.pageSize, ""))
	}
	if r.pageToken != nil {
		localVarQueryParams.Add("page_token", parameterToString(*r.pageToken, ""))
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateSqlv1StatementRequest struct {
	ctx            _context.Context
	ApiService     StatementsSqlV1Api
	organizationId string
	environmentId  string
	statementName  string
	sqlV1Statement *SqlV1Statement
}

func (r ApiUpdateSqlv1StatementRequest) SqlV1Statement(sqlV1Statement SqlV1Statement) ApiUpdateSqlv1StatementRequest {
	r.sqlV1Statement = &sqlV1Statement
	return r
}

func (r ApiUpdateSqlv1StatementRequest) Execute() (*_nethttp.Response, error) {
	return r.ApiService.UpdateSqlv1StatementExecute(r)
}

/*
UpdateSqlv1Statement Update a Statement

[![Preview](https://img.shields.io/badge/Lifecycle%20Stage-Preview-%2300afba)](#section/Versioning/API-Lifecycle-Policy)
Make a request to update a statement.

	@param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@param organizationId The unique identifier for the organization.
	@param environmentId The unique identifier for the environment.
	@param statementName The unique identifier for the statement.
	@return ApiUpdateSqlv1StatementRequest
*/
func (a *StatementsSqlV1ApiService) UpdateSqlv1Statement(ctx _context.Context, organizationId string, environmentId string, statementName string) ApiUpdateSqlv1StatementRequest {
	return ApiUpdateSqlv1StatementRequest{
		ApiService:     a,
		ctx:            ctx,
		organizationId: organizationId,
		environmentId:  environmentId,
		statementName:  statementName,
	}
}

// Execute executes the request
func (a *StatementsSqlV1ApiService) UpdateSqlv1StatementExecute(r ApiUpdateSqlv1StatementRequest) (*_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPut
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "StatementsSqlV1ApiService.UpdateSqlv1Statement")
	if err != nil {
		return nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/sql/v1/organizations/{organization_id}/environments/{environment_id}/statements/{statement_name}"
	localVarPath = strings.Replace(localVarPath, "{"+"organization_id"+"}", _neturl.PathEscape(parameterToString(r.organizationId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"environment_id"+"}", _neturl.PathEscape(parameterToString(r.environmentId, "")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"statement_name"+"}", _neturl.PathEscape(parameterToString(r.statementName, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.sqlV1Statement
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = _ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
			return localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Failure
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarHTTPResponse, newErr
			}
			newErr.model = v
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}
