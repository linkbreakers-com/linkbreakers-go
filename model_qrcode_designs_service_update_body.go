/*
Linkbreakers API

The Linkbreakers REST API: links, QR codes, workflows, visitors and analytics, as JSON over HTTPS. Send a workspace API token as `Authorization: Bearer <token>`.  ## Entry points  `GET https://api.linkbreakers.com/` and `GET https://api.linkbreakers.com/v1` answer without credentials with the API's name, version and links. This document is served at `https://api.linkbreakers.com/openapi.json`.  ## Rate limits  Every response describes where the caller stands in the IETF `RateLimit-Policy` and `RateLimit` header fields (draft-ietf-httpapi-ratelimit-headers), and in `RateLimit-Limit`, `RateLimit-Remaining`, `RateLimit-Reset` and `X-RateLimit-Limit`, `X-RateLimit-Remaining`, `X-RateLimit-Reset` (a Unix timestamp) for clients that read those. A request past the limit gets `429 Too Many Requests` with `Retry-After` in seconds.  - Workspace API token: 720 requests per 60 seconds, per token. - Dashboard session or OAuth token: 600 requests per 60 seconds, per token. - No credentials: 300 requests per 60 seconds, per client IP.  Each plan also includes a number of API calls per billing cycle. Past it, calls answer `429` with the code `MAX_PUBLIC_API_QUERIED_REACHED` until the cycle renews or the plan changes.  ## Pagination  Every list pages with opaque cursors (AIP-158). Pass `pageSize`, then send the `nextPageToken` of each response back as `pageToken` (or its alias `cursor`) to get the next page. An empty or absent `nextPageToken` means the list is complete. Each list operation describes this in `x-pagination`.  ## Batches and background jobs  `POST /v1/links/bulk` creates up to 250 links in one transactional request; `POST /v1/links/batch` takes the same links as a bare JSON array. `POST /v1/links/bulk/jobs` accepts the same batch and answers `202 Accepted` at once with a link job and a `Location` header; poll `GET /v1/links/jobs/{id}` until `status` is `LINK_JOB_STATUS_SUCCEEDED` (the created link ids are in `linksBulkCreate.linkIds`) or `LINK_JOB_STATUS_FAILED` (`error` says why, and nothing was created). Link jobs can be read for 24 hours after they finish.  ## Paused workspaces  A free workspace whose members have not verified any email address is paused 7 days after sign-up. Its calls then answer `403` with the code `WORKSPACE_UNVERIFIED` and change nothing; QR codes and links keep working for visitors. To lift the pause, ask the workspace owner to verify their email: `POST /v1/members/email-verification` sends the verification email again, and works while the workspace is paused. `unverifiedPauseAt` on the workspace says when the pause starts.  ## Versioning and deprecation  The version is in the path: every operation lives under `/v1`. Within a version, changes are additive only: new operations, new optional request fields, new response fields and new enum values. Clients should ignore fields and enum values they do not know. Anything that would break an existing client ships under a new path version instead.  An operation is retired in two steps. First it is deprecated: it keeps working, is flagged `deprecated: true` in this document, and every response carries a `Deprecation` header field (RFC 9745) with the time it was deprecated and a `Link` with `rel=\"deprecation\"` to the policy. Then, at least 180 days later, it may be removed; that date is announced in advance in a `Sunset` header field (RFC 8594) and as `x-sunset` on the operation. The full policy: https://linkbreakers.com/help/article/api-versioning-and-deprecation-policy  A renamed operation keeps answering its former path while clients still call it, until that path's sunset at the latest: a former path no client calls for 14 days is removed. Responses on a former path carry the same `Deprecation`, `Sunset` and `Link` header fields; this document lists only the current path, and the policy lists the former paths that still answer.  ## Idempotent retries  Every POST, PUT and PATCH accepts an optional `Idempotency-Key` header (draft-ietf-httpapi-idempotency-key-header). Send a new unique value, such as a UUID, with each logical request and reuse it only to retry that request. A retry with the same key and the same request returns the first response, marked `Idempotent-Replayed: true`, instead of running again. Keys are scoped to the token and kept for 24 hours. The same key with a different request is refused with `422`, and a retry while the first request is still running gets `409` with `Retry-After`. Server errors (5xx) and refusals such as 401, 403 and 429 are not stored, so retrying them runs the request again.  ## Identifying your client  Send a `Linkbreakers-Client` request header saying which client made the call: `<client>/<version> (<details>)`, for example `my-integration/1.2.0 (python 3.12; linux)`. The official SDKs, the CLI and the MCP bridge set it for you, as `sdk-python/2.3.1 (python 3.12.4; linux)` or `cli/1.4.0 (darwin arm64)`. It is optional and changes nothing about how a request is handled; it tells us which clients and versions are in use, so we know what can be retired and what to fix first. Keep the details short and free of personal data. Browsers may send it too: it is allowed by CORS.

API version: 1.205.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package linkbreakers

import (
	"encoding/json"
)

// checks if the QrcodeDesignsServiceUpdateBody type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QrcodeDesignsServiceUpdateBody{}

// QrcodeDesignsServiceUpdateBody Updates colors, shapes, and media used when regenerating QR assets.
type QrcodeDesignsServiceUpdateBody struct {
	BackgroundOptions *BackgroundOptions `json:"backgroundOptions,omitempty"`
	// Set this to true to delete the central image. It can't be used at the same time as centralImageMediaId.
	CentralImageDelete *bool `json:"centralImageDelete,omitempty"`
	CentralImageMargin *int32 `json:"centralImageMargin,omitempty"`
	// An image from the workspace's media library to use as the central image (UUID). Upload the file first (POST /v1/media/uploads, a PUT of its bytes, then its complete call) or import it from a URL, and pass the media id: the image itself is never sent here. Any READY image of at most 3 MB can be a QR code logo; anything else is refused with MEDIA_NOT_ALLOWED_HERE, and media of another workspace or an upload not completed with QRCODE_DESIGN_INVALID_LOGO_MEDIA. The design keeps its own copy of the image, and the media lists this QR code among its references.
	CentralImageMediaId *string `json:"centralImageMediaId,omitempty"`
	CentralImageSize *float64 `json:"centralImageSize,omitempty"`
	CornersDotOptions *CornersDotOptions `json:"cornersDotOptions,omitempty"`
	CornersSquareOptions *CornersSquareOptions `json:"cornersSquareOptions,omitempty"`
	DotsOptions *DotsOptions `json:"dotsOptions,omitempty"`
	ErrorCorrectionLevel *string `json:"errorCorrectionLevel,omitempty"`
	Height *int32 `json:"height,omitempty"`
	HideBackgroundDots *bool `json:"hideBackgroundDots,omitempty"`
	// - OUTPUT_FILE_FORMAT_UNSPECIFIED: Output format not specified (defaults to PNG server-side)  - OUTPUT_FILE_FORMAT_SVG: Scalable Vector Graphics format for infinite resolution  - OUTPUT_FILE_FORMAT_PNG: Portable Network Graphics format with transparency support  - OUTPUT_FILE_FORMAT_JPEG: Joint Photographic Experts Group format for compressed images  - OUTPUT_FILE_FORMAT_WEBP: WebP format for modern web-optimized compression
	OutputFileFormat *string `json:"outputFileFormat,omitempty"`
	// - SHAPE_UNSPECIFIED: Shape not specified (defaults to square modules)  - SHAPE_SQUARE: Traditional square modules for QR code data points  - SHAPE_CIRCLE: Circular modules for a rounded visual appearance
	Shape *string `json:"shape,omitempty"`
	Width *int32 `json:"width,omitempty"`
}

// NewQrcodeDesignsServiceUpdateBody instantiates a new QrcodeDesignsServiceUpdateBody object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQrcodeDesignsServiceUpdateBody() *QrcodeDesignsServiceUpdateBody {
	this := QrcodeDesignsServiceUpdateBody{}
	return &this
}

// NewQrcodeDesignsServiceUpdateBodyWithDefaults instantiates a new QrcodeDesignsServiceUpdateBody object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQrcodeDesignsServiceUpdateBodyWithDefaults() *QrcodeDesignsServiceUpdateBody {
	this := QrcodeDesignsServiceUpdateBody{}
	return &this
}

// GetBackgroundOptions returns the BackgroundOptions field value if set, zero value otherwise.
func (o *QrcodeDesignsServiceUpdateBody) GetBackgroundOptions() BackgroundOptions {
	if o == nil || IsNil(o.BackgroundOptions) {
		var ret BackgroundOptions
		return ret
	}
	return *o.BackgroundOptions
}

// GetBackgroundOptionsOk returns a tuple with the BackgroundOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QrcodeDesignsServiceUpdateBody) GetBackgroundOptionsOk() (*BackgroundOptions, bool) {
	if o == nil || IsNil(o.BackgroundOptions) {
		return nil, false
	}
	return o.BackgroundOptions, true
}

// HasBackgroundOptions returns a boolean if a field has been set.
func (o *QrcodeDesignsServiceUpdateBody) HasBackgroundOptions() bool {
	if o != nil && !IsNil(o.BackgroundOptions) {
		return true
	}

	return false
}

// SetBackgroundOptions gets a reference to the given BackgroundOptions and assigns it to the BackgroundOptions field.
func (o *QrcodeDesignsServiceUpdateBody) SetBackgroundOptions(v BackgroundOptions) {
	o.BackgroundOptions = &v
}

// GetCentralImageDelete returns the CentralImageDelete field value if set, zero value otherwise.
func (o *QrcodeDesignsServiceUpdateBody) GetCentralImageDelete() bool {
	if o == nil || IsNil(o.CentralImageDelete) {
		var ret bool
		return ret
	}
	return *o.CentralImageDelete
}

// GetCentralImageDeleteOk returns a tuple with the CentralImageDelete field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QrcodeDesignsServiceUpdateBody) GetCentralImageDeleteOk() (*bool, bool) {
	if o == nil || IsNil(o.CentralImageDelete) {
		return nil, false
	}
	return o.CentralImageDelete, true
}

// HasCentralImageDelete returns a boolean if a field has been set.
func (o *QrcodeDesignsServiceUpdateBody) HasCentralImageDelete() bool {
	if o != nil && !IsNil(o.CentralImageDelete) {
		return true
	}

	return false
}

// SetCentralImageDelete gets a reference to the given bool and assigns it to the CentralImageDelete field.
func (o *QrcodeDesignsServiceUpdateBody) SetCentralImageDelete(v bool) {
	o.CentralImageDelete = &v
}

// GetCentralImageMargin returns the CentralImageMargin field value if set, zero value otherwise.
func (o *QrcodeDesignsServiceUpdateBody) GetCentralImageMargin() int32 {
	if o == nil || IsNil(o.CentralImageMargin) {
		var ret int32
		return ret
	}
	return *o.CentralImageMargin
}

// GetCentralImageMarginOk returns a tuple with the CentralImageMargin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QrcodeDesignsServiceUpdateBody) GetCentralImageMarginOk() (*int32, bool) {
	if o == nil || IsNil(o.CentralImageMargin) {
		return nil, false
	}
	return o.CentralImageMargin, true
}

// HasCentralImageMargin returns a boolean if a field has been set.
func (o *QrcodeDesignsServiceUpdateBody) HasCentralImageMargin() bool {
	if o != nil && !IsNil(o.CentralImageMargin) {
		return true
	}

	return false
}

// SetCentralImageMargin gets a reference to the given int32 and assigns it to the CentralImageMargin field.
func (o *QrcodeDesignsServiceUpdateBody) SetCentralImageMargin(v int32) {
	o.CentralImageMargin = &v
}

// GetCentralImageMediaId returns the CentralImageMediaId field value if set, zero value otherwise.
func (o *QrcodeDesignsServiceUpdateBody) GetCentralImageMediaId() string {
	if o == nil || IsNil(o.CentralImageMediaId) {
		var ret string
		return ret
	}
	return *o.CentralImageMediaId
}

// GetCentralImageMediaIdOk returns a tuple with the CentralImageMediaId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QrcodeDesignsServiceUpdateBody) GetCentralImageMediaIdOk() (*string, bool) {
	if o == nil || IsNil(o.CentralImageMediaId) {
		return nil, false
	}
	return o.CentralImageMediaId, true
}

// HasCentralImageMediaId returns a boolean if a field has been set.
func (o *QrcodeDesignsServiceUpdateBody) HasCentralImageMediaId() bool {
	if o != nil && !IsNil(o.CentralImageMediaId) {
		return true
	}

	return false
}

// SetCentralImageMediaId gets a reference to the given string and assigns it to the CentralImageMediaId field.
func (o *QrcodeDesignsServiceUpdateBody) SetCentralImageMediaId(v string) {
	o.CentralImageMediaId = &v
}

// GetCentralImageSize returns the CentralImageSize field value if set, zero value otherwise.
func (o *QrcodeDesignsServiceUpdateBody) GetCentralImageSize() float64 {
	if o == nil || IsNil(o.CentralImageSize) {
		var ret float64
		return ret
	}
	return *o.CentralImageSize
}

// GetCentralImageSizeOk returns a tuple with the CentralImageSize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QrcodeDesignsServiceUpdateBody) GetCentralImageSizeOk() (*float64, bool) {
	if o == nil || IsNil(o.CentralImageSize) {
		return nil, false
	}
	return o.CentralImageSize, true
}

// HasCentralImageSize returns a boolean if a field has been set.
func (o *QrcodeDesignsServiceUpdateBody) HasCentralImageSize() bool {
	if o != nil && !IsNil(o.CentralImageSize) {
		return true
	}

	return false
}

// SetCentralImageSize gets a reference to the given float64 and assigns it to the CentralImageSize field.
func (o *QrcodeDesignsServiceUpdateBody) SetCentralImageSize(v float64) {
	o.CentralImageSize = &v
}

// GetCornersDotOptions returns the CornersDotOptions field value if set, zero value otherwise.
func (o *QrcodeDesignsServiceUpdateBody) GetCornersDotOptions() CornersDotOptions {
	if o == nil || IsNil(o.CornersDotOptions) {
		var ret CornersDotOptions
		return ret
	}
	return *o.CornersDotOptions
}

// GetCornersDotOptionsOk returns a tuple with the CornersDotOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QrcodeDesignsServiceUpdateBody) GetCornersDotOptionsOk() (*CornersDotOptions, bool) {
	if o == nil || IsNil(o.CornersDotOptions) {
		return nil, false
	}
	return o.CornersDotOptions, true
}

// HasCornersDotOptions returns a boolean if a field has been set.
func (o *QrcodeDesignsServiceUpdateBody) HasCornersDotOptions() bool {
	if o != nil && !IsNil(o.CornersDotOptions) {
		return true
	}

	return false
}

// SetCornersDotOptions gets a reference to the given CornersDotOptions and assigns it to the CornersDotOptions field.
func (o *QrcodeDesignsServiceUpdateBody) SetCornersDotOptions(v CornersDotOptions) {
	o.CornersDotOptions = &v
}

// GetCornersSquareOptions returns the CornersSquareOptions field value if set, zero value otherwise.
func (o *QrcodeDesignsServiceUpdateBody) GetCornersSquareOptions() CornersSquareOptions {
	if o == nil || IsNil(o.CornersSquareOptions) {
		var ret CornersSquareOptions
		return ret
	}
	return *o.CornersSquareOptions
}

// GetCornersSquareOptionsOk returns a tuple with the CornersSquareOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QrcodeDesignsServiceUpdateBody) GetCornersSquareOptionsOk() (*CornersSquareOptions, bool) {
	if o == nil || IsNil(o.CornersSquareOptions) {
		return nil, false
	}
	return o.CornersSquareOptions, true
}

// HasCornersSquareOptions returns a boolean if a field has been set.
func (o *QrcodeDesignsServiceUpdateBody) HasCornersSquareOptions() bool {
	if o != nil && !IsNil(o.CornersSquareOptions) {
		return true
	}

	return false
}

// SetCornersSquareOptions gets a reference to the given CornersSquareOptions and assigns it to the CornersSquareOptions field.
func (o *QrcodeDesignsServiceUpdateBody) SetCornersSquareOptions(v CornersSquareOptions) {
	o.CornersSquareOptions = &v
}

// GetDotsOptions returns the DotsOptions field value if set, zero value otherwise.
func (o *QrcodeDesignsServiceUpdateBody) GetDotsOptions() DotsOptions {
	if o == nil || IsNil(o.DotsOptions) {
		var ret DotsOptions
		return ret
	}
	return *o.DotsOptions
}

// GetDotsOptionsOk returns a tuple with the DotsOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QrcodeDesignsServiceUpdateBody) GetDotsOptionsOk() (*DotsOptions, bool) {
	if o == nil || IsNil(o.DotsOptions) {
		return nil, false
	}
	return o.DotsOptions, true
}

// HasDotsOptions returns a boolean if a field has been set.
func (o *QrcodeDesignsServiceUpdateBody) HasDotsOptions() bool {
	if o != nil && !IsNil(o.DotsOptions) {
		return true
	}

	return false
}

// SetDotsOptions gets a reference to the given DotsOptions and assigns it to the DotsOptions field.
func (o *QrcodeDesignsServiceUpdateBody) SetDotsOptions(v DotsOptions) {
	o.DotsOptions = &v
}

// GetErrorCorrectionLevel returns the ErrorCorrectionLevel field value if set, zero value otherwise.
func (o *QrcodeDesignsServiceUpdateBody) GetErrorCorrectionLevel() string {
	if o == nil || IsNil(o.ErrorCorrectionLevel) {
		var ret string
		return ret
	}
	return *o.ErrorCorrectionLevel
}

// GetErrorCorrectionLevelOk returns a tuple with the ErrorCorrectionLevel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QrcodeDesignsServiceUpdateBody) GetErrorCorrectionLevelOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorCorrectionLevel) {
		return nil, false
	}
	return o.ErrorCorrectionLevel, true
}

// HasErrorCorrectionLevel returns a boolean if a field has been set.
func (o *QrcodeDesignsServiceUpdateBody) HasErrorCorrectionLevel() bool {
	if o != nil && !IsNil(o.ErrorCorrectionLevel) {
		return true
	}

	return false
}

// SetErrorCorrectionLevel gets a reference to the given string and assigns it to the ErrorCorrectionLevel field.
func (o *QrcodeDesignsServiceUpdateBody) SetErrorCorrectionLevel(v string) {
	o.ErrorCorrectionLevel = &v
}

// GetHeight returns the Height field value if set, zero value otherwise.
func (o *QrcodeDesignsServiceUpdateBody) GetHeight() int32 {
	if o == nil || IsNil(o.Height) {
		var ret int32
		return ret
	}
	return *o.Height
}

// GetHeightOk returns a tuple with the Height field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QrcodeDesignsServiceUpdateBody) GetHeightOk() (*int32, bool) {
	if o == nil || IsNil(o.Height) {
		return nil, false
	}
	return o.Height, true
}

// HasHeight returns a boolean if a field has been set.
func (o *QrcodeDesignsServiceUpdateBody) HasHeight() bool {
	if o != nil && !IsNil(o.Height) {
		return true
	}

	return false
}

// SetHeight gets a reference to the given int32 and assigns it to the Height field.
func (o *QrcodeDesignsServiceUpdateBody) SetHeight(v int32) {
	o.Height = &v
}

// GetHideBackgroundDots returns the HideBackgroundDots field value if set, zero value otherwise.
func (o *QrcodeDesignsServiceUpdateBody) GetHideBackgroundDots() bool {
	if o == nil || IsNil(o.HideBackgroundDots) {
		var ret bool
		return ret
	}
	return *o.HideBackgroundDots
}

// GetHideBackgroundDotsOk returns a tuple with the HideBackgroundDots field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QrcodeDesignsServiceUpdateBody) GetHideBackgroundDotsOk() (*bool, bool) {
	if o == nil || IsNil(o.HideBackgroundDots) {
		return nil, false
	}
	return o.HideBackgroundDots, true
}

// HasHideBackgroundDots returns a boolean if a field has been set.
func (o *QrcodeDesignsServiceUpdateBody) HasHideBackgroundDots() bool {
	if o != nil && !IsNil(o.HideBackgroundDots) {
		return true
	}

	return false
}

// SetHideBackgroundDots gets a reference to the given bool and assigns it to the HideBackgroundDots field.
func (o *QrcodeDesignsServiceUpdateBody) SetHideBackgroundDots(v bool) {
	o.HideBackgroundDots = &v
}

// GetOutputFileFormat returns the OutputFileFormat field value if set, zero value otherwise.
func (o *QrcodeDesignsServiceUpdateBody) GetOutputFileFormat() string {
	if o == nil || IsNil(o.OutputFileFormat) {
		var ret string
		return ret
	}
	return *o.OutputFileFormat
}

// GetOutputFileFormatOk returns a tuple with the OutputFileFormat field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QrcodeDesignsServiceUpdateBody) GetOutputFileFormatOk() (*string, bool) {
	if o == nil || IsNil(o.OutputFileFormat) {
		return nil, false
	}
	return o.OutputFileFormat, true
}

// HasOutputFileFormat returns a boolean if a field has been set.
func (o *QrcodeDesignsServiceUpdateBody) HasOutputFileFormat() bool {
	if o != nil && !IsNil(o.OutputFileFormat) {
		return true
	}

	return false
}

// SetOutputFileFormat gets a reference to the given string and assigns it to the OutputFileFormat field.
func (o *QrcodeDesignsServiceUpdateBody) SetOutputFileFormat(v string) {
	o.OutputFileFormat = &v
}

// GetShape returns the Shape field value if set, zero value otherwise.
func (o *QrcodeDesignsServiceUpdateBody) GetShape() string {
	if o == nil || IsNil(o.Shape) {
		var ret string
		return ret
	}
	return *o.Shape
}

// GetShapeOk returns a tuple with the Shape field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QrcodeDesignsServiceUpdateBody) GetShapeOk() (*string, bool) {
	if o == nil || IsNil(o.Shape) {
		return nil, false
	}
	return o.Shape, true
}

// HasShape returns a boolean if a field has been set.
func (o *QrcodeDesignsServiceUpdateBody) HasShape() bool {
	if o != nil && !IsNil(o.Shape) {
		return true
	}

	return false
}

// SetShape gets a reference to the given string and assigns it to the Shape field.
func (o *QrcodeDesignsServiceUpdateBody) SetShape(v string) {
	o.Shape = &v
}

// GetWidth returns the Width field value if set, zero value otherwise.
func (o *QrcodeDesignsServiceUpdateBody) GetWidth() int32 {
	if o == nil || IsNil(o.Width) {
		var ret int32
		return ret
	}
	return *o.Width
}

// GetWidthOk returns a tuple with the Width field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QrcodeDesignsServiceUpdateBody) GetWidthOk() (*int32, bool) {
	if o == nil || IsNil(o.Width) {
		return nil, false
	}
	return o.Width, true
}

// HasWidth returns a boolean if a field has been set.
func (o *QrcodeDesignsServiceUpdateBody) HasWidth() bool {
	if o != nil && !IsNil(o.Width) {
		return true
	}

	return false
}

// SetWidth gets a reference to the given int32 and assigns it to the Width field.
func (o *QrcodeDesignsServiceUpdateBody) SetWidth(v int32) {
	o.Width = &v
}

func (o QrcodeDesignsServiceUpdateBody) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QrcodeDesignsServiceUpdateBody) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BackgroundOptions) {
		toSerialize["backgroundOptions"] = o.BackgroundOptions
	}
	if !IsNil(o.CentralImageDelete) {
		toSerialize["centralImageDelete"] = o.CentralImageDelete
	}
	if !IsNil(o.CentralImageMargin) {
		toSerialize["centralImageMargin"] = o.CentralImageMargin
	}
	if !IsNil(o.CentralImageMediaId) {
		toSerialize["centralImageMediaId"] = o.CentralImageMediaId
	}
	if !IsNil(o.CentralImageSize) {
		toSerialize["centralImageSize"] = o.CentralImageSize
	}
	if !IsNil(o.CornersDotOptions) {
		toSerialize["cornersDotOptions"] = o.CornersDotOptions
	}
	if !IsNil(o.CornersSquareOptions) {
		toSerialize["cornersSquareOptions"] = o.CornersSquareOptions
	}
	if !IsNil(o.DotsOptions) {
		toSerialize["dotsOptions"] = o.DotsOptions
	}
	if !IsNil(o.ErrorCorrectionLevel) {
		toSerialize["errorCorrectionLevel"] = o.ErrorCorrectionLevel
	}
	if !IsNil(o.Height) {
		toSerialize["height"] = o.Height
	}
	if !IsNil(o.HideBackgroundDots) {
		toSerialize["hideBackgroundDots"] = o.HideBackgroundDots
	}
	if !IsNil(o.OutputFileFormat) {
		toSerialize["outputFileFormat"] = o.OutputFileFormat
	}
	if !IsNil(o.Shape) {
		toSerialize["shape"] = o.Shape
	}
	if !IsNil(o.Width) {
		toSerialize["width"] = o.Width
	}
	return toSerialize, nil
}

type NullableQrcodeDesignsServiceUpdateBody struct {
	value *QrcodeDesignsServiceUpdateBody
	isSet bool
}

func (v NullableQrcodeDesignsServiceUpdateBody) Get() *QrcodeDesignsServiceUpdateBody {
	return v.value
}

func (v *NullableQrcodeDesignsServiceUpdateBody) Set(val *QrcodeDesignsServiceUpdateBody) {
	v.value = val
	v.isSet = true
}

func (v NullableQrcodeDesignsServiceUpdateBody) IsSet() bool {
	return v.isSet
}

func (v *NullableQrcodeDesignsServiceUpdateBody) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQrcodeDesignsServiceUpdateBody(val *QrcodeDesignsServiceUpdateBody) *NullableQrcodeDesignsServiceUpdateBody {
	return &NullableQrcodeDesignsServiceUpdateBody{value: val, isSet: true}
}

func (v NullableQrcodeDesignsServiceUpdateBody) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQrcodeDesignsServiceUpdateBody) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


