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

	"github.com/stainless-sdks/cloneforce-go/internal/apijson"
	"github.com/stainless-sdks/cloneforce-go/internal/requestconfig"
	"github.com/stainless-sdks/cloneforce-go/option"
	"github.com/stainless-sdks/cloneforce-go/packages/respjson"
)

// Task run history
//
// V1CloneActivityService contains methods and other services that help with
// interacting with the cloneforce API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1CloneActivityService] method instead.
type V1CloneActivityService struct {
	options []option.RequestOption
}

// NewV1CloneActivityService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1CloneActivityService(opts ...option.RequestOption) (r V1CloneActivityService) {
	r = V1CloneActivityService{}
	r.options = opts
	return
}

// Returns a single task run with skill execution details.
func (r *V1CloneActivityService) Get(ctx context.Context, activityID string, query V1CloneActivityGetParams, opts ...option.RequestOption) (res *V1CloneActivityGetResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if query.CloneID == "" {
		err = errors.New("missing required cloneId parameter")
		return nil, err
	}
	if activityID == "" {
		err = errors.New("missing required activityId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/clones/%s/activity/%s", url.PathEscape(query.CloneID), url.PathEscape(activityID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Returns all task run records for a clone, ordered by creation date descending.
func (r *V1CloneActivityService) List(ctx context.Context, cloneID string, opts ...option.RequestOption) (res *V1CloneActivityListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if cloneID == "" {
		err = errors.New("missing required cloneId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/clones/%s/activity", url.PathEscape(cloneID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Delete a task run
func (r *V1CloneActivityService) Delete(ctx context.Context, activityID string, body V1CloneActivityDeleteParams, opts ...option.RequestOption) (res *V1CloneActivityDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if body.CloneID == "" {
		err = errors.New("missing required cloneId parameter")
		return nil, err
	}
	if activityID == "" {
		err = errors.New("missing required activityId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/clones/%s/activity/%s", url.PathEscape(body.CloneID), url.PathEscape(activityID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type V1CloneActivityGetResponse struct {
	ID         string                            `json:"id" api:"required"`
	CreatedAt  time.Time                         `json:"createdAt" api:"required" format:"date-time"`
	IsSuccess  bool                              `json:"isSuccess" api:"required"`
	SkillCount int64                             `json:"skillCount" api:"required"`
	TaskID     string                            `json:"taskId" api:"required"`
	Response   string                            `json:"response"`
	Skills     []V1CloneActivityGetResponseSkill `json:"skills"`
	TaskTitle  string                            `json:"taskTitle"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		IsSuccess   respjson.Field
		SkillCount  respjson.Field
		TaskID      respjson.Field
		Response    respjson.Field
		Skills      respjson.Field
		TaskTitle   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CloneActivityGetResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CloneActivityGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CloneActivityGetResponseSkill struct {
	IsSuccess bool   `json:"isSuccess" api:"required"`
	SkillName string `json:"skillName" api:"required"`
	Result    string `json:"result"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		IsSuccess   respjson.Field
		SkillName   respjson.Field
		Result      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CloneActivityGetResponseSkill) RawJSON() string { return r.JSON.raw }
func (r *V1CloneActivityGetResponseSkill) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CloneActivityListResponse struct {
	Data []V1CloneActivityListResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CloneActivityListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CloneActivityListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CloneActivityListResponseData struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	IsSuccess bool      `json:"isSuccess" api:"required"`
	TaskID    string    `json:"taskId" api:"required"`
	Response  string    `json:"response"`
	TaskTitle string    `json:"taskTitle"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		IsSuccess   respjson.Field
		TaskID      respjson.Field
		Response    respjson.Field
		TaskTitle   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CloneActivityListResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1CloneActivityListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CloneActivityDeleteResponse struct {
	Deleted bool `json:"deleted" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Deleted     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CloneActivityDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CloneActivityDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CloneActivityGetParams struct {
	CloneID string `path:"cloneId" api:"required" json:"-"`
	paramObj
}

type V1CloneActivityDeleteParams struct {
	CloneID string `path:"cloneId" api:"required" json:"-"`
	paramObj
}
