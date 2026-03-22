// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloneforce

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/clone-global/cloneforce-go/internal/apijson"
	"github.com/clone-global/cloneforce-go/internal/apiquery"
	"github.com/clone-global/cloneforce-go/internal/requestconfig"
	"github.com/clone-global/cloneforce-go/option"
	"github.com/clone-global/cloneforce-go/packages/param"
	"github.com/clone-global/cloneforce-go/packages/respjson"
)

// Organization-level connection credential management
//
// V1ConnectionService contains methods and other services that help with
// interacting with the cloneforce API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1ConnectionService] method instead.
type V1ConnectionService struct {
	options []option.RequestOption
}

// NewV1ConnectionService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1ConnectionService(opts ...option.RequestOption) (r V1ConnectionService) {
	r = V1ConnectionService{}
	r.options = opts
	return
}

// Creates a new key-value connection credential.
func (r *V1ConnectionService) New(ctx context.Context, body V1ConnectionNewParams, opts ...option.RequestOption) (res *ConnectionDetail, err error) {
	opts = slices.Concat(r.options, opts)
	path := "public/v1/connections"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get a connection
func (r *V1ConnectionService) Get(ctx context.Context, connectionID string, opts ...option.RequestOption) (res *ConnectionDetail, err error) {
	opts = slices.Concat(r.options, opts)
	if connectionID == "" {
		err = errors.New("missing required connectionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("public/v1/connections/%s", url.PathEscape(connectionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates name, key, and/or value on a standard connection. OAuth connections
// cannot be updated here.
func (r *V1ConnectionService) Update(ctx context.Context, connectionID string, body V1ConnectionUpdateParams, opts ...option.RequestOption) (res *ConnectionDetail, err error) {
	opts = slices.Concat(r.options, opts)
	if connectionID == "" {
		err = errors.New("missing required connectionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("public/v1/connections/%s", url.PathEscape(connectionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

// Returns all connections in the organization.
func (r *V1ConnectionService) List(ctx context.Context, query V1ConnectionListParams, opts ...option.RequestOption) (res *V1ConnectionListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "public/v1/connections"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete a connection
func (r *V1ConnectionService) Delete(ctx context.Context, connectionID string, opts ...option.RequestOption) (res *V1ConnectionDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if connectionID == "" {
		err = errors.New("missing required connectionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("public/v1/connections/%s", url.PathEscape(connectionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

// Creates a pending OAuth connection and returns a provision URL. Present the
// provision URL to the user to complete the OAuth consent flow.
func (r *V1ConnectionService) NewOAuth(ctx context.Context, body V1ConnectionNewOAuthParams, opts ...option.RequestOption) (res *OAuthProvision, err error) {
	opts = slices.Concat(r.options, opts)
	path := "public/v1/connections/oauth"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns lightweight validity and expiration status.
func (r *V1ConnectionService) GetStatus(ctx context.Context, connectionID string, opts ...option.RequestOption) (res *ConnectionStatus, err error) {
	opts = slices.Concat(r.options, opts)
	if connectionID == "" {
		err = errors.New("missing required connectionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("public/v1/connections/%s/status", url.PathEscape(connectionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Refreshes OAuth tokens for a connection. Only supported for refreshable
// connection types.
func (r *V1ConnectionService) Refresh(ctx context.Context, connectionID string, opts ...option.RequestOption) (res *ConnectionStatus, err error) {
	opts = slices.Concat(r.options, opts)
	if connectionID == "" {
		err = errors.New("missing required connectionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("public/v1/connections/%s/refresh", url.PathEscape(connectionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// Returns a new provision URL for an existing OAuth connection to re-authorize.
func (r *V1ConnectionService) ReprovisionOAuth(ctx context.Context, connectionID string, opts ...option.RequestOption) (res *OAuthProvision, err error) {
	opts = slices.Concat(r.options, opts)
	if connectionID == "" {
		err = errors.New("missing required connectionId parameter")
		return nil, err
	}
	path := fmt.Sprintf("public/v1/connections/%s/reprovision", url.PathEscape(connectionID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

type ConnectionDetail struct {
	ID                  string            `json:"id" api:"required"`
	CreatedAt           time.Time         `json:"createdAt" api:"required" format:"date-time"`
	IsExpired           bool              `json:"isExpired" api:"required"`
	IsValid             bool              `json:"isValid" api:"required"`
	Name                string            `json:"name" api:"required"`
	SettingType         string            `json:"settingType" api:"required"`
	UpdatedAt           time.Time         `json:"updatedAt" api:"required" format:"date-time"`
	AcquiredPermissions []string          `json:"acquiredPermissions"`
	ErrorReason         string            `json:"errorReason"`
	ExpiresAt           time.Time         `json:"expiresAt" format:"date-time"`
	Metadata            map[string]string `json:"metadata"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		CreatedAt           respjson.Field
		IsExpired           respjson.Field
		IsValid             respjson.Field
		Name                respjson.Field
		SettingType         respjson.Field
		UpdatedAt           respjson.Field
		AcquiredPermissions respjson.Field
		ErrorReason         respjson.Field
		ExpiresAt           respjson.Field
		Metadata            respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectionDetail) RawJSON() string { return r.JSON.raw }
func (r *ConnectionDetail) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ConnectionStatus struct {
	ID          string    `json:"id" api:"required"`
	IsExpired   bool      `json:"isExpired" api:"required"`
	IsValid     bool      `json:"isValid" api:"required"`
	ErrorReason string    `json:"errorReason"`
	ExpiresAt   time.Time `json:"expiresAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		IsExpired   respjson.Field
		IsValid     respjson.Field
		ErrorReason respjson.Field
		ExpiresAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ConnectionStatus) RawJSON() string { return r.JSON.raw }
func (r *ConnectionStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type OAuthProvision struct {
	ConnectionID string `json:"connectionId" api:"required"`
	// URL to present to the user to complete the OAuth consent flow
	ProvisionURL string `json:"provisionUrl" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ConnectionID respjson.Field
		ProvisionURL respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r OAuthProvision) RawJSON() string { return r.JSON.raw }
func (r *OAuthProvision) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ConnectionListResponse struct {
	Data []V1ConnectionListResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ConnectionListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ConnectionListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ConnectionListResponseData struct {
	ID          string    `json:"id" api:"required"`
	CreatedAt   time.Time `json:"createdAt" api:"required" format:"date-time"`
	IsExpired   bool      `json:"isExpired" api:"required"`
	IsValid     bool      `json:"isValid" api:"required"`
	Name        string    `json:"name" api:"required"`
	SettingType string    `json:"settingType" api:"required"`
	UpdatedAt   time.Time `json:"updatedAt" api:"required" format:"date-time"`
	ErrorReason string    `json:"errorReason"`
	ExpiresAt   time.Time `json:"expiresAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		IsExpired   respjson.Field
		IsValid     respjson.Field
		Name        respjson.Field
		SettingType respjson.Field
		UpdatedAt   respjson.Field
		ErrorReason respjson.Field
		ExpiresAt   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ConnectionListResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1ConnectionListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ConnectionDeleteResponse struct {
	Deleted bool `json:"deleted" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Deleted     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1ConnectionDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *V1ConnectionDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ConnectionNewParams struct {
	Key   string `json:"key" api:"required"`
	Name  string `json:"name" api:"required"`
	Value string `json:"value" api:"required"`
	paramObj
}

func (r V1ConnectionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow V1ConnectionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ConnectionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ConnectionUpdateParams struct {
	Key   param.Opt[string] `json:"key,omitzero"`
	Name  param.Opt[string] `json:"name,omitzero"`
	Value param.Opt[string] `json:"value,omitzero"`
	paramObj
}

func (r V1ConnectionUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow V1ConnectionUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ConnectionUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1ConnectionListParams struct {
	// Filter by setting type (e.g. Standard, or an OAuth provider type)
	Type param.Opt[string] `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1ConnectionListParams]'s query parameters as `url.Values`.
func (r V1ConnectionListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type V1ConnectionNewOAuthParams struct {
	Name string `json:"name" api:"required"`
	// Must be an OAuth provider type
	SettingType string   `json:"settingType" api:"required"`
	Scopes      []string `json:"scopes,omitzero"`
	paramObj
}

func (r V1ConnectionNewOAuthParams) MarshalJSON() (data []byte, err error) {
	type shadow V1ConnectionNewOAuthParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1ConnectionNewOAuthParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
