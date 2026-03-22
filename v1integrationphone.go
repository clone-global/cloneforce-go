// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloneforce

import (
	"context"
	"net/http"
	"net/url"
	"slices"

	"github.com/clone-global/cloneforce-go/internal/apijson"
	"github.com/clone-global/cloneforce-go/internal/apiquery"
	"github.com/clone-global/cloneforce-go/internal/requestconfig"
	"github.com/clone-global/cloneforce-go/option"
	"github.com/clone-global/cloneforce-go/packages/param"
	"github.com/clone-global/cloneforce-go/packages/respjson"
)

// Clone integration management (Slack, Email, MS Teams, Phone)
//
// V1IntegrationPhoneService contains methods and other services that help with
// interacting with the cloneforce API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1IntegrationPhoneService] method instead.
type V1IntegrationPhoneService struct {
	options []option.RequestOption
}

// NewV1IntegrationPhoneService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1IntegrationPhoneService(opts ...option.RequestOption) (r V1IntegrationPhoneService) {
	r = V1IntegrationPhoneService{}
	r.options = opts
	return
}

// Searches for available phone numbers via Twilio that can be purchased.
func (r *V1IntegrationPhoneService) ListAvailable(ctx context.Context, query V1IntegrationPhoneListAvailableParams, opts ...option.RequestOption) (res *V1IntegrationPhoneListAvailableResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "public/v1/integrations/phone/available"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type V1IntegrationPhoneListAvailableResponse struct {
	Data []V1IntegrationPhoneListAvailableResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1IntegrationPhoneListAvailableResponse) RawJSON() string { return r.JSON.raw }
func (r *V1IntegrationPhoneListAvailableResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1IntegrationPhoneListAvailableResponseData struct {
	FriendlyName string `json:"friendlyName" api:"required"`
	Phone        string `json:"phone" api:"required"`
	Country      string `json:"country"`
	Locality     string `json:"locality"`
	PostalCode   string `json:"postalCode"`
	Region       string `json:"region"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		FriendlyName respjson.Field
		Phone        respjson.Field
		Country      respjson.Field
		Locality     respjson.Field
		PostalCode   respjson.Field
		Region       respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1IntegrationPhoneListAvailableResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1IntegrationPhoneListAvailableResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1IntegrationPhoneListAvailableParams struct {
	// Area code filter
	AreaCode param.Opt[int64] `query:"areaCode,omitzero" json:"-"`
	// Country code (default US)
	Country param.Opt[string] `query:"country,omitzero" json:"-"`
	// Max results (default 20)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1IntegrationPhoneListAvailableParams]'s query parameters
// as `url.Values`.
func (r V1IntegrationPhoneListAvailableParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
