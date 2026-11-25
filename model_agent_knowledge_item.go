/*
Linkbreakers API

The Linkbreakers REST API: links, QR codes, workflows, visitors and analytics, as JSON over HTTPS. Send a workspace API token as `Authorization: Bearer <token>`.  ## Entry points  `GET https://api.linkbreakers.com/` and `GET https://api.linkbreakers.com/v1` answer without credentials with the API's name, version and links. This document is served at `https://api.linkbreakers.com/openapi.json`.  ## Rate limits  Every response describes where the caller stands in the IETF `RateLimit-Policy` and `RateLimit` header fields (draft-ietf-httpapi-ratelimit-headers), and in `RateLimit-Limit`, `RateLimit-Remaining`, `RateLimit-Reset` and `X-RateLimit-Limit`, `X-RateLimit-Remaining`, `X-RateLimit-Reset` (a Unix timestamp) for clients that read those. A request past the limit gets `429 Too Many Requests` with `Retry-After` in seconds.  - Workspace API token: 720 requests per 60 seconds, per token. - Dashboard session or OAuth token: 600 requests per 60 seconds, per token. - No credentials: 300 requests per 60 seconds, per client IP.  Each plan also includes a number of API calls per billing cycle. Past it, calls answer `429` with the code `MAX_PUBLIC_API_QUERIED_REACHED` until the cycle renews or the plan changes.  ## Pagination  Every list pages with opaque cursors (AIP-158). Pass `pageSize`, then send the `nextPageToken` of each response back as `pageToken` (or its alias `cursor`) to get the next page. An empty or absent `nextPageToken` means the list is complete. Each list operation describes this in `x-pagination`.  ## Batches and background jobs  `POST /v1/links/bulk` creates up to 250 links in one transactional request; `POST /v1/links/batch` takes the same links as a bare JSON array. `POST /v1/links/bulk/jobs` accepts the same batch and answers `202 Accepted` at once with a link job and a `Location` header; poll `GET /v1/links/jobs/{id}` until `status` is `LINK_JOB_STATUS_SUCCEEDED` (the created link ids are in `linksBulkCreate.linkIds`) or `LINK_JOB_STATUS_FAILED` (`error` says why, and nothing was created). Link jobs can be read for 24 hours after they finish.  ## Paused workspaces  A free workspace whose members have not verified any email address is paused 7 days after sign-up. Its calls then answer `403` with the code `WORKSPACE_UNVERIFIED` and change nothing; QR codes and links keep working for visitors. To lift the pause, ask the workspace owner to verify their email: `POST /v1/members/email-verification` sends the verification email again, and works while the workspace is paused. `unverifiedPauseAt` on the workspace says when the pause starts.  ## Versioning and deprecation  The version is in the path: every operation lives under `/v1`. Within a version, changes are additive only: new operations, new optional request fields, new response fields and new enum values. Clients should ignore fields and enum values they do not know. Anything that would break an existing client ships under a new path version instead.  An operation is retired in two steps. First it is deprecated: it keeps working, is flagged `deprecated: true` in this document, and every response carries a `Deprecation` header field (RFC 9745) with the time it was deprecated and a `Link` with `rel=\"deprecation\"` to the policy. Then, at least 180 days later, it may be removed; that date is announced in advance in a `Sunset` header field (RFC 8594) and as `x-sunset` on the operation. The full policy: https://linkbreakers.com/help/article/api-versioning-and-deprecation-policy  A renamed operation keeps answering its former path while clients still call it, until that path's sunset at the latest: a former path no client calls for 14 days is removed. Responses on a former path carry the same `Deprecation`, `Sunset` and `Link` header fields; this document lists only the current path, and the policy lists the former paths that still answer.  ## Idempotent retries  Every POST, PUT and PATCH accepts an optional `Idempotency-Key` header (draft-ietf-httpapi-idempotency-key-header). Send a new unique value, such as a UUID, with each logical request and reuse it only to retry that request. A retry with the same key and the same request returns the first response, marked `Idempotent-Replayed: true`, instead of running again. Keys are scoped to the token and kept for 24 hours. The same key with a different request is refused with `422`, and a retry while the first request is still running gets `409` with `Retry-After`. Server errors (5xx) and refusals such as 401, 403 and 429 are not stored, so retrying them runs the request again.  ## Identifying your client  Send a `Linkbreakers-Client` request header saying which client made the call: `<client>/<version> (<details>)`, for example `my-integration/1.2.0 (python 3.12; linux)`. The official SDKs, the CLI and the MCP bridge set it for you, as `sdk-python/2.3.1 (python 3.12.4; linux)` or `cli/1.4.0 (darwin arm64)`. It is optional and changes nothing about how a request is handled; it tells us which clients and versions are in use, so we know what can be retired and what to fix first. Keep the details short and free of personal data. Browsers may send it too: it is allowed by CORS.

API version: 1.205.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package linkbreakers

import (
	"encoding/json"
	"time"
)

// checks if the AgentKnowledgeItem type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentKnowledgeItem{}

// AgentKnowledgeItem One thing the agent knows: Markdown, a web page or a media file. The agent answers from the text of its items.
type AgentKnowledgeItem struct {
	// The agent the item belongs to (UUID).
	AgentId *string `json:"agentId,omitempty"`
	// What the agent knows from this item, in Markdown. For a Markdown item, the text written; for a web page or media item, the text read from it, empty until it is ready.
	Content *string `json:"content,omitempty"`
	// When the item was created.
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	// Why the web page or file could not be read. Set only when the status is failed.
	FailureReason *string `json:"failureReason,omitempty"`
	// The unique identifier of the item (UUID).
	Id *string `json:"id,omitempty"`
	// What the item is made of. It cannot change once the item exists.
	Kind *AgentKnowledgeItemKind `json:"kind,omitempty"`
	// The media item read (UUID): a PDF or an image of the workspace's media library. Set for media items only.
	MediaId *string `json:"mediaId,omitempty"`
	// When the web page or file was last read successfully. Empty for Markdown items and for items never read.
	ReadAt *time.Time `json:"readAt,omitempty"`
	// Whether the item's text is part of what the agent knows.
	Status *AgentKnowledgeItemStatus `json:"status,omitempty"`
	// The item's title. A web page or file added without one takes the page's title or the file's name.
	Title *string `json:"title,omitempty"`
	// When the item was last updated.
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// The web page's http or https URL. Set for web page items only.
	Url *string `json:"url,omitempty"`
}

// NewAgentKnowledgeItem instantiates a new AgentKnowledgeItem object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentKnowledgeItem() *AgentKnowledgeItem {
	this := AgentKnowledgeItem{}
	return &this
}

// NewAgentKnowledgeItemWithDefaults instantiates a new AgentKnowledgeItem object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentKnowledgeItemWithDefaults() *AgentKnowledgeItem {
	this := AgentKnowledgeItem{}
	return &this
}

// GetAgentId returns the AgentId field value if set, zero value otherwise.
func (o *AgentKnowledgeItem) GetAgentId() string {
	if o == nil || IsNil(o.AgentId) {
		var ret string
		return ret
	}
	return *o.AgentId
}

// GetAgentIdOk returns a tuple with the AgentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentKnowledgeItem) GetAgentIdOk() (*string, bool) {
	if o == nil || IsNil(o.AgentId) {
		return nil, false
	}
	return o.AgentId, true
}

// HasAgentId returns a boolean if a field has been set.
func (o *AgentKnowledgeItem) HasAgentId() bool {
	if o != nil && !IsNil(o.AgentId) {
		return true
	}

	return false
}

// SetAgentId gets a reference to the given string and assigns it to the AgentId field.
func (o *AgentKnowledgeItem) SetAgentId(v string) {
	o.AgentId = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *AgentKnowledgeItem) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentKnowledgeItem) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *AgentKnowledgeItem) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *AgentKnowledgeItem) SetContent(v string) {
	o.Content = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *AgentKnowledgeItem) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentKnowledgeItem) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *AgentKnowledgeItem) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *AgentKnowledgeItem) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetFailureReason returns the FailureReason field value if set, zero value otherwise.
func (o *AgentKnowledgeItem) GetFailureReason() string {
	if o == nil || IsNil(o.FailureReason) {
		var ret string
		return ret
	}
	return *o.FailureReason
}

// GetFailureReasonOk returns a tuple with the FailureReason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentKnowledgeItem) GetFailureReasonOk() (*string, bool) {
	if o == nil || IsNil(o.FailureReason) {
		return nil, false
	}
	return o.FailureReason, true
}

// HasFailureReason returns a boolean if a field has been set.
func (o *AgentKnowledgeItem) HasFailureReason() bool {
	if o != nil && !IsNil(o.FailureReason) {
		return true
	}

	return false
}

// SetFailureReason gets a reference to the given string and assigns it to the FailureReason field.
func (o *AgentKnowledgeItem) SetFailureReason(v string) {
	o.FailureReason = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AgentKnowledgeItem) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentKnowledgeItem) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AgentKnowledgeItem) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AgentKnowledgeItem) SetId(v string) {
	o.Id = &v
}

// GetKind returns the Kind field value if set, zero value otherwise.
func (o *AgentKnowledgeItem) GetKind() AgentKnowledgeItemKind {
	if o == nil || IsNil(o.Kind) {
		var ret AgentKnowledgeItemKind
		return ret
	}
	return *o.Kind
}

// GetKindOk returns a tuple with the Kind field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentKnowledgeItem) GetKindOk() (*AgentKnowledgeItemKind, bool) {
	if o == nil || IsNil(o.Kind) {
		return nil, false
	}
	return o.Kind, true
}

// HasKind returns a boolean if a field has been set.
func (o *AgentKnowledgeItem) HasKind() bool {
	if o != nil && !IsNil(o.Kind) {
		return true
	}

	return false
}

// SetKind gets a reference to the given AgentKnowledgeItemKind and assigns it to the Kind field.
func (o *AgentKnowledgeItem) SetKind(v AgentKnowledgeItemKind) {
	o.Kind = &v
}

// GetMediaId returns the MediaId field value if set, zero value otherwise.
func (o *AgentKnowledgeItem) GetMediaId() string {
	if o == nil || IsNil(o.MediaId) {
		var ret string
		return ret
	}
	return *o.MediaId
}

// GetMediaIdOk returns a tuple with the MediaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentKnowledgeItem) GetMediaIdOk() (*string, bool) {
	if o == nil || IsNil(o.MediaId) {
		return nil, false
	}
	return o.MediaId, true
}

// HasMediaId returns a boolean if a field has been set.
func (o *AgentKnowledgeItem) HasMediaId() bool {
	if o != nil && !IsNil(o.MediaId) {
		return true
	}

	return false
}

// SetMediaId gets a reference to the given string and assigns it to the MediaId field.
func (o *AgentKnowledgeItem) SetMediaId(v string) {
	o.MediaId = &v
}

// GetReadAt returns the ReadAt field value if set, zero value otherwise.
func (o *AgentKnowledgeItem) GetReadAt() time.Time {
	if o == nil || IsNil(o.ReadAt) {
		var ret time.Time
		return ret
	}
	return *o.ReadAt
}

// GetReadAtOk returns a tuple with the ReadAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentKnowledgeItem) GetReadAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.ReadAt) {
		return nil, false
	}
	return o.ReadAt, true
}

// HasReadAt returns a boolean if a field has been set.
func (o *AgentKnowledgeItem) HasReadAt() bool {
	if o != nil && !IsNil(o.ReadAt) {
		return true
	}

	return false
}

// SetReadAt gets a reference to the given time.Time and assigns it to the ReadAt field.
func (o *AgentKnowledgeItem) SetReadAt(v time.Time) {
	o.ReadAt = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *AgentKnowledgeItem) GetStatus() AgentKnowledgeItemStatus {
	if o == nil || IsNil(o.Status) {
		var ret AgentKnowledgeItemStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentKnowledgeItem) GetStatusOk() (*AgentKnowledgeItemStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *AgentKnowledgeItem) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given AgentKnowledgeItemStatus and assigns it to the Status field.
func (o *AgentKnowledgeItem) SetStatus(v AgentKnowledgeItemStatus) {
	o.Status = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AgentKnowledgeItem) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentKnowledgeItem) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AgentKnowledgeItem) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AgentKnowledgeItem) SetTitle(v string) {
	o.Title = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *AgentKnowledgeItem) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentKnowledgeItem) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *AgentKnowledgeItem) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *AgentKnowledgeItem) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AgentKnowledgeItem) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentKnowledgeItem) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AgentKnowledgeItem) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AgentKnowledgeItem) SetUrl(v string) {
	o.Url = &v
}

func (o AgentKnowledgeItem) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentKnowledgeItem) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AgentId) {
		toSerialize["agentId"] = o.AgentId
	}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.FailureReason) {
		toSerialize["failureReason"] = o.FailureReason
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Kind) {
		toSerialize["kind"] = o.Kind
	}
	if !IsNil(o.MediaId) {
		toSerialize["mediaId"] = o.MediaId
	}
	if !IsNil(o.ReadAt) {
		toSerialize["readAt"] = o.ReadAt
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

type NullableAgentKnowledgeItem struct {
	value *AgentKnowledgeItem
	isSet bool
}

func (v NullableAgentKnowledgeItem) Get() *AgentKnowledgeItem {
	return v.value
}

func (v *NullableAgentKnowledgeItem) Set(val *AgentKnowledgeItem) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentKnowledgeItem) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentKnowledgeItem) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentKnowledgeItem(val *AgentKnowledgeItem) *NullableAgentKnowledgeItem {
	return &NullableAgentKnowledgeItem{value: val, isSet: true}
}

func (v NullableAgentKnowledgeItem) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentKnowledgeItem) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


