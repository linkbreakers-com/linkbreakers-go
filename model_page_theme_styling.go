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

// checks if the PageThemeStyling type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PageThemeStyling{}

// PageThemeStyling The set of visual values a page theme carries.  It exists so styling can travel on its own, detached from any stored theme: the editor sends what is currently on screen, which is what the user means by \"this look\", whether or not it has been saved yet.
type PageThemeStyling struct {
	Badge *MediaSelection `json:"badge,omitempty"`
	BadgeEnabled *bool `json:"badgeEnabled,omitempty"`
	BadgeIconColor *ThemeColor `json:"badgeIconColor,omitempty"`
	BadgeShape *string `json:"badgeShape,omitempty"`
	// How body text is aligned: descriptions, messages, helper text, and answer and button labels. Unset: centered.
	BodyAlign *string `json:"bodyAlign,omitempty"`
	BodyFont *FontConfig `json:"bodyFont,omitempty"`
	CanvasColor *ThemeColor `json:"canvasColor,omitempty"`
	CanvasCover *MediaSelection `json:"canvasCover,omitempty"`
	CanvasCoverMode *string `json:"canvasCoverMode,omitempty"`
	// A texture over the page background: NONE, BRUSHED (paint streaks), GRAIN (film grain) or STARFIELD (stars and Milky Way dust, for dark backgrounds).
	CanvasTexture *string `json:"canvasTexture,omitempty"`
	// How a choice question lays out its answers: AUTO (pills or tiles) or ROWS (full-width rows).
	ChoiceLayout *string `json:"choiceLayout,omitempty"`
	ContainerConfig *ContainerConfig `json:"containerConfig,omitempty"`
	ContainerEnabled *bool `json:"containerEnabled,omitempty"`
	ElementColor *ThemeColor `json:"elementColor,omitempty"`
	ElementFont *FontConfig `json:"elementFont,omitempty"`
	// The look of each element job (SURFACE, SELECTED, ACTION, QUIET). A job not listed follows the shape, the colours, or the look it always had.
	ElementStyles []PageThemeElementStyleOverride `json:"elementStyles,omitempty"`
	FooterLogo *MediaSelection `json:"footerLogo,omitempty"`
	// Small company or brand logo centered at the bottom of every page using the theme.
	FooterLogoEnabled *bool `json:"footerLogoEnabled,omitempty"`
	// How titles and headline questions are aligned. Unset: centered titles, questions as the question style draws them.
	HeadingAlign *string `json:"headingAlign,omitempty"`
	// How a form's questions are written: LABEL (small, in the text style) or HEADLINE (centered, in the title style).
	QuestionStyle *string `json:"questionStyle,omitempty"`
	// The radius and border every job follows unless it has its own.
	Shape *PageThemeShape `json:"shape,omitempty"`
	// Small text: hints, counters, consent. Unset: the body text, smaller and muted.
	SmallText *PageThemeSmallText `json:"smallText,omitempty"`
	// The text colour headings, body and small text follow when their own font has none.
	TextColor *ThemeColor `json:"textColor,omitempty"`
	TitleFont *FontConfig `json:"titleFont,omitempty"`
	// How visitor pages pace their content: CLASSIC shows a whole form at once and a chat as a transcript, FOCUSED one question per screen and a chat as a short fading trail.
	VisitorExperience *string `json:"visitorExperience,omitempty"`
}

// NewPageThemeStyling instantiates a new PageThemeStyling object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPageThemeStyling() *PageThemeStyling {
	this := PageThemeStyling{}
	return &this
}

// NewPageThemeStylingWithDefaults instantiates a new PageThemeStyling object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPageThemeStylingWithDefaults() *PageThemeStyling {
	this := PageThemeStyling{}
	return &this
}

// GetBadge returns the Badge field value if set, zero value otherwise.
func (o *PageThemeStyling) GetBadge() MediaSelection {
	if o == nil || IsNil(o.Badge) {
		var ret MediaSelection
		return ret
	}
	return *o.Badge
}

// GetBadgeOk returns a tuple with the Badge field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageThemeStyling) GetBadgeOk() (*MediaSelection, bool) {
	if o == nil || IsNil(o.Badge) {
		return nil, false
	}
	return o.Badge, true
}

