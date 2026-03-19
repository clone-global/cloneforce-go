// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloneforce

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/clone-global/cloneforce-go/internal/apijson"
	"github.com/clone-global/cloneforce-go/internal/requestconfig"
	"github.com/clone-global/cloneforce-go/option"
	"github.com/clone-global/cloneforce-go/packages/param"
	"github.com/clone-global/cloneforce-go/packages/respjson"
)

// Clone integration management (Slack, Email, MS Teams, Phone)
//
// V1CloneIntegrationMsteamService contains methods and other services that help
// with interacting with the cloneforce API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1CloneIntegrationMsteamService] method instead.
type V1CloneIntegrationMsteamService struct {
	options []option.RequestOption
}

// NewV1CloneIntegrationMsteamService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewV1CloneIntegrationMsteamService(opts ...option.RequestOption) (r V1CloneIntegrationMsteamService) {
	r = V1CloneIntegrationMsteamService{}
	r.options = opts
	return
}

// Adds a team to an existing MS Teams integration. Validates team access via MS
// Graph.
func (r *V1CloneIntegrationMsteamService) Teams(ctx context.Context, integrationID string, params V1CloneIntegrationMsteamTeamsParams, opts ...option.RequestOption) (res *V1CloneIntegrationMsteamTeamsResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if params.CloneID == "" {
		err = errors.New("missing required cloneId parameter")
		return nil, err
	}
	if integrationID == "" {
		err = errors.New("missing required integrationId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/clones/%s/integrations/msteams/%s/teams", url.PathEscape(params.CloneID), url.PathEscape(integrationID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

type MsTeamsTeamRef struct {
	ID   string `json:"id" api:"required"`
	Name string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MsTeamsTeamRef) RawJSON() string { return r.JSON.raw }
func (r *MsTeamsTeamRef) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CloneIntegrationMsteamTeamsResponse struct {
	ID               string           `json:"id" api:"required"`
	Status           string           `json:"status" api:"required"`
	OrganizationName string           `json:"organizationName"`
	Teams            []MsTeamsTeamRef `json:"teams"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		Status           respjson.Field
		OrganizationName respjson.Field
		Teams            respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CloneIntegrationMsteamTeamsResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CloneIntegrationMsteamTeamsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CloneIntegrationMsteamTeamsParams struct {
	CloneID string `path:"cloneId" api:"required" json:"-"`
	TeamID  string `json:"teamId" api:"required"`
	paramObj
}

func (r V1CloneIntegrationMsteamTeamsParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CloneIntegrationMsteamTeamsParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CloneIntegrationMsteamTeamsParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
