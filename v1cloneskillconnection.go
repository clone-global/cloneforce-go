// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloneforce

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/cloneforce-go/internal/apijson"
	"github.com/stainless-sdks/cloneforce-go/internal/requestconfig"
	"github.com/stainless-sdks/cloneforce-go/option"
	"github.com/stainless-sdks/cloneforce-go/packages/param"
	"github.com/stainless-sdks/cloneforce-go/packages/respjson"
)

// Skill marketplace search and clone skill management
//
// V1CloneSkillConnectionService contains methods and other services that help with
// interacting with the cloneforce API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1CloneSkillConnectionService] method instead.
type V1CloneSkillConnectionService struct {
	options []option.RequestOption
}

// NewV1CloneSkillConnectionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1CloneSkillConnectionService(opts ...option.RequestOption) (r V1CloneSkillConnectionService) {
	r = V1CloneSkillConnectionService{}
	r.options = opts
	return
}

// Assigns a connection to a connection-type setting on a clone skill.
func (r *V1CloneSkillConnectionService) Update(ctx context.Context, settingName string, params V1CloneSkillConnectionUpdateParams, opts ...option.RequestOption) (res *SkillConnectionInfo, err error) {
	opts = slices.Concat(r.options, opts)
	if params.CloneID == "" {
		err = errors.New("missing required cloneId parameter")
		return nil, err
	}
	if params.SkillName == "" {
		err = errors.New("missing required skillName parameter")
		return nil, err
	}
	if settingName == "" {
		err = errors.New("missing required settingName parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/clones/%s/skills/%s/connections/%s", url.PathEscape(params.CloneID), url.PathEscape(params.SkillName), url.PathEscape(settingName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPut, path, params, &res, opts...)
	return res, err
}

// Returns connection-type settings for a skill on a clone, with their
// configuration status.
func (r *V1CloneSkillConnectionService) List(ctx context.Context, skillName string, query V1CloneSkillConnectionListParams, opts ...option.RequestOption) (res *V1CloneSkillConnectionListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if query.CloneID == "" {
		err = errors.New("missing required cloneId parameter")
		return nil, err
	}
	if skillName == "" {
		err = errors.New("missing required skillName parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/clones/%s/skills/%s/connections", url.PathEscape(query.CloneID), url.PathEscape(skillName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type SkillConnectionInfo struct {
	IsConfigured bool   `json:"isConfigured" api:"required"`
	SettingName  string `json:"settingName" api:"required"`
	SettingType  string `json:"settingType" api:"required"`
	ConnectionID string `json:"connectionId"`
	IsValid      bool   `json:"isValid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsConfigured respjson.Field
		SettingName  respjson.Field
		SettingType  respjson.Field
		ConnectionID respjson.Field
		IsValid      respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SkillConnectionInfo) RawJSON() string { return r.JSON.raw }
func (r *SkillConnectionInfo) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CloneSkillConnectionListResponse struct {
	Data []SkillConnectionInfo `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CloneSkillConnectionListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CloneSkillConnectionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CloneSkillConnectionUpdateParams struct {
	CloneID      string `path:"cloneId" api:"required" json:"-"`
	SkillName    string `path:"skillName" api:"required" json:"-"`
	ConnectionID string `json:"connectionId" api:"required"`
	paramObj
}

func (r V1CloneSkillConnectionUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CloneSkillConnectionUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CloneSkillConnectionUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CloneSkillConnectionListParams struct {
	CloneID string `path:"cloneId" api:"required" json:"-"`
	paramObj
}
