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

// Scheduled task management
//
// V1CloneTaskService contains methods and other services that help with
// interacting with the cloneforce API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1CloneTaskService] method instead.
type V1CloneTaskService struct {
	options []option.RequestOption
}

// NewV1CloneTaskService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1CloneTaskService(opts ...option.RequestOption) (r V1CloneTaskService) {
	r = V1CloneTaskService{}
	r.options = opts
	return
}

// Creates a new scheduled task for a clone.
func (r *V1CloneTaskService) New(ctx context.Context, cloneID string, body V1CloneTaskNewParams, opts ...option.RequestOption) (res *TaskSummary, err error) {
	opts = slices.Concat(r.options, opts)
	if cloneID == "" {
		err = errors.New("missing required cloneId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/clones/%s/tasks", url.PathEscape(cloneID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get a scheduled task
func (r *V1CloneTaskService) Get(ctx context.Context, taskID string, query V1CloneTaskGetParams, opts ...option.RequestOption) (res *TaskSummary, err error) {
	opts = slices.Concat(r.options, opts)
	if query.CloneID == "" {
		err = errors.New("missing required cloneId parameter")
		return nil, err
	}
	if taskID == "" {
		err = errors.New("missing required taskId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/clones/%s/tasks/%s", url.PathEscape(query.CloneID), url.PathEscape(taskID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates one or more fields on a scheduled task.
func (r *V1CloneTaskService) Update(ctx context.Context, taskID string, params V1CloneTaskUpdateParams, opts ...option.RequestOption) (res *TaskSummary, err error) {
	opts = slices.Concat(r.options, opts)
	if params.CloneID == "" {
		err = errors.New("missing required cloneId parameter")
		return nil, err
	}
	if taskID == "" {
		err = errors.New("missing required taskId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/clones/%s/tasks/%s", url.PathEscape(params.CloneID), url.PathEscape(taskID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, params, &res, opts...)
	return res, err
}

// Returns all scheduled tasks for a clone.
func (r *V1CloneTaskService) List(ctx context.Context, cloneID string, query V1CloneTaskListParams, opts ...option.RequestOption) (res *V1CloneTaskListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if cloneID == "" {
		err = errors.New("missing required cloneId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/clones/%s/tasks", url.PathEscape(cloneID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete a scheduled task
func (r *V1CloneTaskService) Delete(ctx context.Context, taskID string, body V1CloneTaskDeleteParams, opts ...option.RequestOption) (res *V1CloneTaskDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if body.CloneID == "" {
		err = errors.New("missing required cloneId parameter")
		return nil, err
	}
	if taskID == "" {
		err = errors.New("missing required taskId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/clones/%s/tasks/%s", url.PathEscape(body.CloneID), url.PathEscape(taskID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type TaskRecurrence struct {
	Interval int64 `json:"interval" api:"required"`
	// Any of "Minutely", "Hourly", "Daily", "Weekly", "Monthly", "Yearly".
	Pattern TaskRecurrencePattern `json:"pattern" api:"required"`
	EndsAt  time.Time             `json:"endsAt" format:"date-time"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Interval    respjson.Field
		Pattern     respjson.Field
		EndsAt      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskRecurrence) RawJSON() string { return r.JSON.raw }
func (r *TaskRecurrence) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this TaskRecurrence to a TaskRecurrenceParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// TaskRecurrenceParam.Overrides()
func (r TaskRecurrence) ToParam() TaskRecurrenceParam {
	return param.Override[TaskRecurrenceParam](json.RawMessage(r.RawJSON()))
}

type TaskRecurrencePattern string

const (
	TaskRecurrencePatternMinutely TaskRecurrencePattern = "Minutely"
	TaskRecurrencePatternHourly   TaskRecurrencePattern = "Hourly"
	TaskRecurrencePatternDaily    TaskRecurrencePattern = "Daily"
	TaskRecurrencePatternWeekly   TaskRecurrencePattern = "Weekly"
	TaskRecurrencePatternMonthly  TaskRecurrencePattern = "Monthly"
	TaskRecurrencePatternYearly   TaskRecurrencePattern = "Yearly"
)

// The properties Interval, Pattern are required.
type TaskRecurrenceParam struct {
	Interval int64 `json:"interval" api:"required"`
	// Any of "Minutely", "Hourly", "Daily", "Weekly", "Monthly", "Yearly".
	Pattern TaskRecurrencePattern `json:"pattern,omitzero" api:"required"`
	EndsAt  param.Opt[time.Time]  `json:"endsAt,omitzero" format:"date-time"`
	paramObj
}

func (r TaskRecurrenceParam) MarshalJSON() (data []byte, err error) {
	type shadow TaskRecurrenceParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TaskRecurrenceParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TaskSummary struct {
	ID          string         `json:"id" api:"required"`
	CreatedAt   time.Time      `json:"createdAt" api:"required" format:"date-time"`
	IsRecurring bool           `json:"isRecurring" api:"required"`
	Prompt      string         `json:"prompt" api:"required"`
	StartsAt    time.Time      `json:"startsAt" api:"required" format:"date-time"`
	Title       string         `json:"title" api:"required"`
	UpdatedAt   time.Time      `json:"updatedAt" api:"required" format:"date-time"`
	Color       string         `json:"color"`
	LastRanAt   time.Time      `json:"lastRanAt" format:"date-time"`
	Recurrence  TaskRecurrence `json:"recurrence"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		IsRecurring respjson.Field
		Prompt      respjson.Field
		StartsAt    respjson.Field
		Title       respjson.Field
		UpdatedAt   respjson.Field
		Color       respjson.Field
		LastRanAt   respjson.Field
		Recurrence  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r TaskSummary) RawJSON() string { return r.JSON.raw }
func (r *TaskSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CloneTaskListResponse struct {
	Data []TaskSummary `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CloneTaskListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CloneTaskListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CloneTaskDeleteResponse struct {
	Deleted bool `json:"deleted" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Deleted     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CloneTaskDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CloneTaskDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CloneTaskNewParams struct {
	Prompt      string              `json:"prompt" api:"required"`
	StartsAt    time.Time           `json:"startsAt" api:"required" format:"date-time"`
	Title       string              `json:"title" api:"required"`
	Color       param.Opt[string]   `json:"color,omitzero"`
	IsRecurring param.Opt[bool]     `json:"isRecurring,omitzero"`
	Recurrence  TaskRecurrenceParam `json:"recurrence,omitzero"`
	paramObj
}

func (r V1CloneTaskNewParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CloneTaskNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CloneTaskNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CloneTaskGetParams struct {
	CloneID string `path:"cloneId" api:"required" json:"-"`
	paramObj
}

type V1CloneTaskUpdateParams struct {
	CloneID     string               `path:"cloneId" api:"required" json:"-"`
	Color       param.Opt[string]    `json:"color,omitzero"`
	IsRecurring param.Opt[bool]      `json:"isRecurring,omitzero"`
	Prompt      param.Opt[string]    `json:"prompt,omitzero"`
	StartsAt    param.Opt[time.Time] `json:"startsAt,omitzero" format:"date-time"`
	Title       param.Opt[string]    `json:"title,omitzero"`
	Recurrence  TaskRecurrenceParam  `json:"recurrence,omitzero"`
	paramObj
}

func (r V1CloneTaskUpdateParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CloneTaskUpdateParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CloneTaskUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CloneTaskListParams struct {
	// Filter by recurring status
	//
	// Any of "true", "false".
	IsRecurring V1CloneTaskListParamsIsRecurring `query:"isRecurring,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1CloneTaskListParams]'s query parameters as `url.Values`.
func (r V1CloneTaskListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by recurring status
type V1CloneTaskListParamsIsRecurring string

const (
	V1CloneTaskListParamsIsRecurringTrue  V1CloneTaskListParamsIsRecurring = "true"
	V1CloneTaskListParamsIsRecurringFalse V1CloneTaskListParamsIsRecurring = "false"
)

type V1CloneTaskDeleteParams struct {
	CloneID string `path:"cloneId" api:"required" json:"-"`
	paramObj
}