// HasBadge returns a boolean if a field has been set.
func (o *PageThemeStyling) HasBadge() bool {
	if o != nil && !IsNil(o.Badge) {
		return true
	}

	return false
}

// SetBadge gets a reference to the given MediaSelection and assigns it to the Badge field.
func (o *PageThemeStyling) SetBadge(v MediaSelection) {
	o.Badge = &v
}

// GetBadgeEnabled returns the BadgeEnabled field value if set, zero value otherwise.
func (o *PageThemeStyling) GetBadgeEnabled() bool {
	if o == nil || IsNil(o.BadgeEnabled) {
		var ret bool
		return ret
	}
	return *o.BadgeEnabled
}

// GetBadgeEnabledOk returns a tuple with the BadgeEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageThemeStyling) GetBadgeEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.BadgeEnabled) {
		return nil, false
	}
	return o.BadgeEnabled, true
}

// HasBadgeEnabled returns a boolean if a field has been set.
func (o *PageThemeStyling) HasBadgeEnabled() bool {
	if o != nil && !IsNil(o.BadgeEnabled) {
		return true
	}

	return false
}

// SetBadgeEnabled gets a reference to the given bool and assigns it to the BadgeEnabled field.
func (o *PageThemeStyling) SetBadgeEnabled(v bool) {
	o.BadgeEnabled = &v
}

// GetBadgeIconColor returns the BadgeIconColor field value if set, zero value otherwise.
func (o *PageThemeStyling) GetBadgeIconColor() ThemeColor {
	if o == nil || IsNil(o.BadgeIconColor) {
		var ret ThemeColor
		return ret
	}
	return *o.BadgeIconColor
}

// GetBadgeIconColorOk returns a tuple with the BadgeIconColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageThemeStyling) GetBadgeIconColorOk() (*ThemeColor, bool) {
	if o == nil || IsNil(o.BadgeIconColor) {
		return nil, false
	}
	return o.BadgeIconColor, true
}

// HasBadgeIconColor returns a boolean if a field has been set.
func (o *PageThemeStyling) HasBadgeIconColor() bool {
	if o != nil && !IsNil(o.BadgeIconColor) {
		return true
	}

	return false
}

// SetBadgeIconColor gets a reference to the given ThemeColor and assigns it to the BadgeIconColor field.
func (o *PageThemeStyling) SetBadgeIconColor(v ThemeColor) {
	o.BadgeIconColor = &v
}

// GetBadgeShape returns the BadgeShape field value if set, zero value otherwise.
func (o *PageThemeStyling) GetBadgeShape() string {
	if o == nil || IsNil(o.BadgeShape) {
		var ret string
		return ret
	}
	return *o.BadgeShape
}

// GetBadgeShapeOk returns a tuple with the BadgeShape field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageThemeStyling) GetBadgeShapeOk() (*string, bool) {
	if o == nil || IsNil(o.BadgeShape) {
		return nil, false
	}
	return o.BadgeShape, true
}

// HasBadgeShape returns a boolean if a field has been set.
func (o *PageThemeStyling) HasBadgeShape() bool {
	if o != nil && !IsNil(o.BadgeShape) {
		return true
	}

	return false
}

// SetBadgeShape gets a reference to the given string and assigns it to the BadgeShape field.
func (o *PageThemeStyling) SetBadgeShape(v string) {
	o.BadgeShape = &v
}

// GetBodyAlign returns the BodyAlign field value if set, zero value otherwise.
func (o *PageThemeStyling) GetBodyAlign() string {
	if o == nil || IsNil(o.BodyAlign) {
		var ret string
		return ret
	}
	return *o.BodyAlign
}

// GetBodyAlignOk returns a tuple with the BodyAlign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageThemeStyling) GetBodyAlignOk() (*string, bool) {
	if o == nil || IsNil(o.BodyAlign) {
		return nil, false
	}
	return o.BodyAlign, true
}

// HasBodyAlign returns a boolean if a field has been set.
func (o *PageThemeStyling) HasBodyAlign() bool {
	if o != nil && !IsNil(o.BodyAlign) {
		return true
	}

	return false
}

// SetBodyAlign gets a reference to the given string and assigns it to the BodyAlign field.
func (o *PageThemeStyling) SetBodyAlign(v string) {
	o.BodyAlign = &v
}

