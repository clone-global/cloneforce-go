// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package cloneforce

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/clone-global/cloneforce-go/internal/apijson"
	"github.com/clone-global/cloneforce-go/internal/requestconfig"
	"github.com/clone-global/cloneforce-go/option"
	"github.com/clone-global/cloneforce-go/packages/param"
	"github.com/clone-global/cloneforce-go/packages/respjson"
	"github.com/clone-global/cloneforce-go/packages/ssestream"
)

// Chat sessions and completions
//
// V1CloneChatCompletionService contains methods and other services that help with
// interacting with the cloneforce API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewV1CloneChatCompletionService] method instead.
type V1CloneChatCompletionService struct {
	options []option.RequestOption
}

// NewV1CloneChatCompletionService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewV1CloneChatCompletionService(opts ...option.RequestOption) (r V1CloneChatCompletionService) {
	r = V1CloneChatCompletionService{}
	r.options = opts
	return
}

// Sends a user message to the clone and returns the assistant's response.
//
// Set `stream: true` to receive the response as Server-Sent Events (SSE). SSE
// events: `message.delta` (incremental text), `message.completed` (final message),
// `done` (stream end).
func (r *V1CloneChatCompletionService) New(ctx context.Context, chatID string, params V1CloneChatCompletionNewParams, opts ...option.RequestOption) (res *ChatCompletionResponse, err error) {
	opts = slices.Concat(r.options, opts)
	if params.CloneID == "" {
		err = errors.New("missing required cloneId parameter")
		return nil, err
	}
	if chatID == "" {
		err = errors.New("missing required chatId parameter")
		return nil, err
	}
	path := fmt.Sprintf("public/v1/clones/%s/chats/%s/completions", url.PathEscape(params.CloneID), url.PathEscape(chatID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return res, err
}

// Sends a user message to the clone and returns the assistant's response.
//
// Set `stream: true` to receive the response as Server-Sent Events (SSE). SSE
// events: `message.delta` (incremental text), `message.completed` (final message),
// `done` (stream end).
func (r *V1CloneChatCompletionService) NewStreaming(ctx context.Context, chatID string, params V1CloneChatCompletionNewParams, opts ...option.RequestOption) (stream *ssestream.Stream[ChatCompletionChunk]) {
	var (
		raw *http.Response
		err error
	)
	opts = slices.Concat(r.options, opts)
	opts = append(opts, option.WithJSONSet("stream", true))
	if params.CloneID == "" {
		err = errors.New("missing required cloneId parameter")
		return ssestream.NewStream[ChatCompletionChunk](nil, err)
	}
	if chatID == "" {
		err = errors.New("missing required chatId parameter")
		return ssestream.NewStream[ChatCompletionChunk](nil, err)
	}
	path := fmt.Sprintf("public/v1/clones/%s/chats/%s/completions", url.PathEscape(params.CloneID), url.PathEscape(chatID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &raw, opts...)
	return ssestream.NewStream[ChatCompletionChunk](ssestream.NewDecoder(raw), err)
}

// SSE message.delta event payload
type ChatCompletionChunk struct {
	ID    string                   `json:"id" api:"required"`
	Delta ChatCompletionChunkDelta `json:"delta"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		Delta       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionChunk) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionChunk) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ChatCompletionChunkDelta struct {
	Content string `json:"content"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Content     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ChatCompletionChunkDelta) RawJSON() string { return r.JSON.raw }
func (r *ChatCompletionChunkDelta) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type V1CloneChatCompletionNewParams struct {
	CloneID string `path:"cloneId" api:"required" json:"-"`
	// The user message to send to the clone
	Message string `json:"message" api:"required"`
	paramObj
}

func (r V1CloneChatCompletionNewParams) MarshalJSON() (data []byte, err error) {
	type shadow V1CloneChatCompletionNewParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *V1CloneChatCompletionNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
