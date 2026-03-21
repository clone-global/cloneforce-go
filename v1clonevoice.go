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

	"github.com/clone-global/cloneforce-go/internal/apijson"
	shimjson "github.com/clone-global/cloneforce-go/internal/encoding/json"
	"github.com/clone-global/cloneforce-go/internal/requestconfig"
	"github.com/clone-global/cloneforce-go/option"
	"github.com/clone-global/cloneforce-go/packages/param"
	"github.com/clone-global/cloneforce-go/packages/respjson"
)

// Clone profile management and asset generation
//
// V1CloneVoiceService contains methods and other services that help with
// interacting with the cloneforce API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1CloneVoiceService] method instead.
type V1CloneVoiceService struct {
	options []option.RequestOption
}

// NewV1CloneVoiceService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1CloneVoiceService(opts ...option.RequestOption) (r V1CloneVoiceService) {
	r = V1CloneVoiceService{}
	r.options = opts
	return
}

// Triggers asynchronous voice regeneration for a clone. Returns immediately with
// 202 Accepted. Poll `GET /clones/{cloneId}/profile` and check `state.voice` for
// progress.
//
// An optional request body can include `additionalInstructions` to nudge the voice
// style (e.g. "deeper voice", "more energetic"). An empty body performs a standard
// regeneration.
func (r *V1CloneVoiceService) Generate(ctx context.Context, cloneID string, body V1CloneVoiceGenerateParams, opts ...option.RequestOption) (res *GenerationStatus, err error) {
	opts = slices.Concat(r.options, opts)
	if cloneID == "" {
		err = errors.New("missing required cloneId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/clones/%s/voice/generate", url.PathEscape(cloneID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type GenerateRequestParam struct {
	// Optional instructions to nudge the generation (e.g. "deeper voice", "longer
	// hair")
	AdditionalInstructions param.Opt[string] `json:"additionalInstructions,omitzero"`
	paramObj
}

func (r GenerateRequestParam) MarshalJSON() (data []byte, err error) {
	type shadow GenerateRequestParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *GenerateRequestParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type GenerationStatus struct {
	Status string `json:"status" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r GenerationStatus) RawJSON() string { return r.JSON.raw }
func (r *GenerationStatus) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CloneVoiceGenerateParams struct {
	GenerateRequest GenerateRequestParam
	paramObj
}

func (r V1CloneVoiceGenerateParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.GenerateRequest)
}
func (r *V1CloneVoiceGenerateParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.GenerateRequest)
}