// GetBodyFont returns the BodyFont field value if set, zero value otherwise.
func (o *PageThemeStyling) GetBodyFont() FontConfig {
	if o == nil || IsNil(o.BodyFont) {
		var ret FontConfig
		return ret
	}
	return *o.BodyFont
}

// GetBodyFontOk returns a tuple with the BodyFont field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageThemeStyling) GetBodyFontOk() (*FontConfig, bool) {
	if o == nil || IsNil(o.BodyFont) {
		return nil, false
	}
	return o.BodyFont, true
}

// HasBodyFont returns a boolean if a field has been set.
func (o *PageThemeStyling) HasBodyFont() bool {
	if o != nil && !IsNil(o.BodyFont) {
		return true
	}

	return false
}

// SetBodyFont gets a reference to the given FontConfig and assigns it to the BodyFont field.
func (o *PageThemeStyling) SetBodyFont(v FontConfig) {
	o.BodyFont = &v
}

// GetCanvasColor returns the CanvasColor field value if set, zero value otherwise.
func (o *PageThemeStyling) GetCanvasColor() ThemeColor {
	if o == nil || IsNil(o.CanvasColor) {
		var ret ThemeColor
		return ret
	}
	return *o.CanvasColor
}

// GetCanvasColorOk returns a tuple with the CanvasColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageThemeStyling) GetCanvasColorOk() (*ThemeColor, bool) {
	if o == nil || IsNil(o.CanvasColor) {
		return nil, false
	}
	return o.CanvasColor, true
}

// HasCanvasColor returns a boolean if a field has been set.
func (o *PageThemeStyling) HasCanvasColor() bool {
	if o != nil && !IsNil(o.CanvasColor) {
		return true
	}

	return false
}

// SetCanvasColor gets a reference to the given ThemeColor and assigns it to the CanvasColor field.
func (o *PageThemeStyling) SetCanvasColor(v ThemeColor) {
	o.CanvasColor = &v
}

// GetCanvasCover returns the CanvasCover field value if set, zero value otherwise.
func (o *PageThemeStyling) GetCanvasCover() MediaSelection {
	if o == nil || IsNil(o.CanvasCover) {
		var ret MediaSelection
		return ret
	}
	return *o.CanvasCover
}

// GetCanvasCoverOk returns a tuple with the CanvasCover field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageThemeStyling) GetCanvasCoverOk() (*MediaSelection, bool) {
	if o == nil || IsNil(o.CanvasCover) {
		return nil, false
	}
	return o.CanvasCover, true
}

// HasCanvasCover returns a boolean if a field has been set.
func (o *PageThemeStyling) HasCanvasCover() bool {
	if o != nil && !IsNil(o.CanvasCover) {
		return true
	}

	return false
}

// SetCanvasCover gets a reference to the given MediaSelection and assigns it to the CanvasCover field.
func (o *PageThemeStyling) SetCanvasCover(v MediaSelection) {
	o.CanvasCover = &v
}

// GetCanvasCoverMode returns the CanvasCoverMode field value if set, zero value otherwise.
func (o *PageThemeStyling) GetCanvasCoverMode() string {
	if o == nil || IsNil(o.CanvasCoverMode) {
		var ret string
		return ret
	}
	return *o.CanvasCoverMode
}

// GetCanvasCoverModeOk returns a tuple with the CanvasCoverMode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageThemeStyling) GetCanvasCoverModeOk() (*string, bool) {
	if o == nil || IsNil(o.CanvasCoverMode) {
		return nil, false
	}
	return o.CanvasCoverMode, true
}

// HasCanvasCoverMode returns a boolean if a field has been set.
func (o *PageThemeStyling) HasCanvasCoverMode() bool {
	if o != nil && !IsNil(o.CanvasCoverMode) {
		return true
	}

	return false
}

// SetCanvasCoverMode gets a reference to the given string and assigns it to the CanvasCoverMode field.
func (o *PageThemeStyling) SetCanvasCoverMode(v string) {
	o.CanvasCoverMode = &v
}

// GetCanvasTexture returns the CanvasTexture field value if set, zero value otherwise.
func (o *PageThemeStyling) GetCanvasTexture() string {
	if o == nil || IsNil(o.CanvasTexture) {
		var ret string
		return ret
	}
	return *o.CanvasTexture
}

