// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloneforce

import (
	"context"
	"net/http"
	"slices"
	"time"

	"github.com/clone-global/cloneforce-go/internal/apijson"
	"github.com/clone-global/cloneforce-go/internal/requestconfig"
	"github.com/clone-global/cloneforce-go/option"
	"github.com/clone-global/cloneforce-go/packages/respjson"
)

// Clone profile management and asset generation
//
// V1CloneService contains methods and other services that help with interacting
// with the cloneforce API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1CloneService] method instead.
type V1CloneService struct {
	options []option.RequestOption
	// Clone profile management and asset generation
	Profile V1CloneProfileService
	// Clone profile management and asset generation
	Voice V1CloneVoiceService
	// Clone profile management and asset generation
	Headshot V1CloneHeadshotService
	// Skill marketplace search and clone skill management
	Skills V1CloneSkillService
	// Scheduled task management
	Tasks V1CloneTaskService
	// Clone knowledge base file management
	Files V1CloneFileService
	// Clone gallery media management
	Gallery V1CloneGalleryService
	// Clone integration management (Slack, Email, MS Teams, Phone)
	Integrations V1CloneIntegrationService
	// Task run history
	Activity V1CloneActivityService
}

// NewV1CloneService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewV1CloneService(opts ...option.RequestOption) (r V1CloneService) {
	r = V1CloneService{}
	r.options = opts
	r.Profile = NewV1CloneProfileService(opts...)
	r.Voice = NewV1CloneVoiceService(opts...)
	r.Headshot = NewV1CloneHeadshotService(opts...)
	r.Skills = NewV1CloneSkillService(opts...)
	r.Tasks = NewV1CloneTaskService(opts...)
	r.Files = NewV1CloneFileService(opts...)
	r.Gallery = NewV1CloneGalleryService(opts...)
	r.Integrations = NewV1CloneIntegrationService(opts...)
	r.Activity = NewV1CloneActivityService(opts...)
	return
}

// Returns all clones in the organization, ordered by creation date descending.
func (r *V1CloneService) List(ctx context.Context, opts ...option.RequestOption) (res *V1CloneListResponse, err error) {
	opts = slices.Concat(r.options, opts)
	path := "api/v1/clones"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

type V1CloneListResponse struct {
	Data []V1CloneListResponseData `json:"data" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Data        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CloneListResponse) RawJSON() string { return r.JSON.raw }
func (r *V1CloneListResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CloneListResponseData struct {
	ID         string        `json:"id" api:"required"`
	CreatedAt  time.Time     `json:"createdAt" api:"required" format:"date-time"`
	Generation string        `json:"generation" api:"required"`
	Name       string        `json:"name" api:"required"`
	ScreenName string        `json:"screenName" api:"required"`
	Status     string        `json:"status" api:"required"`
	Headshot   CloneHeadshot `json:"headshot"`
	IsEnabled  bool          `json:"isEnabled"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		Generation  respjson.Field
		Name        respjson.Field
		ScreenName  respjson.Field
		Status      respjson.Field
		Headshot    respjson.Field
		IsEnabled   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r V1CloneListResponseData) RawJSON() string { return r.JSON.raw }
func (r *V1CloneListResponseData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
