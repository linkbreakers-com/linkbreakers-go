/*
Linkbreakers API

The Linkbreakers REST API: links, QR codes, workflows, visitors and analytics, as JSON over HTTPS. Send a workspace API token as `Authorization: Bearer <token>`.  ## Entry points  `GET https://api.linkbreakers.com/` and `GET https://api.linkbreakers.com/v1` answer without credentials with the API's name, version and links. This document is served at `https://api.linkbreakers.com/openapi.json`.  ## Rate limits  Every response describes where the caller stands in the IETF `RateLimit-Policy` and `RateLimit` header fields (draft-ietf-httpapi-ratelimit-headers), and in `RateLimit-Limit`, `RateLimit-Remaining`, `RateLimit-Reset` and `X-RateLimit-Limit`, `X-RateLimit-Remaining`, `X-RateLimit-Reset` (a Unix timestamp) for clients that read those. A request past the limit gets `429 Too Many Requests` with `Retry-After` in seconds.  - Workspace API token: 720 requests per 60 seconds, per token. - Dashboard session or OAuth token: 600 requests per 60 seconds, per token. - No credentials: 300 requests per 60 seconds, per client IP.  Each plan also includes a number of API calls per billing cycle. Past it, calls answer `429` with the code `MAX_PUBLIC_API_QUERIED_REACHED` until the cycle renews or the plan changes.  ## Pagination  Every list pages with opaque cursors (AIP-158). Pass `pageSize`, then send the `nextPageToken` of each response back as `pageToken` (or its alias `cursor`) to get the next page. An empty or absent `nextPageToken` means the list is complete. Each list operation describes this in `x-pagination`.  ## Batches and background jobs  `POST /v1/links/bulk` creates up to 250 links in one transactional request; `POST /v1/links/batch` takes the same links as a bare JSON array. `POST /v1/links/bulk/jobs` accepts the same batch and answers `202 Accepted` at once with a link job and a `Location` header; poll `GET /v1/links/jobs/{id}` until `status` is `LINK_JOB_STATUS_SUCCEEDED` (the created link ids are in `linksBulkCreate.linkIds`) or `LINK_JOB_STATUS_FAILED` (`error` says why, and nothing was created). Link jobs can be read for 24 hours after they finish.  ## Paused workspaces  A free workspace whose members have not verified any email address is paused 7 days after sign-up. Its calls then answer `403` with the code `WORKSPACE_UNVERIFIED` and change nothing; QR codes and links keep working for visitors. To lift the pause, ask the workspace owner to verify their email: `POST /v1/members/email-verification` sends the verification email again, and works while the workspace is paused. `unverifiedPauseAt` on the workspace says when the pause starts.  ## Versioning and deprecation  The version is in the path: every operation lives under `/v1`. Within a version, changes are additive only: new operations, new optional request fields, new response fields and new enum values. Clients should ignore fields and enum values they do not know. Anything that would break an existing client ships under a new path version instead.  An operation is retired in two steps. First it is deprecated: it keeps working, is flagged `deprecated: true` in this document, and every response carries a `Deprecation` header field (RFC 9745) with the time it was deprecated and a `Link` with `rel=\"deprecation\"` to the policy. Then, at least 180 days later, it may be removed; that date is announced in advance in a `Sunset` header field (RFC 8594) and as `x-sunset` on the operation. The full policy: https://linkbreakers.com/help/article/api-versioning-and-deprecation-policy  A renamed operation keeps answering its former path while clients still call it, until that path's sunset at the latest: a former path no client calls for 14 days is removed. Responses on a former path carry the same `Deprecation`, `Sunset` and `Link` header fields; this document lists only the current path, and the policy lists the former paths that still answer.  ## Idempotent retries  Every POST, PUT and PATCH accepts an optional `Idempotency-Key` header (draft-ietf-httpapi-idempotency-key-header). Send a new unique value, such as a UUID, with each logical request and reuse it only to retry that request. A retry with the same key and the same request returns the first response, marked `Idempotent-Replayed: true`, instead of running again. Keys are scoped to the token and kept for 24 hours. The same key with a different request is refused with `422`, and a retry while the first request is still running gets `409` with `Retry-After`. Server errors (5xx) and refusals such as 401, 403 and 429 are not stored, so retrying them runs the request again.  ## Identifying your client  Send a `Linkbreakers-Client` request header saying which client made the call: `<client>/<version> (<details>)`, for example `my-integration/1.2.0 (python 3.12; linux)`. The official SDKs, the CLI and the MCP bridge set it for you, as `sdk-python/2.3.1 (python 3.12.4; linux)` or `cli/1.4.0 (darwin arm64)`. It is optional and changes nothing about how a request is handled; it tells us which clients and versions are in use, so we know what can be retired and what to fix first. Keep the details short and free of personal data. Browsers may send it too: it is allowed by CORS.

API version: 1.205.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package linkbreakers

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// AnalyticsDashboardsAPIService AnalyticsDashboardsAPI service
type AnalyticsDashboardsAPIService service

type ApiAnalyticsDashboardsServiceCreateRequest struct {
	ctx context.Context
	ApiService *AnalyticsDashboardsAPIService
	createDashboardRequest *CreateDashboardRequest
	idempotencyKey *string
}

func (r ApiAnalyticsDashboardsServiceCreateRequest) CreateDashboardRequest(createDashboardRequest CreateDashboardRequest) ApiAnalyticsDashboardsServiceCreateRequest {
	r.createDashboardRequest = &createDashboardRequest
	return r
}

// A unique value per request, such as a UUID, that makes retrying it safe: a retry with the same key and body returns the first response instead of acting twice.
func (r ApiAnalyticsDashboardsServiceCreateRequest) IdempotencyKey(idempotencyKey string) ApiAnalyticsDashboardsServiceCreateRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ApiAnalyticsDashboardsServiceCreateRequest) Execute() (*CreateDashboardResponse, *http.Response, error) {
	return r.ApiService.AnalyticsDashboardsServiceCreateExecute(r)
}

/*
AnalyticsDashboardsServiceCreate Create analytics dashboard

Create a dashboard with its widgets, for the whole workspace or scoped to one link.

Requires the `analytics:write` scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAnalyticsDashboardsServiceCreateRequest
*/
func (a *AnalyticsDashboardsAPIService) AnalyticsDashboardsServiceCreate(ctx context.Context) ApiAnalyticsDashboardsServiceCreateRequest {
	return ApiAnalyticsDashboardsServiceCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CreateDashboardResponse
func (a *AnalyticsDashboardsAPIService) AnalyticsDashboardsServiceCreateExecute(r ApiAnalyticsDashboardsServiceCreateRequest) (*CreateDashboardResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreateDashboardResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnalyticsDashboardsAPIService.AnalyticsDashboardsServiceCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/analytics/dashboards"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createDashboardRequest == nil {
		return localVarReturnValue, nil, reportError("createDashboardRequest is required and must be specified")
	}

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
	if r.idempotencyKey != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Idempotency-Key", r.idempotencyKey, "simple", "")
	}
	// body params
	localVarPostBody = r.createDashboardRequest
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
		if localVarHTTPResponse.StatusCode == 429 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiAnalyticsDashboardsServiceDeleteRequest struct {
	ctx context.Context
	ApiService *AnalyticsDashboardsAPIService
	id string
}

