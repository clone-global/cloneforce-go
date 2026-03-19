// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloneforce

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/stainless-sdks/cloneforce-go/internal/apijson"
	"github.com/stainless-sdks/cloneforce-go/internal/apiquery"
	"github.com/stainless-sdks/cloneforce-go/internal/requestconfig"
	"github.com/stainless-sdks/cloneforce-go/option"
	"github.com/stainless-sdks/cloneforce-go/packages/param"
	"github.com/stainless-sdks/cloneforce-go/packages/respjson"
)

// Clone integration management (Slack, Email, MS Teams, Phone)
//
// V1CloneIntegrationService contains methods and other services that help with
// interacting with the cloneforce API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1CloneIntegrationService] method instead.
type V1CloneIntegrationService struct {
	options []option.RequestOption
	// Clone integration management (Slack, Email, MS Teams, Phone)
	Slack V1CloneIntegrationSlackService
	// Clone integration management (Slack, Email, MS Teams, Phone)
	Msteams V1CloneIntegrationMsteamService
}

// NewV1CloneIntegrationService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1CloneIntegrationService(opts ...option.RequestOption) (r V1CloneIntegrationService) {
	r = V1CloneIntegrationService{}
	r.options = opts
	r.Slack = NewV1CloneIntegrationSlackService(opts...)
	r.Msteams = NewV1CloneIntegrationMsteamService(opts...)
	return
}

