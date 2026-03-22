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
// V1CloneIntegrationSlackService contains methods and other services that help
// with interacting with the cloneforce API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1CloneIntegrationSlackService] method instead.
type V1CloneIntegrationSlackService struct {
	options []option.RequestOption
}

// NewV1CloneIntegrationSlackService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1CloneIntegrationSlackService(opts ...option.RequestOption) (r V1CloneIntegrationSlackService) {
	r = V1CloneIntegrationSlackService{}
	r.options = opts
	return
}

// Creates a pending Slack integration and returns a Slack app manifest for
// installation.
func (r *V1CloneIntegrationSlackService) New(ctx context.Context, cloneID string, opts ...option.RequestOption) (res *SlackIntegration, err error) {
	opts = slices.Concat(r.options, opts)
	if cloneID == "" {
		err = errors.New("missing required cloneId parameter")
		return nil, err
	}
	path := fmt.Sprintf("public/v1/clones/%s/integrations/slack", url.PathEscape(cloneID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Configures bot token and signing secret. Tests the connection if the bot token
// changes.
func (r *V1CloneIntegrationSlackService) Update(ctx context.Context, integrationID string, params V1CloneIntegrationSlackUpdateParams, opts ...option.RequestOption) (res *SlackIntegration, err error) {
	opts = slices.Concat(r.options, opts)
	if params.CloneID == "" {
		err = errors.New("missing required cloneId parameter")
		return nil, err
	}
	if integrationID == "" {
		err = errors.New("missing required integrationId parameter")
		return nil, err
	}
	path := fmt.Sprintf("public/v1/clones/%s/integrations/slack/%s", url.PathEscape(params.CloneID), url.PathEscape(integrationID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

type SlackIntegration struct {
	ID       string `json:"id" api:"required"`
	Status   string `json:"status" api:"required"`
	Manifest string `json:"manifest"`
	TeamID   string `json:"teamId"`
	TeamName string `json:"teamName"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Status      respjson.Field
		Manifest    respjson.Field
		TeamID      respjson.Field
		TeamName    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SlackIntegration) RawJSON() string { return r.JSON.raw }
func (r *SlackIntegration) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CloneIntegrationSlackUpdateParams struct {
	CloneID       string            `path:"cloneId" api:"required" json:"-"`
	BotToken      param.Opt[string] `json:"botToken,omitzero"`
	SigningSecret param.Opt[string] `json:"signingSecret,omitzero"`
	paramObj
}

func (r V1CloneIntegrationSlackUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CloneIntegrationSlackUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CloneIntegrationSlackUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
