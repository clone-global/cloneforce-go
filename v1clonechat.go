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

// Chat sessions and completions
//
// V1CloneChatService contains methods and other services that help with
// interacting with the cloneforce API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1CloneChatService] method instead.
type V1CloneChatService struct {
	options []option.RequestOption
	// Chat sessions and completions
	Completions V1CloneChatCompletionService
}

// NewV1CloneChatService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewV1CloneChatService(opts ...option.RequestOption) (r V1CloneChatService) {
	r = V1CloneChatService{}
	r.options = opts
	r.Completions = NewV1CloneChatCompletionService(opts...)
	return
}

// Create a new chat session
func (r *V1CloneChatService) New(ctx context.Context, cloneID string, body V1CloneChatNewParams, opts ...option.RequestOption) (res *CreateChatResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if cloneID == "" {
		err = errors.New("missing required cloneId parameter")
		return nil, err
	}
	path := fmt.Sprintf("public/v1/clones/%s/chats", url.PathEscape(cloneID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

type ChatCompletionResponse struct {
	ID        string    `json:"id" api:"required"`
	ChatID    string    `json:"chatId" api:"required"`
	Content   string    `json:"content" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	Role      string    `json:"role" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		ChatID      respjson.Field
		Content     respjson.Field
		CreatedAt   respjson.Field
		Role        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionResponse) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type CreateChatResponse struct {
	ID        string    `json:"id" api:"required"`
	CloneID   string    `json:"cloneId" api:"required"`
	CreatedAt time.Time `json:"createdAt" api:"required" format:"date-time"`
	Title     string    `json:"title" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CloneID     respjson.Field
		CreatedAt   respjson.Field
		Title       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r CreateChatResponse) RawJSON() string { return r.JSON.raw }
func (r *CreateChatResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CloneChatNewParams struct {
	// Optional title for the chat session
	Title param.Opt[string] `json:"title,omitzero"`
	paramObj
}

func (r V1CloneChatNewParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CloneChatNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CloneChatNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
