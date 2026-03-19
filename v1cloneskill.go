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
	"github.com/stainless-sdks/cloneforce-go/internal/apiquery"
	"github.com/stainless-sdks/cloneforce-go/internal/requestconfig"
	"github.com/stainless-sdks/cloneforce-go/option"
	"github.com/stainless-sdks/cloneforce-go/packages/param"
	"github.com/stainless-sdks/cloneforce-go/packages/respjson"
)

// Skill marketplace search and clone skill management
//
// V1CloneSkillService contains methods and other services that help with
// interacting with the cloneforce API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1CloneSkillService] method instead.
type V1CloneSkillService struct {
	options []option.RequestOption
	// Skill marketplace search and clone skill management
	Connections V1CloneSkillConnectionService
}

// NewV1CloneSkillService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1CloneSkillService(opts ...option.RequestOption) (r V1CloneSkillService) {
	r = V1CloneSkillService{}
	r.options = opts
	r.Connections = NewV1CloneSkillConnectionService(opts...)
	return
}

// Attaches a marketplace skill to a clone. Returns 409 if already attached.
func (r *V1CloneSkillService) New(ctx context.Context, cloneID string, body V1CloneSkillNewParams, opts ...option.RequestOption) (res *SkillSummary, err error) {
	opts = slices.Concat(r.options, opts)
	if cloneID == "" {
		err = errors.New("missing required cloneId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/clones/%s/skills", url.PathEscape(cloneID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Updates settings and/or active status of a skill on a clone.
func (r *V1CloneSkillService) Update(ctx context.Context, skillName string, params V1CloneSkillUpdateParams, opts ...option.RequestOption) (res *SkillSummary, err error) {
	opts = slices.Concat(r.options, opts)
	if params.CloneID == "" {
		err = errors.New("missing required cloneId parameter")
		return nil, err
	}
	if skillName == "" {
		err = errors.New("missing required skillName parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/clones/%s/skills/%s", url.PathEscape(params.CloneID), url.PathEscape(skillName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Returns all skills attached to a clone, excluding hidden system skills.
func (r *V1CloneSkillService) List(ctx context.Context, cloneID string, query V1CloneSkillListParams, opts ...option.RequestOption) (res *V1CloneSkillListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if cloneID == "" {
		err = errors.New("missing required cloneId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/clones/%s/skills", url.PathEscape(cloneID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Detaches a skill from a clone. System skills cannot be removed.
func (r *V1CloneSkillService) Delete(ctx context.Context, skillName string, body V1CloneSkillDeleteParams, opts ...option.RequestOption) (res *V1CloneSkillDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if body.CloneID == "" {
		err = errors.New("missing required cloneId parameter")
		return nil, err
	}
	if skillName == "" {
		err = errors.New("missing required skillName parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/clones/%s/skills/%s", url.PathEscape(body.CloneID), url.PathEscape(skillName))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type SkillSummary struct {
	Accuracy      float64 `json:"accuracy" api:"required"`
	IsActive      bool    `json:"isActive" api:"required"`
	IsSystemSkill bool    `json:"isSystemSkill" api:"required"`
	Name          string  `json:"name" api:"required"`
	TotalRuns     int64   `json:"totalRuns" api:"required"`
	Category      string  `json:"category"`
	Description   string  `json:"description"`
	SkillID       string  `json:"skillId"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Accuracy      respjson.Field
		IsActive      respjson.Field
		IsSystemSkill respjson.Field
		Name          respjson.Field
		TotalRuns     respjson.Field
		Category      respjson.Field
		Description   respjson.Field
		SkillID       respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SkillSummary) RawJSON() string { return r.JSON.raw }
func (r *SkillSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CloneSkillListResponse struct {
	Data []SkillSummary `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CloneSkillListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CloneSkillListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CloneSkillDeleteResponse struct {
	Deleted bool `json:"deleted" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Deleted     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CloneSkillDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CloneSkillDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CloneSkillNewParams struct {
	SkillID  string          `json:"skillId" api:"required"`
	IsActive param.Opt[bool] `json:"isActive,omitzero"`
	paramObj
}

func (r V1CloneSkillNewParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CloneSkillNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CloneSkillNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CloneSkillUpdateParams struct {
	CloneID  string                            `path:"cloneId" api:"required" json:"-"`
	IsActive param.Opt[bool]                   `json:"isActive,omitzero"`
	Settings []V1CloneSkillUpdateParamsSetting `json:"settings,omitzero"`
	paramObj
}

func (r V1CloneSkillUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CloneSkillUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CloneSkillUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The property Name is required.
type V1CloneSkillUpdateParamsSetting struct {
	Name string `json:"name" api:"required"`
	// Required for connection-type settings
	ConnectionID param.Opt[string] `json:"connectionId,omitzero"`
	// Required for non-connection settings
	Value param.Opt[string] `json:"value,omitzero"`
	paramObj
}

func (r V1CloneSkillUpdateParamsSetting) MarshalJSON() (data []byte, err error) {
	type shadow V1CloneSkillUpdateParamsSetting
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CloneSkillUpdateParamsSetting) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CloneSkillListParams struct {
	// Filter by active status
	//
	// Any of "true", "false".
	Active V1CloneSkillListParamsActive `query:"active,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1CloneSkillListParams]'s query parameters as `url.Values`.
func (r V1CloneSkillListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by active status
type V1CloneSkillListParamsActive string

const (
	V1CloneSkillListParamsActiveTrue  V1CloneSkillListParamsActive = "true"
	V1CloneSkillListParamsActiveFalse V1CloneSkillListParamsActive = "false"
)

type V1CloneSkillDeleteParams struct {
	CloneID string `path:"cloneId" api:"required" json:"-"`
	paramObj
}