func (r ApiAnalyticsDashboardsServiceDeleteRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.AnalyticsDashboardsServiceDeleteExecute(r)
}

/*
AnalyticsDashboardsServiceDelete Delete analytics dashboard

Delete a dashboard and its associated assistant conversation.

Requires the `analytics:write` scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiAnalyticsDashboardsServiceDeleteRequest
*/
func (a *AnalyticsDashboardsAPIService) AnalyticsDashboardsServiceDelete(ctx context.Context, id string) ApiAnalyticsDashboardsServiceDeleteRequest {
	return ApiAnalyticsDashboardsServiceDeleteRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *AnalyticsDashboardsAPIService) AnalyticsDashboardsServiceDeleteExecute(r ApiAnalyticsDashboardsServiceDeleteRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnalyticsDashboardsAPIService.AnalyticsDashboardsServiceDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/analytics/dashboards/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
		if localVarHTTPResponse.StatusCode == 429 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiAnalyticsDashboardsServiceGenerateWidgetSummaryRequest struct {
	ctx context.Context
	ApiService *AnalyticsDashboardsAPIService
	dashboardId string
	widgetId string
	analyticsDashboardsServiceGenerateWidgetSummaryBody *AnalyticsDashboardsServiceGenerateWidgetSummaryBody
	idempotencyKey *string
}

func (r ApiAnalyticsDashboardsServiceGenerateWidgetSummaryRequest) AnalyticsDashboardsServiceGenerateWidgetSummaryBody(analyticsDashboardsServiceGenerateWidgetSummaryBody AnalyticsDashboardsServiceGenerateWidgetSummaryBody) ApiAnalyticsDashboardsServiceGenerateWidgetSummaryRequest {
	r.analyticsDashboardsServiceGenerateWidgetSummaryBody = &analyticsDashboardsServiceGenerateWidgetSummaryBody
	return r
}

// A unique value per request, such as a UUID, that makes retrying it safe: a retry with the same key and body returns the first response instead of acting twice.
func (r ApiAnalyticsDashboardsServiceGenerateWidgetSummaryRequest) IdempotencyKey(idempotencyKey string) ApiAnalyticsDashboardsServiceGenerateWidgetSummaryRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ApiAnalyticsDashboardsServiceGenerateWidgetSummaryRequest) Execute() (*GenerateWidgetSummaryResponse, *http.Response, error) {
	return r.ApiService.AnalyticsDashboardsServiceGenerateWidgetSummaryExecute(r)
}

