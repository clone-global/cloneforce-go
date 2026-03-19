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

	shimjson "github.com/stainless-sdks/cloneforce-go/internal/encoding/json"
	"github.com/stainless-sdks/cloneforce-go/internal/requestconfig"
	"github.com/stainless-sdks/cloneforce-go/option"
)

// Clone profile management and asset generation
//
// V1CloneHeadshotService contains methods and other services that help with
// interacting with the cloneforce API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1CloneHeadshotService] method instead.
type V1CloneHeadshotService struct {
	options []option.RequestOption
}

// NewV1CloneHeadshotService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1CloneHeadshotService(opts ...option.RequestOption) (r V1CloneHeadshotService) {
	r = V1CloneHeadshotService{}
	r.options = opts
	return
}

// Triggers asynchronous headshot regeneration for a clone. Returns immediately
// with 202 Accepted. Poll `GET /clones/{cloneId}/profile` and check
// `state.headshot` for progress.
//
// An optional request body can include `additionalInstructions` to nudge the
// appearance (e.g. "longer hair", "wearing glasses"). An empty body performs a
// standard regeneration.
func (r *V1CloneHeadshotService) Generate(ctx context.Context, cloneID string, body V1CloneHeadshotGenerateParams, opts ...option.RequestOption) (res *GenerationStatus, err error) {
	opts = slices.Concat(r.options, opts)
	if cloneID == "" {
		err = errors.New("missing required cloneId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/clones/%s/headshot/generate", url.PathEscape(cloneID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type V1CloneHeadshotGenerateParams struct {
	GenerateRequest GenerateRequestParam
	paramObj
}

func (r V1CloneHeadshotGenerateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.GenerateRequest)
}
func (r *V1CloneHeadshotGenerateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.GenerateRequest)
}