// GetCanvasTextureOk returns a tuple with the CanvasTexture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageThemeStyling) GetCanvasTextureOk() (*string, bool) {
	if o == nil || IsNil(o.CanvasTexture) {
		return nil, false
	}
	return o.CanvasTexture, true
}

// HasCanvasTexture returns a boolean if a field has been set.
func (o *PageThemeStyling) HasCanvasTexture() bool {
	if o != nil && !IsNil(o.CanvasTexture) {
		return true
	}

	return false
}

// SetCanvasTexture gets a reference to the given string and assigns it to the CanvasTexture field.
func (o *PageThemeStyling) SetCanvasTexture(v string) {
	o.CanvasTexture = &v
}

// GetChoiceLayout returns the ChoiceLayout field value if set, zero value otherwise.
func (o *PageThemeStyling) GetChoiceLayout() string {
	if o == nil || IsNil(o.ChoiceLayout) {
		var ret string
		return ret
	}
	return *o.ChoiceLayout
}

// GetChoiceLayoutOk returns a tuple with the ChoiceLayout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageThemeStyling) GetChoiceLayoutOk() (*string, bool) {
	if o == nil || IsNil(o.ChoiceLayout) {
		return nil, false
	}
	return o.ChoiceLayout, true
}

// HasChoiceLayout returns a boolean if a field has been set.
func (o *PageThemeStyling) HasChoiceLayout() bool {
	if o != nil && !IsNil(o.ChoiceLayout) {
		return true
	}

	return false
}

// SetChoiceLayout gets a reference to the given string and assigns it to the ChoiceLayout field.
func (o *PageThemeStyling) SetChoiceLayout(v string) {
	o.ChoiceLayout = &v
}

// GetContainerConfig returns the ContainerConfig field value if set, zero value otherwise.
func (o *PageThemeStyling) GetContainerConfig() ContainerConfig {
	if o == nil || IsNil(o.ContainerConfig) {
		var ret ContainerConfig
		return ret
	}
	return *o.ContainerConfig
}

// GetContainerConfigOk returns a tuple with the ContainerConfig field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageThemeStyling) GetContainerConfigOk() (*ContainerConfig, bool) {
	if o == nil || IsNil(o.ContainerConfig) {
		return nil, false
	}
	return o.ContainerConfig, true
}

// HasContainerConfig returns a boolean if a field has been set.
func (o *PageThemeStyling) HasContainerConfig() bool {
	if o != nil && !IsNil(o.ContainerConfig) {
		return true
	}

	return false
}

// SetContainerConfig gets a reference to the given ContainerConfig and assigns it to the ContainerConfig field.
func (o *PageThemeStyling) SetContainerConfig(v ContainerConfig) {
	o.ContainerConfig = &v
}

// GetContainerEnabled returns the ContainerEnabled field value if set, zero value otherwise.
func (o *PageThemeStyling) GetContainerEnabled() bool {
	if o == nil || IsNil(o.ContainerEnabled) {
		var ret bool
		return ret
	}
	return *o.ContainerEnabled
}

// GetContainerEnabledOk returns a tuple with the ContainerEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageThemeStyling) GetContainerEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.ContainerEnabled) {
		return nil, false
	}
	return o.ContainerEnabled, true
}

// HasContainerEnabled returns a boolean if a field has been set.
func (o *PageThemeStyling) HasContainerEnabled() bool {
	if o != nil && !IsNil(o.ContainerEnabled) {
		return true
	}

	return false
}

// SetContainerEnabled gets a reference to the given bool and assigns it to the ContainerEnabled field.
func (o *PageThemeStyling) SetContainerEnabled(v bool) {
	o.ContainerEnabled = &v
}

// GetElementColor returns the ElementColor field value if set, zero value otherwise.
func (o *PageThemeStyling) GetElementColor() ThemeColor {
	if o == nil || IsNil(o.ElementColor) {
		var ret ThemeColor
		return ret
	}
	return *o.ElementColor
}

// GetElementColorOk returns a tuple with the ElementColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageThemeStyling) GetElementColorOk() (*ThemeColor, bool) {
	if o == nil || IsNil(o.ElementColor) {
		return nil, false
	}
	return o.ElementColor, true
}

// HasElementColor returns a boolean if a field has been set.
func (o *PageThemeStyling) HasElementColor() bool {
	if o != nil && !IsNil(o.ElementColor) {
		return true
	}

	return false
}