// Get an integration
func (r *V1CloneIntegrationService) Get(ctx context.Context, integrationID string, query V1CloneIntegrationGetParams, opts ...option.RequestOption) (res *IntegrationSummary, err error) {
	opts = slices.Concat(r.options, opts)
	if query.CloneID == "" {
		err = errors.New("missing required cloneId parameter")
		return nil, err
	}
	if integrationID == "" {
		err = errors.New("missing required integrationId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/clones/%s/integrations/%s", url.PathEscape(query.CloneID), url.PathEscape(integrationID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns all integrations for a clone (Slack, Email, MS Teams, Phone).
func (r *V1CloneIntegrationService) List(ctx context.Context, cloneID string, query V1CloneIntegrationListParams, opts ...option.RequestOption) (res *V1CloneIntegrationListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if cloneID == "" {
		err = errors.New("missing required cloneId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/clones/%s/integrations", url.PathEscape(cloneID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Deletes an integration and performs type-specific cleanup (e.g. Twilio release
// for phone).
func (r *V1CloneIntegrationService) Delete(ctx context.Context, integrationID string, body V1CloneIntegrationDeleteParams, opts ...option.RequestOption) (res *V1CloneIntegrationDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if body.CloneID == "" {
		err = errors.New("missing required cloneId parameter")
		return nil, err
	}
	if integrationID == "" {
		err = errors.New("missing required integrationId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/clones/%s/integrations/%s", url.PathEscape(body.CloneID), url.PathEscape(integrationID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Purchases a phone number and provisions it for clone voice calls. Requires
// sufficient account credits. Creates a Twilio number, ElevenLabs voice agent, and
// billing subscription.
func (r *V1CloneIntegrationService) Phone(ctx context.Context, cloneID string, body V1CloneIntegrationPhoneParams, opts ...option.RequestOption) (res *V1CloneIntegrationPhoneResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if cloneID == "" {
		err = errors.New("missing required cloneId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/clones/%s/integrations/phone", url.PathEscape(cloneID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns a browser URL for the OAuth-based setup flow. Supported types: `email`,
// `msteams`. Present this URL to the user and poll the integrations list to detect
// completion.
func (r *V1CloneIntegrationService) GetSetup(ctx context.Context, type_ V1CloneIntegrationGetSetupParamsType, query V1CloneIntegrationGetSetupParams, opts ...option.RequestOption) (res *V1CloneIntegrationGetSetupResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if query.CloneID == "" {
		err = errors.New("missing required cloneId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/clones/%s/integrations/%v/setup", url.PathEscape(query.CloneID), type_)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type IntegrationSummary struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	// Any of "Pending", "Connected", "Error", "Provisioning".
	Status IntegrationSummaryStatus `json:"status" api:"required"`
	// Any of "Slack", "Email", "MsTeams", "Phone".
	Type      IntegrationSummaryType `json:"type" api:"required"`
	UpdatedAt time.Time              `json:"updatedAt" api:"required" format:"date-time"`
	// Type-specific integration details
	Detail       IntegrationSummaryDetailUnion `json:"detail"`
	ErrorMessage string                        `json:"errorMessage"`
	Name         string                        `json:"name"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CreatedAt    respjson.Field
		Status       respjson.Field
		Type         respjson.Field
		UpdatedAt    respjson.Field
		Detail       respjson.Field
		ErrorMessage respjson.Field
		Name         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntegrationSummary) RawJSON() string { return r.JSON.raw }
func (r *IntegrationSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IntegrationSummaryStatus string

const (
	IntegrationSummaryStatusPending      IntegrationSummaryStatus = "Pending"
	IntegrationSummaryStatusConnected    IntegrationSummaryStatus = "Connected"
	IntegrationSummaryStatusError        IntegrationSummaryStatus = "Error"
	IntegrationSummaryStatusProvisioning IntegrationSummaryStatus = "Provisioning"
)

type IntegrationSummaryType string

const (
	IntegrationSummaryTypeSlack   IntegrationSummaryType = "Slack"
	IntegrationSummaryTypeEmail   IntegrationSummaryType = "Email"
	IntegrationSummaryTypeMsTeams IntegrationSummaryType = "MsTeams"
	IntegrationSummaryTypePhone   IntegrationSummaryType = "Phone"
)

// IntegrationSummaryDetailUnion contains all possible properties and values from
// [IntegrationSummaryDetailSlackDetail], [IntegrationSummaryDetailEmailDetail],
// [IntegrationSummaryDetailMsTeamsDetail], [IntegrationSummaryDetailPhoneDetail].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type IntegrationSummaryDetailUnion struct {
	// This field is from variant [IntegrationSummaryDetailSlackDetail].
	TeamID string `json:"teamId"`
	// This field is from variant [IntegrationSummaryDetailSlackDetail].
	TeamName string `json:"teamName"`
	// This field is from variant [IntegrationSummaryDetailEmailDetail].
	Email string `json:"email"`
	// This field is from variant [IntegrationSummaryDetailEmailDetail].
	ConnectionType string `json:"connectionType"`
	// This field is from variant [IntegrationSummaryDetailMsTeamsDetail].
	OrganizationName string `json:"organizationName"`
	// This field is from variant [IntegrationSummaryDetailMsTeamsDetail].
	Teams []MsTeamsTeamRef `json:"teams"`
	// This field is from variant [IntegrationSummaryDetailPhoneDetail].
	PhoneNumber string `json:"phoneNumber"`
	JSON        struct {
		TeamID           respjson.Field
		TeamName         respjson.Field
		Email            respjson.Field
		ConnectionType   respjson.Field
		OrganizationName respjson.Field
		Teams            respjson.Field
		PhoneNumber      respjson.Field
		raw              string
	} `json:"-"`
}

func (u IntegrationSummaryDetailUnion) AsIntegrationSummaryDetailSlackDetail() (v IntegrationSummaryDetailSlackDetail) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u IntegrationSummaryDetailUnion) AsIntegrationSummaryDetailEmailDetail() (v IntegrationSummaryDetailEmailDetail) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u IntegrationSummaryDetailUnion) AsIntegrationSummaryDetailMsTeamsDetail() (v IntegrationSummaryDetailMsTeamsDetail) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u IntegrationSummaryDetailUnion) AsIntegrationSummaryDetailPhoneDetail() (v IntegrationSummaryDetailPhoneDetail) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u IntegrationSummaryDetailUnion) RawJSON() string { return u.JSON.raw }

func (r *IntegrationSummaryDetailUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IntegrationSummaryDetailSlackDetail struct {
	TeamID   string `json:"teamId"`
	TeamName string `json:"teamName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TeamID      respjson.Field
		TeamName    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntegrationSummaryDetailSlackDetail) RawJSON() string { return r.JSON.raw }
func (r *IntegrationSummaryDetailSlackDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IntegrationSummaryDetailEmailDetail struct {
	Email          string `json:"email" api:"required"`
	ConnectionType string `json:"connectionType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Email          respjson.Field
		ConnectionType respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntegrationSummaryDetailEmailDetail) RawJSON() string { return r.JSON.raw }
func (r *IntegrationSummaryDetailEmailDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IntegrationSummaryDetailMsTeamsDetail struct {
	OrganizationName string           `json:"organizationName"`
	Teams            []MsTeamsTeamRef `json:"teams"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		OrganizationName respjson.Field
		Teams            respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntegrationSummaryDetailMsTeamsDetail) RawJSON() string { return r.JSON.raw }
func (r *IntegrationSummaryDetailMsTeamsDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type IntegrationSummaryDetailPhoneDetail struct {
	PhoneNumber string `json:"phoneNumber"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		PhoneNumber respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r IntegrationSummaryDetailPhoneDetail) RawJSON() string { return r.JSON.raw }
func (r *IntegrationSummaryDetailPhoneDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CloneIntegrationListResponse struct {
	Data []IntegrationSummary `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CloneIntegrationListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CloneIntegrationListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CloneIntegrationDeleteResponse struct {
	Deleted bool `json:"deleted" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Deleted     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CloneIntegrationDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CloneIntegrationDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CloneIntegrationPhoneResponse struct {
	ID          string `json:"id" api:"required"`
	PhoneNumber string `json:"phoneNumber" api:"required"`
	Status      string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		PhoneNumber respjson.Field
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CloneIntegrationPhoneResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CloneIntegrationPhoneResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CloneIntegrationGetSetupResponse struct {
	SetupURL string `json:"setupUrl" api:"required"`
	Type     string `json:"type" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SetupURL    respjson.Field
		Type        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CloneIntegrationGetSetupResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CloneIntegrationGetSetupResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CloneIntegrationGetParams struct {
	CloneID string `path:"cloneId" api:"required" json:"-"`
	paramObj
}

type V1CloneIntegrationListParams struct {
	// Filter by integration type
	//
	// Any of "Slack", "Email", "MsTeams", "Phone".
	Type V1CloneIntegrationListParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1CloneIntegrationListParams]'s query parameters as
// `url.Values`.
func (r V1CloneIntegrationListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by integration type
type V1CloneIntegrationListParamsType string

const (
	V1CloneIntegrationListParamsTypeSlack   V1CloneIntegrationListParamsType = "Slack"
	V1CloneIntegrationListParamsTypeEmail   V1CloneIntegrationListParamsType = "Email"
	V1CloneIntegrationListParamsTypeMsTeams V1CloneIntegrationListParamsType = "MsTeams"
	V1CloneIntegrationListParamsTypePhone   V1CloneIntegrationListParamsType = "Phone"
)

type V1CloneIntegrationDeleteParams struct {
	CloneID string `path:"cloneId" api:"required" json:"-"`
	paramObj
}

type V1CloneIntegrationPhoneParams struct {
	// Phone number to purchase (from the available numbers search)
	Phone string `json:"phone" api:"required"`
	paramObj
}

func (r V1CloneIntegrationPhoneParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CloneIntegrationPhoneParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CloneIntegrationPhoneParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CloneIntegrationGetSetupParams struct {
	CloneID string `path:"cloneId" api:"required" json:"-"`
	paramObj
}

type V1CloneIntegrationGetSetupParamsType string

const (
	V1CloneIntegrationGetSetupParamsTypeEmail   V1CloneIntegrationGetSetupParamsType = "email"
	V1CloneIntegrationGetSetupParamsTypeMsteams V1CloneIntegrationGetSetupParamsType = "msteams"
)
