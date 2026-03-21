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

// Clone profile management and asset generation
//
// V1CloneProfileService contains methods and other services that help with
// interacting with the cloneforce API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1CloneProfileService] method instead.
type V1CloneProfileService struct {
	options []option.RequestOption
}

// NewV1CloneProfileService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1CloneProfileService(opts ...option.RequestOption) (r V1CloneProfileService) {
	r = V1CloneProfileService{}
	r.options = opts
	return
}

// Returns the full profile and aesthetics for a clone.
func (r *V1CloneProfileService) List(ctx context.Context, cloneID string, opts ...option.RequestOption) (res *CloneProfile, err error) {
	opts = slices.Concat(r.options, opts)
	if cloneID == "" {
		err = errors.New("missing required cloneId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/clones/%s/profile", url.PathEscape(cloneID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// Updates one or more fields on a clone's profile. Only provided fields are
// changed.
func (r *V1CloneProfileService) PatchAll(ctx context.Context, cloneID string, body V1CloneProfilePatchAllParams, opts ...option.RequestOption) (res *CloneProfile, err error) {
	opts = slices.Concat(r.options, opts)
	if cloneID == "" {
		err = errors.New("missing required cloneId parameter")
		return nil, err
	}
	path := fmt.Sprintf("api/v1/clones/%s/profile", url.PathEscape(cloneID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return res, err
}

type CloneHeadshot struct {
	Large  string `json:"large"`
	Medium string `json:"medium"`
	Small  string `json:"small"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Large       respjson.Field
		Medium      respjson.Field
		Small       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CloneHeadshot) RawJSON() string { return r.JSON.raw }
func (r *CloneHeadshot) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CloneProfile struct {
	ID                 string            `json:"id" api:"required"`
	CreatedAt          time.Time         `json:"createdAt" api:"required" format:"date-time"`
	Generation         string            `json:"generation" api:"required"`
	Name               string            `json:"name" api:"required"`
	ScreenName         string            `json:"screenName" api:"required"`
	Status             string            `json:"status" api:"required"`
	UpdatedAt          time.Time         `json:"updatedAt" api:"required" format:"date-time"`
	AppearanceDesc     string            `json:"appearanceDesc"`
	Birthdate          time.Time         `json:"birthdate" format:"date-time"`
	CareerDesc         string            `json:"careerDesc"`
	ChildhoodDesc      string            `json:"childhoodDesc"`
	CurrentCity        string            `json:"currentCity"`
	EducationDesc      string            `json:"educationDesc"`
	FamilyDesc         string            `json:"familyDesc"`
	FinancesDesc       string            `json:"financesDesc"`
	Gender             string            `json:"gender"`
	Headshot           CloneHeadshot     `json:"headshot"`
	Hometown           string            `json:"hometown"`
	InterestsDesc      string            `json:"interestsDesc"`
	IsEnabled          bool              `json:"isEnabled"`
	Language           string            `json:"language"`
	LifestyleDesc      string            `json:"lifestyleDesc"`
	MaritalStatus      string            `json:"maritalStatus"`
	Nationality        string            `json:"nationality"`
	Occupation         string            `json:"occupation"`
	PrinciplesDesc     string            `json:"principlesDesc"`
	Races              []string          `json:"races"`
	SocialCircleDesc   string            `json:"socialCircleDesc"`
	SocialHobbiesDesc  string            `json:"socialHobbiesDesc"`
	SportsDesc         string            `json:"sportsDesc"`
	State              CloneProfileState `json:"state"`
	TravelDesc         string            `json:"travelDesc"`
	UnusualHobbiesDesc string            `json:"unusualHobbiesDesc"`
	VoiceURL           string            `json:"voiceUrl"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                 respjson.Field
		CreatedAt          respjson.Field
		Generation         respjson.Field
		Name               respjson.Field
		ScreenName         respjson.Field
		Status             respjson.Field
		UpdatedAt          respjson.Field
		AppearanceDesc     respjson.Field
		Birthdate          respjson.Field
		CareerDesc         respjson.Field
		ChildhoodDesc      respjson.Field
		CurrentCity        respjson.Field
		EducationDesc      respjson.Field
		FamilyDesc         respjson.Field
		FinancesDesc       respjson.Field
		Gender             respjson.Field
		Headshot           respjson.Field
		Hometown           respjson.Field
		InterestsDesc      respjson.Field
		IsEnabled          respjson.Field
		Language           respjson.Field
		LifestyleDesc      respjson.Field
		MaritalStatus      respjson.Field
		Nationality        respjson.Field
		Occupation         respjson.Field
		PrinciplesDesc     respjson.Field
		Races              respjson.Field
		SocialCircleDesc   respjson.Field
		SocialHobbiesDesc  respjson.Field
		SportsDesc         respjson.Field
		State              respjson.Field
		TravelDesc         respjson.Field
		UnusualHobbiesDesc respjson.Field
		VoiceURL           respjson.Field
		ExtraFields        map[string]respjson.Field
		raw                string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CloneProfile) RawJSON() string { return r.JSON.raw }
func (r *CloneProfile) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CloneProfileState struct {
	// Generation status: none, generating, ready, error
	//
	// Any of "none", "generating", "ready", "error".
	Headshot string `json:"headshot" api:"required"`
	// Generation status: none, generating, ready, error
	//
	// Any of "none", "generating", "ready", "error".
	Voice string `json:"voice" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Headshot    respjson.Field
		Voice       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CloneProfileState) RawJSON() string { return r.JSON.raw }
func (r *CloneProfileState) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CloneProfilePatchAllParams struct {
	AppearanceDesc     param.Opt[string]    `json:"appearanceDesc,omitzero"`
	Birthdate          param.Opt[time.Time] `json:"birthdate,omitzero" format:"date-time"`
	CareerDesc         param.Opt[string]    `json:"careerDesc,omitzero"`
	ChildhoodDesc      param.Opt[string]    `json:"childhoodDesc,omitzero"`
	CurrentCity        param.Opt[string]    `json:"currentCity,omitzero"`
	EducationDesc      param.Opt[string]    `json:"educationDesc,omitzero"`
	FamilyDesc         param.Opt[string]    `json:"familyDesc,omitzero"`
	FinancesDesc       param.Opt[string]    `json:"financesDesc,omitzero"`
	Gender             param.Opt[string]    `json:"gender,omitzero"`
	Hometown           param.Opt[string]    `json:"hometown,omitzero"`
	InterestsDesc      param.Opt[string]    `json:"interestsDesc,omitzero"`
	IsEnabled          param.Opt[bool]      `json:"isEnabled,omitzero"`
	Language           param.Opt[string]    `json:"language,omitzero"`
	LifestyleDesc      param.Opt[string]    `json:"lifestyleDesc,omitzero"`
	MaritalStatus      param.Opt[string]    `json:"maritalStatus,omitzero"`
	Name               param.Opt[string]    `json:"name,omitzero"`
	Nationality        param.Opt[string]    `json:"nationality,omitzero"`
	Occupation         param.Opt[string]    `json:"occupation,omitzero"`
	PrinciplesDesc     param.Opt[string]    `json:"principlesDesc,omitzero"`
	ScreenName         param.Opt[string]    `json:"screenName,omitzero"`
	SocialCircleDesc   param.Opt[string]    `json:"socialCircleDesc,omitzero"`
	SocialHobbiesDesc  param.Opt[string]    `json:"socialHobbiesDesc,omitzero"`
	SportsDesc         param.Opt[string]    `json:"sportsDesc,omitzero"`
	TravelDesc         param.Opt[string]    `json:"travelDesc,omitzero"`
	UnusualHobbiesDesc param.Opt[string]    `json:"unusualHobbiesDesc,omitzero"`
	Races              []string             `json:"races,omitzero"`
	paramObj
}

func (r V1CloneProfilePatchAllParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CloneProfilePatchAllParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CloneProfilePatchAllParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
