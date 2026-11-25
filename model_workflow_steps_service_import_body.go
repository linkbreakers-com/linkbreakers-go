/*
Linkbreakers API

The Linkbreakers REST API: links, QR codes, workflows, visitors and analytics, as JSON over HTTPS. Send a workspace API token as `Authorization: Bearer <token>`.  ## Entry points  `GET https://api.linkbreakers.com/` and `GET https://api.linkbreakers.com/v1` answer without credentials with the API's name, version and links. This document is served at `https://api.linkbreakers.com/openapi.json`.  ## Rate limits  Every response describes where the caller stands in the IETF `RateLimit-Policy` and `RateLimit` header fields (draft-ietf-httpapi-ratelimit-headers), and in `RateLimit-Limit`, `RateLimit-Remaining`, `RateLimit-Reset` and `X-RateLimit-Limit`, `X-RateLimit-Remaining`, `X-RateLimit-Reset` (a Unix timestamp) for clients that read those. A request past the limit gets `429 Too Many Requests` with `Retry-After` in seconds.  - Workspace API token: 720 requests per 60 seconds, per token. - Dashboard session or OAuth token: 600 requests per 60 seconds, per token. - No credentials: 300 requests per 60 seconds, per client IP.  Each plan also includes a number of API calls per billing cycle. Past it, calls answer `429` with the code `MAX_PUBLIC_API_QUERIED_REACHED` until the cycle renews or the plan changes.  ## Pagination  Every list pages with opaque cursors (AIP-158). Pass `pageSize`, then send the `nextPageToken` of each response back as `pageToken` (or its alias `cursor`) to get the next page. An empty or absent `nextPageToken` means the list is complete. Each list operation describes this in `x-pagination`.  ## Batches and background jobs  `POST /v1/links/bulk` creates up to 250 links in one transactional request; `POST /v1/links/batch` takes the same links as a bare JSON array. `POST /v1/links/bulk/jobs` accepts the same batch and answers `202 Accepted` at once with a link job and a `Location` header; poll `GET /v1/links/jobs/{id}` until `status` is `LINK_JOB_STATUS_SUCCEEDED` (the created link ids are in `linksBulkCreate.linkIds`) or `LINK_JOB_STATUS_FAILED` (`error` says why, and nothing was created). Link jobs can be read for 24 hours after they finish.  ## Paused workspaces  A free workspace whose members have not verified any email address is paused 7 days after sign-up. Its calls then answer `403` with the code `WORKSPACE_UNVERIFIED` and change nothing; QR codes and links keep working for visitors. To lift the pause, ask the workspace owner to verify their email: `POST /v1/members/email-verification` sends the verification email again, and works while the workspace is paused. `unverifiedPauseAt` on the workspace says when the pause starts.  ## Versioning and deprecation  The version is in the path: every operation lives under `/v1`. Within a version, changes are additive only: new operations, new optional request fields, new response fields and new enum values. Clients should ignore fields and enum values they do not know. Anything that would break an existing client ships under a new path version instead.  An operation is retired in two steps. First it is deprecated: it keeps working, is flagged `deprecated: true` in this document, and every response carries a `Deprecation` header field (RFC 9745) with the time it was deprecated and a `Link` with `rel=\"deprecation\"` to the policy. Then, at least 180 days later, it may be removed; that date is announced in advance in a `Sunset` header field (RFC 8594) and as `x-sunset` on the operation. The full policy: https://linkbreakers.com/help/article/api-versioning-and-deprecation-policy  A renamed operation keeps answering its former path while clients still call it, until that path's sunset at the latest: a former path no client calls for 14 days is removed. Responses on a former path carry the same `Deprecation`, `Sunset` and `Link` header fields; this document lists only the current path, and the policy lists the former paths that still answer.  ## Idempotent retries  Every POST, PUT and PATCH accepts an optional `Idempotency-Key` header (draft-ietf-httpapi-idempotency-key-header). Send a new unique value, such as a UUID, with each logical request and reuse it only to retry that request. A retry with the same key and the same request returns the first response, marked `Idempotent-Replayed: true`, instead of running again. Keys are scoped to the token and kept for 24 hours. The same key with a different request is refused with `422`, and a retry while the first request is still running gets `409` with `Retry-After`. Server errors (5xx) and refusals such as 401, 403 and 429 are not stored, so retrying them runs the request again.  ## Identifying your client  Send a `Linkbreakers-Client` request header saying which client made the call: `<client>/<version> (<details>)`, for example `my-integration/1.2.0 (python 3.12; linux)`. The official SDKs, the CLI and the MCP bridge set it for you, as `sdk-python/2.3.1 (python 3.12.4; linux)` or `cli/1.4.0 (darwin arm64)`. It is optional and changes nothing about how a request is handled; it tells us which clients and versions are in use, so we know what can be retired and what to fix first. Keep the details short and free of personal data. Browsers may send it too: it is allowed by CORS.

API version: 1.205.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package linkbreakers

import (
	"encoding/json"
)

// checks if the WorkflowStepsServiceImportBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowStepsServiceImportBody{}

// WorkflowStepsServiceImportBody Imports a document into a link's workflow.
type WorkflowStepsServiceImportBody struct {
	// Where the top left corner of the group goes on the canvas. Left out, the steps go next to the existing ones.
	CanvasPosition *CanvasPosition `json:"canvasPosition,omitempty"`
	// The steps to create: a document made by Export or copied from the dashboard. Leave it out to copy from `sourceLinkId` instead.
	Document *WorkflowStepsDocument `json:"document,omitempty"`
	// Leave this link's entry as it is, whatever the document's entry led to: for steps copied from this same link. Copying from `sourceLinkId` equal to `linkId` does it on its own.
	IgnoreEntry *bool `json:"ignoreEntry,omitempty"`
	// Instead of a document, the link of this workspace to copy steps from, with `sourceStepIds`: the same as importing what Export returns for them.
	SourceLinkId *string `json:"sourceLinkId,omitempty"`
	// The steps of `sourceLinkId` to copy, up to 100. Include its entry to connect this link's entry to the copy of what it leads to.
	SourceStepIds []string `json:"sourceStepIds,omitempty"`
}

// NewWorkflowStepsServiceImportBody instantiates a new WorkflowStepsServiceImportBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowStepsServiceImportBody() *WorkflowStepsServiceImportBody {
	this := WorkflowStepsServiceImportBody{}
	return &this
}

// NewWorkflowStepsServiceImportBodyWithDefaults instantiates a new WorkflowStepsServiceImportBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowStepsServiceImportBodyWithDefaults() *WorkflowStepsServiceImportBody {
	this := WorkflowStepsServiceImportBody{}
	return &this
}

// GetCanvasPosition returns the CanvasPosition field value if set, zero value otherwise.
func (o *WorkflowStepsServiceImportBody) GetCanvasPosition() CanvasPosition {
	if o == nil || IsNil(o.CanvasPosition) {
		var ret CanvasPosition
		return ret
	}
	return *o.CanvasPosition
}

// GetCanvasPositionOk returns a tuple with the CanvasPosition field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStepsServiceImportBody) GetCanvasPositionOk() (*CanvasPosition, bool) {
	if o == nil || IsNil(o.CanvasPosition) {
		return nil, false
	}
	return o.CanvasPosition, true
}

// HasCanvasPosition returns a boolean if a field has been set.
func (o *WorkflowStepsServiceImportBody) HasCanvasPosition() bool {
	if o != nil && !IsNil(o.CanvasPosition) {
		return true
	}

	return false
}

// SetCanvasPosition gets a reference to the given CanvasPosition and assigns it to the CanvasPosition field.
func (o *WorkflowStepsServiceImportBody) SetCanvasPosition(v CanvasPosition) {
	o.CanvasPosition = &v
}

// GetDocument returns the Document field value if set, zero value otherwise.
func (o *WorkflowStepsServiceImportBody) GetDocument() WorkflowStepsDocument {
	if o == nil || IsNil(o.Document) {
		var ret WorkflowStepsDocument
		return ret
	}
	return *o.Document
}

// GetDocumentOk returns a tuple with the Document field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStepsServiceImportBody) GetDocumentOk() (*WorkflowStepsDocument, bool) {
	if o == nil || IsNil(o.Document) {
		return nil, false
	}
	return o.Document, true
}

// HasDocument returns a boolean if a field has been set.
func (o *WorkflowStepsServiceImportBody) HasDocument() bool {
	if o != nil && !IsNil(o.Document) {
		return true
	}

	return false
}

// SetDocument gets a reference to the given WorkflowStepsDocument and assigns it to the Document field.
func (o *WorkflowStepsServiceImportBody) SetDocument(v WorkflowStepsDocument) {
	o.Document = &v
}

// GetIgnoreEntry returns the IgnoreEntry field value if set, zero value otherwise.
func (o *WorkflowStepsServiceImportBody) GetIgnoreEntry() bool {
	if o == nil || IsNil(o.IgnoreEntry) {
		var ret bool
		return ret
	}
	return *o.IgnoreEntry
}

// GetIgnoreEntryOk returns a tuple with the IgnoreEntry field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStepsServiceImportBody) GetIgnoreEntryOk() (*bool, bool) {
	if o == nil || IsNil(o.IgnoreEntry) {
		return nil, false
	}
	return o.IgnoreEntry, true
}

// HasIgnoreEntry returns a boolean if a field has been set.
func (o *WorkflowStepsServiceImportBody) HasIgnoreEntry() bool {
	if o != nil && !IsNil(o.IgnoreEntry) {
		return true
	}

	return false
}

// SetIgnoreEntry gets a reference to the given bool and assigns it to the IgnoreEntry field.
func (o *WorkflowStepsServiceImportBody) SetIgnoreEntry(v bool) {
	o.IgnoreEntry = &v
}

// GetSourceLinkId returns the SourceLinkId field value if set, zero value otherwise.
func (o *WorkflowStepsServiceImportBody) GetSourceLinkId() string {
	if o == nil || IsNil(o.SourceLinkId) {
		var ret string
		return ret
	}
	return *o.SourceLinkId
}

// GetSourceLinkIdOk returns a tuple with the SourceLinkId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStepsServiceImportBody) GetSourceLinkIdOk() (*string, bool) {
	if o == nil || IsNil(o.SourceLinkId) {
		return nil, false
	}
	return o.SourceLinkId, true
}

// HasSourceLinkId returns a boolean if a field has been set.
func (o *WorkflowStepsServiceImportBody) HasSourceLinkId() bool {
	if o != nil && !IsNil(o.SourceLinkId) {
		return true
	}

	return false
}

// SetSourceLinkId gets a reference to the given string and assigns it to the SourceLinkId field.
func (o *WorkflowStepsServiceImportBody) SetSourceLinkId(v string) {
	o.SourceLinkId = &v
}

// GetSourceStepIds returns the SourceStepIds field value if set, zero value otherwise.
func (o *WorkflowStepsServiceImportBody) GetSourceStepIds() []string {
	if o == nil || IsNil(o.SourceStepIds) {
		var ret []string
		return ret
	}
	return o.SourceStepIds
}

// GetSourceStepIdsOk returns a tuple with the SourceStepIds field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStepsServiceImportBody) GetSourceStepIdsOk() ([]string, bool) {
	if o == nil || IsNil(o.SourceStepIds) {
		return nil, false
	}
	return o.SourceStepIds, true
}

// HasSourceStepIds returns a boolean if a field has been set.
func (o *WorkflowStepsServiceImportBody) HasSourceStepIds() bool {
	if o != nil && !IsNil(o.SourceStepIds) {
		return true
	}

	return false
}

// SetSourceStepIds gets a reference to the given []string and assigns it to the SourceStepIds field.
func (o *WorkflowStepsServiceImportBody) SetSourceStepIds(v []string) {
	o.SourceStepIds = v
}

func (o WorkflowStepsServiceImportBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowStepsServiceImportBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CanvasPosition) {
		toSerialize["canvasPosition"] = o.CanvasPosition
	}
	if !IsNil(o.Document) {
		toSerialize["document"] = o.Document
	}
	if !IsNil(o.IgnoreEntry) {
		toSerialize["ignoreEntry"] = o.IgnoreEntry
	}
	if !IsNil(o.SourceLinkId) {
		toSerialize["sourceLinkId"] = o.SourceLinkId
	}
	if !IsNil(o.SourceStepIds) {
		toSerialize["sourceStepIds"] = o.SourceStepIds
	}
	return toSerialize, nil
}

type NullableWorkflowStepsServiceImportBody struct {
	value *WorkflowStepsServiceImportBody
	isSet bool
}

func (v NullableWorkflowStepsServiceImportBody) Get() *WorkflowStepsServiceImportBody {
	return v.value
}

func (v *NullableWorkflowStepsServiceImportBody) Set(val *WorkflowStepsServiceImportBody) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowStepsServiceImportBody) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowStepsServiceImportBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowStepsServiceImportBody(val *WorkflowStepsServiceImportBody) *NullableWorkflowStepsServiceImportBody {
	return &NullableWorkflowStepsServiceImportBody{value: val, isSet: true}
}

func (v NullableWorkflowStepsServiceImportBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowStepsServiceImportBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