// SetElementColor gets a reference to the given ThemeColor and assigns it to the ElementColor field.
func (o *PageThemeStyling) SetElementColor(v ThemeColor) {
	o.ElementColor = &v
}

// GetElementFont returns the ElementFont field value if set, zero value otherwise.
func (o *PageThemeStyling) GetElementFont() FontConfig {
	if o == nil || IsNil(o.ElementFont) {
		var ret FontConfig
		return ret
	}
	return *o.ElementFont
}

// GetElementFontOk returns a tuple with the ElementFont field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageThemeStyling) GetElementFontOk() (*FontConfig, bool) {
	if o == nil || IsNil(o.ElementFont) {
		return nil, false
	}
	return o.ElementFont, true
}

// HasElementFont returns a boolean if a field has been set.
func (o *PageThemeStyling) HasElementFont() bool {
	if o != nil && !IsNil(o.ElementFont) {
		return true
	}

	return false
}

// SetElementFont gets a reference to the given FontConfig and assigns it to the ElementFont field.
func (o *PageThemeStyling) SetElementFont(v FontConfig) {
	o.ElementFont = &v
}

// GetElementStyles returns the ElementStyles field value if set, zero value otherwise.
func (o *PageThemeStyling) GetElementStyles() []PageThemeElementStyleOverride {
	if o == nil || IsNil(o.ElementStyles) {
		var ret []PageThemeElementStyleOverride
		return ret
	}
	return o.ElementStyles
}

// GetElementStylesOk returns a tuple with the ElementStyles field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageThemeStyling) GetElementStylesOk() ([]PageThemeElementStyleOverride, bool) {
	if o == nil || IsNil(o.ElementStyles) {
		return nil, false
	}
	return o.ElementStyles, true
}

// HasElementStyles returns a boolean if a field has been set.
func (o *PageThemeStyling) HasElementStyles() bool {
	if o != nil && !IsNil(o.ElementStyles) {
		return true
	}

	return false
}

// SetElementStyles gets a reference to the given []PageThemeElementStyleOverride and assigns it to the ElementStyles field.
func (o *PageThemeStyling) SetElementStyles(v []PageThemeElementStyleOverride) {
	o.ElementStyles = v
}

// GetFooterLogo returns the FooterLogo field value if set, zero value otherwise.
func (o *PageThemeStyling) GetFooterLogo() MediaSelection {
	if o == nil || IsNil(o.FooterLogo) {
		var ret MediaSelection
		return ret
	}
	return *o.FooterLogo
}

// GetFooterLogoOk returns a tuple with the FooterLogo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageThemeStyling) GetFooterLogoOk() (*MediaSelection, bool) {
	if o == nil || IsNil(o.FooterLogo) {
		return nil, false
	}
	return o.FooterLogo, true
}

// HasFooterLogo returns a boolean if a field has been set.
func (o *PageThemeStyling) HasFooterLogo() bool {
	if o != nil && !IsNil(o.FooterLogo) {
		return true
	}

	return false
}

// SetFooterLogo gets a reference to the given MediaSelection and assigns it to the FooterLogo field.
func (o *PageThemeStyling) SetFooterLogo(v MediaSelection) {
	o.FooterLogo = &v
}

// GetFooterLogoEnabled returns the FooterLogoEnabled field value if set, zero value otherwise.
func (o *PageThemeStyling) GetFooterLogoEnabled() bool {
	if o == nil || IsNil(o.FooterLogoEnabled) {
		var ret bool
		return ret
	}
	return *o.FooterLogoEnabled
}

// GetFooterLogoEnabledOk returns a tuple with the FooterLogoEnabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageThemeStyling) GetFooterLogoEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.FooterLogoEnabled) {
		return nil, false
	}
	return o.FooterLogoEnabled, true
}

// HasFooterLogoEnabled returns a boolean if a field has been set.
func (o *PageThemeStyling) HasFooterLogoEnabled() bool {
	if o != nil && !IsNil(o.FooterLogoEnabled) {
		return true
	}

	return false
}

// SetFooterLogoEnabled gets a reference to the given bool and assigns it to the FooterLogoEnabled field.
func (o *PageThemeStyling) SetFooterLogoEnabled(v bool) {
	o.FooterLogoEnabled = &v
}

