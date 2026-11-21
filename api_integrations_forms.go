/*
Linkbreakers API

The Linkbreakers REST API: links, QR codes, workflows, visitors and analytics, as JSON over HTTPS. Send a workspace API token as `Authorization: Bearer <token>`.  ## Entry points  `GET https://api.linkbreakers.com/` and `GET https://api.linkbreakers.com/v1` answer without credentials with the API's name, version and links. This document is served at `https://api.linkbreakers.com/openapi.json`.  ## Rate limits  Every response describes where the caller stands in the IETF `RateLimit-Policy` and `RateLimit` header fields (draft-ietf-httpapi-ratelimit-headers), and in `RateLimit-Limit`, `RateLimit-Remaining`, `RateLimit-Reset` and `X-RateLimit-Limit`, `X-RateLimit-Remaining`, `X-RateLimit-Reset` (a Unix timestamp) for clients that read those. A request past the limit gets `429 Too Many Requests` with `Retry-After` in seconds.  - Workspace API token: 720 requests per 60 seconds, per token. - Dashboard session or OAuth token: 600 requests per 60 seconds, per token. - No credentials: 300 requests per 60 seconds, per client IP.  Each plan also includes a number of API calls per billing cycle. Past it, calls answer `429` with the code `MAX_PUBLIC_API_QUERIED_REACHED` until the cycle renews or the plan changes.  ## Pagination  Every list pages with opaque cursors (AIP-158). Pass `pageSize`, then send the `nextPageToken` of each response back as `pageToken` (or its alias `cursor`) to get the next page. An empty or absent `nextPageToken` means the list is complete. Each list operation describes this in `x-pagination`.  ## Batches and background jobs  `POST /v1/links/bulk` creates up to 250 links in one transactional request; `POST /v1/links/batch` takes the same links as a bare JSON array. `POST /v1/links/bulk/jobs` accepts the same batch and answers `202 Accepted` at once with a job and a `Location` header; poll `GET /v1/jobs/{id}` until `status` is `JOB_STATUS_SUCCEEDED` (the created link ids are in `linksBulkCreate.linkIds`) or `JOB_STATUS_FAILED` (`error` says why, and nothing was created). Jobs can be read for 24 hours after they finish.  ## Paused workspaces  A free workspace whose members have not verified any email address is paused 7 days after sign-up. Its calls then answer `403` with the code `WORKSPACE_UNVERIFIED` and change nothing; QR codes and links keep working for visitors. To lift the pause, ask the workspace owner to verify their email: `POST /v1/members/email-verification` sends the verification email again, and works while the workspace is paused. `unverifiedPauseAt` on the workspace says when the pause starts.  ## Versioning and deprecation  The version is in the path: every operation lives under `/v1`. Within a version, changes are additive only: new operations, new optional request fields, new response fields and new enum values. Clients should ignore fields and enum values they do not know. Anything that would break an existing client ships under a new path version instead.  An operation is retired in two steps. First it is deprecated: it keeps working, is flagged `deprecated: true` in this document, and every response carries a `Deprecation` header field (RFC 9745) with the time it was deprecated and a `Link` with `rel=\"deprecation\"` to the policy. Then, at least 180 days later, it may be removed; that date is announced in advance in a `Sunset` header field (RFC 8594) and as `x-sunset` on the operation. The full policy: https://linkbreakers.com/help/article/api-versioning-and-deprecation-policy  A renamed operation keeps answering its former path while clients still call it, until that path's sunset at the latest: a former path no client calls for 14 days is removed. Responses on a former path carry the same `Deprecation`, `Sunset` and `Link` header fields; this document lists only the current path, and the policy lists the former paths that still answer.  ## Idempotent retries  Every POST, PUT and PATCH accepts an optional `Idempotency-Key` header (draft-ietf-httpapi-idempotency-key-header). Send a new unique value, such as a UUID, with each logical request and reuse it only to retry that request. A retry with the same key and the same request returns the first response, marked `Idempotent-Replayed: true`, instead of running again. Keys are scoped to the token and kept for 24 hours. The same key with a different request is refused with `422`, and a retry while the first request is still running gets `409` with `Retry-After`. Server errors (5xx) and refusals such as 401, 403 and 429 are not stored, so retrying them runs the request again.  ## Identifying your client  Send a `Linkbreakers-Client` request header saying which client made the call: `<client>/<version> (<details>)`, for example `my-integration/1.2.0 (python 3.12; linux)`. The official SDKs, the CLI and the MCP bridge set it for you, as `sdk-python/2.3.1 (python 3.12.4; linux)` or `cli/1.4.0 (darwin arm64)`. It is optional and changes nothing about how a request is handled; it tells us which clients and versions are in use, so we know what can be retired and what to fix first. Keep the details short and free of personal data. Browsers may send it too: it is allowed by CORS.

API version: 1.202.1
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


// IntegrationsFormsAPIService IntegrationsFormsAPI service
type IntegrationsFormsAPIService service

type ApiFormIntegrationsServiceChooseFormLinkRequest struct {
	ctx context.Context
	ApiService *IntegrationsFormsAPIService
	tool string
	formId string
	formIntegrationsServiceChooseFormLinkBody *FormIntegrationsServiceChooseFormLinkBody
	idempotencyKey *string
}

func (r ApiFormIntegrationsServiceChooseFormLinkRequest) FormIntegrationsServiceChooseFormLinkBody(formIntegrationsServiceChooseFormLinkBody FormIntegrationsServiceChooseFormLinkBody) ApiFormIntegrationsServiceChooseFormLinkRequest {
	r.formIntegrationsServiceChooseFormLinkBody = &formIntegrationsServiceChooseFormLinkBody
	return r
}

// A unique value per request, such as a UUID, that makes retrying it safe: a retry with the same key and body returns the first response instead of acting twice.
func (r ApiFormIntegrationsServiceChooseFormLinkRequest) IdempotencyKey(idempotencyKey string) ApiFormIntegrationsServiceChooseFormLinkRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ApiFormIntegrationsServiceChooseFormLinkRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.FormIntegrationsServiceChooseFormLinkExecute(r)
}

/*
FormIntegrationsServiceChooseFormLink Choose the link of a shared form

When the same form is behind several links and the tool cannot always tell which link a response came from, record the link its responses belong to. Only Google Forms needs it; the other tools' responses carry the visit, and they fail with `FORM_LINK_CHOICE_NOT_NEEDED`. Fails with `GOOGLE_FORMS_LINK_NOT_BEHIND_FORM` when the link does not open the form.

Requires the `integrations:write` scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tool
 @param formId The form's ID at the tool, as listed in the forms found behind links.
 @return ApiFormIntegrationsServiceChooseFormLinkRequest
*/
func (a *IntegrationsFormsAPIService) FormIntegrationsServiceChooseFormLink(ctx context.Context, tool string, formId string) ApiFormIntegrationsServiceChooseFormLinkRequest {
	return ApiFormIntegrationsServiceChooseFormLinkRequest{
		ApiService: a,
		ctx: ctx,
		tool: tool,
		formId: formId,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *IntegrationsFormsAPIService) FormIntegrationsServiceChooseFormLinkExecute(r ApiFormIntegrationsServiceChooseFormLinkRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IntegrationsFormsAPIService.FormIntegrationsServiceChooseFormLink")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/form-integrations/{tool}/forms/{formId}/link"
	localVarPath = strings.Replace(localVarPath, "{"+"tool"+"}", url.PathEscape(parameterValueToString(r.tool, "tool")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"formId"+"}", url.PathEscape(parameterValueToString(r.formId, "formId")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.formIntegrationsServiceChooseFormLinkBody == nil {
		return localVarReturnValue, nil, reportError("formIntegrationsServiceChooseFormLinkBody is required and must be specified")
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
	localVarPostBody = r.formIntegrationsServiceChooseFormLinkBody
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

type ApiFormIntegrationsServiceConnectRequest struct {
	ctx context.Context
	ApiService *IntegrationsFormsAPIService
	tool string
	formIntegrationsServiceConnectBody *FormIntegrationsServiceConnectBody
	idempotencyKey *string
}

func (r ApiFormIntegrationsServiceConnectRequest) FormIntegrationsServiceConnectBody(formIntegrationsServiceConnectBody FormIntegrationsServiceConnectBody) ApiFormIntegrationsServiceConnectRequest {
	r.formIntegrationsServiceConnectBody = &formIntegrationsServiceConnectBody
	return r
}

// A unique value per request, such as a UUID, that makes retrying it safe: a retry with the same key and body returns the first response instead of acting twice.
func (r ApiFormIntegrationsServiceConnectRequest) IdempotencyKey(idempotencyKey string) ApiFormIntegrationsServiceConnectRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ApiFormIntegrationsServiceConnectRequest) Execute() (*ConnectFormToolResponse, *http.Response, error) {
	return r.ApiService.FormIntegrationsServiceConnectExecute(r)
}

/*
FormIntegrationsServiceConnect Connect a form tool with an API key

Connect Tally, Jotform or Paperform with the account's API key. The forms behind the workspace's links are prepared at once: a hidden lbid field and a webhook. Typeform and Google Forms connect by signing in, in the dashboard; Google Forms fails here with `FORM_TOOL_SIGN_IN_REQUIRED`.

Requires the `integrations:write` scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tool
 @return ApiFormIntegrationsServiceConnectRequest
*/
func (a *IntegrationsFormsAPIService) FormIntegrationsServiceConnect(ctx context.Context, tool string) ApiFormIntegrationsServiceConnectRequest {
	return ApiFormIntegrationsServiceConnectRequest{
		ApiService: a,
		ctx: ctx,
		tool: tool,
	}
}

// Execute executes the request
//  @return ConnectFormToolResponse
func (a *IntegrationsFormsAPIService) FormIntegrationsServiceConnectExecute(r ApiFormIntegrationsServiceConnectRequest) (*ConnectFormToolResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ConnectFormToolResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IntegrationsFormsAPIService.FormIntegrationsServiceConnect")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/form-integrations/{tool}/connect"
	localVarPath = strings.Replace(localVarPath, "{"+"tool"+"}", url.PathEscape(parameterValueToString(r.tool, "tool")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.formIntegrationsServiceConnectBody == nil {
		return localVarReturnValue, nil, reportError("formIntegrationsServiceConnectBody is required and must be specified")
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
	localVarPostBody = r.formIntegrationsServiceConnectBody
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

type ApiFormIntegrationsServiceDisconnectRequest struct {
	ctx context.Context
	ApiService *IntegrationsFormsAPIService
	tool string
}

func (r ApiFormIntegrationsServiceDisconnectRequest) Execute() (map[string]interface{}, *http.Response, error) {
	return r.ApiService.FormIntegrationsServiceDisconnectExecute(r)
}

/*
FormIntegrationsServiceDisconnect Disconnect a form tool

Remove our webhooks from its forms, or for Google Forms revoke our access to the Google account, and forget the credential.

Requires the `integrations:write` scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tool
 @return ApiFormIntegrationsServiceDisconnectRequest
*/
func (a *IntegrationsFormsAPIService) FormIntegrationsServiceDisconnect(ctx context.Context, tool string) ApiFormIntegrationsServiceDisconnectRequest {
	return ApiFormIntegrationsServiceDisconnectRequest{
		ApiService: a,
		ctx: ctx,
		tool: tool,
	}
}

// Execute executes the request
//  @return map[string]interface{}
func (a *IntegrationsFormsAPIService) FormIntegrationsServiceDisconnectExecute(r ApiFormIntegrationsServiceDisconnectRequest) (map[string]interface{}, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodDelete
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  map[string]interface{}
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IntegrationsFormsAPIService.FormIntegrationsServiceDisconnect")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/form-integrations/{tool}"
	localVarPath = strings.Replace(localVarPath, "{"+"tool"+"}", url.PathEscape(parameterValueToString(r.tool, "tool")), -1)

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

type ApiFormIntegrationsServiceGetRequest struct {
	ctx context.Context
	ApiService *IntegrationsFormsAPIService
	tool string
}

func (r ApiFormIntegrationsServiceGetRequest) Execute() (*GetFormConnectionResponse, *http.Response, error) {
	return r.ApiService.FormIntegrationsServiceGetExecute(r)
}

/*
FormIntegrationsServiceGet Get a form tool connection

Return one form tool's connection, empty when it is not connected.

Requires the `integrations:read` scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tool
 @return ApiFormIntegrationsServiceGetRequest
*/
func (a *IntegrationsFormsAPIService) FormIntegrationsServiceGet(ctx context.Context, tool string) ApiFormIntegrationsServiceGetRequest {
	return ApiFormIntegrationsServiceGetRequest{
		ApiService: a,
		ctx: ctx,
		tool: tool,
	}
}

// Execute executes the request
//  @return GetFormConnectionResponse
func (a *IntegrationsFormsAPIService) FormIntegrationsServiceGetExecute(r ApiFormIntegrationsServiceGetRequest) (*GetFormConnectionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *GetFormConnectionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IntegrationsFormsAPIService.FormIntegrationsServiceGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/form-integrations/{tool}"
	localVarPath = strings.Replace(localVarPath, "{"+"tool"+"}", url.PathEscape(parameterValueToString(r.tool, "tool")), -1)

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

type ApiFormIntegrationsServiceListRequest struct {
	ctx context.Context
	ApiService *IntegrationsFormsAPIService
}

func (r ApiFormIntegrationsServiceListRequest) Execute() (*ListFormConnectionsResponse, *http.Response, error) {
	return r.ApiService.FormIntegrationsServiceListExecute(r)
}

/*
FormIntegrationsServiceList List form tool connections

Return the workspace's connected form tools: Google Forms, Tally, Typeform, Jotform and Paperform.

Requires the `integrations:read` scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiFormIntegrationsServiceListRequest
*/
func (a *IntegrationsFormsAPIService) FormIntegrationsServiceList(ctx context.Context) ApiFormIntegrationsServiceListRequest {
	return ApiFormIntegrationsServiceListRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ListFormConnectionsResponse
func (a *IntegrationsFormsAPIService) FormIntegrationsServiceListExecute(r ApiFormIntegrationsServiceListRequest) (*ListFormConnectionsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListFormConnectionsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IntegrationsFormsAPIService.FormIntegrationsServiceList")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/form-integrations"

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

type ApiFormIntegrationsServiceListFormsRequest struct {
	ctx context.Context
	ApiService *IntegrationsFormsAPIService
	tool string
}

func (r ApiFormIntegrationsServiceListFormsRequest) Execute() (*ListFormToolFormsResponse, *http.Response, error) {
	return r.ApiService.FormIntegrationsServiceListFormsExecute(r)
}

/*
FormIntegrationsServiceListForms List the forms found behind links

Return the forms of a connected tool behind the workspace's links, each with whether it is ready.

Requires the `integrations:read` scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tool
 @return ApiFormIntegrationsServiceListFormsRequest
*/
func (a *IntegrationsFormsAPIService) FormIntegrationsServiceListForms(ctx context.Context, tool string) ApiFormIntegrationsServiceListFormsRequest {
	return ApiFormIntegrationsServiceListFormsRequest{
		ApiService: a,
		ctx: ctx,
		tool: tool,
	}
}

// Execute executes the request
//  @return ListFormToolFormsResponse
func (a *IntegrationsFormsAPIService) FormIntegrationsServiceListFormsExecute(r ApiFormIntegrationsServiceListFormsRequest) (*ListFormToolFormsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListFormToolFormsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IntegrationsFormsAPIService.FormIntegrationsServiceListForms")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/form-integrations/{tool}/forms"
	localVarPath = strings.Replace(localVarPath, "{"+"tool"+"}", url.PathEscape(parameterValueToString(r.tool, "tool")), -1)

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

type ApiFormIntegrationsServicePrepareFormsRequest struct {
	ctx context.Context
	ApiService *IntegrationsFormsAPIService
	tool string
	body *map[string]interface{}
	idempotencyKey *string
}

func (r ApiFormIntegrationsServicePrepareFormsRequest) Body(body map[string]interface{}) ApiFormIntegrationsServicePrepareFormsRequest {
	r.body = &body
	return r
}

// A unique value per request, such as a UUID, that makes retrying it safe: a retry with the same key and body returns the first response instead of acting twice.
func (r ApiFormIntegrationsServicePrepareFormsRequest) IdempotencyKey(idempotencyKey string) ApiFormIntegrationsServicePrepareFormsRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ApiFormIntegrationsServicePrepareFormsRequest) Execute() (*ListFormToolFormsResponse, *http.Response, error) {
	return r.ApiService.FormIntegrationsServicePrepareFormsExecute(r)
}

/*
FormIntegrationsServicePrepareForms Check the forms again

Add the hidden field and webhook to every form behind a link now, rather than at the next daily check. For Google Forms, which does not send responses as they come, this also reads the responses that arrived since the last read, rather than at the next one a few minutes later.

Requires the `integrations:write` scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tool
 @return ApiFormIntegrationsServicePrepareFormsRequest
*/
func (a *IntegrationsFormsAPIService) FormIntegrationsServicePrepareForms(ctx context.Context, tool string) ApiFormIntegrationsServicePrepareFormsRequest {
	return ApiFormIntegrationsServicePrepareFormsRequest{
		ApiService: a,
		ctx: ctx,
		tool: tool,
	}
}

// Execute executes the request
//  @return ListFormToolFormsResponse
func (a *IntegrationsFormsAPIService) FormIntegrationsServicePrepareFormsExecute(r ApiFormIntegrationsServicePrepareFormsRequest) (*ListFormToolFormsResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ListFormToolFormsResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IntegrationsFormsAPIService.FormIntegrationsServicePrepareForms")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/form-integrations/{tool}/forms/prepare"
	localVarPath = strings.Replace(localVarPath, "{"+"tool"+"}", url.PathEscape(parameterValueToString(r.tool, "tool")), -1)

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

type ApiFormIntegrationsServiceUpdateRequest struct {
	ctx context.Context
	ApiService *IntegrationsFormsAPIService
	tool string
	formIntegrationsServiceUpdateBody *FormIntegrationsServiceUpdateBody
	idempotencyKey *string
}

func (r ApiFormIntegrationsServiceUpdateRequest) FormIntegrationsServiceUpdateBody(formIntegrationsServiceUpdateBody FormIntegrationsServiceUpdateBody) ApiFormIntegrationsServiceUpdateRequest {
	r.formIntegrationsServiceUpdateBody = &formIntegrationsServiceUpdateBody
	return r
}

// A unique value per request, such as a UUID, that makes retrying it safe: a retry with the same key and body returns the first response instead of acting twice.
func (r ApiFormIntegrationsServiceUpdateRequest) IdempotencyKey(idempotencyKey string) ApiFormIntegrationsServiceUpdateRequest {
	r.idempotencyKey = &idempotencyKey
	return r
}

func (r ApiFormIntegrationsServiceUpdateRequest) Execute() (*UpdateFormConnectionResponse, *http.Response, error) {
	return r.ApiService.FormIntegrationsServiceUpdateExecute(r)
}

/*
FormIntegrationsServiceUpdate Update a form tool connection

Pause or resume it, or change which links' forms bring answers in and which answers are kept.

Requires the `integrations:write` scope.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param tool
 @return ApiFormIntegrationsServiceUpdateRequest
*/
func (a *IntegrationsFormsAPIService) FormIntegrationsServiceUpdate(ctx context.Context, tool string) ApiFormIntegrationsServiceUpdateRequest {
	return ApiFormIntegrationsServiceUpdateRequest{
		ApiService: a,
		ctx: ctx,
		tool: tool,
	}
}

// Execute executes the request
//  @return UpdateFormConnectionResponse
func (a *IntegrationsFormsAPIService) FormIntegrationsServiceUpdateExecute(r ApiFormIntegrationsServiceUpdateRequest) (*UpdateFormConnectionResponse, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPatch
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *UpdateFormConnectionResponse
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "IntegrationsFormsAPIService.FormIntegrationsServiceUpdate")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/v1/form-integrations/{tool}"
	localVarPath = strings.Replace(localVarPath, "{"+"tool"+"}", url.PathEscape(parameterValueToString(r.tool, "tool")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.formIntegrationsServiceUpdateBody == nil {
		return localVarReturnValue, nil, reportError("formIntegrationsServiceUpdateBody is required and must be specified")
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
	localVarPostBody = r.formIntegrationsServiceUpdateBody
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
