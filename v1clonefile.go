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
	"github.com/clone-global/cloneforce-go/internal/requestconfig"
	"github.com/clone-global/cloneforce-go/option"
	"github.com/clone-global/cloneforce-go/packages/param"
	"github.com/clone-global/cloneforce-go/packages/respjson"
)

// Clone knowledge base file management
//
// V1CloneFileService contains methods and other services that help with
// interacting with the cloneforce API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1CloneFileService] method instead.
type V1CloneFileService struct {
	options []option.RequestOption
}

// NewV1CloneFileService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1CloneFileService(opts ...option.RequestOption) (r V1CloneFileService) {
	r = V1CloneFileService{}
	r.options = opts
	return
}

// **Not yet implemented.** Returns 501. Pending Pinecone migration.
func (r *V1CloneFileService) New(ctx context.Context, cloneID string, body V1CloneFileNewParams, opts ...option.RequestOption) (res *KBFileSummary, err error) {
	opts = slices.Concat(r.options, opts)
	if cloneID == "" {
		err = errors.New("missing required cloneId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/clones/%s/files", url.PathEscape(cloneID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get a knowledge base file
func (r *V1CloneFileService) Get(ctx context.Context, fileID string, query V1CloneFileGetParams, opts ...option.RequestOption) (res *KBFileSummary, err error) {
	opts = slices.Concat(r.options, opts)
	if query.CloneID == "" {
		err = errors.New("missing required cloneId parameter")
		return nil, err
	}
	if fileID == "" {
		err = errors.New("missing required fileId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/clones/%s/files/%s", url.PathEscape(query.CloneID), url.PathEscape(fileID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List knowledge base files
func (r *V1CloneFileService) List(ctx context.Context, cloneID string, opts ...option.RequestOption) (res *V1CloneFileListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if cloneID == "" {
		err = errors.New("missing required cloneId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/clones/%s/files", url.PathEscape(cloneID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// **Not yet implemented.** Returns 501. Pending Pinecone migration.
func (r *V1CloneFileService) Delete(ctx context.Context, fileID string, body V1CloneFileDeleteParams, opts ...option.RequestOption) (res *V1CloneFileDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if body.CloneID == "" {
		err = errors.New("missing required cloneId parameter")
		return nil, err
	}
	if fileID == "" {
		err = errors.New("missing required fileId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/clones/%s/files/%s", url.PathEscape(body.CloneID), url.PathEscape(fileID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type KBFileSummary struct {
	ID          string    `json:"id" api:"required"`
	ContentType string    `json:"contentType" api:"required"`
	CreatedAt   time.Time `json:"createdAt" api:"required" format:"date-time"`
	Name        string    `json:"name" api:"required"`
	URL         string    `json:"url" api:"required"`
	UploadURL   string    `json:"uploadUrl"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ContentType respjson.Field
		CreatedAt   respjson.Field
		Name        respjson.Field
		URL         respjson.Field
		UploadURL   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r KBFileSummary) RawJSON() string { return r.JSON.raw }
func (r *KBFileSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CloneFileListResponse struct {
	Data []KBFileSummary `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CloneFileListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CloneFileListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CloneFileDeleteResponse struct {
	Deleted bool `json:"deleted" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Deleted     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CloneFileDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CloneFileDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CloneFileNewParams struct {
	URL         string            `json:"url" api:"required"`
	ContentType param.Opt[string] `json:"contentType,omitzero"`
	Name        param.Opt[string] `json:"name,omitzero"`
	paramObj
}

func (r V1CloneFileNewParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CloneFileNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CloneFileNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CloneFileGetParams struct {
	CloneID string `path:"cloneId" api:"required" json:"-"`
	paramObj
}

type V1CloneFileDeleteParams struct {
	CloneID string `path:"cloneId" api:"required" json:"-"`
	paramObj
}
