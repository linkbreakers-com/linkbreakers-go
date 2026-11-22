/*
Linkbreakers API

The Linkbreakers REST API: links, QR codes, workflows, visitors and analytics, as JSON over HTTPS. Send a workspace API token as `Authorization: Bearer <token>`.  ## Entry points  `GET https://api.linkbreakers.com/` and `GET https://api.linkbreakers.com/v1` answer without credentials with the API's name, version and links. This document is served at `https://api.linkbreakers.com/openapi.json`.  ## Rate limits  Every response describes where the caller stands in the IETF `RateLimit-Policy` and `RateLimit` header fields (draft-ietf-httpapi-ratelimit-headers), and in `RateLimit-Limit`, `RateLimit-Remaining`, `RateLimit-Reset` and `X-RateLimit-Limit`, `X-RateLimit-Remaining`, `X-RateLimit-Reset` (a Unix timestamp) for clients that read those. A request past the limit gets `429 Too Many Requests` with `Retry-After` in seconds.  - Workspace API token: 720 requests per 60 seconds, per token. - Dashboard session or OAuth token: 600 requests per 60 seconds, per token. - No credentials: 300 requests per 60 seconds, per client IP.  Each plan also includes a number of API calls per billing cycle. Past it, calls answer `429` with the code `MAX_PUBLIC_API_QUERIED_REACHED` until the cycle renews or the plan changes.  ## Pagination  Every list pages with opaque cursors (AIP-158). Pass `pageSize`, then send the `nextPageToken` of each response back as `pageToken` (or its alias `cursor`) to get the next page. An empty or absent `nextPageToken` means the list is complete. Each list operation describes this in `x-pagination`.  ## Batches and background jobs  `POST /v1/links/bulk` creates up to 250 links in one transactional request; `POST /v1/links/batch` takes the same links as a bare JSON array. `POST /v1/links/bulk/jobs` accepts the same batch and answers `202 Accepted` at once with a link job and a `Location` header; poll `GET /v1/links/jobs/{id}` until `status` is `LINK_JOB_STATUS_SUCCEEDED` (the created link ids are in `linksBulkCreate.linkIds`) or `LINK_JOB_STATUS_FAILED` (`error` says why, and nothing was created). Link jobs can be read for 24 hours after they finish.  ## Paused workspaces  A free workspace whose members have not verified any email address is paused 7 days after sign-up. Its calls then answer `403` with the code `WORKSPACE_UNVERIFIED` and change nothing; QR codes and links keep working for visitors. To lift the pause, ask the workspace owner to verify their email: `POST /v1/members/email-verification` sends the verification email again, and works while the workspace is paused. `unverifiedPauseAt` on the workspace says when the pause starts.  ## Versioning and deprecation  The version is in the path: every operation lives under `/v1`. Within a version, changes are additive only: new operations, new optional request fields, new response fields and new enum values. Clients should ignore fields and enum values they do not know. Anything that would break an existing client ships under a new path version instead.  An operation is retired in two steps. First it is deprecated: it keeps working, is flagged `deprecated: true` in this document, and every response carries a `Deprecation` header field (RFC 9745) with the time it was deprecated and a `Link` with `rel=\"deprecation\"` to the policy. Then, at least 180 days later, it may be removed; that date is announced in advance in a `Sunset` header field (RFC 8594) and as `x-sunset` on the operation. The full policy: https://linkbreakers.com/help/article/api-versioning-and-deprecation-policy  A renamed operation keeps answering its former path while clients still call it, until that path's sunset at the latest: a former path no client calls for 14 days is removed. Responses on a former path carry the same `Deprecation`, `Sunset` and `Link` header fields; this document lists only the current path, and the policy lists the former paths that still answer.  ## Idempotent retries  Every POST, PUT and PATCH accepts an optional `Idempotency-Key` header (draft-ietf-httpapi-idempotency-key-header). Send a new unique value, such as a UUID, with each logical request and reuse it only to retry that request. A retry with the same key and the same request returns the first response, marked `Idempotent-Replayed: true`, instead of running again. Keys are scoped to the token and kept for 24 hours. The same key with a different request is refused with `422`, and a retry while the first request is still running gets `409` with `Retry-After`. Server errors (5xx) and refusals such as 401, 403 and 429 are not stored, so retrying them runs the request again.  ## Identifying your client  Send a `Linkbreakers-Client` request header saying which client made the call: `<client>/<version> (<details>)`, for example `my-integration/1.2.0 (python 3.12; linux)`. The official SDKs, the CLI and the MCP bridge set it for you, as `sdk-python/2.3.1 (python 3.12.4; linux)` or `cli/1.4.0 (darwin arm64)`. It is optional and changes nothing about how a request is handled; it tells us which clients and versions are in use, so we know what can be retired and what to fix first. Keep the details short and free of personal data. Browsers may send it too: it is allowed by CORS.

API version: 1.205.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package linkbreakers

import (
	"encoding/json"
	"time"
)

// checks if the PageTheme type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PageTheme{}

// PageTheme struct for PageTheme
type PageTheme struct {
	BadgeEnabled *bool `json:"badgeEnabled,omitempty"`
	BadgeIconColor *ThemeColor `json:"badgeIconColor,omitempty"`
	BadgeId *string `json:"badgeId,omitempty"`
	BadgeShape *string `json:"badgeShape,omitempty"`
	// How body text is aligned: descriptions, messages, helper text, and answer and button labels. Unset: centered. UNSPECIFIED on an update clears it.
	BodyAlign *string `json:"bodyAlign,omitempty"`
	BodyFont *FontConfig `json:"bodyFont,omitempty"`
	CanvasColor *ThemeColor `json:"canvasColor,omitempty"`
	CanvasCoverId *string `json:"canvasCoverId,omitempty"`
	CanvasCoverMode *string `json:"canvasCoverMode,omitempty"`
	// A texture laid over the page background. Unset reads as NONE.
	CanvasTexture *string `json:"canvasTexture,omitempty"`
	// How a choice question lays out its answers: AUTO (pills or tiles) or ROWS. Unset reads as AUTO.
	ChoiceLayout *string `json:"choiceLayout,omitempty"`
	ContainerConfig *ContainerConfig `json:"containerConfig,omitempty"`
	ContainerEnabled *bool `json:"containerEnabled,omitempty"`
	CreatedAt *time.Time `json:"createdAt,omitempty"`
	ElementColor *ThemeColor `json:"elementColor,omitempty"`
	ElementFont *FontConfig `json:"elementFont,omitempty"`
	// The look of each element job (SURFACE, SELECTED, ACTION, QUIET). A job not listed follows the shape, the colours, or the look it always had.
	ElementStyles []PageThemeElementStyleOverride `json:"elementStyles,omitempty"`
	// Whether the footer logo (a small company or brand logo centered at the bottom of every page using this theme) is shown.  Turning it off keeps footerLogoId, so turning it back on restores the same image.
	FooterLogoEnabled *bool `json:"footerLogoEnabled,omitempty"`
	FooterLogoId *string `json:"footerLogoId,omitempty"`
	// How titles and headline questions are aligned. Unset: centered titles, questions as the question style draws them. UNSPECIFIED on an update clears it.
	HeadingAlign *string `json:"headingAlign,omitempty"`
	Id *string `json:"id,omitempty"`
	// How a form's questions are written: LABEL (small, in the body text) or HEADLINE (a heading, in the title font). Unset reads as LABEL.
	QuestionStyle *string `json:"questionStyle,omitempty"`
	// The radius and border every job follows unless it has its own.
	Shape *PageThemeShape `json:"shape,omitempty"`
	// Small text: hints, counters, consent. Unset: the body text, smaller and muted.
	SmallText *PageThemeSmallText `json:"smallText,omitempty"`
	// The text colour headings, body and small text follow when their own font has none. On an update, a colour without colours clears it.
	TextColor *ThemeColor `json:"textColor,omitempty"`
	TitleFont *FontConfig `json:"titleFont,omitempty"`
	UpdatedAt *time.Time `json:"updatedAt,omitempty"`
	// How visitor pages pace their content. Unset reads as CLASSIC.
	VisitorExperience *string `json:"visitorExperience,omitempty"`
	WorkspaceId *string `json:"workspaceId,omitempty"`
}

// NewPageTheme instantiates a new PageTheme object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPageTheme() *PageTheme {
	this := PageTheme{}
	return &this
}

// NewPageThemeWithDefaults instantiates a new PageTheme object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPageThemeWithDefaults() *PageTheme {
	this := PageTheme{}
	return &this
}

// GetBadgeEnabled returns the BadgeEnabled field value if set, zero value otherwise.
func (o *PageTheme) GetBadgeEnabled() bool {
	if o == nil || IsNil(o.BadgeEnabled) {
		var ret bool
		return ret
	}
	return *o.BadgeEnabled
}

// GetBadgeEnabledOk returns a tuple with the BadgeEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageTheme) GetBadgeEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.BadgeEnabled) {
		return nil, false
	}
	return o.BadgeEnabled, true
}

// HasBadgeEnabled returns a boolean if a field has been set.
func (o *PageTheme) HasBadgeEnabled() bool {
	if o != nil && !IsNil(o.BadgeEnabled) {
		return true
	}

	return false
}

// SetBadgeEnabled gets a reference to the given bool and assigns it to the BadgeEnabled field.
func (o *PageTheme) SetBadgeEnabled(v bool) {
	o.BadgeEnabled = &v
}

// GetBadgeIconColor returns the BadgeIconColor field value if set, zero value otherwise.
func (o *PageTheme) GetBadgeIconColor() ThemeColor {
	if o == nil || IsNil(o.BadgeIconColor) {
		var ret ThemeColor
		return ret
	}
	return *o.BadgeIconColor
}

// GetBadgeIconColorOk returns a tuple with the BadgeIconColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageTheme) GetBadgeIconColorOk() (*ThemeColor, bool) {
	if o == nil || IsNil(o.BadgeIconColor) {
		return nil, false
	}
	return o.BadgeIconColor, true
}

// HasBadgeIconColor returns a boolean if a field has been set.
func (o *PageTheme) HasBadgeIconColor() bool {
	if o != nil && !IsNil(o.BadgeIconColor) {
		return true
	}

	return false
}

// SetBadgeIconColor gets a reference to the given ThemeColor and assigns it to the BadgeIconColor field.
func (o *PageTheme) SetBadgeIconColor(v ThemeColor) {
	o.BadgeIconColor = &v
}

// GetBadgeId returns the BadgeId field value if set, zero value otherwise.
func (o *PageTheme) GetBadgeId() string {
	if o == nil || IsNil(o.BadgeId) {
		var ret string
		return ret
	}
	return *o.BadgeId
}

// GetBadgeIdOk returns a tuple with the BadgeId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageTheme) GetBadgeIdOk() (*string, bool) {
	if o == nil || IsNil(o.BadgeId) {
		return nil, false
	}
	return o.BadgeId, true
}

// HasBadgeId returns a boolean if a field has been set.
func (o *PageTheme) HasBadgeId() bool {
	if o != nil && !IsNil(o.BadgeId) {
		return true
	}

	return false
}

// SetBadgeId gets a reference to the given string and assigns it to the BadgeId field.
func (o *PageTheme) SetBadgeId(v string) {
	o.BadgeId = &v
}

// GetBadgeShape returns the BadgeShape field value if set, zero value otherwise.
func (o *PageTheme) GetBadgeShape() string {
	if o == nil || IsNil(o.BadgeShape) {
		var ret string
		return ret
	}
	return *o.BadgeShape
}

// GetBadgeShapeOk returns a tuple with the BadgeShape field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageTheme) GetBadgeShapeOk() (*string, bool) {
	if o == nil || IsNil(o.BadgeShape) {
		return nil, false
	}
	return o.BadgeShape, true
}

// HasBadgeShape returns a boolean if a field has been set.
func (o *PageTheme) HasBadgeShape() bool {
	if o != nil && !IsNil(o.BadgeShape) {
		return true
	}

	return false
}

// SetBadgeShape gets a reference to the given string and assigns it to the BadgeShape field.
func (o *PageTheme) SetBadgeShape(v string) {
	o.BadgeShape = &v
}

// GetBodyAlign returns the BodyAlign field value if set, zero value otherwise.
func (o *PageTheme) GetBodyAlign() string {
	if o == nil || IsNil(o.BodyAlign) {
		var ret string
		return ret
	}
	return *o.BodyAlign
}

// GetBodyAlignOk returns a tuple with the BodyAlign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageTheme) GetBodyAlignOk() (*string, bool) {
	if o == nil || IsNil(o.BodyAlign) {
		return nil, false
	}
	return o.BodyAlign, true
}

// HasBodyAlign returns a boolean if a field has been set.
func (o *PageTheme) HasBodyAlign() bool {
	if o != nil && !IsNil(o.BodyAlign) {
		return true
	}

	return false
}

// SetBodyAlign gets a reference to the given string and assigns it to the BodyAlign field.
func (o *PageTheme) SetBodyAlign(v string) {
	o.BodyAlign = &v
}

// GetBodyFont returns the BodyFont field value if set, zero value otherwise.
func (o *PageTheme) GetBodyFont() FontConfig {
	if o == nil || IsNil(o.BodyFont) {
		var ret FontConfig
		return ret
	}
	return *o.BodyFont
}

// GetBodyFontOk returns a tuple with the BodyFont field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageTheme) GetBodyFontOk() (*FontConfig, bool) {
	if o == nil || IsNil(o.BodyFont) {
		return nil, false
	}
	return o.BodyFont, true
}

// HasBodyFont returns a boolean if a field has been set.
func (o *PageTheme) HasBodyFont() bool {
	if o != nil && !IsNil(o.BodyFont) {
		return true
	}

	return false
}

// SetBodyFont gets a reference to the given FontConfig and assigns it to the BodyFont field.
func (o *PageTheme) SetBodyFont(v FontConfig) {
	o.BodyFont = &v
}

// GetCanvasColor returns the CanvasColor field value if set, zero value otherwise.
func (o *PageTheme) GetCanvasColor() ThemeColor {
	if o == nil || IsNil(o.CanvasColor) {
		var ret ThemeColor
		return ret
	}
	return *o.CanvasColor
}

// GetCanvasColorOk returns a tuple with the CanvasColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageTheme) GetCanvasColorOk() (*ThemeColor, bool) {
	if o == nil || IsNil(o.CanvasColor) {
		return nil, false
	}
	return o.CanvasColor, true
}

// HasCanvasColor returns a boolean if a field has been set.
func (o *PageTheme) HasCanvasColor() bool {
	if o != nil && !IsNil(o.CanvasColor) {
		return true
	}

	return false
}

// SetCanvasColor gets a reference to the given ThemeColor and assigns it to the CanvasColor field.
func (o *PageTheme) SetCanvasColor(v ThemeColor) {
	o.CanvasColor = &v
}

// GetCanvasCoverId returns the CanvasCoverId field value if set, zero value otherwise.
func (o *PageTheme) GetCanvasCoverId() string {
	if o == nil || IsNil(o.CanvasCoverId) {
		var ret string
		return ret
	}
	return *o.CanvasCoverId
}

// GetCanvasCoverIdOk returns a tuple with the CanvasCoverId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageTheme) GetCanvasCoverIdOk() (*string, bool) {
	if o == nil || IsNil(o.CanvasCoverId) {
		return nil, false
	}
	return o.CanvasCoverId, true
}

// HasCanvasCoverId returns a boolean if a field has been set.
func (o *PageTheme) HasCanvasCoverId() bool {
	if o != nil && !IsNil(o.CanvasCoverId) {
		return true
	}

	return false
}

// SetCanvasCoverId gets a reference to the given string and assigns it to the CanvasCoverId field.
func (o *PageTheme) SetCanvasCoverId(v string) {
	o.CanvasCoverId = &v
}

// GetCanvasCoverMode returns the CanvasCoverMode field value if set, zero value otherwise.
func (o *PageTheme) GetCanvasCoverMode() string {
	if o == nil || IsNil(o.CanvasCoverMode) {
		var ret string
		return ret
	}
	return *o.CanvasCoverMode
}

// GetCanvasCoverModeOk returns a tuple with the CanvasCoverMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageTheme) GetCanvasCoverModeOk() (*string, bool) {
	if o == nil || IsNil(o.CanvasCoverMode) {
		return nil, false
	}
	return o.CanvasCoverMode, true
}

// HasCanvasCoverMode returns a boolean if a field has been set.
func (o *PageTheme) HasCanvasCoverMode() bool {
	if o != nil && !IsNil(o.CanvasCoverMode) {
		return true
	}

	return false
}

// SetCanvasCoverMode gets a reference to the given string and assigns it to the CanvasCoverMode field.
func (o *PageTheme) SetCanvasCoverMode(v string) {
	o.CanvasCoverMode = &v
}

// GetCanvasTexture returns the CanvasTexture field value if set, zero value otherwise.
func (o *PageTheme) GetCanvasTexture() string {
	if o == nil || IsNil(o.CanvasTexture) {
		var ret string
		return ret
	}
	return *o.CanvasTexture
}

// GetCanvasTextureOk returns a tuple with the CanvasTexture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageTheme) GetCanvasTextureOk() (*string, bool) {
	if o == nil || IsNil(o.CanvasTexture) {
		return nil, false
	}
	return o.CanvasTexture, true
}

// HasCanvasTexture returns a boolean if a field has been set.
func (o *PageTheme) HasCanvasTexture() bool {
	if o != nil && !IsNil(o.CanvasTexture) {
		return true
	}

	return false
}

// SetCanvasTexture gets a reference to the given string and assigns it to the CanvasTexture field.
func (o *PageTheme) SetCanvasTexture(v string) {
	o.CanvasTexture = &v
}

// GetChoiceLayout returns the ChoiceLayout field value if set, zero value otherwise.
func (o *PageTheme) GetChoiceLayout() string {
	if o == nil || IsNil(o.ChoiceLayout) {
		var ret string
		return ret
	}
	return *o.ChoiceLayout
}

// GetChoiceLayoutOk returns a tuple with the ChoiceLayout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageTheme) GetChoiceLayoutOk() (*string, bool) {
	if o == nil || IsNil(o.ChoiceLayout) {
		return nil, false
	}
	return o.ChoiceLayout, true
}

// HasChoiceLayout returns a boolean if a field has been set.
func (o *PageTheme) HasChoiceLayout() bool {
	if o != nil && !IsNil(o.ChoiceLayout) {
		return true
	}

	return false
}

// SetChoiceLayout gets a reference to the given string and assigns it to the ChoiceLayout field.
func (o *PageTheme) SetChoiceLayout(v string) {
	o.ChoiceLayout = &v
}

// GetContainerConfig returns the ContainerConfig field value if set, zero value otherwise.
func (o *PageTheme) GetContainerConfig() ContainerConfig {
	if o == nil || IsNil(o.ContainerConfig) {
		var ret ContainerConfig
		return ret
	}
	return *o.ContainerConfig
}

// GetContainerConfigOk returns a tuple with the ContainerConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageTheme) GetContainerConfigOk() (*ContainerConfig, bool) {
	if o == nil || IsNil(o.ContainerConfig) {
		return nil, false
	}
	return o.ContainerConfig, true
}

// HasContainerConfig returns a boolean if a field has been set.
func (o *PageTheme) HasContainerConfig() bool {
	if o != nil && !IsNil(o.ContainerConfig) {
		return true
	}

	return false
}

// SetContainerConfig gets a reference to the given ContainerConfig and assigns it to the ContainerConfig field.
func (o *PageTheme) SetContainerConfig(v ContainerConfig) {
	o.ContainerConfig = &v
}

// GetContainerEnabled returns the ContainerEnabled field value if set, zero value otherwise.
func (o *PageTheme) GetContainerEnabled() bool {
	if o == nil || IsNil(o.ContainerEnabled) {
		var ret bool
		return ret
	}
	return *o.ContainerEnabled
}

// GetContainerEnabledOk returns a tuple with the ContainerEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageTheme) GetContainerEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ContainerEnabled) {
		return nil, false
	}
	return o.ContainerEnabled, true
}

// HasContainerEnabled returns a boolean if a field has been set.
func (o *PageTheme) HasContainerEnabled() bool {
	if o != nil && !IsNil(o.ContainerEnabled) {
		return true
	}

	return false
}

// SetContainerEnabled gets a reference to the given bool and assigns it to the ContainerEnabled field.
func (o *PageTheme) SetContainerEnabled(v bool) {
	o.ContainerEnabled = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *PageTheme) GetCreatedAt() time.Time {
	if o == nil || IsNil(o.CreatedAt) {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageTheme) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *PageTheme) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *PageTheme) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetElementColor returns the ElementColor field value if set, zero value otherwise.
func (o *PageTheme) GetElementColor() ThemeColor {
	if o == nil || IsNil(o.ElementColor) {
		var ret ThemeColor
		return ret
	}
	return *o.ElementColor
}

// GetElementColorOk returns a tuple with the ElementColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageTheme) GetElementColorOk() (*ThemeColor, bool) {
	if o == nil || IsNil(o.ElementColor) {
		return nil, false
	}
	return o.ElementColor, true
}

// HasElementColor returns a boolean if a field has been set.
func (o *PageTheme) HasElementColor() bool {
	if o != nil && !IsNil(o.ElementColor) {
		return true
	}

	return false
}

// SetElementColor gets a reference to the given ThemeColor and assigns it to the ElementColor field.
func (o *PageTheme) SetElementColor(v ThemeColor) {
	o.ElementColor = &v
}

// GetElementFont returns the ElementFont field value if set, zero value otherwise.
func (o *PageTheme) GetElementFont() FontConfig {
	if o == nil || IsNil(o.ElementFont) {
		var ret FontConfig
		return ret
	}
	return *o.ElementFont
}

// GetElementFontOk returns a tuple with the ElementFont field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageTheme) GetElementFontOk() (*FontConfig, bool) {
	if o == nil || IsNil(o.ElementFont) {
		return nil, false
	}
	return o.ElementFont, true
}

// HasElementFont returns a boolean if a field has been set.
func (o *PageTheme) HasElementFont() bool {
	if o != nil && !IsNil(o.ElementFont) {
		return true
	}

	return false
}

// SetElementFont gets a reference to the given FontConfig and assigns it to the ElementFont field.
func (o *PageTheme) SetElementFont(v FontConfig) {
	o.ElementFont = &v
}

// GetElementStyles returns the ElementStyles field value if set, zero value otherwise.
func (o *PageTheme) GetElementStyles() []PageThemeElementStyleOverride {
	if o == nil || IsNil(o.ElementStyles) {
		var ret []PageThemeElementStyleOverride
		return ret
	}
	return o.ElementStyles
}

// GetElementStylesOk returns a tuple with the ElementStyles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageTheme) GetElementStylesOk() ([]PageThemeElementStyleOverride, bool) {
	if o == nil || IsNil(o.ElementStyles) {
		return nil, false
	}
	return o.ElementStyles, true
}

// HasElementStyles returns a boolean if a field has been set.
func (o *PageTheme) HasElementStyles() bool {
	if o != nil && !IsNil(o.ElementStyles) {
		return true
	}

	return false
}

// SetElementStyles gets a reference to the given []PageThemeElementStyleOverride and assigns it to the ElementStyles field.
func (o *PageTheme) SetElementStyles(v []PageThemeElementStyleOverride) {
	o.ElementStyles = v
}

// GetFooterLogoEnabled returns the FooterLogoEnabled field value if set, zero value otherwise.
func (o *PageTheme) GetFooterLogoEnabled() bool {
	if o == nil || IsNil(o.FooterLogoEnabled) {
		var ret bool
		return ret
	}
	return *o.FooterLogoEnabled
}

// GetFooterLogoEnabledOk returns a tuple with the FooterLogoEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageTheme) GetFooterLogoEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.FooterLogoEnabled) {
		return nil, false
	}
	return o.FooterLogoEnabled, true
}

// HasFooterLogoEnabled returns a boolean if a field has been set.
func (o *PageTheme) HasFooterLogoEnabled() bool {
	if o != nil && !IsNil(o.FooterLogoEnabled) {
		return true
	}

	return false
}

// SetFooterLogoEnabled gets a reference to the given bool and assigns it to the FooterLogoEnabled field.
func (o *PageTheme) SetFooterLogoEnabled(v bool) {
	o.FooterLogoEnabled = &v
}

// GetFooterLogoId returns the FooterLogoId field value if set, zero value otherwise.
func (o *PageTheme) GetFooterLogoId() string {
	if o == nil || IsNil(o.FooterLogoId) {
		var ret string
		return ret
	}
	return *o.FooterLogoId
}

// GetFooterLogoIdOk returns a tuple with the FooterLogoId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageTheme) GetFooterLogoIdOk() (*string, bool) {
	if o == nil || IsNil(o.FooterLogoId) {
		return nil, false
	}
	return o.FooterLogoId, true
}

// HasFooterLogoId returns a boolean if a field has been set.
func (o *PageTheme) HasFooterLogoId() bool {
	if o != nil && !IsNil(o.FooterLogoId) {
		return true
	}

	return false
}

// SetFooterLogoId gets a reference to the given string and assigns it to the FooterLogoId field.
func (o *PageTheme) SetFooterLogoId(v string) {
	o.FooterLogoId = &v
}

// GetHeadingAlign returns the HeadingAlign field value if set, zero value otherwise.
func (o *PageTheme) GetHeadingAlign() string {
	if o == nil || IsNil(o.HeadingAlign) {
		var ret string
		return ret
	}
	return *o.HeadingAlign
}

// GetHeadingAlignOk returns a tuple with the HeadingAlign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageTheme) GetHeadingAlignOk() (*string, bool) {
	if o == nil || IsNil(o.HeadingAlign) {
		return nil, false
	}
	return o.HeadingAlign, true
}

// HasHeadingAlign returns a boolean if a field has been set.
func (o *PageTheme) HasHeadingAlign() bool {
	if o != nil && !IsNil(o.HeadingAlign) {
		return true
	}

	return false
}

// SetHeadingAlign gets a reference to the given string and assigns it to the HeadingAlign field.
func (o *PageTheme) SetHeadingAlign(v string) {
	o.HeadingAlign = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PageTheme) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageTheme) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PageTheme) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PageTheme) SetId(v string) {
	o.Id = &v
}

// GetQuestionStyle returns the QuestionStyle field value if set, zero value otherwise.
func (o *PageTheme) GetQuestionStyle() string {
	if o == nil || IsNil(o.QuestionStyle) {
		var ret string
		return ret
	}
	return *o.QuestionStyle
}

// GetQuestionStyleOk returns a tuple with the QuestionStyle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageTheme) GetQuestionStyleOk() (*string, bool) {
	if o == nil || IsNil(o.QuestionStyle) {
		return nil, false
	}
	return o.QuestionStyle, true
}

// HasQuestionStyle returns a boolean if a field has been set.
func (o *PageTheme) HasQuestionStyle() bool {
	if o != nil && !IsNil(o.QuestionStyle) {
		return true
	}

	return false
}

// SetQuestionStyle gets a reference to the given string and assigns it to the QuestionStyle field.
func (o *PageTheme) SetQuestionStyle(v string) {
	o.QuestionStyle = &v
}

// GetShape returns the Shape field value if set, zero value otherwise.
func (o *PageTheme) GetShape() PageThemeShape {
	if o == nil || IsNil(o.Shape) {
		var ret PageThemeShape
		return ret
	}
	return *o.Shape
}

// GetShapeOk returns a tuple with the Shape field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageTheme) GetShapeOk() (*PageThemeShape, bool) {
	if o == nil || IsNil(o.Shape) {
		return nil, false
	}
	return o.Shape, true
}

// HasShape returns a boolean if a field has been set.
func (o *PageTheme) HasShape() bool {
	if o != nil && !IsNil(o.Shape) {
		return true
	}

	return false
}

// SetShape gets a reference to the given PageThemeShape and assigns it to the Shape field.
func (o *PageTheme) SetShape(v PageThemeShape) {
	o.Shape = &v
}

// GetSmallText returns the SmallText field value if set, zero value otherwise.
func (o *PageTheme) GetSmallText() PageThemeSmallText {
	if o == nil || IsNil(o.SmallText) {
		var ret PageThemeSmallText
		return ret
	}
	return *o.SmallText
}

// GetSmallTextOk returns a tuple with the SmallText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageTheme) GetSmallTextOk() (*PageThemeSmallText, bool) {
	if o == nil || IsNil(o.SmallText) {
		return nil, false
	}
	return o.SmallText, true
}

// HasSmallText returns a boolean if a field has been set.
func (o *PageTheme) HasSmallText() bool {
	if o != nil && !IsNil(o.SmallText) {
		return true
	}

	return false
}

// SetSmallText gets a reference to the given PageThemeSmallText and assigns it to the SmallText field.
func (o *PageTheme) SetSmallText(v PageThemeSmallText) {
	o.SmallText = &v
}

// GetTextColor returns the TextColor field value if set, zero value otherwise.
func (o *PageTheme) GetTextColor() ThemeColor {
	if o == nil || IsNil(o.TextColor) {
		var ret ThemeColor
		return ret
	}
	return *o.TextColor
}

// GetTextColorOk returns a tuple with the TextColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageTheme) GetTextColorOk() (*ThemeColor, bool) {
	if o == nil || IsNil(o.TextColor) {
		return nil, false
	}
	return o.TextColor, true
}

// HasTextColor returns a boolean if a field has been set.
func (o *PageTheme) HasTextColor() bool {
	if o != nil && !IsNil(o.TextColor) {
		return true
	}

	return false
}

// SetTextColor gets a reference to the given ThemeColor and assigns it to the TextColor field.
func (o *PageTheme) SetTextColor(v ThemeColor) {
	o.TextColor = &v
}

// GetTitleFont returns the TitleFont field value if set, zero value otherwise.
func (o *PageTheme) GetTitleFont() FontConfig {
	if o == nil || IsNil(o.TitleFont) {
		var ret FontConfig
		return ret
	}
	return *o.TitleFont
}

// GetTitleFontOk returns a tuple with the TitleFont field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageTheme) GetTitleFontOk() (*FontConfig, bool) {
	if o == nil || IsNil(o.TitleFont) {
		return nil, false
	}
	return o.TitleFont, true
}

// HasTitleFont returns a boolean if a field has been set.
func (o *PageTheme) HasTitleFont() bool {
	if o != nil && !IsNil(o.TitleFont) {
		return true
	}

	return false
}

// SetTitleFont gets a reference to the given FontConfig and assigns it to the TitleFont field.
func (o *PageTheme) SetTitleFont(v FontConfig) {
	o.TitleFont = &v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *PageTheme) GetUpdatedAt() time.Time {
	if o == nil || IsNil(o.UpdatedAt) {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageTheme) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || IsNil(o.UpdatedAt) {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *PageTheme) HasUpdatedAt() bool {
	if o != nil && !IsNil(o.UpdatedAt) {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *PageTheme) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetVisitorExperience returns the VisitorExperience field value if set, zero value otherwise.
func (o *PageTheme) GetVisitorExperience() string {
	if o == nil || IsNil(o.VisitorExperience) {
		var ret string
		return ret
	}
	return *o.VisitorExperience
}

// GetVisitorExperienceOk returns a tuple with the VisitorExperience field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageTheme) GetVisitorExperienceOk() (*string, bool) {
	if o == nil || IsNil(o.VisitorExperience) {
		return nil, false
	}
	return o.VisitorExperience, true
}

// HasVisitorExperience returns a boolean if a field has been set.
func (o *PageTheme) HasVisitorExperience() bool {
	if o != nil && !IsNil(o.VisitorExperience) {
		return true
	}

	return false
}

// SetVisitorExperience gets a reference to the given string and assigns it to the VisitorExperience field.
func (o *PageTheme) SetVisitorExperience(v string) {
	o.VisitorExperience = &v
}

// GetWorkspaceId returns the WorkspaceId field value if set, zero value otherwise.
func (o *PageTheme) GetWorkspaceId() string {
	if o == nil || IsNil(o.WorkspaceId) {
		var ret string
		return ret
	}
	return *o.WorkspaceId
}

// GetWorkspaceIdOk returns a tuple with the WorkspaceId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageTheme) GetWorkspaceIdOk() (*string, bool) {
	if o == nil || IsNil(o.WorkspaceId) {
		return nil, false
	}
	return o.WorkspaceId, true
}

// HasWorkspaceId returns a boolean if a field has been set.
func (o *PageTheme) HasWorkspaceId() bool {
	if o != nil && !IsNil(o.WorkspaceId) {
		return true
	}

	return false
}

// SetWorkspaceId gets a reference to the given string and assigns it to the WorkspaceId field.
func (o *PageTheme) SetWorkspaceId(v string) {
	o.WorkspaceId = &v
}

func (o PageTheme) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PageTheme) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BadgeEnabled) {
		toSerialize["badgeEnabled"] = o.BadgeEnabled
	}
	if !IsNil(o.BadgeIconColor) {
		toSerialize["badgeIconColor"] = o.BadgeIconColor
	}
	if !IsNil(o.BadgeId) {
		toSerialize["badgeId"] = o.BadgeId
	}
	if !IsNil(o.BadgeShape) {
		toSerialize["badgeShape"] = o.BadgeShape
	}
	if !IsNil(o.BodyAlign) {
		toSerialize["bodyAlign"] = o.BodyAlign
	}
	if !IsNil(o.BodyFont) {
		toSerialize["bodyFont"] = o.BodyFont
	}
	if !IsNil(o.CanvasColor) {
		toSerialize["canvasColor"] = o.CanvasColor
	}
	if !IsNil(o.CanvasCoverId) {
		toSerialize["canvasCoverId"] = o.CanvasCoverId
	}
	if !IsNil(o.CanvasCoverMode) {
		toSerialize["canvasCoverMode"] = o.CanvasCoverMode
	}
	if !IsNil(o.CanvasTexture) {
		toSerialize["canvasTexture"] = o.CanvasTexture
	}
	if !IsNil(o.ChoiceLayout) {
		toSerialize["choiceLayout"] = o.ChoiceLayout
	}
	if !IsNil(o.ContainerConfig) {
		toSerialize["containerConfig"] = o.ContainerConfig
	}
	if !IsNil(o.ContainerEnabled) {
		toSerialize["containerEnabled"] = o.ContainerEnabled
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["createdAt"] = o.CreatedAt
	}
	if !IsNil(o.ElementColor) {
		toSerialize["elementColor"] = o.ElementColor
	}
	if !IsNil(o.ElementFont) {
		toSerialize["elementFont"] = o.ElementFont
	}
	if !IsNil(o.ElementStyles) {
		toSerialize["elementStyles"] = o.ElementStyles
	}
	if !IsNil(o.FooterLogoEnabled) {
		toSerialize["footerLogoEnabled"] = o.FooterLogoEnabled
	}
	if !IsNil(o.FooterLogoId) {
		toSerialize["footerLogoId"] = o.FooterLogoId
	}
	if !IsNil(o.HeadingAlign) {
		toSerialize["headingAlign"] = o.HeadingAlign
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.QuestionStyle) {
		toSerialize["questionStyle"] = o.QuestionStyle
	}
	if !IsNil(o.Shape) {
		toSerialize["shape"] = o.Shape
	}
	if !IsNil(o.SmallText) {
		toSerialize["smallText"] = o.SmallText
	}
	if !IsNil(o.TextColor) {
		toSerialize["textColor"] = o.TextColor
	}
	if !IsNil(o.TitleFont) {
		toSerialize["titleFont"] = o.TitleFont
	}
	if !IsNil(o.UpdatedAt) {
		toSerialize["updatedAt"] = o.UpdatedAt
	}
	if !IsNil(o.VisitorExperience) {
		toSerialize["visitorExperience"] = o.VisitorExperience
	}
	if !IsNil(o.WorkspaceId) {
		toSerialize["workspaceId"] = o.WorkspaceId
	}
	return toSerialize, nil
}

type NullablePageTheme struct {
	value *PageTheme
	isSet bool
}

func (v NullablePageTheme) Get() *PageTheme {
	return v.value
}

func (v *NullablePageTheme) Set(val *PageTheme) {
	v.value = val
	v.isSet = true
}

func (v NullablePageTheme) IsSet() bool {
	return v.isSet
}

func (v *NullablePageTheme) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePageTheme(val *PageTheme) *NullablePageTheme {
	return &NullablePageTheme{value: val, isSet: true}
}

func (v NullablePageTheme) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePageTheme) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


