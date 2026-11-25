/*
Linkbreakers API

The Linkbreakers REST API: links, QR codes, workflows, visitors and analytics, as JSON over HTTPS. Send a workspace API token as `Authorization: Bearer <token>`.  ## Entry points  `GET https://api.linkbreakers.com/` and `GET https://api.linkbreakers.com/v1` answer without credentials with the API's name, version and links. This document is served at `https://api.linkbreakers.com/openapi.json`.  ## Rate limits  Every response describes where the caller stands in the IETF `RateLimit-Policy` and `RateLimit` header fields (draft-ietf-httpapi-ratelimit-headers), and in `RateLimit-Limit`, `RateLimit-Remaining`, `RateLimit-Reset` and `X-RateLimit-Limit`, `X-RateLimit-Remaining`, `X-RateLimit-Reset` (a Unix timestamp) for clients that read those. A request past the limit gets `429 Too Many Requests` with `Retry-After` in seconds.  - Workspace API token: 720 requests per 60 seconds, per token. - Dashboard session or OAuth token: 600 requests per 60 seconds, per token. - No credentials: 300 requests per 60 seconds, per client IP.  Each plan also includes a number of API calls per billing cycle. Past it, calls answer `429` with the code `MAX_PUBLIC_API_QUERIED_REACHED` until the cycle renews or the plan changes.  ## Pagination  Every list pages with opaque cursors (AIP-158). Pass `pageSize`, then send the `nextPageToken` of each response back as `pageToken` (or its alias `cursor`) to get the next page. An empty or absent `nextPageToken` means the list is complete. Each list operation describes this in `x-pagination`.  ## Batches and background jobs  `POST /v1/links/bulk` creates up to 250 links in one transactional request; `POST /v1/links/batch` takes the same links as a bare JSON array. `POST /v1/links/bulk/jobs` accepts the same batch and answers `202 Accepted` at once with a link job and a `Location` header; poll `GET /v1/links/jobs/{id}` until `status` is `LINK_JOB_STATUS_SUCCEEDED` (the created link ids are in `linksBulkCreate.linkIds`) or `LINK_JOB_STATUS_FAILED` (`error` says why, and nothing was created). Link jobs can be read for 24 hours after they finish.  ## Paused workspaces  A free workspace whose members have not verified any email address is paused 7 days after sign-up. Its calls then answer `403` with the code `WORKSPACE_UNVERIFIED` and change nothing; QR codes and links keep working for visitors. To lift the pause, ask the workspace owner to verify their email: `POST /v1/members/email-verification` sends the verification email again, and works while the workspace is paused. `unverifiedPauseAt` on the workspace says when the pause starts.  ## Versioning and deprecation  The version is in the path: every operation lives under `/v1`. Within a version, changes are additive only: new operations, new optional request fields, new response fields and new enum values. Clients should ignore fields and enum values they do not know. Anything that would break an existing client ships under a new path version instead.  An operation is retired in two steps. First it is deprecated: it keeps working, is flagged `deprecated: true` in this document, and every response carries a `Deprecation` header field (RFC 9745) with the time it was deprecated and a `Link` with `rel=\"deprecation\"` to the policy. Then, at least 180 days later, it may be removed; that date is announced in advance in a `Sunset` header field (RFC 8594) and as `x-sunset` on the operation. The full policy: https://linkbreakers.com/help/article/api-versioning-and-deprecation-policy  A renamed operation keeps answering its former path while clients still call it, until that path's sunset at the latest: a former path no client calls for 14 days is removed. Responses on a former path carry the same `Deprecation`, `Sunset` and `Link` header fields; this document lists only the current path, and the policy lists the former paths that still answer.  ## Idempotent retries  Every POST, PUT and PATCH accepts an optional `Idempotency-Key` header (draft-ietf-httpapi-idempotency-key-header). Send a new unique value, such as a UUID, with each logical request and reuse it only to retry that request. A retry with the same key and the same request returns the first response, marked `Idempotent-Replayed: true`, instead of running again. Keys are scoped to the token and kept for 24 hours. The same key with a different request is refused with `422`, and a retry while the first request is still running gets `409` with `Retry-After`. Server errors (5xx) and refusals such as 401, 403 and 429 are not stored, so retrying them runs the request again.  ## Identifying your client  Send a `Linkbreakers-Client` request header saying which client made the call: `<client>/<version> (<details>)`, for example `my-integration/1.2.0 (python 3.12; linux)`. The official SDKs, the CLI and the MCP bridge set it for you, as `sdk-python/2.3.1 (python 3.12.4; linux)` or `cli/1.4.0 (darwin arm64)`. It is optional and changes nothing about how a request is handled; it tells us which clients and versions are in use, so we know what can be retired and what to fix first. Keep the details short and free of personal data. Browsers may send it too: it is allowed by CORS.

API version: 1.205.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package linkbreakers

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"time"
	"reflect"
)


// AnalyticsAPIService AnalyticsAPI service
type AnalyticsAPIService service

type ApiAnalyticsServiceGetAttributeDistributionRequest struct {
	ctx context.Context
	ApiService *AnalyticsAPIService
	getAttributeDistributionRequest *GetAttributeDistributionRequest
	idempotencyKey *string
}

func (r ApiAnalyticsServiceGetAttributeDistributionRequest) GetAttributeDistributionRequest(getAttributeDistributionRequest GetAttributeDistributionRequest) ApiAnalyticsServiceGetAttributeDistributionRequest {
	r.getAttributeDistributionRequest = &getAttributeDistributionRequest
	return r
}

// A unique value per request, such as a UUID, that makes retrying it safe: a retry with the same key and body returns the first response instead of acting twice.
func (r ApiAnalyticsServiceGetAttributeDistributionRequest) IdempotencyKey(idempotencyKey string) ApiAnalyticsServiceGetAttributeDistributionRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ApiAnalyticsServiceGetAttributeDistributionRequest) Execute() (*GetAttributeDistributionResponse, *http.Response, error) {
	return r.ApiService.AnalyticsServiceGetAttributeDistributionExecute(r)
}

/*
AnalyticsServiceGetAttributeDistribution Get visitor attribute distribution

Returns how visitors are spread across the values of one attribute, for example how many gave each rating.

Requires the `analytics:read` scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAnalyticsServiceGetAttributeDistributionRequest
*/
func (a *AnalyticsAPIService) AnalyticsServiceGetAttributeDistribution(ctx context.Context) ApiAnalyticsServiceGetAttributeDistributionRequest {
	return ApiAnalyticsServiceGetAttributeDistributionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetAttributeDistributionResponse
func (a *AnalyticsAPIService) AnalyticsServiceGetAttributeDistributionExecute(r ApiAnalyticsServiceGetAttributeDistributionRequest) (*GetAttributeDistributionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetAttributeDistributionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnalyticsAPIService.AnalyticsServiceGetAttributeDistribution")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/analytics/attributes/distribution"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.getAttributeDistributionRequest == nil {
		return localVarReturnValue, nil, reportError("getAttributeDistributionRequest is required and must be specified")
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
	localVarPostBody = r.getAttributeDistributionRequest
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

type ApiAnalyticsServiceGetAttributeMetricRequest struct {
	ctx context.Context
	ApiService *AnalyticsAPIService
	getAttributeMetricRequest *GetAttributeMetricRequest
	idempotencyKey *string
}

func (r ApiAnalyticsServiceGetAttributeMetricRequest) GetAttributeMetricRequest(getAttributeMetricRequest GetAttributeMetricRequest) ApiAnalyticsServiceGetAttributeMetricRequest {
	r.getAttributeMetricRequest = &getAttributeMetricRequest
	return r
}

// A unique value per request, such as a UUID, that makes retrying it safe: a retry with the same key and body returns the first response instead of acting twice.
func (r ApiAnalyticsServiceGetAttributeMetricRequest) IdempotencyKey(idempotencyKey string) ApiAnalyticsServiceGetAttributeMetricRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ApiAnalyticsServiceGetAttributeMetricRequest) Execute() (*GetAttributeMetricResponse, *http.Response, error) {
	return r.ApiService.AnalyticsServiceGetAttributeMetricExecute(r)
}

/*
AnalyticsServiceGetAttributeMetric Get visitor attribute metric

Returns count, average, minimum and maximum for a visitor attribute, plus how much of the audience carries it.

Requires the `analytics:read` scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAnalyticsServiceGetAttributeMetricRequest
*/
func (a *AnalyticsAPIService) AnalyticsServiceGetAttributeMetric(ctx context.Context) ApiAnalyticsServiceGetAttributeMetricRequest {
	return ApiAnalyticsServiceGetAttributeMetricRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetAttributeMetricResponse
func (a *AnalyticsAPIService) AnalyticsServiceGetAttributeMetricExecute(r ApiAnalyticsServiceGetAttributeMetricRequest) (*GetAttributeMetricResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetAttributeMetricResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnalyticsAPIService.AnalyticsServiceGetAttributeMetric")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/analytics/attributes/metric"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.getAttributeMetricRequest == nil {
		return localVarReturnValue, nil, reportError("getAttributeMetricRequest is required and must be specified")
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
	localVarPostBody = r.getAttributeMetricRequest
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

type ApiAnalyticsServiceGetAttributeValuesRequest struct {
	ctx context.Context
	ApiService *AnalyticsAPIService
	getAttributeValuesRequest *GetAttributeValuesRequest
	idempotencyKey *string
}

func (r ApiAnalyticsServiceGetAttributeValuesRequest) GetAttributeValuesRequest(getAttributeValuesRequest GetAttributeValuesRequest) ApiAnalyticsServiceGetAttributeValuesRequest {
	r.getAttributeValuesRequest = &getAttributeValuesRequest
	return r
}

// A unique value per request, such as a UUID, that makes retrying it safe: a retry with the same key and body returns the first response instead of acting twice.
func (r ApiAnalyticsServiceGetAttributeValuesRequest) IdempotencyKey(idempotencyKey string) ApiAnalyticsServiceGetAttributeValuesRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ApiAnalyticsServiceGetAttributeValuesRequest) Execute() (*GetAttributeValuesResponse, *http.Response, error) {
	return r.ApiService.AnalyticsServiceGetAttributeValuesExecute(r)
}

/*
AnalyticsServiceGetAttributeValues List visitor attribute values

Returns the individual recorded values of one attribute, most recent first. Intended for free text such as comments.

Requires the `analytics:read` scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAnalyticsServiceGetAttributeValuesRequest
*/
func (a *AnalyticsAPIService) AnalyticsServiceGetAttributeValues(ctx context.Context) ApiAnalyticsServiceGetAttributeValuesRequest {
	return ApiAnalyticsServiceGetAttributeValuesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetAttributeValuesResponse
func (a *AnalyticsAPIService) AnalyticsServiceGetAttributeValuesExecute(r ApiAnalyticsServiceGetAttributeValuesRequest) (*GetAttributeValuesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetAttributeValuesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnalyticsAPIService.AnalyticsServiceGetAttributeValues")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/analytics/attributes/values"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.getAttributeValuesRequest == nil {
		return localVarReturnValue, nil, reportError("getAttributeValuesRequest is required and must be specified")
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
	localVarPostBody = r.getAttributeValuesRequest
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

type ApiAnalyticsServiceGetBreakdownRequest struct {
	ctx context.Context
	ApiService *AnalyticsAPIService
	getBreakdownRequest *GetBreakdownRequest
	idempotencyKey *string
}

func (r ApiAnalyticsServiceGetBreakdownRequest) GetBreakdownRequest(getBreakdownRequest GetBreakdownRequest) ApiAnalyticsServiceGetBreakdownRequest {
	r.getBreakdownRequest = &getBreakdownRequest
	return r
}

// A unique value per request, such as a UUID, that makes retrying it safe: a retry with the same key and body returns the first response instead of acting twice.
func (r ApiAnalyticsServiceGetBreakdownRequest) IdempotencyKey(idempotencyKey string) ApiAnalyticsServiceGetBreakdownRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ApiAnalyticsServiceGetBreakdownRequest) Execute() (*GetBreakdownResponse, *http.Response, error) {
	return r.ApiService.AnalyticsServiceGetBreakdownExecute(r)
}

/*
AnalyticsServiceGetBreakdown Get dimension breakdown

Returns a ranked breakdown of events by any dimension (browser, platform, device type, country, city, visit type, in-app browser, in-app browser version).

Requires the `analytics:read` scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAnalyticsServiceGetBreakdownRequest
*/
func (a *AnalyticsAPIService) AnalyticsServiceGetBreakdown(ctx context.Context) ApiAnalyticsServiceGetBreakdownRequest {
	return ApiAnalyticsServiceGetBreakdownRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetBreakdownResponse
func (a *AnalyticsAPIService) AnalyticsServiceGetBreakdownExecute(r ApiAnalyticsServiceGetBreakdownRequest) (*GetBreakdownResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetBreakdownResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnalyticsAPIService.AnalyticsServiceGetBreakdown")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/analytics/breakdown"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.getBreakdownRequest == nil {
		return localVarReturnValue, nil, reportError("getBreakdownRequest is required and must be specified")
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
	localVarPostBody = r.getBreakdownRequest
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

type ApiAnalyticsServiceGetDeviceDistributionRequest struct {
	ctx context.Context
	ApiService *AnalyticsAPIService
	getDeviceDistributionRequest *GetDeviceDistributionRequest
	idempotencyKey *string
}

func (r ApiAnalyticsServiceGetDeviceDistributionRequest) GetDeviceDistributionRequest(getDeviceDistributionRequest GetDeviceDistributionRequest) ApiAnalyticsServiceGetDeviceDistributionRequest {
	r.getDeviceDistributionRequest = &getDeviceDistributionRequest
	return r
}

// A unique value per request, such as a UUID, that makes retrying it safe: a retry with the same key and body returns the first response instead of acting twice.
func (r ApiAnalyticsServiceGetDeviceDistributionRequest) IdempotencyKey(idempotencyKey string) ApiAnalyticsServiceGetDeviceDistributionRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ApiAnalyticsServiceGetDeviceDistributionRequest) Execute() (*GetDeviceDistributionResponse, *http.Response, error) {
	return r.ApiService.AnalyticsServiceGetDeviceDistributionExecute(r)
}

/*
AnalyticsServiceGetDeviceDistribution Get device distribution

Retrieve event distribution across device types.

Requires the `analytics:read` scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAnalyticsServiceGetDeviceDistributionRequest
*/
func (a *AnalyticsAPIService) AnalyticsServiceGetDeviceDistribution(ctx context.Context) ApiAnalyticsServiceGetDeviceDistributionRequest {
	return ApiAnalyticsServiceGetDeviceDistributionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetDeviceDistributionResponse
func (a *AnalyticsAPIService) AnalyticsServiceGetDeviceDistributionExecute(r ApiAnalyticsServiceGetDeviceDistributionRequest) (*GetDeviceDistributionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetDeviceDistributionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnalyticsAPIService.AnalyticsServiceGetDeviceDistribution")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/analytics/devices"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.getDeviceDistributionRequest == nil {
		return localVarReturnValue, nil, reportError("getDeviceDistributionRequest is required and must be specified")
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
	localVarPostBody = r.getDeviceDistributionRequest
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

type ApiAnalyticsServiceGetEventsMetricRequest struct {
	ctx context.Context
	ApiService *AnalyticsAPIService
	startDate *time.Time
	endDate *time.Time
	linkId *string
	directoryId *string
	tagId *string
	countries *[]string
}

// Inclusive start timestamp for the aggregation
func (r ApiAnalyticsServiceGetEventsMetricRequest) StartDate(startDate time.Time) ApiAnalyticsServiceGetEventsMetricRequest {
	r.startDate = &startDate
	return r
}

// Inclusive end timestamp for the aggregation
func (r ApiAnalyticsServiceGetEventsMetricRequest) EndDate(endDate time.Time) ApiAnalyticsServiceGetEventsMetricRequest {
	r.endDate = &endDate
	return r
}

// Optional link filter (UUID); when omitted aggregates workspace-wide  UUID format
func (r ApiAnalyticsServiceGetEventsMetricRequest) LinkId(linkId string) ApiAnalyticsServiceGetEventsMetricRequest {
	r.linkId = &linkId
	return r
}

// Optional directory filter (UUID): the links of this directory and of every directory below it. Send linkId or directoryId, not both; when both are omitted aggregates workspace-wide
func (r ApiAnalyticsServiceGetEventsMetricRequest) DirectoryId(directoryId string) ApiAnalyticsServiceGetEventsMetricRequest {
	r.directoryId = &directoryId
	return r
}

// Optional tag filter (UUID): every link carrying this tag. Send at most one of linkId, directoryId and tagId
func (r ApiAnalyticsServiceGetEventsMetricRequest) TagId(tagId string) ApiAnalyticsServiceGetEventsMetricRequest {
	r.tagId = &tagId
	return r
}

// Optional country filter: visits from any of these countries (ISO 3166-1 alpha-2 codes, e.g. \&quot;FR\&quot;). Combines with the link, directory or tag filter
func (r ApiAnalyticsServiceGetEventsMetricRequest) Countries(countries []string) ApiAnalyticsServiceGetEventsMetricRequest {
	r.countries = &countries
	return r
}

func (r ApiAnalyticsServiceGetEventsMetricRequest) Execute() (*EventsMetric, *http.Response, error) {
	return r.ApiService.AnalyticsServiceGetEventsMetricExecute(r)
}

/*
AnalyticsServiceGetEventsMetric Get events metric

Retrieve aggregated event analytics for the provided time window and optional link filter.

Requires the `analytics:read` scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAnalyticsServiceGetEventsMetricRequest
*/
func (a *AnalyticsAPIService) AnalyticsServiceGetEventsMetric(ctx context.Context) ApiAnalyticsServiceGetEventsMetricRequest {
	return ApiAnalyticsServiceGetEventsMetricRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return EventsMetric
func (a *AnalyticsAPIService) AnalyticsServiceGetEventsMetricExecute(r ApiAnalyticsServiceGetEventsMetricRequest) (*EventsMetric, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *EventsMetric
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnalyticsAPIService.AnalyticsServiceGetEventsMetric")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/metrics/events"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startDate", r.startDate, "form", "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "endDate", r.endDate, "form", "")
	}
	if r.linkId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "linkId", r.linkId, "form", "")
	}
	if r.directoryId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "directoryId", r.directoryId, "form", "")
	}
	if r.tagId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "tagId", r.tagId, "form", "")
	}
	if r.countries != nil {
		t := *r.countries
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "countries", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "countries", t, "form", "multi")
		}
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

type ApiAnalyticsServiceGetGeoBreakdownRequest struct {
	ctx context.Context
	ApiService *AnalyticsAPIService
	getGeoBreakdownRequest *GetGeoBreakdownRequest
	idempotencyKey *string
}

func (r ApiAnalyticsServiceGetGeoBreakdownRequest) GetGeoBreakdownRequest(getGeoBreakdownRequest GetGeoBreakdownRequest) ApiAnalyticsServiceGetGeoBreakdownRequest {
	r.getGeoBreakdownRequest = &getGeoBreakdownRequest
	return r
}

// A unique value per request, such as a UUID, that makes retrying it safe: a retry with the same key and body returns the first response instead of acting twice.
func (r ApiAnalyticsServiceGetGeoBreakdownRequest) IdempotencyKey(idempotencyKey string) ApiAnalyticsServiceGetGeoBreakdownRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ApiAnalyticsServiceGetGeoBreakdownRequest) Execute() (*GetGeoBreakdownResponse, *http.Response, error) {
	return r.ApiService.AnalyticsServiceGetGeoBreakdownExecute(r)
}

/*
AnalyticsServiceGetGeoBreakdown Get geographic breakdown

Retrieve geographic distribution of events grouped by country or city.

Requires the `analytics:read` scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAnalyticsServiceGetGeoBreakdownRequest
*/
func (a *AnalyticsAPIService) AnalyticsServiceGetGeoBreakdown(ctx context.Context) ApiAnalyticsServiceGetGeoBreakdownRequest {
	return ApiAnalyticsServiceGetGeoBreakdownRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetGeoBreakdownResponse
func (a *AnalyticsAPIService) AnalyticsServiceGetGeoBreakdownExecute(r ApiAnalyticsServiceGetGeoBreakdownRequest) (*GetGeoBreakdownResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetGeoBreakdownResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnalyticsAPIService.AnalyticsServiceGetGeoBreakdown")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/analytics/geo"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.getGeoBreakdownRequest == nil {
		return localVarReturnValue, nil, reportError("getGeoBreakdownRequest is required and must be specified")
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
	localVarPostBody = r.getGeoBreakdownRequest
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

type ApiAnalyticsServiceGetGroupedTimeSeriesRequest struct {
	ctx context.Context
	ApiService *AnalyticsAPIService
	getGroupedTimeSeriesRequest *GetGroupedTimeSeriesRequest
	idempotencyKey *string
}

func (r ApiAnalyticsServiceGetGroupedTimeSeriesRequest) GetGroupedTimeSeriesRequest(getGroupedTimeSeriesRequest GetGroupedTimeSeriesRequest) ApiAnalyticsServiceGetGroupedTimeSeriesRequest {
	r.getGroupedTimeSeriesRequest = &getGroupedTimeSeriesRequest
	return r
}

// A unique value per request, such as a UUID, that makes retrying it safe: a retry with the same key and body returns the first response instead of acting twice.
func (r ApiAnalyticsServiceGetGroupedTimeSeriesRequest) IdempotencyKey(idempotencyKey string) ApiAnalyticsServiceGetGroupedTimeSeriesRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ApiAnalyticsServiceGetGroupedTimeSeriesRequest) Execute() (*GetGroupedTimeSeriesResponse, *http.Response, error) {
	return r.ApiService.AnalyticsServiceGetGroupedTimeSeriesExecute(r)
}

/*
AnalyticsServiceGetGroupedTimeSeries Get grouped time series

Retrieve time-bucketed event counts grouped by a dimension like country, city, device type, or visit type.

Requires the `analytics:read` scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAnalyticsServiceGetGroupedTimeSeriesRequest
*/
func (a *AnalyticsAPIService) AnalyticsServiceGetGroupedTimeSeries(ctx context.Context) ApiAnalyticsServiceGetGroupedTimeSeriesRequest {
	return ApiAnalyticsServiceGetGroupedTimeSeriesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetGroupedTimeSeriesResponse
func (a *AnalyticsAPIService) AnalyticsServiceGetGroupedTimeSeriesExecute(r ApiAnalyticsServiceGetGroupedTimeSeriesRequest) (*GetGroupedTimeSeriesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetGroupedTimeSeriesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnalyticsAPIService.AnalyticsServiceGetGroupedTimeSeries")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/analytics/timeseries"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.getGroupedTimeSeriesRequest == nil {
		return localVarReturnValue, nil, reportError("getGroupedTimeSeriesRequest is required and must be specified")
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
	localVarPostBody = r.getGroupedTimeSeriesRequest
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

type ApiAnalyticsServiceGetHourOfDayDistributionRequest struct {
	ctx context.Context
	ApiService *AnalyticsAPIService
	getHourOfDayDistributionRequest *GetHourOfDayDistributionRequest
	idempotencyKey *string
}

func (r ApiAnalyticsServiceGetHourOfDayDistributionRequest) GetHourOfDayDistributionRequest(getHourOfDayDistributionRequest GetHourOfDayDistributionRequest) ApiAnalyticsServiceGetHourOfDayDistributionRequest {
	r.getHourOfDayDistributionRequest = &getHourOfDayDistributionRequest
	return r
}

// A unique value per request, such as a UUID, that makes retrying it safe: a retry with the same key and body returns the first response instead of acting twice.
func (r ApiAnalyticsServiceGetHourOfDayDistributionRequest) IdempotencyKey(idempotencyKey string) ApiAnalyticsServiceGetHourOfDayDistributionRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ApiAnalyticsServiceGetHourOfDayDistributionRequest) Execute() (*GetHourOfDayDistributionResponse, *http.Response, error) {
	return r.ApiService.AnalyticsServiceGetHourOfDayDistributionExecute(r)
}

/*
AnalyticsServiceGetHourOfDayDistribution Get hour of day distribution

Retrieve event counts aggregated by hour of day, in the workspace's time zone, for identifying peak activity times.

Requires the `analytics:read` scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAnalyticsServiceGetHourOfDayDistributionRequest
*/
func (a *AnalyticsAPIService) AnalyticsServiceGetHourOfDayDistribution(ctx context.Context) ApiAnalyticsServiceGetHourOfDayDistributionRequest {
	return ApiAnalyticsServiceGetHourOfDayDistributionRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetHourOfDayDistributionResponse
func (a *AnalyticsAPIService) AnalyticsServiceGetHourOfDayDistributionExecute(r ApiAnalyticsServiceGetHourOfDayDistributionRequest) (*GetHourOfDayDistributionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetHourOfDayDistributionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnalyticsAPIService.AnalyticsServiceGetHourOfDayDistribution")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/analytics/hours"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.getHourOfDayDistributionRequest == nil {
		return localVarReturnValue, nil, reportError("getHourOfDayDistributionRequest is required and must be specified")
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
	localVarPostBody = r.getHourOfDayDistributionRequest
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

type ApiAnalyticsServiceGetSharedLocationsRequest struct {
	ctx context.Context
	ApiService *AnalyticsAPIService
	startDate *time.Time
	endDate *time.Time
	limit *int32
	linkId *string
	directoryId *string
	tagId *string
	countries *[]string
}

// Inclusive start timestamp. Without startDate and endDate the request covers all time.
func (r ApiAnalyticsServiceGetSharedLocationsRequest) StartDate(startDate time.Time) ApiAnalyticsServiceGetSharedLocationsRequest {
	r.startDate = &startDate
	return r
}

// Exclusive end timestamp
func (r ApiAnalyticsServiceGetSharedLocationsRequest) EndDate(endDate time.Time) ApiAnalyticsServiceGetSharedLocationsRequest {
	r.endDate = &endDate
	return r
}

// Maximum number of positions returned, newest first. Defaults to and may not exceed 2000. total still counts every match.
func (r ApiAnalyticsServiceGetSharedLocationsRequest) Limit(limit int32) ApiAnalyticsServiceGetSharedLocationsRequest {
	r.limit = &limit
	return r
}

// Optional link filter (UUID); when omitted covers the whole workspace
func (r ApiAnalyticsServiceGetSharedLocationsRequest) LinkId(linkId string) ApiAnalyticsServiceGetSharedLocationsRequest {
	r.linkId = &linkId
	return r
}

// Optional directory filter (UUID): the links of this directory and of every directory below it. Send at most one of linkId, directoryId and tagId
func (r ApiAnalyticsServiceGetSharedLocationsRequest) DirectoryId(directoryId string) ApiAnalyticsServiceGetSharedLocationsRequest {
	r.directoryId = &directoryId
	return r
}

// Optional tag filter (UUID): every link carrying this tag. Send at most one of linkId, directoryId and tagId
func (r ApiAnalyticsServiceGetSharedLocationsRequest) TagId(tagId string) ApiAnalyticsServiceGetSharedLocationsRequest {
	r.tagId = &tagId
	return r
}

// Optional country filter: visits from any of these countries (ISO 3166-1 alpha-2 codes, e.g. \&quot;FR\&quot;). Combines with the link, directory or tag filter
func (r ApiAnalyticsServiceGetSharedLocationsRequest) Countries(countries []string) ApiAnalyticsServiceGetSharedLocationsRequest {
	r.countries = &countries
	return r
}

func (r ApiAnalyticsServiceGetSharedLocationsRequest) Execute() (*GetSharedLocationsResponse, *http.Response, error) {
	return r.ApiService.AnalyticsServiceGetSharedLocationsExecute(r)
}

/*
AnalyticsServiceGetSharedLocations List shared locations

Returns the precise positions visitors shared through a Location request workflow step, newest first, capped at 2000, with the total count. Without startDate and endDate it covers all time.

Requires the `analytics:read` scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAnalyticsServiceGetSharedLocationsRequest
*/
func (a *AnalyticsAPIService) AnalyticsServiceGetSharedLocations(ctx context.Context) ApiAnalyticsServiceGetSharedLocationsRequest {
	return ApiAnalyticsServiceGetSharedLocationsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetSharedLocationsResponse
func (a *AnalyticsAPIService) AnalyticsServiceGetSharedLocationsExecute(r ApiAnalyticsServiceGetSharedLocationsRequest) (*GetSharedLocationsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetSharedLocationsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnalyticsAPIService.AnalyticsServiceGetSharedLocations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/analytics/shared-locations"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.startDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "startDate", r.startDate, "form", "")
	}
	if r.endDate != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "endDate", r.endDate, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.linkId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "linkId", r.linkId, "form", "")
	}
	if r.directoryId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "directoryId", r.directoryId, "form", "")
	}
	if r.tagId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "tagId", r.tagId, "form", "")
	}
	if r.countries != nil {
		t := *r.countries
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "countries", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "countries", t, "form", "multi")
		}
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

type ApiAnalyticsServiceGetStepFunnelRequest struct {
	ctx context.Context
	ApiService *AnalyticsAPIService
	getStepFunnelRequest *GetStepFunnelRequest
	idempotencyKey *string
}

func (r ApiAnalyticsServiceGetStepFunnelRequest) GetStepFunnelRequest(getStepFunnelRequest GetStepFunnelRequest) ApiAnalyticsServiceGetStepFunnelRequest {
	r.getStepFunnelRequest = &getStepFunnelRequest
	return r
}

// A unique value per request, such as a UUID, that makes retrying it safe: a retry with the same key and body returns the first response instead of acting twice.
func (r ApiAnalyticsServiceGetStepFunnelRequest) IdempotencyKey(idempotencyKey string) ApiAnalyticsServiceGetStepFunnelRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ApiAnalyticsServiceGetStepFunnelRequest) Execute() (*GetStepFunnelResponse, *http.Response, error) {
	return r.ApiService.AnalyticsServiceGetStepFunnelExecute(r)
}

/*
AnalyticsServiceGetStepFunnel Get workflow step funnel

Returns the workflow steps visitors reached, in distinct visitors, ordered from arrival to exit, with the conversion rate and drop-off at each step.

Requires the `analytics:read` scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAnalyticsServiceGetStepFunnelRequest
*/
func (a *AnalyticsAPIService) AnalyticsServiceGetStepFunnel(ctx context.Context) ApiAnalyticsServiceGetStepFunnelRequest {
	return ApiAnalyticsServiceGetStepFunnelRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetStepFunnelResponse
func (a *AnalyticsAPIService) AnalyticsServiceGetStepFunnelExecute(r ApiAnalyticsServiceGetStepFunnelRequest) (*GetStepFunnelResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetStepFunnelResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnalyticsAPIService.AnalyticsServiceGetStepFunnel")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/analytics/funnel"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.getStepFunnelRequest == nil {
		return localVarReturnValue, nil, reportError("getStepFunnelRequest is required and must be specified")
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
	localVarPostBody = r.getStepFunnelRequest
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

type ApiAnalyticsServiceGetTopLinksRequest struct {
	ctx context.Context
	ApiService *AnalyticsAPIService
	getTopLinksRequest *GetTopLinksRequest
	idempotencyKey *string
}

func (r ApiAnalyticsServiceGetTopLinksRequest) GetTopLinksRequest(getTopLinksRequest GetTopLinksRequest) ApiAnalyticsServiceGetTopLinksRequest {
	r.getTopLinksRequest = &getTopLinksRequest
	return r
}

// A unique value per request, such as a UUID, that makes retrying it safe: a retry with the same key and body returns the first response instead of acting twice.
func (r ApiAnalyticsServiceGetTopLinksRequest) IdempotencyKey(idempotencyKey string) ApiAnalyticsServiceGetTopLinksRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ApiAnalyticsServiceGetTopLinksRequest) Execute() (*GetTopLinksResponse, *http.Response, error) {
	return r.ApiService.AnalyticsServiceGetTopLinksExecute(r)
}

/*
AnalyticsServiceGetTopLinks Get top performing links

Retrieve links ranked by event count for identifying best performing QR codes and links.

Requires the `analytics:read` scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAnalyticsServiceGetTopLinksRequest
*/
func (a *AnalyticsAPIService) AnalyticsServiceGetTopLinks(ctx context.Context) ApiAnalyticsServiceGetTopLinksRequest {
	return ApiAnalyticsServiceGetTopLinksRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetTopLinksResponse
func (a *AnalyticsAPIService) AnalyticsServiceGetTopLinksExecute(r ApiAnalyticsServiceGetTopLinksRequest) (*GetTopLinksResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetTopLinksResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnalyticsAPIService.AnalyticsServiceGetTopLinks")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/analytics/top-links"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.getTopLinksRequest == nil {
		return localVarReturnValue, nil, reportError("getTopLinksRequest is required and must be specified")
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
	localVarPostBody = r.getTopLinksRequest
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

type ApiAnalyticsServiceGetTrafficSplitResultsRequest struct {
	ctx context.Context
	ApiService *AnalyticsAPIService
	getTrafficSplitResultsRequest *GetTrafficSplitResultsRequest
	idempotencyKey *string
}

func (r ApiAnalyticsServiceGetTrafficSplitResultsRequest) GetTrafficSplitResultsRequest(getTrafficSplitResultsRequest GetTrafficSplitResultsRequest) ApiAnalyticsServiceGetTrafficSplitResultsRequest {
	r.getTrafficSplitResultsRequest = &getTrafficSplitResultsRequest
	return r
}

// A unique value per request, such as a UUID, that makes retrying it safe: a retry with the same key and body returns the first response instead of acting twice.
func (r ApiAnalyticsServiceGetTrafficSplitResultsRequest) IdempotencyKey(idempotencyKey string) ApiAnalyticsServiceGetTrafficSplitResultsRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ApiAnalyticsServiceGetTrafficSplitResultsRequest) Execute() (*GetTrafficSplitResultsResponse, *http.Response, error) {
	return r.ApiService.AnalyticsServiceGetTrafficSplitResultsExecute(r)
}

/*
AnalyticsServiceGetTrafficSplitResults Get traffic split results

Returns visits and visitors per variant of a traffic split step, leaving out bots, link previews and automated browsers, with whether there is enough data (100 per variant) and a sample ratio mismatch warning. A sticky split, the default, compares visitors against the weights. With sticky off every visit is assigned afresh, so it compares visits and one person may appear under several variants. It does not test significance or name a winner.

Requires the `analytics:read` scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAnalyticsServiceGetTrafficSplitResultsRequest
*/
func (a *AnalyticsAPIService) AnalyticsServiceGetTrafficSplitResults(ctx context.Context) ApiAnalyticsServiceGetTrafficSplitResultsRequest {
	return ApiAnalyticsServiceGetTrafficSplitResultsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetTrafficSplitResultsResponse
func (a *AnalyticsAPIService) AnalyticsServiceGetTrafficSplitResultsExecute(r ApiAnalyticsServiceGetTrafficSplitResultsRequest) (*GetTrafficSplitResultsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetTrafficSplitResultsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnalyticsAPIService.AnalyticsServiceGetTrafficSplitResults")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/analytics/traffic-split"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.getTrafficSplitResultsRequest == nil {
		return localVarReturnValue, nil, reportError("getTrafficSplitResultsRequest is required and must be specified")
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
	localVarPostBody = r.getTrafficSplitResultsRequest
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

type ApiAnalyticsServiceGetVisitorTrendsRequest struct {
	ctx context.Context
	ApiService *AnalyticsAPIService
	getVisitorTrendsRequest *GetVisitorTrendsRequest
	idempotencyKey *string
}

func (r ApiAnalyticsServiceGetVisitorTrendsRequest) GetVisitorTrendsRequest(getVisitorTrendsRequest GetVisitorTrendsRequest) ApiAnalyticsServiceGetVisitorTrendsRequest {
	r.getVisitorTrendsRequest = &getVisitorTrendsRequest
	return r
}

// A unique value per request, such as a UUID, that makes retrying it safe: a retry with the same key and body returns the first response instead of acting twice.
func (r ApiAnalyticsServiceGetVisitorTrendsRequest) IdempotencyKey(idempotencyKey string) ApiAnalyticsServiceGetVisitorTrendsRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ApiAnalyticsServiceGetVisitorTrendsRequest) Execute() (*GetVisitorTrendsResponse, *http.Response, error) {
	return r.ApiService.AnalyticsServiceGetVisitorTrendsExecute(r)
}

/*
AnalyticsServiceGetVisitorTrends Get visitor trends

Retrieve new vs returning visitor counts over time.

Requires the `analytics:read` scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAnalyticsServiceGetVisitorTrendsRequest
*/
func (a *AnalyticsAPIService) AnalyticsServiceGetVisitorTrends(ctx context.Context) ApiAnalyticsServiceGetVisitorTrendsRequest {
	return ApiAnalyticsServiceGetVisitorTrendsRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GetVisitorTrendsResponse
func (a *AnalyticsAPIService) AnalyticsServiceGetVisitorTrendsExecute(r ApiAnalyticsServiceGetVisitorTrendsRequest) (*GetVisitorTrendsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetVisitorTrendsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnalyticsAPIService.AnalyticsServiceGetVisitorTrends")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/analytics/visitors"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.getVisitorTrendsRequest == nil {
		return localVarReturnValue, nil, reportError("getVisitorTrendsRequest is required and must be specified")
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
	localVarPostBody = r.getVisitorTrendsRequest
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

type ApiAnalyticsServiceGetWorkspaceMetricRequest struct {
	ctx context.Context
	ApiService *AnalyticsAPIService
}

func (r ApiAnalyticsServiceGetWorkspaceMetricRequest) Execute() (*WorkspaceMetric, *http.Response, error) {
	return r.ApiService.AnalyticsServiceGetWorkspaceMetricExecute(r)
}

/*
AnalyticsServiceGetWorkspaceMetric Get workspace metric

Return high-level workspace totals covering members, links, and activity volume.

Requires the `analytics:read` scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAnalyticsServiceGetWorkspaceMetricRequest
*/
func (a *AnalyticsAPIService) AnalyticsServiceGetWorkspaceMetric(ctx context.Context) ApiAnalyticsServiceGetWorkspaceMetricRequest {
	return ApiAnalyticsServiceGetWorkspaceMetricRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return WorkspaceMetric
func (a *AnalyticsAPIService) AnalyticsServiceGetWorkspaceMetricExecute(r ApiAnalyticsServiceGetWorkspaceMetricRequest) (*WorkspaceMetric, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *WorkspaceMetric
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnalyticsAPIService.AnalyticsServiceGetWorkspaceMetric")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/metrics/workspace"

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

type ApiAnalyticsServiceListVisitorCountriesRequest struct {
	ctx context.Context
	ApiService *AnalyticsAPIService
}

func (r ApiAnalyticsServiceListVisitorCountriesRequest) Execute() (*ListVisitorCountriesResponse, *http.Response, error) {
	return r.ApiService.AnalyticsServiceListVisitorCountriesExecute(r)
}

/*
AnalyticsServiceListVisitorCountries List visitor countries

Every country the workspace's visits came from, all time, most visited first, without counts. Meant for country filters. For visit numbers per country, use the geo breakdown.

Requires the `analytics:read` scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiAnalyticsServiceListVisitorCountriesRequest
*/
func (a *AnalyticsAPIService) AnalyticsServiceListVisitorCountries(ctx context.Context) ApiAnalyticsServiceListVisitorCountriesRequest {
	return ApiAnalyticsServiceListVisitorCountriesRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListVisitorCountriesResponse
func (a *AnalyticsAPIService) AnalyticsServiceListVisitorCountriesExecute(r ApiAnalyticsServiceListVisitorCountriesRequest) (*ListVisitorCountriesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListVisitorCountriesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AnalyticsAPIService.AnalyticsServiceListVisitorCountries")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/analytics/visitor-countries"

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
