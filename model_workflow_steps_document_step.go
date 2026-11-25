/*
Linkbreakers API

The Linkbreakers REST API: links, QR codes, workflows, visitors and analytics, as JSON over HTTPS. Send a workspace API token as `Authorization: Bearer <token>`.  ## Entry points  `GET https://api.linkbreakers.com/` and `GET https://api.linkbreakers.com/v1` answer without credentials with the API's name, version and links. This document is served at `https://api.linkbreakers.com/openapi.json`.  ## Rate limits  Every response describes where the caller stands in the IETF `RateLimit-Policy` and `RateLimit` header fields (draft-ietf-httpapi-ratelimit-headers), and in `RateLimit-Limit`, `RateLimit-Remaining`, `RateLimit-Reset` and `X-RateLimit-Limit`, `X-RateLimit-Remaining`, `X-RateLimit-Reset` (a Unix timestamp) for clients that read those. A request past the limit gets `429 Too Many Requests` with `Retry-After` in seconds.  - Workspace API token: 720 requests per 60 seconds, per token. - Dashboard session or OAuth token: 600 requests per 60 seconds, per token. - No credentials: 300 requests per 60 seconds, per client IP.  Each plan also includes a number of API calls per billing cycle. Past it, calls answer `429` with the code `MAX_PUBLIC_API_QUERIED_REACHED` until the cycle renews or the plan changes.  ## Pagination  Every list pages with opaque cursors (AIP-158). Pass `pageSize`, then send the `nextPageToken` of each response back as `pageToken` (or its alias `cursor`) to get the next page. An empty or absent `nextPageToken` means the list is complete. Each list operation describes this in `x-pagination`.  ## Batches and background jobs  `POST /v1/links/bulk` creates up to 250 links in one transactional request; `POST /v1/links/batch` takes the same links as a bare JSON array. `POST /v1/links/bulk/jobs` accepts the same batch and answers `202 Accepted` at once with a link job and a `Location` header; poll `GET /v1/links/jobs/{id}` until `status` is `LINK_JOB_STATUS_SUCCEEDED` (the created link ids are in `linksBulkCreate.linkIds`) or `LINK_JOB_STATUS_FAILED` (`error` says why, and nothing was created). Link jobs can be read for 24 hours after they finish.  ## Paused workspaces  A free workspace whose members have not verified any email address is paused 7 days after sign-up. Its calls then answer `403` with the code `WORKSPACE_UNVERIFIED` and change nothing; QR codes and links keep working for visitors. To lift the pause, ask the workspace owner to verify their email: `POST /v1/members/email-verification` sends the verification email again, and works while the workspace is paused. `unverifiedPauseAt` on the workspace says when the pause starts.  ## Versioning and deprecation  The version is in the path: every operation lives under `/v1`. Within a version, changes are additive only: new operations, new optional request fields, new response fields and new enum values. Clients should ignore fields and enum values they do not know. Anything that would break an existing client ships under a new path version instead.  An operation is retired in two steps. First it is deprecated: it keeps working, is flagged `deprecated: true` in this document, and every response carries a `Deprecation` header field (RFC 9745) with the time it was deprecated and a `Link` with `rel=\"deprecation\"` to the policy. Then, at least 180 days later, it may be removed; that date is announced in advance in a `Sunset` header field (RFC 8594) and as `x-sunset` on the operation. The full policy: https://linkbreakers.com/help/article/api-versioning-and-deprecation-policy  A renamed operation keeps answering its former path while clients still call it, until that path's sunset at the latest: a former path no client calls for 14 days is removed. Responses on a former path carry the same `Deprecation`, `Sunset` and `Link` header fields; this document lists only the current path, and the policy lists the former paths that still answer.  ## Idempotent retries  Every POST, PUT and PATCH accepts an optional `Idempotency-Key` header (draft-ietf-httpapi-idempotency-key-header). Send a new unique value, such as a UUID, with each logical request and reuse it only to retry that request. A retry with the same key and the same request returns the first response, marked `Idempotent-Replayed: true`, instead of running again. Keys are scoped to the token and kept for 24 hours. The same key with a different request is refused with `422`, and a retry while the first request is still running gets `409` with `Retry-After`. Server errors (5xx) and refusals such as 401, 403 and 429 are not stored, so retrying them runs the request again.  ## Identifying your client  Send a `Linkbreakers-Client` request header saying which client made the call: `<client>/<version> (<details>)`, for example `my-integration/1.2.0 (python 3.12; linux)`. The official SDKs, the CLI and the MCP bridge set it for you, as `sdk-python/2.3.1 (python 3.12.4; linux)` or `cli/1.4.0 (darwin arm64)`. It is optional and changes nothing about how a request is handled; it tells us which clients and versions are in use, so we know what can be retired and what to fix first. Keep the details short and free of personal data. Browsers may send it too: it is allowed by CORS.

API version: 1.205.1
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package linkbreakers

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the WorkflowStepsDocumentStep type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WorkflowStepsDocumentStep{}

// WorkflowStepsDocumentStep One step of a document.
type WorkflowStepsDocumentStep struct {
	// The name of the agent a chat bot step talks through, in place of its id, which the payload leaves empty. Imported, the step talks through the workspace's agent of that name, or else its first agent.
	AgentName *string `json:"agentName,omitempty"`
	// What the step saves to the visitor's profile.
	AttributeCapture *AttributeCapture `json:"attributeCapture,omitempty"`
	// The page's back button.
	BackButton *BackButton `json:"backButton,omitempty"`
	// - EVENT_ACTION_UNSPECIFIED: Event action not specified  - EVENT_ACTION_CLIENT_REDIRECT: Redirect visitors to a destination URL  - EVENT_ACTION_LINK_VISIT: Entry step triggered when the link is visited.  - EVENT_ACTION_PASSWORD_VERIFY: Require visitors to enter a password before proceeding  - EVENT_ACTION_FORM_SUBMIT: Present a data collection form to visitors  - EVENT_ACTION_CONTACT_CARD: Generate a vCard (VCF) response with contact information  - EVENT_ACTION_MULTI_LINK: A page of buttons where each one continues into more workflow: the button leads to a next step, not straight out. Records which button the visitor pressed, so use it over SOCIAL_LINKS_PAGE when that choice is worth measuring or when a branch needs to do more than redirect. This is the menu step: for several documents or PDFs, give each one a button connected to an OPEN_MEDIA step for media uploaded to the media library, or to a CLIENT_REDIRECT when the document is already hosted at a URL.  - EVENT_ACTION_VISIT_TYPE_CONDITION: Route visitors based on whether it's a first or returning visit  - EVENT_ACTION_COUNTRY_CONDITION: Route visitors based on their country location  - EVENT_ACTION_DAY_OF_WEEK_CONDITION: Route visitors based on day of week  - EVENT_ACTION_TIME_OF_DAY_CONDITION: Route visitors based on time of day  - EVENT_ACTION_DAY_OF_MONTH_CONDITION: Route visitors based on day of month  - EVENT_ACTION_SPECIFIC_DATE_CONDITION: Route visitors based on specific calendar dates  - EVENT_ACTION_SOCIAL_LINKS_PAGE: A page of buttons that each leave for an external URL. Social profiles are the common case, not the only one: app stores, docs, booking pages and storefronts all work, and the name is the only thing about it that is social. It is an exit step, so the visitor leaves from here.  - EVENT_ACTION_DEVICE_TYPE_CONDITION: Route visitors based on device type (e.g. MOBILE, DESKTOP, TABLET)  - EVENT_ACTION_DEVICE_PLATFORM_CONDITION: Route visitors based on device platform / operating system (e.g. iOS, Android, macOS)  - EVENT_ACTION_DEVICE_BROWSER_CONDITION: Route visitors based on their browser: a browser family (Chrome, Safari, Firefox...) or the in-app browser of the app a link was opened in (Instagram, TikTok, Facebook, LinkedIn...)  - EVENT_ACTION_DEVICE_BRAND_CONDITION: Route visitors based on hardware brand (e.g. Apple, Samsung, Google)  - EVENT_ACTION_DEVICE_ACTOR_CONDITION: Route visitors based on actor classification: HUMAN, BOT, AGENT, AUTOMATED or UNKNOWN. AGENT is an AI agent that names itself (GPTBot, Claude-User, Perplexity-User...). AUTOMATED is a browser driven by software, whoever is behind it: headless Chrome from its request, and a visible browser under automation once a link with the Advanced data collection level has read navigator.webdriver, from that device's next visit on.  - EVENT_ACTION_VISITOR_DATA_CONDITION: Route visitors based on their profile data or custom attributes  - EVENT_ACTION_THANK_YOU_PAGE: Display a final thank you / completion message (exit step)  - EVENT_ACTION_MESSAGE_PAGE: Text with ONE way onward: a single button, or a timer that moves the visitor automatically. It carries no list of destinations, so a page that should offer a choice between two or more places is not this step. Use SOCIAL_LINKS_PAGE when the choices leave for external URLs, or MULTI_LINK when each choice continues into more workflow. Never put links, documents or PDFs on it: it is a message, nothing else.  - EVENT_ACTION_OPEN_MEDIA: Hand the visitor one piece of uploaded media, opened in their browser (exit step): a PDF, an image (PNG, JPEG, WebP, SVG) or an MP4 video. The visitor is sent straight to the media, with no page of its own. Any media of the library can be opened: it is uploaded once and referenced by its mediaId; swapping it means pointing the step at another media item, and the QR code never changes.  - EVENT_ACTION_LOCATION_REQUEST: Ask the visitor for their precise location (GPS) through the browser, with their permission (composite: two branches, \"location\" when it was shared, \"else\" when it was refused, skipped or is unavailable).  - EVENT_ACTION_LANGUAGE_CONDITION: Route visitors based on their phone or browser language (the first language it sends, matched on its two-letter ISO 639-1 base, so \"fr\" catches fr-FR and fr-CA)  - EVENT_ACTION_CHAT_BOT: A conversation with one of the workspace's agents that collects visitor attributes (composite: two branches, \"completed\" when every required objective is known, \"else\" otherwise)  - EVENT_ACTION_TRAFFIC_SPLIT: Split visitors between 2 to 20 variants by percentage, for an A/B test (composite: one branch per variant, no \"else\"). Assignment is deterministic and sticky per device by default. See TrafficSplitPayload.
	EventAction string `json:"eventAction"`
	// The step's name within the document, which edges use.
	Key *string `json:"key,omitempty"`
	// The step's file, in place of its media id, which the payload leaves empty: the one an open media step hands out, or a contact card's photo. Imported into another workspace, the file is added to its media library.
	Media *WorkflowStepsDocumentMedia `json:"media,omitempty"`
	// A readable name for the step, such as \"Form: Join the list\". Ignored on import.
	Name *string `json:"name,omitempty"`
	Payload *WorkflowStepPayload `json:"payload,omitempty"`
	// Where the step sits relative to the top left corner of the copied group.
	Position *CanvasPosition `json:"position,omitempty"`
}

type _WorkflowStepsDocumentStep WorkflowStepsDocumentStep

// NewWorkflowStepsDocumentStep instantiates a new WorkflowStepsDocumentStep object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowStepsDocumentStep(eventAction string) *WorkflowStepsDocumentStep {
	this := WorkflowStepsDocumentStep{}
	this.EventAction = eventAction
	return &this
}

// NewWorkflowStepsDocumentStepWithDefaults instantiates a new WorkflowStepsDocumentStep object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowStepsDocumentStepWithDefaults() *WorkflowStepsDocumentStep {
	this := WorkflowStepsDocumentStep{}
	return &this
}

// GetAgentName returns the AgentName field value if set, zero value otherwise.
func (o *WorkflowStepsDocumentStep) GetAgentName() string {
	if o == nil || IsNil(o.AgentName) {
		var ret string
		return ret
	}
	return *o.AgentName
}

// GetAgentNameOk returns a tuple with the AgentName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStepsDocumentStep) GetAgentNameOk() (*string, bool) {
	if o == nil || IsNil(o.AgentName) {
		return nil, false
	}
	return o.AgentName, true
}

// HasAgentName returns a boolean if a field has been set.
func (o *WorkflowStepsDocumentStep) HasAgentName() bool {
	if o != nil && !IsNil(o.AgentName) {
		return true
	}

	return false
}

// SetAgentName gets a reference to the given string and assigns it to the AgentName field.
func (o *WorkflowStepsDocumentStep) SetAgentName(v string) {
	o.AgentName = &v
}

// GetAttributeCapture returns the AttributeCapture field value if set, zero value otherwise.
func (o *WorkflowStepsDocumentStep) GetAttributeCapture() AttributeCapture {
	if o == nil || IsNil(o.AttributeCapture) {
		var ret AttributeCapture
		return ret
	}
	return *o.AttributeCapture
}

// GetAttributeCaptureOk returns a tuple with the AttributeCapture field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStepsDocumentStep) GetAttributeCaptureOk() (*AttributeCapture, bool) {
	if o == nil || IsNil(o.AttributeCapture) {
		return nil, false
	}
	return o.AttributeCapture, true
}

// HasAttributeCapture returns a boolean if a field has been set.
func (o *WorkflowStepsDocumentStep) HasAttributeCapture() bool {
	if o != nil && !IsNil(o.AttributeCapture) {
		return true
	}

	return false
}

// SetAttributeCapture gets a reference to the given AttributeCapture and assigns it to the AttributeCapture field.
func (o *WorkflowStepsDocumentStep) SetAttributeCapture(v AttributeCapture) {
	o.AttributeCapture = &v
}

// GetBackButton returns the BackButton field value if set, zero value otherwise.
func (o *WorkflowStepsDocumentStep) GetBackButton() BackButton {
	if o == nil || IsNil(o.BackButton) {
		var ret BackButton
		return ret
	}
	return *o.BackButton
}

// GetBackButtonOk returns a tuple with the BackButton field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStepsDocumentStep) GetBackButtonOk() (*BackButton, bool) {
	if o == nil || IsNil(o.BackButton) {
		return nil, false
	}
	return o.BackButton, true
}

// HasBackButton returns a boolean if a field has been set.
func (o *WorkflowStepsDocumentStep) HasBackButton() bool {
	if o != nil && !IsNil(o.BackButton) {
		return true
	}

	return false
}

// SetBackButton gets a reference to the given BackButton and assigns it to the BackButton field.
func (o *WorkflowStepsDocumentStep) SetBackButton(v BackButton) {
	o.BackButton = &v
}

// GetEventAction returns the EventAction field value
func (o *WorkflowStepsDocumentStep) GetEventAction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EventAction
}

// GetEventActionOk returns a tuple with the EventAction field value
// and a boolean to check if the value has been set.
func (o *WorkflowStepsDocumentStep) GetEventActionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EventAction, true
}

// SetEventAction sets field value
func (o *WorkflowStepsDocumentStep) SetEventAction(v string) {
	o.EventAction = v
}

// GetKey returns the Key field value if set, zero value otherwise.
func (o *WorkflowStepsDocumentStep) GetKey() string {
	if o == nil || IsNil(o.Key) {
		var ret string
		return ret
	}
	return *o.Key
}

// GetKeyOk returns a tuple with the Key field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStepsDocumentStep) GetKeyOk() (*string, bool) {
	if o == nil || IsNil(o.Key) {
		return nil, false
	}
	return o.Key, true
}

// HasKey returns a boolean if a field has been set.
func (o *WorkflowStepsDocumentStep) HasKey() bool {
	if o != nil && !IsNil(o.Key) {
		return true
	}

	return false
}

// SetKey gets a reference to the given string and assigns it to the Key field.
func (o *WorkflowStepsDocumentStep) SetKey(v string) {
	o.Key = &v
}

// GetMedia returns the Media field value if set, zero value otherwise.
func (o *WorkflowStepsDocumentStep) GetMedia() WorkflowStepsDocumentMedia {
	if o == nil || IsNil(o.Media) {
		var ret WorkflowStepsDocumentMedia
		return ret
	}
	return *o.Media
}

// GetMediaOk returns a tuple with the Media field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStepsDocumentStep) GetMediaOk() (*WorkflowStepsDocumentMedia, bool) {
	if o == nil || IsNil(o.Media) {
		return nil, false
	}
	return o.Media, true
}

// HasMedia returns a boolean if a field has been set.
func (o *WorkflowStepsDocumentStep) HasMedia() bool {
	if o != nil && !IsNil(o.Media) {
		return true
	}

	return false
}

// SetMedia gets a reference to the given WorkflowStepsDocumentMedia and assigns it to the Media field.
func (o *WorkflowStepsDocumentStep) SetMedia(v WorkflowStepsDocumentMedia) {
	o.Media = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowStepsDocumentStep) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStepsDocumentStep) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowStepsDocumentStep) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowStepsDocumentStep) SetName(v string) {
	o.Name = &v
}

// GetPayload returns the Payload field value if set, zero value otherwise.
func (o *WorkflowStepsDocumentStep) GetPayload() WorkflowStepPayload {
	if o == nil || IsNil(o.Payload) {
		var ret WorkflowStepPayload
		return ret
	}
	return *o.Payload
}

// GetPayloadOk returns a tuple with the Payload field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStepsDocumentStep) GetPayloadOk() (*WorkflowStepPayload, bool) {
	if o == nil || IsNil(o.Payload) {
		return nil, false
	}
	return o.Payload, true
}

// HasPayload returns a boolean if a field has been set.
func (o *WorkflowStepsDocumentStep) HasPayload() bool {
	if o != nil && !IsNil(o.Payload) {
		return true
	}

	return false
}

// SetPayload gets a reference to the given WorkflowStepPayload and assigns it to the Payload field.
func (o *WorkflowStepsDocumentStep) SetPayload(v WorkflowStepPayload) {
	o.Payload = &v
}

// GetPosition returns the Position field value if set, zero value otherwise.
func (o *WorkflowStepsDocumentStep) GetPosition() CanvasPosition {
	if o == nil || IsNil(o.Position) {
		var ret CanvasPosition
		return ret
	}
	return *o.Position
}

// GetPositionOk returns a tuple with the Position field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowStepsDocumentStep) GetPositionOk() (*CanvasPosition, bool) {
	if o == nil || IsNil(o.Position) {
		return nil, false
	}
	return o.Position, true
}

// HasPosition returns a boolean if a field has been set.
func (o *WorkflowStepsDocumentStep) HasPosition() bool {
	if o != nil && !IsNil(o.Position) {
		return true
	}

	return false
}

// SetPosition gets a reference to the given CanvasPosition and assigns it to the Position field.
func (o *WorkflowStepsDocumentStep) SetPosition(v CanvasPosition) {
	o.Position = &v
}

func (o WorkflowStepsDocumentStep) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WorkflowStepsDocumentStep) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AgentName) {
		toSerialize["agentName"] = o.AgentName
	}
	if !IsNil(o.AttributeCapture) {
		toSerialize["attributeCapture"] = o.AttributeCapture
	}
	if !IsNil(o.BackButton) {
		toSerialize["backButton"] = o.BackButton
	}
	toSerialize["eventAction"] = o.EventAction
	if !IsNil(o.Key) {
		toSerialize["key"] = o.Key
	}
	if !IsNil(o.Media) {
		toSerialize["media"] = o.Media
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Payload) {
		toSerialize["payload"] = o.Payload
	}
	if !IsNil(o.Position) {
		toSerialize["position"] = o.Position
	}
	return toSerialize, nil
}

func (o *WorkflowStepsDocumentStep) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"eventAction",
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

	varWorkflowStepsDocumentStep := _WorkflowStepsDocumentStep{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varWorkflowStepsDocumentStep)

	if err != nil {
		return err
	}

	*o = WorkflowStepsDocumentStep(varWorkflowStepsDocumentStep)

	return err
}

type NullableWorkflowStepsDocumentStep struct {
	value *WorkflowStepsDocumentStep
	isSet bool
}

func (v NullableWorkflowStepsDocumentStep) Get() *WorkflowStepsDocumentStep {
	return v.value
}

func (v *NullableWorkflowStepsDocumentStep) Set(val *WorkflowStepsDocumentStep) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowStepsDocumentStep) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowStepsDocumentStep) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowStepsDocumentStep(val *WorkflowStepsDocumentStep) *NullableWorkflowStepsDocumentStep {
	return &NullableWorkflowStepsDocumentStep{value: val, isSet: true}
}

func (v NullableWorkflowStepsDocumentStep) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowStepsDocumentStep) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