// GetHeadingAlign returns the HeadingAlign field value if set, zero value otherwise.
func (o *PageThemeStyling) GetHeadingAlign() string {
	if o == nil || IsNil(o.HeadingAlign) {
		var ret string
		return ret
	}
	return *o.HeadingAlign
}

// GetHeadingAlignOk returns a tuple with the HeadingAlign field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageThemeStyling) GetHeadingAlignOk() (*string, bool) {
	if o == nil || IsNil(o.HeadingAlign) {
		return nil, false
	}
	return o.HeadingAlign, true
}

// HasHeadingAlign returns a boolean if a field has been set.
func (o *PageThemeStyling) HasHeadingAlign() bool {
	if o != nil && !IsNil(o.HeadingAlign) {
		return true
	}

	return false
}

// SetHeadingAlign gets a reference to the given string and assigns it to the HeadingAlign field.
func (o *PageThemeStyling) SetHeadingAlign(v string) {
	o.HeadingAlign = &v
}

// GetQuestionStyle returns the QuestionStyle field value if set, zero value otherwise.
func (o *PageThemeStyling) GetQuestionStyle() string {
	if o == nil || IsNil(o.QuestionStyle) {
		var ret string
		return ret
	}
	return *o.QuestionStyle
}

// GetQuestionStyleOk returns a tuple with the QuestionStyle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageThemeStyling) GetQuestionStyleOk() (*string, bool) {
	if o == nil || IsNil(o.QuestionStyle) {
		return nil, false
	}
	return o.QuestionStyle, true
}

// HasQuestionStyle returns a boolean if a field has been set.
func (o *PageThemeStyling) HasQuestionStyle() bool {
	if o != nil && !IsNil(o.QuestionStyle) {
		return true
	}

	return false
}

// SetQuestionStyle gets a reference to the given string and assigns it to the QuestionStyle field.
func (o *PageThemeStyling) SetQuestionStyle(v string) {
	o.QuestionStyle = &v
}

// GetShape returns the Shape field value if set, zero value otherwise.
func (o *PageThemeStyling) GetShape() PageThemeShape {
	if o == nil || IsNil(o.Shape) {
		var ret PageThemeShape
		return ret
	}
	return *o.Shape
}

// GetShapeOk returns a tuple with the Shape field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageThemeStyling) GetShapeOk() (*PageThemeShape, bool) {
	if o == nil || IsNil(o.Shape) {
		return nil, false
	}
	return o.Shape, true
}

// HasShape returns a boolean if a field has been set.
func (o *PageThemeStyling) HasShape() bool {
	if o != nil && !IsNil(o.Shape) {
		return true
	}

	return false
}

// SetShape gets a reference to the given PageThemeShape and assigns it to the Shape field.
func (o *PageThemeStyling) SetShape(v PageThemeShape) {
	o.Shape = &v
}

// GetSmallText returns the SmallText field value if set, zero value otherwise.
func (o *PageThemeStyling) GetSmallText() PageThemeSmallText {
	if o == nil || IsNil(o.SmallText) {
		var ret PageThemeSmallText
		return ret
	}
	return *o.SmallText
}

// GetSmallTextOk returns a tuple with the SmallText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageThemeStyling) GetSmallTextOk() (*PageThemeSmallText, bool) {
	if o == nil || IsNil(o.SmallText) {
		return nil, false
	}
	return o.SmallText, true
}

// HasSmallText returns a boolean if a field has been set.
func (o *PageThemeStyling) HasSmallText() bool {
	if o != nil && !IsNil(o.SmallText) {
		return true
	}

	return false
}

// SetSmallText gets a reference to the given PageThemeSmallText and assigns it to the SmallText field.
func (o *PageThemeStyling) SetSmallText(v PageThemeSmallText) {
	o.SmallText = &v
}

// GetTextColor returns the TextColor field value if set, zero value otherwise.
func (o *PageThemeStyling) GetTextColor() ThemeColor {
	if o == nil || IsNil(o.TextColor) {
		var ret ThemeColor
		return ret
	}
	return *o.TextColor
}

// GetTextColorOk returns a tuple with the TextColor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageThemeStyling) GetTextColorOk() (*ThemeColor, bool) {
	if o == nil || IsNil(o.TextColor) {
		return nil, false
	}
	return o.TextColor, true
}

