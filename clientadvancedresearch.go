// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package swarms

import (
	"context"
	"net/http"
	"slices"

	"github.com/The-Swarm-Corporation/swarms-client-go/internal/apijson"
	"github.com/The-Swarm-Corporation/swarms-client-go/internal/requestconfig"
	"github.com/The-Swarm-Corporation/swarms-client-go/option"
	"github.com/The-Swarm-Corporation/swarms-client-go/packages/param"
	"github.com/The-Swarm-Corporation/swarms-client-go/packages/respjson"
)

// ClientAdvancedResearchService contains methods and other services that help with
// interacting with the swarms-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClientAdvancedResearchService] method instead.
type ClientAdvancedResearchService struct {
	Options []option.RequestOption
	Batch   ClientAdvancedResearchBatchService
}

// NewClientAdvancedResearchService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewClientAdvancedResearchService(opts ...option.RequestOption) (r ClientAdvancedResearchService) {
	r = ClientAdvancedResearchService{}
	r.Options = opts
	r.Batch = NewClientAdvancedResearchBatchService(opts...)
	return
}

// Execute comprehensive research sessions with multi-source data collection,
// analysis, and synthesis capabilities.
func (r *ClientAdvancedResearchService) NewCompletion(ctx context.Context, body ClientAdvancedResearchNewCompletionParams, opts ...option.RequestOption) (res *ClientAdvancedResearchNewCompletionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/advanced-research/completions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ClientAdvancedResearchNewCompletionResponse struct {
	// The id of the advanced research session
	ID string `json:"id,required"`
	// The number of characters per source used for the advanced research session
	CharactersPerSource int64 `json:"characters_per_source,required"`
	// The description of the advanced research session
	Description string `json:"description,required"`
	// The name of the advanced research session
	Name string `json:"name,required"`
	// The outputs of the advanced research session
	Outputs any `json:"outputs,required"`
	// The number of sources used for the advanced research session
	Sources int64 `json:"sources,required"`
	// The timestamp of the advanced research session
	Timestamp string `json:"timestamp,required"`
	// The usage of the advanced research session
	Usage map[string]any `json:"usage,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID                  respjson.Field
		CharactersPerSource respjson.Field
		Description         respjson.Field
		Name                respjson.Field
		Outputs             respjson.Field
		Sources             respjson.Field
		Timestamp           respjson.Field
		Usage               respjson.Field
		ExtraFields         map[string]respjson.Field
		raw                 string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClientAdvancedResearchNewCompletionResponse) RawJSON() string { return r.JSON.raw }
func (r *ClientAdvancedResearchNewCompletionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClientAdvancedResearchNewCompletionParams struct {
	// The task to be completed
	Task param.Opt[string] `json:"task,omitzero,required"`
	// The configuration for the advanced research
	Config ClientAdvancedResearchNewCompletionParamsConfig `json:"config,omitzero,required"`
	// The image to be used for the advanced research
	Img param.Opt[string] `json:"img,omitzero"`
	paramObj
}

func (r ClientAdvancedResearchNewCompletionParams) MarshalJSON() (data []byte, err error) {
	type shadow ClientAdvancedResearchNewCompletionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClientAdvancedResearchNewCompletionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The configuration for the advanced research
type ClientAdvancedResearchNewCompletionParamsConfig struct {
	// Description of the advanced research session
	Description param.Opt[string] `json:"description,omitzero"`
	// Name of the director agent
	DirectorAgentName param.Opt[string] `json:"director_agent_name,omitzero"`
	// Maximum loops for the director agent
	DirectorMaxLoops param.Opt[int64] `json:"director_max_loops,omitzero"`
	// Maximum tokens for the director agent's output
	DirectorMaxTokens param.Opt[int64] `json:"director_max_tokens,omitzero"`
	// Model name for the director agent
	DirectorModelName param.Opt[string] `json:"director_model_name,omitzero"`
	// Maximum characters to return from the Exa search tool
	ExaSearchMaxCharacters param.Opt[int64] `json:"exa_search_max_characters,omitzero"`
	// Number of results to return from the Exa search tool
	ExaSearchNumResults param.Opt[int64] `json:"exa_search_num_results,omitzero"`
	// Number of research loops to run
	MaxLoops param.Opt[int64] `json:"max_loops,omitzero"`
	// Name of the advanced research session
	Name param.Opt[string] `json:"name,omitzero"`
	// Model name for worker agents
	WorkerModelName param.Opt[string] `json:"worker_model_name,omitzero"`
	paramObj
}

func (r ClientAdvancedResearchNewCompletionParamsConfig) MarshalJSON() (data []byte, err error) {
	type shadow ClientAdvancedResearchNewCompletionParamsConfig
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClientAdvancedResearchNewCompletionParamsConfig) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
