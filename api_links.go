/*
Linkbreakers API

The Linkbreakers REST API: links, QR codes, workflows, visitors and analytics, as JSON over HTTPS. Send a workspace API token as `Authorization: Bearer <token>`.  ## Entry points  `GET https://api.linkbreakers.com/` and `GET https://api.linkbreakers.com/v1` answer without credentials with the API's name, version and links. This document is served at `https://api.linkbreakers.com/openapi.json`.  ## Rate limits  Every response describes where the caller stands in the IETF `RateLimit-Policy` and `RateLimit` header fields (draft-ietf-httpapi-ratelimit-headers), and in `RateLimit-Limit`, `RateLimit-Remaining`, `RateLimit-Reset` and `X-RateLimit-Limit`, `X-RateLimit-Remaining`, `X-RateLimit-Reset` (a Unix timestamp) for clients that read those. A request past the limit gets `429 Too Many Requests` with `Retry-After` in seconds.  - Workspace API token: 720 requests per 60 seconds, per token. - Dashboard session or OAuth token: 600 requests per 60 seconds, per token. - No credentials: 300 requests per 60 seconds, per client IP.  Each plan also includes a number of API calls per billing cycle. Past it, calls answer `429` with the code `MAX_PUBLIC_API_QUERIED_REACHED` until the cycle renews or the plan changes.  ## Pagination  Every list pages with opaque cursors (AIP-158). Pass `pageSize`, then send the `nextPageToken` of each response back as `pageToken` (or its alias `cursor`) to get the next page. An empty or absent `nextPageToken` means the list is complete. Each list operation describes this in `x-pagination`.  ## Batches and background jobs  `POST /v1/links/bulk` creates up to 250 links in one transactional request; `POST /v1/links/batch` takes the same links as a bare JSON array. `POST /v1/links/bulk/jobs` accepts the same batch and answers `202 Accepted` at once with a link job and a `Location` header; poll `GET /v1/links/jobs/{id}` until `status` is `LINK_JOB_STATUS_SUCCEEDED` (the created link ids are in `linksBulkCreate.linkIds`) or `LINK_JOB_STATUS_FAILED` (`error` says why, and nothing was created). Link jobs can be read for 24 hours after they finish.  ## Paused workspaces  A free workspace whose members have not verified any email address is paused 7 days after sign-up. Its calls then answer `403` with the code `WORKSPACE_UNVERIFIED` and change nothing; QR codes and links keep working for visitors. To lift the pause, ask the workspace owner to verify their email: `POST /v1/members/email-verification` sends the verification email again, and works while the workspace is paused. `unverifiedPauseAt` on the workspace says when the pause starts.  ## Versioning and deprecation  The version is in the path: every operation lives under `/v1`. Within a version, changes are additive only: new operations, new optional request fields, new response fields and new enum values. Clients should ignore fields and enum values they do not know. Anything that would break an existing client ships under a new path version instead.  An operation is retired in two steps. First it is deprecated: it keeps working, is flagged `deprecated: true` in this document, and every response carries a `Deprecation` header field (RFC 9745) with the time it was deprecated and a `Link` with `rel=\"deprecation\"` to the policy. Then, at least 180 days later, it may be removed; that date is announced in advance in a `Sunset` header field (RFC 8594) and as `x-sunset` on the operation. The full policy: https://linkbreakers.com/help/article/api-versioning-and-deprecation-policy  A renamed operation keeps answering its former path while clients still call it, until that path's sunset at the latest: a former path no client calls for 14 days is removed. Responses on a former path carry the same `Deprecation`, `Sunset` and `Link` header fields; this document lists only the current path, and the policy lists the former paths that still answer.  ## Idempotent retries  Every POST, PUT and PATCH accepts an optional `Idempotency-Key` header (draft-ietf-httpapi-idempotency-key-header). Send a new unique value, such as a UUID, with each logical request and reuse it only to retry that request. A retry with the same key and the same request returns the first response, marked `Idempotent-Replayed: true`, instead of running again. Keys are scoped to the token and kept for 24 hours. The same key with a different request is refused with `422`, and a retry while the first request is still running gets `409` with `Retry-After`. Server errors (5xx) and refusals such as 401, 403 and 429 are not stored, so retrying them runs the request again.  ## Identifying your client  Send a `Linkbreakers-Client` request header saying which client made the call: `<client>/<version> (<details>)`, for example `my-integration/1.2.0 (python 3.12; linux)`. The official SDKs, the CLI and the MCP bridge set it for you, as `sdk-python/2.3.1 (python 3.12.4; linux)` or `cli/1.4.0 (darwin arm64)`. It is optional and changes nothing about how a request is handled; it tells us which clients and versions are in use, so we know what can be retired and what to fix first. Keep the details short and free of personal data. Browsers may send it too: it is allowed by CORS.

API version: 1.205.0
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
	"reflect"
)


// LinksAPIService LinksAPI service
type LinksAPIService service

type ApiLinksServiceCheckShortlinkAvailabilityRequest struct {
	ctx context.Context
	ApiService *LinksAPIService
	shortlink *string
	customDomainId *string
}

// The shortlink to check availability for.
func (r ApiLinksServiceCheckShortlinkAvailabilityRequest) Shortlink(shortlink string) ApiLinksServiceCheckShortlinkAvailabilityRequest {
	r.shortlink = &shortlink
	return r
}

// The ID of the custom domain to check against (UUID) If not provided, checks against the default domain
func (r ApiLinksServiceCheckShortlinkAvailabilityRequest) CustomDomainId(customDomainId string) ApiLinksServiceCheckShortlinkAvailabilityRequest {
	r.customDomainId = &customDomainId
	return r
}

func (r ApiLinksServiceCheckShortlinkAvailabilityRequest) Execute() (*CheckShortlinkAvailabilityResponse, *http.Response, error) {
	return r.ApiService.LinksServiceCheckShortlinkAvailabilityExecute(r)
}

/*
LinksServiceCheckShortlinkAvailability Check shortlink availability

Checks if a shortlink is available for use to create a new link. A custom domain can be provided to check against.
It typically checks the shortlink slugs and shortlink IDs already present within the system.

Requires the `links:read` scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiLinksServiceCheckShortlinkAvailabilityRequest
*/
func (a *LinksAPIService) LinksServiceCheckShortlinkAvailability(ctx context.Context) ApiLinksServiceCheckShortlinkAvailabilityRequest {
	return ApiLinksServiceCheckShortlinkAvailabilityRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CheckShortlinkAvailabilityResponse
func (a *LinksAPIService) LinksServiceCheckShortlinkAvailabilityExecute(r ApiLinksServiceCheckShortlinkAvailabilityRequest) (*CheckShortlinkAvailabilityResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CheckShortlinkAvailabilityResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinksAPIService.LinksServiceCheckShortlinkAvailability")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/links/shortlink-availability"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.shortlink != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "shortlink", r.shortlink, "form", "")
	}
	if r.customDomainId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "customDomainId", r.customDomainId, "form", "")
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

