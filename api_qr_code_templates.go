/*
Linkbreakers API

The Linkbreakers REST API: links, QR codes, workflows, visitors and analytics, as JSON over HTTPS. Send a workspace API token as `Authorization: Bearer <token>`.  ## Entry points  `GET https://api.linkbreakers.com/` and `GET https://api.linkbreakers.com/v1` answer without credentials with the API's name, version and links. This document is served at `https://api.linkbreakers.com/openapi.json`.  ## Rate limits  Every response describes where the caller stands in the IETF `RateLimit-Policy` and `RateLimit` header fields (draft-ietf-httpapi-ratelimit-headers), and in `RateLimit-Limit`, `RateLimit-Remaining`, `RateLimit-Reset` and `X-RateLimit-Limit`, `X-RateLimit-Remaining`, `X-RateLimit-Reset` (a Unix timestamp) for clients that read those. A request past the limit gets `429 Too Many Requests` with `Retry-After` in seconds.  - Workspace API token: 720 requests per 60 seconds, per token. - Dashboard session or OAuth token: 600 requests per 60 seconds, per token. - No credentials: 300 requests per 60 seconds, per client IP.  Each plan also includes a number of API calls per billing cycle. Past it, calls answer `429` with the code `MAX_PUBLIC_API_QUERIED_REACHED` until the cycle renews or the plan changes.  ## Pagination  Every list pages with opaque cursors (AIP-158). Pass `pageSize`, then send the `nextPageToken` of each response back as `pageToken` (or its alias `cursor`) to get the next page. An empty or absent `nextPageToken` means the list is complete. Each list operation describes this in `x-pagination`.  ## Batches and background jobs  `POST /v1/links/bulk` creates up to 250 links in one transactional request; `POST /v1/links/batch` takes the same links as a bare JSON array. `POST /v1/links/bulk/jobs` accepts the same batch and answers `202 Accepted` at once with a job and a `Location` header; poll `GET /v1/jobs/{id}` until `status` is `JOB_STATUS_SUCCEEDED` (the created link ids are in `linksBulkCreate.linkIds`) or `JOB_STATUS_FAILED` (`error` says why, and nothing was created). Jobs can be read for 24 hours after they finish.  ## Paused workspaces  A free workspace whose members have not verified any email address is paused 7 days after sign-up. Its calls then answer `403` with the code `WORKSPACE_UNVERIFIED` and change nothing; QR codes and links keep working for visitors. To lift the pause, ask the workspace owner to verify their email: `POST /v1/members/email-verification` sends the verification email again, and works while the workspace is paused. `unverifiedPauseAt` on the workspace says when the pause starts.  ## Versioning and deprecation  The version is in the path: every operation lives under `/v1`. Within a version, changes are additive only: new operations, new optional request fields, new response fields and new enum values. Clients should ignore fields and enum values they do not know. Anything that would break an existing client ships under a new path version instead.  An operation is retired in two steps. First it is deprecated: it keeps working, is flagged `deprecated: true` in this document, and every response carries a `Deprecation` header field (RFC 9745) with the time it was deprecated and a `Link` with `rel=\"deprecation\"` to the policy. Then, at least 180 days later, it may be removed; that date is announced in advance in a `Sunset` header field (RFC 8594) and as `x-sunset` on the operation. The full policy: https://linkbreakers.com/help/article/api-versioning-and-deprecation-policy  A renamed operation keeps answering its former path while clients still call it, until that path's sunset at the latest: a former path no client calls for 14 days is removed. Responses on a former path carry the same `Deprecation`, `Sunset` and `Link` header fields; this document lists only the current path, and the policy lists the former paths that still answer.  ## Idempotent retries  Every POST, PUT and PATCH accepts an optional `Idempotency-Key` header (draft-ietf-httpapi-idempotency-key-header). Send a new unique value, such as a UUID, with each logical request and reuse it only to retry that request. A retry with the same key and the same request returns the first response, marked `Idempotent-Replayed: true`, instead of running again. Keys are scoped to the token and kept for 24 hours. The same key with a different request is refused with `422`, and a retry while the first request is still running gets `409` with `Retry-After`. Server errors (5xx) and refusals such as 401, 403 and 429 are not stored, so retrying them runs the request again.  ## Identifying your client  Send a `Linkbreakers-Client` request header saying which client made the call: `<client>/<version> (<details>)`, for example `my-integration/1.2.0 (python 3.12; linux)`. The official SDKs, the CLI and the MCP bridge set it for you, as `sdk-python/2.3.1 (python 3.12.4; linux)` or `cli/1.4.0 (darwin arm64)`. It is optional and changes nothing about how a request is handled; it tells us which clients and versions are in use, so we know what can be retired and what to fix first. Keep the details short and free of personal data. Browsers may send it too: it is allowed by CORS.

API version: 1.201.0
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


// QRCodeTemplatesAPIService QRCodeTemplatesAPI service
type QRCodeTemplatesAPIService service

type ApiQrcodeTemplatesServiceCreateRequest struct {
	ctx context.Context
	ApiService *QRCodeTemplatesAPIService
	createQrcodeTemplateRequest *CreateQrcodeTemplateRequest
	idempotencyKey *string
}

func (r ApiQrcodeTemplatesServiceCreateRequest) CreateQrcodeTemplateRequest(createQrcodeTemplateRequest CreateQrcodeTemplateRequest) ApiQrcodeTemplatesServiceCreateRequest {
	r.createQrcodeTemplateRequest = &createQrcodeTemplateRequest
	return r
}

// A unique value per request, such as a UUID, that makes retrying it safe: a retry with the same key and body returns the first response instead of acting twice.
func (r ApiQrcodeTemplatesServiceCreateRequest) IdempotencyKey(idempotencyKey string) ApiQrcodeTemplatesServiceCreateRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ApiQrcodeTemplatesServiceCreateRequest) Execute() (*CreateQrcodeTemplateResponse, *http.Response, error) {
	return r.ApiService.QrcodeTemplatesServiceCreateExecute(r)
}

/*
QrcodeTemplatesServiceCreate Create a new QR code template

Create a new QR code template using a fresh design. The design must not be used by any other links or templates.

Requires the `qrcodes:write` scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiQrcodeTemplatesServiceCreateRequest
*/
func (a *QRCodeTemplatesAPIService) QrcodeTemplatesServiceCreate(ctx context.Context) ApiQrcodeTemplatesServiceCreateRequest {
	return ApiQrcodeTemplatesServiceCreateRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CreateQrcodeTemplateResponse
func (a *QRCodeTemplatesAPIService) QrcodeTemplatesServiceCreateExecute(r ApiQrcodeTemplatesServiceCreateRequest) (*CreateQrcodeTemplateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CreateQrcodeTemplateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QRCodeTemplatesAPIService.QrcodeTemplatesServiceCreate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/qrcode-templates"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createQrcodeTemplateRequest == nil {
		return localVarReturnValue, nil, reportError("createQrcodeTemplateRequest is required and must be specified")
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
	localVarPostBody = r.createQrcodeTemplateRequest
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

type ApiQrcodeTemplatesServiceDeleteRequest struct {
	ctx context.Context
	ApiService *QRCodeTemplatesAPIService
	id string
}

func (r ApiQrcodeTemplatesServiceDeleteRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.QrcodeTemplatesServiceDeleteExecute(r)
}

/*
QrcodeTemplatesServiceDelete Delete a workspace QR code template

Remove a QR code template belonging to the authenticated caller's workspace.

Requires the `qrcodes:write` scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiQrcodeTemplatesServiceDeleteRequest
*/
func (a *QRCodeTemplatesAPIService) QrcodeTemplatesServiceDelete(ctx context.Context, id string) ApiQrcodeTemplatesServiceDeleteRequest {
	return ApiQrcodeTemplatesServiceDeleteRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *QRCodeTemplatesAPIService) QrcodeTemplatesServiceDeleteExecute(r ApiQrcodeTemplatesServiceDeleteRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QRCodeTemplatesAPIService.QrcodeTemplatesServiceDelete")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/qrcode-templates/{id}"
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

type ApiQrcodeTemplatesServiceGetRequest struct {
	ctx context.Context
	ApiService *QRCodeTemplatesAPIService
	id string
	include *[]string
}

func (r ApiQrcodeTemplatesServiceGetRequest) Include(include []string) ApiQrcodeTemplatesServiceGetRequest {
	r.include = &include
	return r
}

func (r ApiQrcodeTemplatesServiceGetRequest) Execute() (*QrcodeTemplate, *http.Response, error) {
	return r.ApiService.QrcodeTemplatesServiceGetExecute(r)
}

/*
QrcodeTemplatesServiceGet Get a workspace QR code template

Retrieve a QR code template that belongs to the authenticated caller's workspace.

Requires the `qrcodes:read` scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id
 @return ApiQrcodeTemplatesServiceGetRequest
*/
func (a *QRCodeTemplatesAPIService) QrcodeTemplatesServiceGet(ctx context.Context, id string) ApiQrcodeTemplatesServiceGetRequest {
	return ApiQrcodeTemplatesServiceGetRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return QrcodeTemplate
func (a *QRCodeTemplatesAPIService) QrcodeTemplatesServiceGetExecute(r ApiQrcodeTemplatesServiceGetRequest) (*QrcodeTemplate, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *QrcodeTemplate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QRCodeTemplatesAPIService.QrcodeTemplatesServiceGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/qrcode-templates/{id}"
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

type ApiQrcodeTemplatesServiceGetPublicRequest struct {
	ctx context.Context
	ApiService *QRCodeTemplatesAPIService
	slug string
	include *[]string
}

// Related objects to include in the response  List of related objects to include. Supported values: &#39;workspace&#39;
func (r ApiQrcodeTemplatesServiceGetPublicRequest) Include(include []string) ApiQrcodeTemplatesServiceGetPublicRequest {
	r.include = &include
	return r
}

func (r ApiQrcodeTemplatesServiceGetPublicRequest) Execute() (*PublicQrcodeTemplate, *http.Response, error) {
	return r.ApiService.QrcodeTemplatesServiceGetPublicExecute(r)
}

/*
QrcodeTemplatesServiceGetPublic Get a public QR code template

Retrieve a single published QR code template that is visible in the public gallery.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param slug
 @return ApiQrcodeTemplatesServiceGetPublicRequest
*/
func (a *QRCodeTemplatesAPIService) QrcodeTemplatesServiceGetPublic(ctx context.Context, slug string) ApiQrcodeTemplatesServiceGetPublicRequest {
	return ApiQrcodeTemplatesServiceGetPublicRequest{
		ApiService: a,
		ctx: ctx,
		slug: slug,
	}
}

// Execute executes the request
//  @return PublicQrcodeTemplate
func (a *QRCodeTemplatesAPIService) QrcodeTemplatesServiceGetPublicExecute(r ApiQrcodeTemplatesServiceGetPublicRequest) (*PublicQrcodeTemplate, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PublicQrcodeTemplate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QRCodeTemplatesAPIService.QrcodeTemplatesServiceGetPublic")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/public/qrcode-templates/{slug}"
	localVarPath = strings.Replace(localVarPath, "{"+"slug"+"}", url.PathEscape(parameterValueToString(r.slug, "slug")), -1)

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

type ApiQrcodeTemplatesServiceImportRequest struct {
	ctx context.Context
	ApiService *QRCodeTemplatesAPIService
	importQrcodeTemplateRequest *ImportQrcodeTemplateRequest
	idempotencyKey *string
}

func (r ApiQrcodeTemplatesServiceImportRequest) ImportQrcodeTemplateRequest(importQrcodeTemplateRequest ImportQrcodeTemplateRequest) ApiQrcodeTemplatesServiceImportRequest {
	r.importQrcodeTemplateRequest = &importQrcodeTemplateRequest
	return r
}

// A unique value per request, such as a UUID, that makes retrying it safe: a retry with the same key and body returns the first response instead of acting twice.
func (r ApiQrcodeTemplatesServiceImportRequest) IdempotencyKey(idempotencyKey string) ApiQrcodeTemplatesServiceImportRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ApiQrcodeTemplatesServiceImportRequest) Execute() (*ImportQrcodeTemplateResponse, *http.Response, error) {
	return r.ApiService.QrcodeTemplatesServiceImportExecute(r)
}

/*
QrcodeTemplatesServiceImport Import a public QR code template

Import a public QR code template into your workspace by slug. Creates a private copy owned by your workspace.

Requires the `qrcodes:write` scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiQrcodeTemplatesServiceImportRequest
*/
func (a *QRCodeTemplatesAPIService) QrcodeTemplatesServiceImport(ctx context.Context) ApiQrcodeTemplatesServiceImportRequest {
	return ApiQrcodeTemplatesServiceImportRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ImportQrcodeTemplateResponse
func (a *QRCodeTemplatesAPIService) QrcodeTemplatesServiceImportExecute(r ApiQrcodeTemplatesServiceImportRequest) (*ImportQrcodeTemplateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ImportQrcodeTemplateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QRCodeTemplatesAPIService.QrcodeTemplatesServiceImport")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/qrcode-templates/import"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.importQrcodeTemplateRequest == nil {
		return localVarReturnValue, nil, reportError("importQrcodeTemplateRequest is required and must be specified")
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
	localVarPostBody = r.importQrcodeTemplateRequest
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

type ApiQrcodeTemplatesServiceListRequest struct {
	ctx context.Context
	ApiService *QRCodeTemplatesAPIService
	include *[]string
	search *string
	pageSize *int64
	pageToken *string
	scope *string
	cursor *string
}

func (r ApiQrcodeTemplatesServiceListRequest) Include(include []string) ApiQrcodeTemplatesServiceListRequest {
	r.include = &include
	return r
}

func (r ApiQrcodeTemplatesServiceListRequest) Search(search string) ApiQrcodeTemplatesServiceListRequest {
	r.search = &search
	return r
}

func (r ApiQrcodeTemplatesServiceListRequest) PageSize(pageSize int64) ApiQrcodeTemplatesServiceListRequest {
	r.pageSize = &pageSize
	return r
}

func (r ApiQrcodeTemplatesServiceListRequest) PageToken(pageToken string) ApiQrcodeTemplatesServiceListRequest {
	r.pageToken = &pageToken
	return r
}

// Which templates to return. WORKSPACE (default) returns only templates owned by the caller&#39;s workspace. AVAILABLE also returns built-in templates and templates other workspaces have published, ordered by the workspace&#39;s own usage first, then its own templates, then the rest.   - QRCODE_TEMPLATE_SCOPE_UNSPECIFIED: Defaults to QRCODE_TEMPLATE_SCOPE_WORKSPACE.  - QRCODE_TEMPLATE_SCOPE_WORKSPACE: Only templates owned by the caller&#39;s workspace.  - QRCODE_TEMPLATE_SCOPE_AVAILABLE: Every template the workspace may use: its own, built-in ones, and templates published by other workspaces.
func (r ApiQrcodeTemplatesServiceListRequest) Scope(scope string) ApiQrcodeTemplatesServiceListRequest {
	r.scope = &scope
	return r
}

// Opaque cursor from the previous page&#39;s nextPageToken. Same as pageToken; when both are sent, pageToken wins.
func (r ApiQrcodeTemplatesServiceListRequest) Cursor(cursor string) ApiQrcodeTemplatesServiceListRequest {
	r.cursor = &cursor
	return r
}

func (r ApiQrcodeTemplatesServiceListRequest) Execute() (*ListQrcodeTemplatesResponse, *http.Response, error) {
	return r.ApiService.QrcodeTemplatesServiceListExecute(r)
}

/*
QrcodeTemplatesServiceList List my QR code templates

Retrieve QR code templates that belong to the authenticated caller's workspace.

Requires the `qrcodes:read` scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiQrcodeTemplatesServiceListRequest
*/
func (a *QRCodeTemplatesAPIService) QrcodeTemplatesServiceList(ctx context.Context) ApiQrcodeTemplatesServiceListRequest {
	return ApiQrcodeTemplatesServiceListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListQrcodeTemplatesResponse
func (a *QRCodeTemplatesAPIService) QrcodeTemplatesServiceListExecute(r ApiQrcodeTemplatesServiceListRequest) (*ListQrcodeTemplatesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListQrcodeTemplatesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QRCodeTemplatesAPIService.QrcodeTemplatesServiceList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/qrcode-templates"

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
	if r.search != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search", r.search, "form", "")
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "form", "")
	}
	if r.pageToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageToken", r.pageToken, "form", "")
	}
	if r.scope != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "scope", r.scope, "form", "")
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

type ApiQrcodeTemplatesServiceListPublicRequest struct {
	ctx context.Context
	ApiService *QRCodeTemplatesAPIService
	include *[]string
	search *string
	pageSize *int64
	pageToken *string
	workspaceId *string
	cursor *string
}

// Related objects to include in the response  List of related objects to include. Supported values: &#39;workspace&#39;
func (r ApiQrcodeTemplatesServiceListPublicRequest) Include(include []string) ApiQrcodeTemplatesServiceListPublicRequest {
	r.include = &include
	return r
}

// Search query applied to template name/description
func (r ApiQrcodeTemplatesServiceListPublicRequest) Search(search string) ApiQrcodeTemplatesServiceListPublicRequest {
	r.search = &search
	return r
}

// Maximum number of templates to return (defaults to 20)
func (r ApiQrcodeTemplatesServiceListPublicRequest) PageSize(pageSize int64) ApiQrcodeTemplatesServiceListPublicRequest {
	r.pageSize = &pageSize
	return r
}

// Cursor token to retrieve the next page of results
func (r ApiQrcodeTemplatesServiceListPublicRequest) PageToken(pageToken string) ApiQrcodeTemplatesServiceListPublicRequest {
	r.pageToken = &pageToken
	return r
}

// Optional workspace ID to filter templates by workspace  Filter public templates by workspace. Only published templates for that workspace will be returned.
func (r ApiQrcodeTemplatesServiceListPublicRequest) WorkspaceId(workspaceId string) ApiQrcodeTemplatesServiceListPublicRequest {
	r.workspaceId = &workspaceId
	return r
}

// Opaque cursor from the previous page&#39;s nextPageToken. Same as pageToken; when both are sent, pageToken wins.
func (r ApiQrcodeTemplatesServiceListPublicRequest) Cursor(cursor string) ApiQrcodeTemplatesServiceListPublicRequest {
	r.cursor = &cursor
	return r
}

func (r ApiQrcodeTemplatesServiceListPublicRequest) Execute() (*PublicListQrcodeTemplatesResponse, *http.Response, error) {
	return r.ApiService.QrcodeTemplatesServiceListPublicExecute(r)
}

/*
QrcodeTemplatesServiceListPublic List QR code templates

Retrieve available QR code templates, optionally filtered by search terms or link destination.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiQrcodeTemplatesServiceListPublicRequest
*/
func (a *QRCodeTemplatesAPIService) QrcodeTemplatesServiceListPublic(ctx context.Context) ApiQrcodeTemplatesServiceListPublicRequest {
	return ApiQrcodeTemplatesServiceListPublicRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return PublicListQrcodeTemplatesResponse
func (a *QRCodeTemplatesAPIService) QrcodeTemplatesServiceListPublicExecute(r ApiQrcodeTemplatesServiceListPublicRequest) (*PublicListQrcodeTemplatesResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PublicListQrcodeTemplatesResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QRCodeTemplatesAPIService.QrcodeTemplatesServiceListPublic")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/public/qrcode-templates"

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
	if r.search != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "search", r.search, "form", "")
	}
	if r.pageSize != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageSize", r.pageSize, "form", "")
	}
	if r.pageToken != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "pageToken", r.pageToken, "form", "")
	}
	if r.workspaceId != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "workspaceId", r.workspaceId, "form", "")
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

type ApiQrcodeTemplatesServicePublishRequest struct {
	ctx context.Context
	ApiService *QRCodeTemplatesAPIService
	id string
	body *map[string]interface{}
	idempotencyKey *string
}

func (r ApiQrcodeTemplatesServicePublishRequest) Body(body map[string]interface{}) ApiQrcodeTemplatesServicePublishRequest {
	r.body = &body
	return r
}

// A unique value per request, such as a UUID, that makes retrying it safe: a retry with the same key and body returns the first response instead of acting twice.
func (r ApiQrcodeTemplatesServicePublishRequest) IdempotencyKey(idempotencyKey string) ApiQrcodeTemplatesServicePublishRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ApiQrcodeTemplatesServicePublishRequest) Execute() (*PublishQrcodeTemplateResponse, *http.Response, error) {
	return r.ApiService.QrcodeTemplatesServicePublishExecute(r)
}

/*
QrcodeTemplatesServicePublish Publish a workspace QR code template

Publish a QR code template to make it available in the public gallery. Requires workspace to be public and template to have slug, description, and name.

Requires the `qrcodes:write` scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The ID of the QR code template to publish (UUID)
 @return ApiQrcodeTemplatesServicePublishRequest
*/
func (a *QRCodeTemplatesAPIService) QrcodeTemplatesServicePublish(ctx context.Context, id string) ApiQrcodeTemplatesServicePublishRequest {
	return ApiQrcodeTemplatesServicePublishRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return PublishQrcodeTemplateResponse
func (a *QRCodeTemplatesAPIService) QrcodeTemplatesServicePublishExecute(r ApiQrcodeTemplatesServicePublishRequest) (*PublishQrcodeTemplateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *PublishQrcodeTemplateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QRCodeTemplatesAPIService.QrcodeTemplatesServicePublish")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/qrcode-templates/{id}/publish"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
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
	localVarPostBody = r.body
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

type ApiQrcodeTemplatesServiceUpdateRequest struct {
	ctx context.Context
	ApiService *QRCodeTemplatesAPIService
	id string
	qrcodeTemplatesServiceUpdateBody *QrcodeTemplatesServiceUpdateBody
	idempotencyKey *string
}

func (r ApiQrcodeTemplatesServiceUpdateRequest) QrcodeTemplatesServiceUpdateBody(qrcodeTemplatesServiceUpdateBody QrcodeTemplatesServiceUpdateBody) ApiQrcodeTemplatesServiceUpdateRequest {
	r.qrcodeTemplatesServiceUpdateBody = &qrcodeTemplatesServiceUpdateBody
	return r
}

// A unique value per request, such as a UUID, that makes retrying it safe: a retry with the same key and body returns the first response instead of acting twice.
func (r ApiQrcodeTemplatesServiceUpdateRequest) IdempotencyKey(idempotencyKey string) ApiQrcodeTemplatesServiceUpdateRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ApiQrcodeTemplatesServiceUpdateRequest) Execute() (*UpdateQrcodeTemplateResponse, *http.Response, error) {
	return r.ApiService.QrcodeTemplatesServiceUpdateExecute(r)
}

/*
QrcodeTemplatesServiceUpdate Update a workspace QR code template

Modify the name or description of a QR code template owned by the authenticated workspace.

Requires the `qrcodes:write` scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param id The ID of the QR code template to update (UUID)
 @return ApiQrcodeTemplatesServiceUpdateRequest
*/
func (a *QRCodeTemplatesAPIService) QrcodeTemplatesServiceUpdate(ctx context.Context, id string) ApiQrcodeTemplatesServiceUpdateRequest {
	return ApiQrcodeTemplatesServiceUpdateRequest{
		ApiService: a,
		ctx: ctx,
		id: id,
	}
}

// Execute executes the request
//  @return UpdateQrcodeTemplateResponse
func (a *QRCodeTemplatesAPIService) QrcodeTemplatesServiceUpdateExecute(r ApiQrcodeTemplatesServiceUpdateRequest) (*UpdateQrcodeTemplateResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UpdateQrcodeTemplateResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QRCodeTemplatesAPIService.QrcodeTemplatesServiceUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/qrcode-templates/{id}"
	localVarPath = strings.Replace(localVarPath, "{"+"id"+"}", url.PathEscape(parameterValueToString(r.id, "id")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.qrcodeTemplatesServiceUpdateBody == nil {
		return localVarReturnValue, nil, reportError("qrcodeTemplatesServiceUpdateBody is required and must be specified")
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
	localVarPostBody = r.qrcodeTemplatesServiceUpdateBody
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
