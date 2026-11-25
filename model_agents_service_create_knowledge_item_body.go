/*
Linkbreakers API

The Linkbreakers REST API: links, QR codes, workflows, visitors and analytics, as JSON over HTTPS. Send a workspace API token as `Authorization: Bearer <token>`.  ## Entry points  `GET https://api.linkbreakers.com/` and `GET https://api.linkbreakers.com/v1` answer without credentials with the API's name, version and links. This document is served at `https://api.linkbreakers.com/openapi.json`.  ## Rate limits  Every response describes where the caller stands in the IETF `RateLimit-Policy` and `RateLimit` header fields (draft-ietf-httpapi-ratelimit-headers), and in `RateLimit-Limit`, `RateLimit-Remaining`, `RateLimit-Reset` and `X-RateLimit-Limit`, `X-RateLimit-Remaining`, `X-RateLimit-Reset` (a Unix timestamp) for clients that read those. A request past the limit gets `429 Too Many Requests` with `Retry-After` in seconds.  - Workspace API token: 720 requests per 60 seconds, per token. - Dashboard session or OAuth token: 600 requests per 60 seconds, per token. - No credentials: 300 requests per 60 seconds, per client IP.  Each plan also includes a number of API calls per billing cycle. Past it, calls answer `429` with the code `MAX_PUBLIC_API_QUERIED_REACHED` until the cycle renews or the plan changes.  ## Pagination  Every list pages with opaque cursors (AIP-158). Pass `pageSize`, then send the `nextPageToken` of each response back as `pageToken` (or its alias `cursor`) to get the next page. An empty or absent `nextPageToken` means the list is complete. Each list operation describes this in `x-pagination`.  ## Batches and background jobs  `POST /v1/links/bulk` creates up to 250 links in one transactional request; `POST /v1/links/batch` takes the same links as a bare JSON array. `POST /v1/links/bulk/jobs` accepts the same batch and answers `202 Accepted` at once with a link job and a `Location` header; poll `GET /v1/links/jobs/{id}` until `status` is `LINK_JOB_STATUS_SUCCEEDED` (the created link ids are in `linksBulkCreate.linkIds`) or `LINK_JOB_STATUS_FAILED` (`error` says why, and nothing was created). Link jobs can be read for 24 hours after they finish.  ## Paused workspaces  A free workspace whose members have not verified any email address is paused 7 days after sign-up. Its calls then answer `403` with the code `WORKSPACE_UNVERIFIED` and change nothing; QR codes and links keep working for visitors. To lift the pause, ask the workspace owner to verify their email: `POST /v1/members/email-verification` sends the verification email again, and works while the workspace is paused. `unverifiedPauseAt` on the workspace says when the pause starts.  ## Versioning and deprecation  The version is in the path: every operation lives under `/v1`. Within a version, changes are additive only: new operations, new optional request fields, new response fields and new enum values. Clients should ignore fields and enum values they do not know. Anything that would break an existing client ships under a new path version instead.  An operation is retired in two steps. First it is deprecated: it keeps working, is flagged `deprecated: true` in this document, and every response carries a `Deprecation` header field (RFC 9745) with the time it was deprecated and a `Link` with `rel=\"deprecation\"` to the policy. Then, at least 180 days later, it may be removed; that date is announced in advance in a `Sunset` header field (RFC 8594) and as `x-sunset` on the operation. The full policy: https://linkbreakers.com/help/article/api-versioning-and-deprecation-policy  A renamed operation keeps answering its former path while clients still call it, until that path's sunset at the latest: a former path no client calls for 14 days is removed. Responses on a former path carry the same `Deprecation`, `Sunset` and `Link` header fields; this document lists only the current path, and the policy lists the former paths that still answer.  ## Idempotent retries  Every POST, PUT and PATCH accepts an optional `Idempotency-Key` header (draft-ietf-httpapi-idempotency-key-header). Send a new unique value, such as a UUID, with each logical request and reuse it only to retry that request. A retry with the same key and the same request returns the first response, marked `Idempotent-Replayed: true`, instead of running again. Keys are scoped to the token and kept for 24 hours. The same key with a different request is refused with `422`, and a retry while the first request is still running gets `409` with `Retry-After`. Server errors (5xx) and refusals such as 401, 403 and 429 are not stored, so retrying them runs the request again.  ## Identifying your client  Send a `Linkbreakers-Client` request header saying which client made the call: `<client>/<version> (<details>)`, for example `my-integration/1.2.0 (python 3.12; linux)`. The official SDKs, the CLI and the MCP bridge set it for you, as `sdk-python/2.3.1 (python 3.12.4; linux)` or `cli/1.4.0 (darwin arm64)`. It is optional and changes nothing about how a request is handled; it tells us which clients and versions are in use, so we know what can be retired and what to fix first. Keep the details short and free of personal data. Browsers may send it too: it is allowed by CORS.

API version: 1.205.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package linkbreakers

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the AgentsServiceCreateKnowledgeItemBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentsServiceCreateKnowledgeItemBody{}

// AgentsServiceCreateKnowledgeItemBody Adds a knowledge item to an agent: Markdown written in it, a web page read from its URL, or a PDF or image of the media library.
type AgentsServiceCreateKnowledgeItemBody struct {
	// The Markdown text, for a Markdown item only.
	Content *string `json:"content,omitempty"`
	// What the item is made of: MARKDOWN takes content, WEB_PAGE takes url, MEDIA takes mediaId.
	Kind AgentKnowledgeItemKind `json:"kind"`
	// The media item to read (UUID), for a media item only: a PDF or an image of the workspace's media library.
	MediaId *string `json:"mediaId,omitempty"`
	// The item's title. Required for a Markdown item; a web page or file left without one takes the page's title or the file's name.
	Title *string `json:"title,omitempty"`
	// The web page's http or https URL, for a web page item only. One item reads one page.
	Url *string `json:"url,omitempty"`
}

type _AgentsServiceCreateKnowledgeItemBody AgentsServiceCreateKnowledgeItemBody

// NewAgentsServiceCreateKnowledgeItemBody instantiates a new AgentsServiceCreateKnowledgeItemBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentsServiceCreateKnowledgeItemBody(kind AgentKnowledgeItemKind) *AgentsServiceCreateKnowledgeItemBody {
	this := AgentsServiceCreateKnowledgeItemBody{}
	this.Kind = kind
	return &this
}

// NewAgentsServiceCreateKnowledgeItemBodyWithDefaults instantiates a new AgentsServiceCreateKnowledgeItemBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentsServiceCreateKnowledgeItemBodyWithDefaults() *AgentsServiceCreateKnowledgeItemBody {
	this := AgentsServiceCreateKnowledgeItemBody{}
	return &this
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *AgentsServiceCreateKnowledgeItemBody) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentsServiceCreateKnowledgeItemBody) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *AgentsServiceCreateKnowledgeItemBody) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *AgentsServiceCreateKnowledgeItemBody) SetContent(v string) {
	o.Content = &v
}

// GetKind returns the Kind field value
func (o *AgentsServiceCreateKnowledgeItemBody) GetKind() AgentKnowledgeItemKind {
	if o == nil {
		var ret AgentKnowledgeItemKind
		return ret
	}

	return o.Kind
}

// GetKindOk returns a tuple with the Kind field value
// and a boolean to check if the value has been set.
func (o *AgentsServiceCreateKnowledgeItemBody) GetKindOk() (*AgentKnowledgeItemKind, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Kind, true
}

// SetKind sets field value
func (o *AgentsServiceCreateKnowledgeItemBody) SetKind(v AgentKnowledgeItemKind) {
	o.Kind = v
}

// GetMediaId returns the MediaId field value if set, zero value otherwise.
func (o *AgentsServiceCreateKnowledgeItemBody) GetMediaId() string {
	if o == nil || IsNil(o.MediaId) {
		var ret string
		return ret
	}
	return *o.MediaId
}

// GetMediaIdOk returns a tuple with the MediaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentsServiceCreateKnowledgeItemBody) GetMediaIdOk() (*string, bool) {
	if o == nil || IsNil(o.MediaId) {
		return nil, false
	}
	return o.MediaId, true
}

// HasMediaId returns a boolean if a field has been set.
func (o *AgentsServiceCreateKnowledgeItemBody) HasMediaId() bool {
	if o != nil && !IsNil(o.MediaId) {
		return true
	}

	return false
}

// SetMediaId gets a reference to the given string and assigns it to the MediaId field.
func (o *AgentsServiceCreateKnowledgeItemBody) SetMediaId(v string) {
	o.MediaId = &v
}

// GetTitle returns the Title field value if set, zero value otherwise.
func (o *AgentsServiceCreateKnowledgeItemBody) GetTitle() string {
	if o == nil || IsNil(o.Title) {
		var ret string
		return ret
	}
	return *o.Title
}

// GetTitleOk returns a tuple with the Title field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentsServiceCreateKnowledgeItemBody) GetTitleOk() (*string, bool) {
	if o == nil || IsNil(o.Title) {
		return nil, false
	}
	return o.Title, true
}

// HasTitle returns a boolean if a field has been set.
func (o *AgentsServiceCreateKnowledgeItemBody) HasTitle() bool {
	if o != nil && !IsNil(o.Title) {
		return true
	}

	return false
}

// SetTitle gets a reference to the given string and assigns it to the Title field.
func (o *AgentsServiceCreateKnowledgeItemBody) SetTitle(v string) {
	o.Title = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *AgentsServiceCreateKnowledgeItemBody) GetUrl() string {
	if o == nil || IsNil(o.Url) {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentsServiceCreateKnowledgeItemBody) GetUrlOk() (*string, bool) {
	if o == nil || IsNil(o.Url) {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *AgentsServiceCreateKnowledgeItemBody) HasUrl() bool {
	if o != nil && !IsNil(o.Url) {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *AgentsServiceCreateKnowledgeItemBody) SetUrl(v string) {
	o.Url = &v
}

func (o AgentsServiceCreateKnowledgeItemBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentsServiceCreateKnowledgeItemBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	toSerialize["kind"] = o.Kind
	if !IsNil(o.MediaId) {
		toSerialize["mediaId"] = o.MediaId
	}
	if !IsNil(o.Title) {
		toSerialize["title"] = o.Title
	}
	if !IsNil(o.Url) {
		toSerialize["url"] = o.Url
	}
	return toSerialize, nil
}

func (o *AgentsServiceCreateKnowledgeItemBody) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"kind",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varAgentsServiceCreateKnowledgeItemBody := _AgentsServiceCreateKnowledgeItemBody{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varAgentsServiceCreateKnowledgeItemBody)

	if err != nil {
		return err
	}

	*o = AgentsServiceCreateKnowledgeItemBody(varAgentsServiceCreateKnowledgeItemBody)

	return err
}

type NullableAgentsServiceCreateKnowledgeItemBody struct {
	value *AgentsServiceCreateKnowledgeItemBody
	isSet bool
}

func (v NullableAgentsServiceCreateKnowledgeItemBody) Get() *AgentsServiceCreateKnowledgeItemBody {
	return v.value
}

func (v *NullableAgentsServiceCreateKnowledgeItemBody) Set(val *AgentsServiceCreateKnowledgeItemBody) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentsServiceCreateKnowledgeItemBody) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentsServiceCreateKnowledgeItemBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentsServiceCreateKnowledgeItemBody(val *AgentsServiceCreateKnowledgeItemBody) *NullableAgentsServiceCreateKnowledgeItemBody {
	return &NullableAgentsServiceCreateKnowledgeItemBody{value: val, isSet: true}
}

func (v NullableAgentsServiceCreateKnowledgeItemBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentsServiceCreateKnowledgeItemBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


