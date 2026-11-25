/*
Linkbreakers API

The Linkbreakers REST API: links, QR codes, workflows, visitors and analytics, as JSON over HTTPS. Send a workspace API token as `Authorization: Bearer <token>`.  ## Entry points  `GET https://api.linkbreakers.com/` and `GET https://api.linkbreakers.com/v1` answer without credentials with the API's name, version and links. This document is served at `https://api.linkbreakers.com/openapi.json`.  ## Rate limits  Every response describes where the caller stands in the IETF `RateLimit-Policy` and `RateLimit` header fields (draft-ietf-httpapi-ratelimit-headers), and in `RateLimit-Limit`, `RateLimit-Remaining`, `RateLimit-Reset` and `X-RateLimit-Limit`, `X-RateLimit-Remaining`, `X-RateLimit-Reset` (a Unix timestamp) for clients that read those. A request past the limit gets `429 Too Many Requests` with `Retry-After` in seconds.  - Workspace API token: 720 requests per 60 seconds, per token. - Dashboard session or OAuth token: 600 requests per 60 seconds, per token. - No credentials: 300 requests per 60 seconds, per client IP.  Each plan also includes a number of API calls per billing cycle. Past it, calls answer `429` with the code `MAX_PUBLIC_API_QUERIED_REACHED` until the cycle renews or the plan changes.  ## Pagination  Every list pages with opaque cursors (AIP-158). Pass `pageSize`, then send the `nextPageToken` of each response back as `pageToken` (or its alias `cursor`) to get the next page. An empty or absent `nextPageToken` means the list is complete. Each list operation describes this in `x-pagination`.  ## Batches and background jobs  `POST /v1/links/bulk` creates up to 250 links in one transactional request; `POST /v1/links/batch` takes the same links as a bare JSON array. `POST /v1/links/bulk/jobs` accepts the same batch and answers `202 Accepted` at once with a link job and a `Location` header; poll `GET /v1/links/jobs/{id}` until `status` is `LINK_JOB_STATUS_SUCCEEDED` (the created link ids are in `linksBulkCreate.linkIds`) or `LINK_JOB_STATUS_FAILED` (`error` says why, and nothing was created). Link jobs can be read for 24 hours after they finish.  ## Paused workspaces  A free workspace whose members have not verified any email address is paused 7 days after sign-up. Its calls then answer `403` with the code `WORKSPACE_UNVERIFIED` and change nothing; QR codes and links keep working for visitors. To lift the pause, ask the workspace owner to verify their email: `POST /v1/members/email-verification` sends the verification email again, and works while the workspace is paused. `unverifiedPauseAt` on the workspace says when the pause starts.  ## Versioning and deprecation  The version is in the path: every operation lives under `/v1`. Within a version, changes are additive only: new operations, new optional request fields, new response fields and new enum values. Clients should ignore fields and enum values they do not know. Anything that would break an existing client ships under a new path version instead.  An operation is retired in two steps. First it is deprecated: it keeps working, is flagged `deprecated: true` in this document, and every response carries a `Deprecation` header field (RFC 9745) with the time it was deprecated and a `Link` with `rel=\"deprecation\"` to the policy. Then, at least 180 days later, it may be removed; that date is announced in advance in a `Sunset` header field (RFC 8594) and as `x-sunset` on the operation. The full policy: https://linkbreakers.com/help/article/api-versioning-and-deprecation-policy  A renamed operation keeps answering its former path while clients still call it, until that path's sunset at the latest: a former path no client calls for 14 days is removed. Responses on a former path carry the same `Deprecation`, `Sunset` and `Link` header fields; this document lists only the current path, and the policy lists the former paths that still answer.  ## Idempotent retries  Every POST, PUT and PATCH accepts an optional `Idempotency-Key` header (draft-ietf-httpapi-idempotency-key-header). Send a new unique value, such as a UUID, with each logical request and reuse it only to retry that request. A retry with the same key and the same request returns the first response, marked `Idempotent-Replayed: true`, instead of running again. Keys are scoped to the token and kept for 24 hours. The same key with a different request is refused with `422`, and a retry while the first request is still running gets `409` with `Retry-After`. Server errors (5xx) and refusals such as 401, 403 and 429 are not stored, so retrying them runs the request again.  ## Identifying your client  Send a `Linkbreakers-Client` request header saying which client made the call: `<client>/<version> (<details>)`, for example `my-integration/1.2.0 (python 3.12; linux)`. The official SDKs, the CLI and the MCP bridge set it for you, as `sdk-python/2.3.1 (python 3.12.4; linux)` or `cli/1.4.0 (darwin arm64)`. It is optional and changes nothing about how a request is handled; it tells us which clients and versions are in use, so we know what can be retired and what to fix first. Keep the details short and free of personal data. Browsers may send it too: it is allowed by CORS.

API version: 1.205.2
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package linkbreakers

import (
	"encoding/json"
)

// checks if the PageThemeElementStyle type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PageThemeElementStyle{}

// PageThemeElementStyle The look of one job. Every setting is optional: one left unset follows the theme's shape (radius, border) and colours, or the look the element always had.
type PageThemeElementStyle struct {
	// Border colour.
	BorderColor *ThemeColor `json:"borderColor,omitempty"`
	// Border width in pixels, 0 to 12. 0 removes the border.
	BorderSize *int32 `json:"borderSize,omitempty"`
	// Its background. An rgba colour makes it translucent (glass).
	Fill *ThemeColor `json:"fill,omitempty"`
	// Corner radius in pixels, 0 to 999 (999 draws a pill).
	Radius *int32 `json:"radius,omitempty"`
	// Its text colour.
	TextColor *ThemeColor `json:"textColor,omitempty"`
}

// NewPageThemeElementStyle instantiates a new PageThemeElementStyle object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPageThemeElementStyle() *PageThemeElementStyle {
	this := PageThemeElementStyle{}
	return &this
}

// NewPageThemeElementStyleWithDefaults instantiates a new PageThemeElementStyle object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPageThemeElementStyleWithDefaults() *PageThemeElementStyle {
	this := PageThemeElementStyle{}
	return &this
}

// GetBorderColor returns the BorderColor field value if set, zero value otherwise.
func (o *PageThemeElementStyle) GetBorderColor() ThemeColor {
	if o == nil || IsNil(o.BorderColor) {
		var ret ThemeColor
		return ret
	}
	return *o.BorderColor
}

// GetBorderColorOk returns a tuple with the BorderColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageThemeElementStyle) GetBorderColorOk() (*ThemeColor, bool) {
	if o == nil || IsNil(o.BorderColor) {
		return nil, false
	}
	return o.BorderColor, true
}

// HasBorderColor returns a boolean if a field has been set.
func (o *PageThemeElementStyle) HasBorderColor() bool {
	if o != nil && !IsNil(o.BorderColor) {
		return true
	}

	return false
}

// SetBorderColor gets a reference to the given ThemeColor and assigns it to the BorderColor field.
func (o *PageThemeElementStyle) SetBorderColor(v ThemeColor) {
	o.BorderColor = &v
}

// GetBorderSize returns the BorderSize field value if set, zero value otherwise.
func (o *PageThemeElementStyle) GetBorderSize() int32 {
	if o == nil || IsNil(o.BorderSize) {
		var ret int32
		return ret
	}
	return *o.BorderSize
}

// GetBorderSizeOk returns a tuple with the BorderSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageThemeElementStyle) GetBorderSizeOk() (*int32, bool) {
	if o == nil || IsNil(o.BorderSize) {
		return nil, false
	}
	return o.BorderSize, true
}

// HasBorderSize returns a boolean if a field has been set.
func (o *PageThemeElementStyle) HasBorderSize() bool {
	if o != nil && !IsNil(o.BorderSize) {
		return true
	}

	return false
}

// SetBorderSize gets a reference to the given int32 and assigns it to the BorderSize field.
func (o *PageThemeElementStyle) SetBorderSize(v int32) {
	o.BorderSize = &v
}

// GetFill returns the Fill field value if set, zero value otherwise.
func (o *PageThemeElementStyle) GetFill() ThemeColor {
	if o == nil || IsNil(o.Fill) {
		var ret ThemeColor
		return ret
	}
	return *o.Fill
}

// GetFillOk returns a tuple with the Fill field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageThemeElementStyle) GetFillOk() (*ThemeColor, bool) {
	if o == nil || IsNil(o.Fill) {
		return nil, false
	}
	return o.Fill, true
}

// HasFill returns a boolean if a field has been set.
func (o *PageThemeElementStyle) HasFill() bool {
	if o != nil && !IsNil(o.Fill) {
		return true
	}

	return false
}

// SetFill gets a reference to the given ThemeColor and assigns it to the Fill field.
func (o *PageThemeElementStyle) SetFill(v ThemeColor) {
	o.Fill = &v
}

// GetRadius returns the Radius field value if set, zero value otherwise.
func (o *PageThemeElementStyle) GetRadius() int32 {
	if o == nil || IsNil(o.Radius) {
		var ret int32
		return ret
	}
	return *o.Radius
}

// GetRadiusOk returns a tuple with the Radius field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageThemeElementStyle) GetRadiusOk() (*int32, bool) {
	if o == nil || IsNil(o.Radius) {
		return nil, false
	}
	return o.Radius, true
}

// HasRadius returns a boolean if a field has been set.
func (o *PageThemeElementStyle) HasRadius() bool {
	if o != nil && !IsNil(o.Radius) {
		return true
	}

	return false
}

// SetRadius gets a reference to the given int32 and assigns it to the Radius field.
func (o *PageThemeElementStyle) SetRadius(v int32) {
	o.Radius = &v
}

// GetTextColor returns the TextColor field value if set, zero value otherwise.
func (o *PageThemeElementStyle) GetTextColor() ThemeColor {
	if o == nil || IsNil(o.TextColor) {
		var ret ThemeColor
		return ret
	}
	return *o.TextColor
}

// GetTextColorOk returns a tuple with the TextColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageThemeElementStyle) GetTextColorOk() (*ThemeColor, bool) {
	if o == nil || IsNil(o.TextColor) {
		return nil, false
	}
	return o.TextColor, true
}

// HasTextColor returns a boolean if a field has been set.
func (o *PageThemeElementStyle) HasTextColor() bool {
	if o != nil && !IsNil(o.TextColor) {
		return true
	}

	return false
}

// SetTextColor gets a reference to the given ThemeColor and assigns it to the TextColor field.
func (o *PageThemeElementStyle) SetTextColor(v ThemeColor) {
	o.TextColor = &v
}

func (o PageThemeElementStyle) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PageThemeElementStyle) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BorderColor) {
		toSerialize["borderColor"] = o.BorderColor
	}
	if !IsNil(o.BorderSize) {
		toSerialize["borderSize"] = o.BorderSize
	}
	if !IsNil(o.Fill) {
		toSerialize["fill"] = o.Fill
	}
	if !IsNil(o.Radius) {
		toSerialize["radius"] = o.Radius
	}
	if !IsNil(o.TextColor) {
		toSerialize["textColor"] = o.TextColor
	}
	return toSerialize, nil
}

type NullablePageThemeElementStyle struct {
	value *PageThemeElementStyle
	isSet bool
}

func (v NullablePageThemeElementStyle) Get() *PageThemeElementStyle {
	return v.value
}

func (v *NullablePageThemeElementStyle) Set(val *PageThemeElementStyle) {
	v.value = val
	v.isSet = true
}

func (v NullablePageThemeElementStyle) IsSet() bool {
	return v.isSet
}

func (v *NullablePageThemeElementStyle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePageThemeElementStyle(val *PageThemeElementStyle) *NullablePageThemeElementStyle {
	return &NullablePageThemeElementStyle{value: val, isSet: true}
}

func (v NullablePageThemeElementStyle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePageThemeElementStyle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


