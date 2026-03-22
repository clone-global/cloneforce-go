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

// Clone gallery media management
//
// V1CloneGalleryService contains methods and other services that help with
// interacting with the cloneforce API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1CloneGalleryService] method instead.
type V1CloneGalleryService struct {
	options []option.RequestOption
}

// NewV1CloneGalleryService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1CloneGalleryService(opts ...option.RequestOption) (r V1CloneGalleryService) {
	r = V1CloneGalleryService{}
	r.options = opts
	return
}

// Downloads media from the provided URL and adds it to the clone's gallery.
func (r *V1CloneGalleryService) New(ctx context.Context, cloneID string, body V1CloneGalleryNewParams, opts ...option.RequestOption) (res *GalleryItemSummary, err error) {
	opts = slices.Concat(r.options, opts)
	if cloneID == "" {
		err = errors.New("missing required cloneId parameter")
		return nil, err
	}
	path := fmt.Sprintf("public/v1/clones/%s/gallery", url.PathEscape(cloneID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Get a gallery item
func (r *V1CloneGalleryService) Get(ctx context.Context, itemID string, query V1CloneGalleryGetParams, opts ...option.RequestOption) (res *GalleryItemSummary, err error) {
	opts = slices.Concat(r.options, opts)
	if query.CloneID == "" {
		err = errors.New("missing required cloneId parameter")
		return nil, err
	}
	if itemID == "" {
		err = errors.New("missing required itemId parameter")
		return nil, err
	}
	path := fmt.Sprintf("public/v1/clones/%s/gallery/%s", url.PathEscape(query.CloneID), url.PathEscape(itemID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// List gallery items
func (r *V1CloneGalleryService) List(ctx context.Context, cloneID string, query V1CloneGalleryListParams, opts ...option.RequestOption) (res *V1CloneGalleryListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if cloneID == "" {
		err = errors.New("missing required cloneId parameter")
		return nil, err
	}
	path := fmt.Sprintf("public/v1/clones/%s/gallery", url.PathEscape(cloneID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

// Delete a gallery item
func (r *V1CloneGalleryService) Delete(ctx context.Context, itemID string, body V1CloneGalleryDeleteParams, opts ...option.RequestOption) (res *V1CloneGalleryDeleteResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if body.CloneID == "" {
		err = errors.New("missing required cloneId parameter")
		return nil, err
	}
	if itemID == "" {
		err = errors.New("missing required itemId parameter")
		return nil, err
	}
	path := fmt.Sprintf("public/v1/clones/%s/gallery/%s", url.PathEscape(body.CloneID), url.PathEscape(itemID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, &res, opts...)
	return res, err
}

type GalleryItemSummary struct {
	ID        string    `json:"id" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	Name      string    `json:"name" api:"required"`
	// Any of "Video", "Audio", "Image".
	Type         GalleryItemSummaryType `json:"type" api:"required"`
	UpdatedAt    time.Time              `json:"updatedAt" api:"required" format:"date-time"`
	URL          string                 `json:"url" api:"required"`
	Description  string                 `json:"description"`
	IsHeroVideo  bool                   `json:"isHeroVideo"`
	ThumbnailURL string                 `json:"thumbnailUrl"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID           respjson.Field
		CreatedAt    respjson.Field
		Name         respjson.Field
		Type         respjson.Field
		UpdatedAt    respjson.Field
		URL          respjson.Field
		Description  respjson.Field
		IsHeroVideo  respjson.Field
		ThumbnailURL respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GalleryItemSummary) RawJSON() string { return r.JSON.raw }
func (r *GalleryItemSummary) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GalleryItemSummaryType string

const (
	GalleryItemSummaryTypeVideo GalleryItemSummaryType = "Video"
	GalleryItemSummaryTypeAudio GalleryItemSummaryType = "Audio"
	GalleryItemSummaryTypeImage GalleryItemSummaryType = "Image"
)

type V1CloneGalleryListResponse struct {
	Data []GalleryItemSummary `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CloneGalleryListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CloneGalleryListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CloneGalleryDeleteResponse struct {
	Deleted bool `json:"deleted" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Deleted     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CloneGalleryDeleteResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CloneGalleryDeleteResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CloneGalleryNewParams struct {
	// URL of the media file to download and add to the gallery
	MediaURL string `json:"mediaUrl" api:"required"`
	// Any of "Video", "Audio", "Image".
	Type        V1CloneGalleryNewParamsType `json:"type,omitzero" api:"required"`
	Description param.Opt[string]           `json:"description,omitzero"`
	IsHeroVideo param.Opt[bool]             `json:"isHeroVideo,omitzero"`
	paramObj
}

func (r V1CloneGalleryNewParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CloneGalleryNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CloneGalleryNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CloneGalleryNewParamsType string

const (
	V1CloneGalleryNewParamsTypeVideo V1CloneGalleryNewParamsType = "Video"
	V1CloneGalleryNewParamsTypeAudio V1CloneGalleryNewParamsType = "Audio"
	V1CloneGalleryNewParamsTypeImage V1CloneGalleryNewParamsType = "Image"
)

type V1CloneGalleryGetParams struct {
	CloneID string `path:"cloneId" api:"required" json:"-"`
	paramObj
}

type V1CloneGalleryListParams struct {
	// Filter by media type
	//
	// Any of "Video", "Audio", "Image".
	Type V1CloneGalleryListParamsType `query:"type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [V1CloneGalleryListParams]'s query parameters as
// `url.Values`.
func (r V1CloneGalleryListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Filter by media type
type V1CloneGalleryListParamsType string

const (
	V1CloneGalleryListParamsTypeVideo V1CloneGalleryListParamsType = "Video"
	V1CloneGalleryListParamsTypeAudio V1CloneGalleryListParamsType = "Audio"
	V1CloneGalleryListParamsTypeImage V1CloneGalleryListParamsType = "Image"
)

type V1CloneGalleryDeleteParams struct {
	CloneID string `path:"cloneId" api:"required" json:"-"`
	paramObj
}