type ApiLinksServiceCreateRequest struct {
	ctx context.Context
	ApiService *LinksAPIService
	createLinkRequest *CreateLinkRequest
	idempotencyKey *string
}

func (r ApiLinksServiceCreateRequest) CreateLinkRequest(createLinkRequest CreateLinkRequest) ApiLinksServiceCreateRequest {
	r.createLinkRequest = &createLinkRequest
	return r
}

// A unique value per request, such as a UUID, that makes retrying it safe: a retry with the same key and body returns the first response instead of acting twice.
func (r ApiLinksServiceCreateRequest) IdempotencyKey(idempotencyKey string) ApiLinksServiceCreateRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ApiLinksServiceCreateRequest) Execute() (*CreateLinkResponse, *http.Response, error) {
	return r.ApiService.LinksServiceCreateExecute(r)
}

/*
LinksServiceCreate Create a new link

This endpoint allows users to create a new shortened URL link in their Linkbreakers workspace.
The URL will be checked for malicious content before being shortened.

Requires the `links:write` scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiLinksServiceCreateRequest
*/
func (a *LinksAPIService) LinksServiceCreate(ctx context.Context) ApiLinksServiceCreateRequest {
	return ApiLinksServiceCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CreateLinkResponse
func (a *LinksAPIService) LinksServiceCreateExecute(r ApiLinksServiceCreateRequest) (*CreateLinkResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreateLinkResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinksAPIService.LinksServiceCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/links"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createLinkRequest == nil {
		return localVarReturnValue, nil, reportError("createLinkRequest is required and must be specified")
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
	localVarPostBody = r.createLinkRequest
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

type ApiLinksServiceCreateBatchRequest struct {
	ctx context.Context
	ApiService *LinksAPIService
	createLinkRequest *[]CreateLinkRequest
	idempotencyKey *string
}

func (r ApiLinksServiceCreateBatchRequest) CreateLinkRequest(createLinkRequest []CreateLinkRequest) ApiLinksServiceCreateBatchRequest {
	r.createLinkRequest = &createLinkRequest
	return r
}

// A unique value per request, such as a UUID, that makes retrying it safe: a retry with the same key and body returns the first response instead of acting twice.
func (r ApiLinksServiceCreateBatchRequest) IdempotencyKey(idempotencyKey string) ApiLinksServiceCreateBatchRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ApiLinksServiceCreateBatchRequest) Execute() (*CreateBulkLinksResponse, *http.Response, error) {
	return r.ApiService.LinksServiceCreateBatchExecute(r)
}

/*
LinksServiceCreateBatch Create multiple links from a bare JSON array

Creates links in batch using the same per-link options as Create.
This operation is transactional: if one link fails validation or insertion,
no links from the batch are created.

Requires the `links:write` scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiLinksServiceCreateBatchRequest
*/
func (a *LinksAPIService) LinksServiceCreateBatch(ctx context.Context) ApiLinksServiceCreateBatchRequest {
	return ApiLinksServiceCreateBatchRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CreateBulkLinksResponse
func (a *LinksAPIService) LinksServiceCreateBatchExecute(r ApiLinksServiceCreateBatchRequest) (*CreateBulkLinksResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreateBulkLinksResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinksAPIService.LinksServiceCreateBatch")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/links/batch"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createLinkRequest == nil {
		return localVarReturnValue, nil, reportError("createLinkRequest is required and must be specified")
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
	localVarPostBody = r.createLinkRequest
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

type ApiLinksServiceCreateBulkRequest struct {
	ctx context.Context
	ApiService *LinksAPIService
	createBulkLinksRequest *CreateBulkLinksRequest
	idempotencyKey *string
}

func (r ApiLinksServiceCreateBulkRequest) CreateBulkLinksRequest(createBulkLinksRequest CreateBulkLinksRequest) ApiLinksServiceCreateBulkRequest {
	r.createBulkLinksRequest = &createBulkLinksRequest
	return r
}

// A unique value per request, such as a UUID, that makes retrying it safe: a retry with the same key and body returns the first response instead of acting twice.
func (r ApiLinksServiceCreateBulkRequest) IdempotencyKey(idempotencyKey string) ApiLinksServiceCreateBulkRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ApiLinksServiceCreateBulkRequest) Execute() (*CreateBulkLinksResponse, *http.Response, error) {
	return r.ApiService.LinksServiceCreateBulkExecute(r)
}

/*
LinksServiceCreateBulk Create multiple links in a single request

Creates links in batch using the same per-link options as Create.
This operation is transactional: if one link fails validation or insertion,
no links from the batch are created.

Requires the `links:write` scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiLinksServiceCreateBulkRequest
*/
func (a *LinksAPIService) LinksServiceCreateBulk(ctx context.Context) ApiLinksServiceCreateBulkRequest {
	return ApiLinksServiceCreateBulkRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CreateBulkLinksResponse
func (a *LinksAPIService) LinksServiceCreateBulkExecute(r ApiLinksServiceCreateBulkRequest) (*CreateBulkLinksResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreateBulkLinksResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinksAPIService.LinksServiceCreateBulk")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/links/bulk"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createBulkLinksRequest == nil {
		return localVarReturnValue, nil, reportError("createBulkLinksRequest is required and must be specified")
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
	localVarPostBody = r.createBulkLinksRequest
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

type ApiLinksServiceCreateBulkAsyncRequest struct {
	ctx context.Context
	ApiService *LinksAPIService
	createBulkLinksRequest *CreateBulkLinksRequest
	idempotencyKey *string
}

func (r ApiLinksServiceCreateBulkAsyncRequest) CreateBulkLinksRequest(createBulkLinksRequest CreateBulkLinksRequest) ApiLinksServiceCreateBulkAsyncRequest {
	r.createBulkLinksRequest = &createBulkLinksRequest
	return r
}

// A unique value per request, such as a UUID, that makes retrying it safe: a retry with the same key and body returns the first response instead of acting twice.
func (r ApiLinksServiceCreateBulkAsyncRequest) IdempotencyKey(idempotencyKey string) ApiLinksServiceCreateBulkAsyncRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ApiLinksServiceCreateBulkAsyncRequest) Execute() (*LinkJob, *http.Response, error) {
	return r.ApiService.LinksServiceCreateBulkAsyncExecute(r)
}

/*
LinksServiceCreateBulkAsync Create links in bulk in the background

Accepts the same batch as POST /v1/links/bulk, answers 202 Accepted at
once with a link job and a Location header, and creates the links in the
background. Poll GET /v1/links/jobs/{id} for the result. The batch is
transactional: when the job fails, no link was created.

Requires the `links:write` scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiLinksServiceCreateBulkAsyncRequest
*/
func (a *LinksAPIService) LinksServiceCreateBulkAsync(ctx context.Context) ApiLinksServiceCreateBulkAsyncRequest {
	return ApiLinksServiceCreateBulkAsyncRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return LinkJob
func (a *LinksAPIService) LinksServiceCreateBulkAsyncExecute(r ApiLinksServiceCreateBulkAsyncRequest) (*LinkJob, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *LinkJob
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinksAPIService.LinksServiceCreateBulkAsync")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/links/bulk/jobs"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createBulkLinksRequest == nil {
		return localVarReturnValue, nil, reportError("createBulkLinksRequest is required and must be specified")
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
	localVarPostBody = r.createBulkLinksRequest
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

type ApiLinksServiceCreateContactLinkRequest struct {
	ctx context.Context
	ApiService *LinksAPIService
	createContactLinkRequest *CreateContactLinkRequest
	idempotencyKey *string
}

func (r ApiLinksServiceCreateContactLinkRequest) CreateContactLinkRequest(createContactLinkRequest CreateContactLinkRequest) ApiLinksServiceCreateContactLinkRequest {
	r.createContactLinkRequest = &createContactLinkRequest
	return r
}

// A unique value per request, such as a UUID, that makes retrying it safe: a retry with the same key and body returns the first response instead of acting twice.
func (r ApiLinksServiceCreateContactLinkRequest) IdempotencyKey(idempotencyKey string) ApiLinksServiceCreateContactLinkRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ApiLinksServiceCreateContactLinkRequest) Execute() (*CreateContactLinkResponse, *http.Response, error) {
	return r.ApiService.LinksServiceCreateContactLinkExecute(r)
}

/*
LinksServiceCreateContactLink Create a new contact card link

This endpoint allows users to create a new vCard contact link in their Linkbreakers workspace.
The link will generate a vCard file for contact information sharing.

Requires the `links:write` scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiLinksServiceCreateContactLinkRequest
*/
func (a *LinksAPIService) LinksServiceCreateContactLink(ctx context.Context) ApiLinksServiceCreateContactLinkRequest {
	return ApiLinksServiceCreateContactLinkRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CreateContactLinkResponse
func (a *LinksAPIService) LinksServiceCreateContactLinkExecute(r ApiLinksServiceCreateContactLinkRequest) (*CreateContactLinkResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreateContactLinkResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinksAPIService.LinksServiceCreateContactLink")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/links/contact"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createContactLinkRequest == nil {
		return localVarReturnValue, nil, reportError("createContactLinkRequest is required and must be specified")
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
	localVarPostBody = r.createContactLinkRequest
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

type ApiLinksServiceCreateOpenMediaLinkRequest struct {
	ctx context.Context
	ApiService *LinksAPIService
	createOpenMediaLinkRequest *CreateOpenMediaLinkRequest
	idempotencyKey *string
}

func (r ApiLinksServiceCreateOpenMediaLinkRequest) CreateOpenMediaLinkRequest(createOpenMediaLinkRequest CreateOpenMediaLinkRequest) ApiLinksServiceCreateOpenMediaLinkRequest {
	r.createOpenMediaLinkRequest = &createOpenMediaLinkRequest
	return r
}

// A unique value per request, such as a UUID, that makes retrying it safe: a retry with the same key and body returns the first response instead of acting twice.
func (r ApiLinksServiceCreateOpenMediaLinkRequest) IdempotencyKey(idempotencyKey string) ApiLinksServiceCreateOpenMediaLinkRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ApiLinksServiceCreateOpenMediaLinkRequest) Execute() (*CreateOpenMediaLinkResponse, *http.Response, error) {
	return r.ApiService.LinksServiceCreateOpenMediaLinkExecute(r)
}

/*
LinksServiceCreateOpenMediaLink Create a new media link

Creates a link, and its QR code, that hands visitors one piece of media from the media
library: a PDF (a menu or brochure), an image (PNG, JPEG, WebP, SVG) or an MP4 video. Upload
the media first (POST /v1/media/uploads, a PUT of the file, then its complete call; no usage is needed),
then pass its mediaId. Any media of the library can be opened. The Open media step is the exit step. Swap the media later by
pointing the step at another media item.

Requires the `links:write` scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiLinksServiceCreateOpenMediaLinkRequest
*/
func (a *LinksAPIService) LinksServiceCreateOpenMediaLink(ctx context.Context) ApiLinksServiceCreateOpenMediaLinkRequest {
	return ApiLinksServiceCreateOpenMediaLinkRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CreateOpenMediaLinkResponse
func (a *LinksAPIService) LinksServiceCreateOpenMediaLinkExecute(r ApiLinksServiceCreateOpenMediaLinkRequest) (*CreateOpenMediaLinkResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreateOpenMediaLinkResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinksAPIService.LinksServiceCreateOpenMediaLink")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/links/open-media"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createOpenMediaLinkRequest == nil {
		return localVarReturnValue, nil, reportError("createOpenMediaLinkRequest is required and must be specified")
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
	localVarPostBody = r.createOpenMediaLinkRequest
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

type ApiLinksServiceCreateThankYouLinkRequest struct {
	ctx context.Context
	ApiService *LinksAPIService
	createThankYouLinkRequest *CreateThankYouLinkRequest
	idempotencyKey *string
}

func (r ApiLinksServiceCreateThankYouLinkRequest) CreateThankYouLinkRequest(createThankYouLinkRequest CreateThankYouLinkRequest) ApiLinksServiceCreateThankYouLinkRequest {
	r.createThankYouLinkRequest = &createThankYouLinkRequest
	return r
}

// A unique value per request, such as a UUID, that makes retrying it safe: a retry with the same key and body returns the first response instead of acting twice.
func (r ApiLinksServiceCreateThankYouLinkRequest) IdempotencyKey(idempotencyKey string) ApiLinksServiceCreateThankYouLinkRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ApiLinksServiceCreateThankYouLinkRequest) Execute() (*CreateThankYouLinkResponse, *http.Response, error) {
	return r.ApiService.LinksServiceCreateThankYouLinkExecute(r)
}

/*
LinksServiceCreateThankYouLink Create a new thank you page link

This endpoint allows users to create a new link with a thank you page as the exit step.

Requires the `links:write` scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiLinksServiceCreateThankYouLinkRequest
*/
func (a *LinksAPIService) LinksServiceCreateThankYouLink(ctx context.Context) ApiLinksServiceCreateThankYouLinkRequest {
	return ApiLinksServiceCreateThankYouLinkRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CreateThankYouLinkResponse
func (a *LinksAPIService) LinksServiceCreateThankYouLinkExecute(r ApiLinksServiceCreateThankYouLinkRequest) (*CreateThankYouLinkResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreateThankYouLinkResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinksAPIService.LinksServiceCreateThankYouLink")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/links/thank-you"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createThankYouLinkRequest == nil {
		return localVarReturnValue, nil, reportError("createThankYouLinkRequest is required and must be specified")
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
	localVarPostBody = r.createThankYouLinkRequest
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

type ApiLinksServiceDeleteRequest struct {
	ctx context.Context
	ApiService *LinksAPIService
	id string
}

func (r ApiLinksServiceDeleteRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.LinksServiceDeleteExecute(r)
}

/*
LinksServiceDelete Delete a link

Removes a shortened link from the workspace. This operation cannot be undone.

Requires the `links:write` scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The ID of the link to delete (UUID)
 @return ApiLinksServiceDeleteRequest
*/
func (a *LinksAPIService) LinksServiceDelete(ctx context.Context, id string) ApiLinksServiceDeleteRequest {
	return ApiLinksServiceDeleteRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *LinksAPIService) LinksServiceDeleteExecute(r ApiLinksServiceDeleteRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinksAPIService.LinksServiceDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/links/{id}"
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

type ApiLinksServiceGenerateShortlinkRequest struct {
	ctx context.Context
	ApiService *LinksAPIService
}

func (r ApiLinksServiceGenerateShortlinkRequest) Execute() (*GenerateShortlinkResponse, *http.Response, error) {
	return r.ApiService.LinksServiceGenerateShortlinkExecute(r)
}

/*
LinksServiceGenerateShortlink Generate a shortlink

Generates a random shortlink that is still available, to use for a new link.

Requires the `links:read` scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiLinksServiceGenerateShortlinkRequest
*/
func (a *LinksAPIService) LinksServiceGenerateShortlink(ctx context.Context) ApiLinksServiceGenerateShortlinkRequest {
	return ApiLinksServiceGenerateShortlinkRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return GenerateShortlinkResponse
func (a *LinksAPIService) LinksServiceGenerateShortlinkExecute(r ApiLinksServiceGenerateShortlinkRequest) (*GenerateShortlinkResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GenerateShortlinkResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinksAPIService.LinksServiceGenerateShortlink")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/links/shortlink-random"

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

type ApiLinksServiceGetRequest struct {
	ctx context.Context
	ApiService *LinksAPIService
	id string
	include *[]string
}

// List of related resources to include in the response. Supported values: \&quot;customDomain\&quot;, \&quot;qrcodeDesign\&quot;, \&quot;qrcodeSignedUrl\&quot;, \&quot;tags\&quot;, \&quot;pageTheme\&quot;. Use \&quot;qrcodeSignedUrl\&quot; to fetch signed QR image URLs—they are omitted by default to keep requests fast.
func (r ApiLinksServiceGetRequest) Include(include []string) ApiLinksServiceGetRequest {
	r.include = &include
	return r
}

func (r ApiLinksServiceGetRequest) Execute() (*GetLinkResponse, *http.Response, error) {
	return r.ApiService.LinksServiceGetExecute(r)
}

/*
LinksServiceGet Get link details

Retrieves information about a shortened link including its metadata and metrics.

Requires the `links:read` scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The ID of the link to retrieve (UUID)
 @return ApiLinksServiceGetRequest
*/
func (a *LinksAPIService) LinksServiceGet(ctx context.Context, id string) ApiLinksServiceGetRequest {
	return ApiLinksServiceGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return GetLinkResponse
func (a *LinksAPIService) LinksServiceGetExecute(r ApiLinksServiceGetRequest) (*GetLinkResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetLinkResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinksAPIService.LinksServiceGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/links/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.include != nil {
		t := *r.include
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "include", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "include", t, "form", "multi")
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

type ApiLinksServiceListRequest struct {
	ctx context.Context
	ApiService *LinksAPIService
	pageSize *int64
	pageToken *string
	include *[]string
	tags *[]string
	search *string
	sortBy *string
	sortDirection *string
	view *string
	directoryId *string
	cursor *string
}

// The maximum number of links to return.
func (r ApiLinksServiceListRequest) PageSize(pageSize int64) ApiLinksServiceListRequest {
	r.pageSize = &pageSize
	return r
}

// The token for retrieving the next page of results If not provided, the first page of results will be returned
func (r ApiLinksServiceListRequest) PageToken(pageToken string) ApiLinksServiceListRequest {
	r.pageToken = &pageToken
	return r
}

// List of related resources to include in the response. Supported values: \&quot;customDomain\&quot;, \&quot;qrcodeDesign\&quot;, \&quot;qrcodeSignedUrl\&quot;, \&quot;tags\&quot;, \&quot;pageTheme\&quot;. Use \&quot;qrcodeSignedUrl\&quot; to fetch signed QR image URLs—they are omitted by default to keep requests fast.
func (r ApiLinksServiceListRequest) Include(include []string) ApiLinksServiceListRequest {
	r.include = &include
	return r
}

// Tags to filter links by (matches links containing any of the provided tags)
func (r ApiLinksServiceListRequest) Tags(tags []string) ApiLinksServiceListRequest {
	r.tags = &tags
	return r
}

// Search query to filter links by name or shortlink  Search query to filter links by name or shortlink
func (r ApiLinksServiceListRequest) Search(search string) ApiLinksServiceListRequest {
	r.search = &search
	return r
}

// Field to sort links by (defaults to updatedAt if not specified)  Field to sort links by   - LINK_SORT_FIELD_UNSPECIFIED: Sort field not specified (defaults to updatedAt)  - LINK_SORT_FIELD_CREATED_AT: Sort by creation date  - LINK_SORT_FIELD_UPDATED_AT: Sort by last update date  - LINK_SORT_FIELD_SHORTLINK: Sort by shortlink name  - LINK_SORT_FIELD_NAME: Sort by link name  - LINK_SORT_FIELD_EVENT_COUNT: Sort by event count
func (r ApiLinksServiceListRequest) SortBy(sortBy string) ApiLinksServiceListRequest {
	r.sortBy = &sortBy
	return r
}

// Sort direction (defaults to descending if not specified)  Sort direction (ascending or descending)   - SORT_DIRECTION_UNSPECIFIED: Sort direction not specified (defaults to descending)  - SORT_DIRECTION_ASC: Sort in ascending order  - SORT_DIRECTION_DESC: Sort in descending order
func (r ApiLinksServiceListRequest) SortDirection(sortDirection string) ApiLinksServiceListRequest {
	r.sortDirection = &sortDirection
	return r
}

// View mode. Default returns all links. LINK_VIEW_DIRECTORY scopes to a directory (or root when directoryId is empty).   - LINK_VIEW_UNSPECIFIED: Return all links regardless of directory (default)  - LINK_VIEW_DIRECTORY: Return links scoped to a directory. When directoryId is set, returns links in that directory. When directoryId is empty, returns root-level links (no directory).
func (r ApiLinksServiceListRequest) View(view string) ApiLinksServiceListRequest {
	r.view = &view
	return r
}

// Directory ID to scope to when view is LINK_VIEW_DIRECTORY. When empty with LINK_VIEW_DIRECTORY, returns root-level links.
func (r ApiLinksServiceListRequest) DirectoryId(directoryId string) ApiLinksServiceListRequest {
	r.directoryId = &directoryId
	return r
}

// Opaque cursor from the previous page&#39;s nextPageToken. Same as pageToken; when both are sent, pageToken wins.
func (r ApiLinksServiceListRequest) Cursor(cursor string) ApiLinksServiceListRequest {
	r.cursor = &cursor
	return r
}

func (r ApiLinksServiceListRequest) Execute() (*ListLinksResponse, *http.Response, error) {
	return r.ApiService.LinksServiceListExecute(r)
}

/*
LinksServiceList List links

Retrieves the shortened links for your workspace

Requires the `links:read` scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiLinksServiceListRequest
*/
func (a *LinksAPIService) LinksServiceList(ctx context.Context) ApiLinksServiceListRequest {
	return ApiLinksServiceListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListLinksResponse
func (a *LinksAPIService) LinksServiceListExecute(r ApiLinksServiceListRequest) (*ListLinksResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListLinksResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinksAPIService.LinksServiceList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/links"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.pageSize == nil {
		return localVarReturnValue, nil, reportError("pageSize is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "form", "")
	if r.pageToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageToken", r.pageToken, "form", "")
	}
	if r.include != nil {
		t := *r.include
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "include", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "include", t, "form", "multi")
		}
	}
	if r.tags != nil {
		t := *r.tags
		if reflect.TypeOf(t).Kind() == reflect.Slice {
			s := reflect.ValueOf(t)
			for i := 0; i < s.Len(); i++ {
				parameterAddToHeaderOrQuery(localVarQueryParams, "tags", s.Index(i).Interface(), "form", "multi")
			}
		} else {
			parameterAddToHeaderOrQuery(localVarQueryParams, "tags", t, "form", "multi")
		}
	}
	if r.search != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search", r.search, "form", "")
	}
	if r.sortBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sortBy", r.sortBy, "form", "")
	}
	if r.sortDirection != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sortDirection", r.sortDirection, "form", "")
	}
	if r.view != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "view", r.view, "form", "")
	}
	if r.directoryId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "directoryId", r.directoryId, "form", "")
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

type ApiLinksServiceListIntegrationsRequest struct {
	ctx context.Context
	ApiService *LinksAPIService
	linkId string
}

func (r ApiLinksServiceListIntegrationsRequest) Execute() (*ListLinkIntegrationsResponse, *http.Response, error) {
	return r.ApiService.LinksServiceListIntegrationsExecute(r)
}

/*
LinksServiceListIntegrations List the integrations active on a link

Returns every integration that is running and whose settings cover this link: Slack, Google Analytics, form tools, contact tools, automation platforms such as Zapier or Make, webhooks and retargeting pixels. Paused integrations and those limited to other links are left out.

Each one says when it runs: at every workflow step of the kinds in `moments`, and at the steps named in `stepIds`.

Requires the `integrations:read` scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param linkId The link to list the integrations of (UUID).
 @return ApiLinksServiceListIntegrationsRequest
*/
func (a *LinksAPIService) LinksServiceListIntegrations(ctx context.Context, linkId string) ApiLinksServiceListIntegrationsRequest {
	return ApiLinksServiceListIntegrationsRequest{
		ApiService: a,
		ctx: ctx,
		linkId: linkId,
	}
}

// Execute executes the request
//  @return ListLinkIntegrationsResponse
func (a *LinksAPIService) LinksServiceListIntegrationsExecute(r ApiLinksServiceListIntegrationsRequest) (*ListLinkIntegrationsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListLinkIntegrationsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinksAPIService.LinksServiceListIntegrations")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/links/{linkId}/integrations"
	localVarPath = strings.Replace(localVarPath, "{"+"linkId"+"}", url.PathEscape(parameterValueToString(r.linkId, "linkId")), -1)

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

type ApiLinksServiceRenderQrcodeImagesRequest struct {
	ctx context.Context
	ApiService *LinksAPIService
	id string
	linksServiceRenderQrcodeImagesBody *LinksServiceRenderQrcodeImagesBody
	idempotencyKey *string
}

func (r ApiLinksServiceRenderQrcodeImagesRequest) LinksServiceRenderQrcodeImagesBody(linksServiceRenderQrcodeImagesBody LinksServiceRenderQrcodeImagesBody) ApiLinksServiceRenderQrcodeImagesRequest {
	r.linksServiceRenderQrcodeImagesBody = &linksServiceRenderQrcodeImagesBody
	return r
}

// A unique value per request, such as a UUID, that makes retrying it safe: a retry with the same key and body returns the first response instead of acting twice.
func (r ApiLinksServiceRenderQrcodeImagesRequest) IdempotencyKey(idempotencyKey string) ApiLinksServiceRenderQrcodeImagesRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ApiLinksServiceRenderQrcodeImagesRequest) Execute() (*RenderQrcodeImagesResponse, *http.Response, error) {
	return r.ApiService.LinksServiceRenderQrcodeImagesExecute(r)
}

/*
LinksServiceRenderQrcodeImages Render QR codes for a link

Draw the link's QR code in its saved design for each requested address and return signed image URLs. A dynamic link takes one item per code, each with a value for every {name} variable of its shortlink (up to 100 per call), so a single link can print as hundreds of serialised or GS1 Digital Link codes. A regular link takes one item without values. The address encoded in each code carries lbsrc=qrcode so its scans count as QR code scans. Images are cached per design, values, format and size; the URLs expire, so fetch them when you need them.

Requires the `qrcodes:read` scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The link's ID (UUID).
 @return ApiLinksServiceRenderQrcodeImagesRequest
*/
func (a *LinksAPIService) LinksServiceRenderQrcodeImages(ctx context.Context, id string) ApiLinksServiceRenderQrcodeImagesRequest {
	return ApiLinksServiceRenderQrcodeImagesRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return RenderQrcodeImagesResponse
func (a *LinksAPIService) LinksServiceRenderQrcodeImagesExecute(r ApiLinksServiceRenderQrcodeImagesRequest) (*RenderQrcodeImagesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RenderQrcodeImagesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinksAPIService.LinksServiceRenderQrcodeImages")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/links/{id}/qrcode-images:render"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.linksServiceRenderQrcodeImagesBody == nil {
		return localVarReturnValue, nil, reportError("linksServiceRenderQrcodeImagesBody is required and must be specified")
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
	localVarPostBody = r.linksServiceRenderQrcodeImagesBody
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

type ApiLinksServiceUpdateRequest struct {
	ctx context.Context
	ApiService *LinksAPIService
	id string
	linksServiceUpdateBody *LinksServiceUpdateBody
	idempotencyKey *string
}

func (r ApiLinksServiceUpdateRequest) LinksServiceUpdateBody(linksServiceUpdateBody LinksServiceUpdateBody) ApiLinksServiceUpdateRequest {
	r.linksServiceUpdateBody = &linksServiceUpdateBody
	return r
}

// A unique value per request, such as a UUID, that makes retrying it safe: a retry with the same key and body returns the first response instead of acting twice.
func (r ApiLinksServiceUpdateRequest) IdempotencyKey(idempotencyKey string) ApiLinksServiceUpdateRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ApiLinksServiceUpdateRequest) Execute() (*UpdateLinkResponse, *http.Response, error) {
	return r.ApiService.LinksServiceUpdateExecute(r)
}

/*
LinksServiceUpdate Update a link

Updates an existing shortened link in your Linkbreakers workspace.

Requires the `links:write` scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The ID of the link to update (UUID)
 @return ApiLinksServiceUpdateRequest
*/
func (a *LinksAPIService) LinksServiceUpdate(ctx context.Context, id string) ApiLinksServiceUpdateRequest {
	return ApiLinksServiceUpdateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return UpdateLinkResponse
func (a *LinksAPIService) LinksServiceUpdateExecute(r ApiLinksServiceUpdateRequest) (*UpdateLinkResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UpdateLinkResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "LinksAPIService.LinksServiceUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/links/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.linksServiceUpdateBody == nil {
		return localVarReturnValue, nil, reportError("linksServiceUpdateBody is required and must be specified")
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
	localVarPostBody = r.linksServiceUpdateBody
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
