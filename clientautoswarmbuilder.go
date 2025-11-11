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

// ClientAutoSwarmBuilderService contains methods and other services that help with
// interacting with the swarms-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClientAutoSwarmBuilderService] method instead.
type ClientAutoSwarmBuilderService struct {
	Options []option.RequestOption
}

// NewClientAutoSwarmBuilderService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewClientAutoSwarmBuilderService(opts ...option.RequestOption) (r ClientAutoSwarmBuilderService) {
	r = ClientAutoSwarmBuilderService{}
	r.Options = opts
	return
}

// Generate and orchestrate agent swarms autonomously using AI-powered swarm
// composition and task decomposition.
func (r *ClientAutoSwarmBuilderService) NewCompletion(ctx context.Context, body ClientAutoSwarmBuilderNewCompletionParams, opts ...option.RequestOption) (res *ClientAutoSwarmBuilderNewCompletionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/auto-swarm-builder/completions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve all available execution types and return formats for the Auto Swarm
// Builder endpoint.
func (r *ClientAutoSwarmBuilderService) ListExecutionTypes(ctx context.Context, opts ...option.RequestOption) (res *[]string, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/auto-swarm-builder/execution-types"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Schema for the Auto Swarm Builder API response.
//
// Attributes: success (bool): Whether the swarm was built successfully. job_id
// (Optional[str]): The job ID of the swarm. outputs (Optional[dict]): The outputs
// of the auto swarms builder. type (Optional[str]): The type of the swarm
// execution. timestamp (Optional[str]): The timestamp of the swarm execution.
// usage (Optional[dict]): The usage statistics of the swarm execution.
type ClientAutoSwarmBuilderNewCompletionResponse struct {
	// Whether the swarm was built successfully.
	Success bool `json:"success,required"`
	// The job ID of the swarm.
	JobID string `json:"job_id,nullable"`
	// The outputs of the auto swarms builder.
	Outputs map[string]any `json:"outputs,nullable"`
	// The timestamp of the swarm execution.
	Timestamp string `json:"timestamp,nullable"`
	// The type of the swarm execution.
	Type string `json:"type,nullable"`
	// The usage of the swarm execution.
	Usage map[string]any `json:"usage,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		JobID       respjson.Field
		Outputs     respjson.Field
		Timestamp   respjson.Field
		Type        respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClientAutoSwarmBuilderNewCompletionResponse) RawJSON() string { return r.JSON.raw }
func (r *ClientAutoSwarmBuilderNewCompletionResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClientAutoSwarmBuilderNewCompletionParams struct {
	// A description of the swarm.
	Description param.Opt[string] `json:"description,omitzero"`
	// Maximum number of loops to run.
	MaxLoops param.Opt[int64] `json:"max_loops,omitzero"`
	// The maximum number of tokens to use for the swarm.
	MaxTokens param.Opt[int64] `json:"max_tokens,omitzero"`
	// The model name to use for the swarm.
	ModelName param.Opt[string] `json:"model_name,omitzero"`
	// The name of the swarm.
	Name param.Opt[string] `json:"name,omitzero"`
	// The task for the swarm, if any.
	Task param.Opt[string] `json:"task,omitzero"`
	// The type of execution to perform.
	ExecutionType []any `json:"execution_type,omitzero"`
	paramObj
}

func (r ClientAutoSwarmBuilderNewCompletionParams) MarshalJSON() (data []byte, err error) {
	type shadow ClientAutoSwarmBuilderNewCompletionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClientAutoSwarmBuilderNewCompletionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