/*
AnalyticsDashboardsServiceGenerateWidgetSummary Generate a widget assistant summary

Analyses the widget's current values and stores the result. Consumes platform credits, so it runs only when requested.

Requires the `analytics:write` scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dashboardId Dashboard holding the widget
 @param widgetId Widget whose values should be analysed
 @return ApiAnalyticsDashboardsServiceGenerateWidgetSummaryRequest
*/
func (a *AnalyticsDashboardsAPIService) AnalyticsDashboardsServiceGenerateWidgetSummary(ctx context.Context, dashboardId string, widgetId string) ApiAnalyticsDashboardsServiceGenerateWidgetSummaryRequest {
	return ApiAnalyticsDashboardsServiceGenerateWidgetSummaryRequest{
		ApiService: a,
		ctx: ctx,
		dashboardId: dashboardId,
		widgetId: widgetId,
	}
}

// Execute executes the request
//  @return GenerateWidgetSummaryResponse
func (a *AnalyticsDashboardsAPIService) AnalyticsDashboardsServiceGenerateWidgetSummaryExecute(r ApiAnalyticsDashboardsServiceGenerateWidgetSummaryRequest) (*GenerateWidgetSummaryResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GenerateWidgetSummaryResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnalyticsDashboardsAPIService.AnalyticsDashboardsServiceGenerateWidgetSummary")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/analytics/dashboards/{dashboardId}/widgets/{widgetId}/summary"
	localVarPath = strings.Replace(localVarPath, "{"+"dashboardId"+"}", url.PathEscape(parameterValueToString(r.dashboardId, "dashboardId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"widgetId"+"}", url.PathEscape(parameterValueToString(r.widgetId, "widgetId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.analyticsDashboardsServiceGenerateWidgetSummaryBody == nil {
		return localVarReturnValue, nil, reportError("analyticsDashboardsServiceGenerateWidgetSummaryBody is required and must be specified")
	}

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
	if r.idempotencyKey != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Idempotency-Key", r.idempotencyKey, "simple", "")
	}
	// body params
	localVarPostBody = r.analyticsDashboardsServiceGenerateWidgetSummaryBody
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
		if localVarHTTPResponse.StatusCode == 429 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiAnalyticsDashboardsServiceGetWidgetSummaryRequest struct {
	ctx context.Context
	ApiService *AnalyticsDashboardsAPIService
	dashboardId string
	widgetId string
}

func (r ApiAnalyticsDashboardsServiceGetWidgetSummaryRequest) Execute() (*GetWidgetSummaryResponse, *http.Response, error) {
	return r.ApiService.AnalyticsDashboardsServiceGetWidgetSummaryExecute(r)
}

/*
AnalyticsDashboardsServiceGetWidgetSummary Get a stored widget assistant summary

Returns the last generated analysis, and whether the underlying values have changed since.

Requires the `analytics:read` scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param dashboardId
 @param widgetId
 @return ApiAnalyticsDashboardsServiceGetWidgetSummaryRequest
*/
func (a *AnalyticsDashboardsAPIService) AnalyticsDashboardsServiceGetWidgetSummary(ctx context.Context, dashboardId string, widgetId string) ApiAnalyticsDashboardsServiceGetWidgetSummaryRequest {
	return ApiAnalyticsDashboardsServiceGetWidgetSummaryRequest{
		ApiService: a,
		ctx: ctx,
		dashboardId: dashboardId,
		widgetId: widgetId,
	}
}

// Execute executes the request
//  @return GetWidgetSummaryResponse
func (a *AnalyticsDashboardsAPIService) AnalyticsDashboardsServiceGetWidgetSummaryExecute(r ApiAnalyticsDashboardsServiceGetWidgetSummaryRequest) (*GetWidgetSummaryResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetWidgetSummaryResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnalyticsDashboardsAPIService.AnalyticsDashboardsServiceGetWidgetSummary")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/analytics/dashboards/{dashboardId}/widgets/{widgetId}/summary"
	localVarPath = strings.Replace(localVarPath, "{"+"dashboardId"+"}", url.PathEscape(parameterValueToString(r.dashboardId, "dashboardId")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"widgetId"+"}", url.PathEscape(parameterValueToString(r.widgetId, "widgetId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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
		if localVarHTTPResponse.StatusCode == 429 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiAnalyticsDashboardsServiceListRequest struct {
	ctx context.Context
	ApiService *AnalyticsDashboardsAPIService
	linkId *string
	scope *string
	pageSize *int64
	pageToken *string
	cursor *string
}

// Restricts the result to the dashboard scoped to this link. Takes precedence over scope when both are set.
func (r ApiAnalyticsDashboardsServiceListRequest) LinkId(linkId string) ApiAnalyticsDashboardsServiceListRequest {
	r.linkId = &linkId
	return r
}

// Restricts the result by scope. Unset returns every dashboard.   - DASHBOARD_SCOPE_UNSPECIFIED: Every dashboard in the workspace, link-scoped or not.  - DASHBOARD_SCOPE_WORKSPACE: Only the workspace-wide dashboard, the one with no link.  - DASHBOARD_SCOPE_LINK: Only dashboards tied to a link.
func (r ApiAnalyticsDashboardsServiceListRequest) Scope(scope string) ApiAnalyticsDashboardsServiceListRequest {
	r.scope = &scope
	return r
}

// Maximum number to return. Defaults to 100, capped at 1000.
func (r ApiAnalyticsDashboardsServiceListRequest) PageSize(pageSize int64) ApiAnalyticsDashboardsServiceListRequest {
	r.pageSize = &pageSize
	return r
}

// Cursor from a previous response&#39;s nextPageToken. Empty starts at the first page. Pages are ordered by (createdAt, id), which never change, so a record cannot move between pages while a client is walking them.
func (r ApiAnalyticsDashboardsServiceListRequest) PageToken(pageToken string) ApiAnalyticsDashboardsServiceListRequest {
	r.pageToken = &pageToken
	return r
}

// Opaque cursor from the previous page&#39;s nextPageToken. Same as pageToken; when both are sent, pageToken wins.
func (r ApiAnalyticsDashboardsServiceListRequest) Cursor(cursor string) ApiAnalyticsDashboardsServiceListRequest {
	r.cursor = &cursor
	return r
}

func (r ApiAnalyticsDashboardsServiceListRequest) Execute() (*ListDashboardsResponse, *http.Response, error) {
	return r.ApiService.AnalyticsDashboardsServiceListExecute(r)
}

/*
AnalyticsDashboardsServiceList List analytics dashboards

List the workspace's analytics dashboards: the workspace-wide one and those scoped to a link. Filter by link or by scope.

Requires the `analytics:read` scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAnalyticsDashboardsServiceListRequest
*/
func (a *AnalyticsDashboardsAPIService) AnalyticsDashboardsServiceList(ctx context.Context) ApiAnalyticsDashboardsServiceListRequest {
	return ApiAnalyticsDashboardsServiceListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListDashboardsResponse
func (a *AnalyticsDashboardsAPIService) AnalyticsDashboardsServiceListExecute(r ApiAnalyticsDashboardsServiceListRequest) (*ListDashboardsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListDashboardsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnalyticsDashboardsAPIService.AnalyticsDashboardsServiceList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/analytics/dashboards"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.linkId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "linkId", r.linkId, "form", "")
	}
	if r.scope != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "scope", r.scope, "form", "")
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "form", "")
	}
	if r.pageToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageToken", r.pageToken, "form", "")
	}
	if r.cursor != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "cursor", r.cursor, "form", "")
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
		if localVarHTTPResponse.StatusCode == 429 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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

type ApiAnalyticsDashboardsServiceUpdateRequest struct {
	ctx context.Context
	ApiService *AnalyticsDashboardsAPIService
	id string
	analyticsDashboardsServiceUpdateBody *AnalyticsDashboardsServiceUpdateBody
	idempotencyKey *string
}

func (r ApiAnalyticsDashboardsServiceUpdateRequest) AnalyticsDashboardsServiceUpdateBody(analyticsDashboardsServiceUpdateBody AnalyticsDashboardsServiceUpdateBody) ApiAnalyticsDashboardsServiceUpdateRequest {
	r.analyticsDashboardsServiceUpdateBody = &analyticsDashboardsServiceUpdateBody
	return r
}

// A unique value per request, such as a UUID, that makes retrying it safe: a retry with the same key and body returns the first response instead of acting twice.
func (r ApiAnalyticsDashboardsServiceUpdateRequest) IdempotencyKey(idempotencyKey string) ApiAnalyticsDashboardsServiceUpdateRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ApiAnalyticsDashboardsServiceUpdateRequest) Execute() (*UpdateDashboardResponse, *http.Response, error) {
	return r.ApiService.AnalyticsDashboardsServiceUpdateExecute(r)
}

/*
AnalyticsDashboardsServiceUpdate Update analytics dashboard

Replace a dashboard's name and its widget configuration. Each widget names its own data source and visualization.

Requires the `analytics:write` scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id ID of the dashboard to update.
 @return ApiAnalyticsDashboardsServiceUpdateRequest
*/
func (a *AnalyticsDashboardsAPIService) AnalyticsDashboardsServiceUpdate(ctx context.Context, id string) ApiAnalyticsDashboardsServiceUpdateRequest {
	return ApiAnalyticsDashboardsServiceUpdateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return UpdateDashboardResponse
func (a *AnalyticsDashboardsAPIService) AnalyticsDashboardsServiceUpdateExecute(r ApiAnalyticsDashboardsServiceUpdateRequest) (*UpdateDashboardResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPut
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UpdateDashboardResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnalyticsDashboardsAPIService.AnalyticsDashboardsServiceUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/analytics/dashboards/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.analyticsDashboardsServiceUpdateBody == nil {
		return localVarReturnValue, nil, reportError("analyticsDashboardsServiceUpdateBody is required and must be specified")
	}

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
	if r.idempotencyKey != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "Idempotency-Key", r.idempotencyKey, "simple", "")
	}
	// body params
	localVarPostBody = r.analyticsDashboardsServiceUpdateBody
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
		if localVarHTTPResponse.StatusCode == 429 {
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
			var v ErrorResponse
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
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
