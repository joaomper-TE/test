/*
Administrative API

## Overview Manage users, accounts, and account groups in the ThousandEyes platform using the Administrative API.  This API provides the following endpoints that define the operations to manage your organization:     * `/account-groups`: Account groups are used to divide an organization into different sections. These endpoints can be used to create, retrieve, update and delete account groups.   * `/users`: Create, retrieve, update and delete users within an organization.    * `/roles`: Create, retrieve and update roles for the current user.    * `/permissions`: Retrieve all assignable permissions. Used in the context of modifying roles.    * `/audit-user-events`: Retrieve all activity log events.

API version: 7.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package thousandeyes

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"time"
)


// UserEventsAPIService UserEventsAPI service
type UserEventsAPIService service

type ApiGetUserEventsRequest struct {
	ctx context.Context
	ApiService *UserEventsAPIService
	aid *string
	window *string
	startDate *time.Time
	endDate *time.Time
	cursor *string
}

// A unique identifier associated with your account group. You can retrieve your &#x60;AccountGroupId&#x60; from the &#x60;/account-groups&#x60; endpoint. You must be assigned to the target account group. Specifying this parameter without being assigned to the target account results in an error response.
func (r ApiGetUserEventsRequest) Aid(aid string) ApiGetUserEventsRequest {
	r.aid = &aid
	return r
}

// The &#x60;window&#x60; parameter helps you get data from a specific time in the past up to the present request time. You specify an amount of time back, like a number followed by an optional time interval type: &#x60;s&#x60; for seconds, &#x60;m&#x60; for minutes, &#x60;h&#x60; for hours, &#x60;d&#x60; for days, and &#x60;w&#x60; for weeks. If no type is given, it&#39;s assumed to be seconds. If you need a precise range, use startDate and endDate instead.
func (r ApiGetUserEventsRequest) Window(window string) ApiGetUserEventsRequest {
	r.window = &window
	return r
}

// To focus on specific data, use the &#x60;startDate&#x60; and &#x60;endDate&#x60; parameters. Dates should follow the ISO 8601 date-time format, with hyphens for dates and colons for times. You need to include the complete time (hours, minutes, and seconds). The date and time can be separated by either a space or the letter T. For instance: &#x60;2012-01-01T00:00:00Z&#x60;. The date range includes both start and end dates. The time zone is UTC. Please note that this parameter can&#39;t be used alongside &#x60;window&#x60;.
func (r ApiGetUserEventsRequest) StartDate(startDate time.Time) ApiGetUserEventsRequest {
	r.startDate = &startDate
	return r
}

// The endDate parameter is optional; use alongside the &#x60;startDate&#x60; parameter. If not provided, the current time (at the request moment) is assumed. This parameter is incompatible with &#x60;window&#x60;.
func (r ApiGetUserEventsRequest) EndDate(endDate time.Time) ApiGetUserEventsRequest {
	r.endDate = &endDate
	return r
}

// Optional opaque cursor that is used for pagination. Clients  must not use this parameter directly and should instead use the &#x60;_links&#x60; provided by the api.
func (r ApiGetUserEventsRequest) Cursor(cursor string) ApiGetUserEventsRequest {
	r.cursor = &cursor
	return r
}

func (r ApiGetUserEventsRequest) Execute() (*GetUserEvents200Response, *http.Response, error) {
	return r.ApiService.GetUserEventsExecute(r)
}

/*
GetUserEvents List activity log events

Returns a list of activity log events. Users with the `View activity log for all users in account group` permission can see all activity log events in the current account group. Users with the `View own activity log` permission can see their own activity log events in the current account group. 

For more information about changing the account group context, see [Account Context](https://developer.thousandeyes.com/v7/#/accountcontext).

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetUserEventsRequest
*/
func (a *UserEventsAPIService) GetUserEvents(ctx context.Context) ApiGetUserEventsRequest {
	return ApiGetUserEventsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetUserEvents200Response
func (a *UserEventsAPIService) GetUserEventsExecute(r ApiGetUserEventsRequest) (*GetUserEvents200Response, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetUserEvents200Response
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "UserEventsAPIService.GetUserEvents")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v7/audit-user-events/"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.aid != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "aid", r.aid, "")
	}
	if r.window != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "window", r.window, "")
	}
	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startDate", r.startDate, "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "endDate", r.endDate, "")
	}
	if r.cursor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cursor", r.cursor, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/hal+json", "application/problem+json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v UnauthorizedError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 403 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 429 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 500 {
			var v Error
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}