// HasTextColor returns a boolean if a field has been set.
func (o *PageThemeStyling) HasTextColor() bool {
	if o != nil && !IsNil(o.TextColor) {
		return true
	}

	return false
}

// SetTextColor gets a reference to the given ThemeColor and assigns it to the TextColor field.
func (o *PageThemeStyling) SetTextColor(v ThemeColor) {
	o.TextColor = &v
}

// GetTitleFont returns the TitleFont field value if set, zero value otherwise.
func (o *PageThemeStyling) GetTitleFont() FontConfig {
	if o == nil || IsNil(o.TitleFont) {
		var ret FontConfig
		return ret
	}
	return *o.TitleFont
}

// GetTitleFontOk returns a tuple with the TitleFont field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageThemeStyling) GetTitleFontOk() (*FontConfig, bool) {
	if o == nil || IsNil(o.TitleFont) {
		return nil, false
	}
	return o.TitleFont, true
}

// HasTitleFont returns a boolean if a field has been set.
func (o *PageThemeStyling) HasTitleFont() bool {
	if o != nil && !IsNil(o.TitleFont) {
		return true
	}

	return false
}

// SetTitleFont gets a reference to the given FontConfig and assigns it to the TitleFont field.
func (o *PageThemeStyling) SetTitleFont(v FontConfig) {
	o.TitleFont = &v
}

// GetVisitorExperience returns the VisitorExperience field value if set, zero value otherwise.
func (o *PageThemeStyling) GetVisitorExperience() string {
	if o == nil || IsNil(o.VisitorExperience) {
		var ret string
		return ret
	}
	return *o.VisitorExperience
}

// GetVisitorExperienceOk returns a tuple with the VisitorExperience field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PageThemeStyling) GetVisitorExperienceOk() (*string, bool) {
	if o == nil || IsNil(o.VisitorExperience) {
		return nil, false
	}
	return o.VisitorExperience, true
}

// HasVisitorExperience returns a boolean if a field has been set.
func (o *PageThemeStyling) HasVisitorExperience() bool {
	if o != nil && !IsNil(o.VisitorExperience) {
		return true
	}

	return false
}

// SetVisitorExperience gets a reference to the given string and assigns it to the VisitorExperience field.
func (o *PageThemeStyling) SetVisitorExperience(v string) {
	o.VisitorExperience = &v
}

func (o PageThemeStyling) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PageThemeStyling) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Badge) {
		toSerialize["badge"] = o.Badge
	}
	if !IsNil(o.BadgeEnabled) {
		toSerialize["badgeEnabled"] = o.BadgeEnabled
	}
	if !IsNil(o.BadgeIconColor) {
		toSerialize["badgeIconColor"] = o.BadgeIconColor
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
	if !IsNil(o.CanvasCover) {
		toSerialize["canvasCover"] = o.CanvasCover
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
	if !IsNil(o.ElementColor) {
		toSerialize["elementColor"] = o.ElementColor
	}
	if !IsNil(o.ElementFont) {
		toSerialize["elementFont"] = o.ElementFont
	}
	if !IsNil(o.ElementStyles) {
		toSerialize["elementStyles"] = o.ElementStyles
	}
	if !IsNil(o.FooterLogo) {
		toSerialize["footerLogo"] = o.FooterLogo
	}
	if !IsNil(o.FooterLogoEnabled) {
		toSerialize["footerLogoEnabled"] = o.FooterLogoEnabled
	}
	if !IsNil(o.HeadingAlign) {
		toSerialize["headingAlign"] = o.HeadingAlign
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
	if !IsNil(o.VisitorExperience) {
		toSerialize["visitorExperience"] = o.VisitorExperience
	}
	return toSerialize, nil
}

type NullablePageThemeStyling struct {
	value *PageThemeStyling
	isSet bool
}

func (v NullablePageThemeStyling) Get() *PageThemeStyling {
	return v.value
}

func (v *NullablePageThemeStyling) Set(val *PageThemeStyling) {
	v.value = val
	v.isSet = true
}

func (v NullablePageThemeStyling) IsSet() bool {
	return v.isSet
}

func (v *NullablePageThemeStyling) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePageThemeStyling(val *PageThemeStyling) *NullablePageThemeStyling {
	return &NullablePageThemeStyling{value: val, isSet: true}
}

func (v NullablePageThemeStyling) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePageThemeStyling) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


