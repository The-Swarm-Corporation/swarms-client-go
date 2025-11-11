// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package swarms

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"

	"github.com/The-Swarm-Corporation/swarms-client-go/internal/apijson"
	"github.com/The-Swarm-Corporation/swarms-client-go/internal/requestconfig"
	"github.com/The-Swarm-Corporation/swarms-client-go/option"
	"github.com/The-Swarm-Corporation/swarms-client-go/packages/respjson"
)

// ClientMarketplaceService contains methods and other services that help with
// interacting with the swarms-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClientMarketplaceService] method instead.
type ClientMarketplaceService struct {
	Options []option.RequestOption
}

// NewClientMarketplaceService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewClientMarketplaceService(opts ...option.RequestOption) (r ClientMarketplaceService) {
	r = ClientMarketplaceService{}
	r.Options = opts
	return
}

// Retrieve free agents from the marketplace.
func (r *ClientMarketplaceService) ListAgents(ctx context.Context, query ClientMarketplaceListAgentsParams, opts ...option.RequestOption) (res *ClientMarketplaceListAgentsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/marketplace/agents"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return
}

// Response schema for marketplace prompts endpoint.
type ClientMarketplaceListAgentsResponse struct {
	// List of marketplace prompts
	Prompts []ClientMarketplaceListAgentsResponsePrompt `json:"prompts,required"`
	// Total number of prompts available
	TotalCount int64 `json:"total_count,required"`
	// The status of the marketplace prompts response.
	Status string `json:"status,nullable"`
	// The timestamp of the marketplace prompts response.
	Timestamp string `json:"timestamp,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Prompts     respjson.Field
		TotalCount  respjson.Field
		Status      respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClientMarketplaceListAgentsResponse) RawJSON() string { return r.JSON.raw }
func (r *ClientMarketplaceListAgentsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Schema for marketplace prompts from the swarms_cloud_prompts table.
type ClientMarketplaceListAgentsResponsePrompt struct {
	// Unique identifier for the prompt
	ID string `json:"id,required"`
	// Timestamp when the prompt was created
	CreatedAt string `json:"created_at,required"`
	// ID of the user who created the prompt
	UserID string `json:"user_id,required"`
	// Category name(s) - can be string or list
	Category ClientMarketplaceListAgentsResponsePromptCategoryUnion `json:"category,nullable"`
	// Description of the prompt
	Description string `json:"description,nullable"`
	// Associated links - can be list of dicts or strings
	Links ClientMarketplaceListAgentsResponsePromptLinksUnion `json:"links,nullable"`
	// Name of the prompt
	Name string `json:"name,nullable"`
	// The actual prompt text
	Prompt string `json:"prompt,nullable"`
	// Status of the prompt
	Status string `json:"status,nullable"`
	// Tags associated with the prompt
	Tags string `json:"tags,nullable"`
	// Use cases - can be dict or list of dicts
	UseCases ClientMarketplaceListAgentsResponsePromptUseCasesUnion `json:"use_cases,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID          respjson.Field
		CreatedAt   respjson.Field
		UserID      respjson.Field
		Category    respjson.Field
		Description respjson.Field
		Links       respjson.Field
		Name        respjson.Field
		Prompt      respjson.Field
		Status      respjson.Field
		Tags        respjson.Field
		UseCases    respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClientMarketplaceListAgentsResponsePrompt) RawJSON() string { return r.JSON.raw }
func (r *ClientMarketplaceListAgentsResponsePrompt) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ClientMarketplaceListAgentsResponsePromptCategoryUnion contains all possible
// properties and values from [string], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfString OfStringArray]
type ClientMarketplaceListAgentsResponsePromptCategoryUnion struct {
	// This field will be present if the value is a [string] instead of an object.
	OfString string `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfString      respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

func (u ClientMarketplaceListAgentsResponsePromptCategoryUnion) AsString() (v string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ClientMarketplaceListAgentsResponsePromptCategoryUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ClientMarketplaceListAgentsResponsePromptCategoryUnion) RawJSON() string { return u.JSON.raw }

func (r *ClientMarketplaceListAgentsResponsePromptCategoryUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ClientMarketplaceListAgentsResponsePromptLinksUnion contains all possible
// properties and values from [[]map[string]any], [[]string].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfMapOfAnyMap OfStringArray]
type ClientMarketplaceListAgentsResponsePromptLinksUnion struct {
	// This field will be present if the value is a [[]map[string]any] instead of an
	// object.
	OfMapOfAnyMap []map[string]any `json:",inline"`
	// This field will be present if the value is a [[]string] instead of an object.
	OfStringArray []string `json:",inline"`
	JSON          struct {
		OfMapOfAnyMap respjson.Field
		OfStringArray respjson.Field
		raw           string
	} `json:"-"`
}

func (u ClientMarketplaceListAgentsResponsePromptLinksUnion) AsMapOfAnyMap() (v []map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ClientMarketplaceListAgentsResponsePromptLinksUnion) AsStringArray() (v []string) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ClientMarketplaceListAgentsResponsePromptLinksUnion) RawJSON() string { return u.JSON.raw }

func (r *ClientMarketplaceListAgentsResponsePromptLinksUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ClientMarketplaceListAgentsResponsePromptUseCasesUnion contains all possible
// properties and values from [map[string]any], [[]map[string]any].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfClientMarketplaceListAgentsResponsePromptUseCasesMapItem
// OfMapOfAnyMap]
type ClientMarketplaceListAgentsResponsePromptUseCasesUnion struct {
	// This field will be present if the value is a [any] instead of an object.
	OfClientMarketplaceListAgentsResponsePromptUseCasesMapItem any `json:",inline"`
	// This field will be present if the value is a [[]map[string]any] instead of an
	// object.
	OfMapOfAnyMap []map[string]any `json:",inline"`
	JSON          struct {
		OfClientMarketplaceListAgentsResponsePromptUseCasesMapItem respjson.Field
		OfMapOfAnyMap                                              respjson.Field
		raw                                                        string
	} `json:"-"`
}

func (u ClientMarketplaceListAgentsResponsePromptUseCasesUnion) AsAnyMap() (v map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u ClientMarketplaceListAgentsResponsePromptUseCasesUnion) AsMapOfAnyMap() (v []map[string]any) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u ClientMarketplaceListAgentsResponsePromptUseCasesUnion) RawJSON() string { return u.JSON.raw }

func (r *ClientMarketplaceListAgentsResponsePromptUseCasesUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClientMarketplaceListAgentsParams struct {
	paramObj
}
