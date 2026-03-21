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
	"github.com/clone-global/cloneforce-go/internal/apiquery"
	"github.com/clone-global/cloneforce-go/internal/requestconfig"
	"github.com/clone-global/cloneforce-go/option"
	"github.com/clone-global/cloneforce-go/packages/param"
	"github.com/clone-global/cloneforce-go/packages/respjson"
)

// Skill marketplace search and clone skill management
//
// V1SkillService contains methods and other services that help with interacting
// with the cloneforce API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1SkillService] method instead.
type V1SkillService struct {
	options []option.RequestOption
}

// NewV1SkillService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV1SkillService(opts ...option.RequestOption) (r V1SkillService) {
	r = V1SkillService{}
	r.options = opts
	return
}

// Returns full skill details including setting definitions. If `cloneId` is
// provided, also returns attachment and configuration status.
func (r *V1SkillService) Get(ctx context.Context, skillID string, query V1SkillGetParams, opts ...option.RequestOption) (res *V1SkillGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if skillID == "" {
		err = errors.New("missing required skillId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/skills/%s", url.PathEscape(skillID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Searches the skill marketplace. Optionally marks results already attached to a
// clone.
func (r *V1SkillService) Search(ctx context.Context, query V1SkillSearchParams, opts ...option.RequestOption) (res *V1SkillSearchResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/skills/search"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type V1SkillGetResponse struct {
	IsAttached            bool                                   `json:"isAttached" api:"required"`
	LogicType             string                                 `json:"logicType" api:"required"`
	Name                  string                                 `json:"name" api:"required"`
	RequiresConnections   bool                                   `json:"requiresConnections" api:"required"`
	SkillID               string                                 `json:"skillId" api:"required"`
	Category              string                                 `json:"category"`
	ConnectionsConfigured bool                                   `json:"connectionsConfigured"`
	Description           string                                 `json:"description"`
	IsActive              bool                                   `json:"isActive"`
	SettingDefinitions    []V1SkillGetResponseSettingDefinition  `json:"settingDefinitions"`
	SettingsConfigured    []V1SkillGetResponseSettingsConfigured `json:"settingsConfigured"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsAttached            respjson.Field
		LogicType             respjson.Field
		Name                  respjson.Field
		RequiresConnections   respjson.Field
		SkillID               respjson.Field
		Category              respjson.Field
		ConnectionsConfigured respjson.Field
		Description           respjson.Field
		IsActive              respjson.Field
		SettingDefinitions    respjson.Field
		SettingsConfigured    respjson.Field
		ExtraFields           map[string]respjson.Field
		raw                   string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SkillGetResponse) RawJSON() string { return r.JSON.raw }
func (r *V1SkillGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SkillGetResponseSettingDefinition struct {
	IsConnection bool   `json:"isConnection" api:"required"`
	IsRequired   bool   `json:"isRequired" api:"required"`
	Name         string `json:"name" api:"required"`
	Description  string `json:"description"`
	SettingType  string `json:"settingType"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsConnection respjson.Field
		IsRequired   respjson.Field
		Name         respjson.Field
		Description  respjson.Field
		SettingType  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SkillGetResponseSettingDefinition) RawJSON() string { return r.JSON.raw }
func (r *V1SkillGetResponseSettingDefinition) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SkillGetResponseSettingsConfigured struct {
	IsConfigured bool   `json:"isConfigured" api:"required"`
	IsConnection bool   `json:"isConnection" api:"required"`
	Name         string `json:"name" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsConfigured respjson.Field
		IsConnection respjson.Field
		Name         respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SkillGetResponseSettingsConfigured) RawJSON() string { return r.JSON.raw }
func (r *V1SkillGetResponseSettingsConfigured) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SkillSearchResponse struct {
	Data []V1SkillSearchResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SkillSearchResponse) RawJSON() string { return r.JSON.raw }
func (r *V1SkillSearchResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SkillSearchResponseData struct {
	IsAlreadyAttached bool    `json:"isAlreadyAttached" api:"required"`
	Name              string  `json:"name" api:"required"`
	SkillID           string  `json:"skillId" api:"required"`
	Description       string  `json:"description"`
	Score             float64 `json:"score"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsAlreadyAttached respjson.Field
		Name              respjson.Field
		SkillID           respjson.Field
		Description       respjson.Field
		Score             respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1SkillSearchResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1SkillSearchResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1SkillGetParams struct {
	// Clone ID to check attachment and configuration status
	CloneID param.Opt[string] `query:"cloneId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1SkillGetParams]'s query parameters as `url.Values`.
func (r V1SkillGetParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1SkillSearchParams struct {
	// Search query
	Q string `query:"q" api:"required" json:"-"`
	// Clone ID to check attachment status against
	CloneID param.Opt[string] `query:"cloneId,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1SkillSearchParams]'s query parameters as `url.Values`.
func (r V1SkillSearchParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
