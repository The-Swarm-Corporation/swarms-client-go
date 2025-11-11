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

// ClientBatchedGridWorkflowService contains methods and other services that help
// with interacting with the swarms-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClientBatchedGridWorkflowService] method instead.
type ClientBatchedGridWorkflowService struct {
	Options []option.RequestOption
}

// NewClientBatchedGridWorkflowService generates a new service that applies the
// given options to each request. These options are applied after the parent
// client's options (if there is one), and before any request-specific options.
func NewClientBatchedGridWorkflowService(opts ...option.RequestOption) (r ClientBatchedGridWorkflowService) {
	r = ClientBatchedGridWorkflowService{}
	r.Options = opts
	return
}

// Complete a batched grid workflow with the specified input data. Enables you to
// run a grid workflow with multiple agents and tasks in a single request.
func (r *ClientBatchedGridWorkflowService) CompleteWorkflow(ctx context.Context, body ClientBatchedGridWorkflowCompleteWorkflowParams, opts ...option.RequestOption) (res *ClientBatchedGridWorkflowCompleteWorkflowResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/batched-grid-workflow/completions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type ClientBatchedGridWorkflowCompleteWorkflowResponse struct {
	// The description of the batched grid workflow.
	Description string `json:"description,required"`
	// The job ID of the batched grid workflow.
	JobID string `json:"job_id,required"`
	// The name of the batched grid workflow.
	Name string `json:"name,required"`
	// The outputs of the batched grid workflow.
	Outputs any `json:"outputs,required"`
	// The status of the batched grid workflow.
	Status string `json:"status,required"`
	// The timestamp of the batched grid workflow.
	Timestamp string `json:"timestamp,required"`
	// The usage of the batched grid workflow.
	Usage ClientBatchedGridWorkflowCompleteWorkflowResponseUsage `json:"usage,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		JobID       respjson.Field
		Name        respjson.Field
		Outputs     respjson.Field
		Status      respjson.Field
		Timestamp   respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClientBatchedGridWorkflowCompleteWorkflowResponse) RawJSON() string { return r.JSON.raw }
func (r *ClientBatchedGridWorkflowCompleteWorkflowResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The usage of the batched grid workflow.
type ClientBatchedGridWorkflowCompleteWorkflowResponseUsage struct {
	// The cost in credits for the agents.
	CostPerAgent float64 `json:"cost_per_agent,required"`
	// The number of input tokens.
	InputTokens int64 `json:"input_tokens,required"`
	// The number of output tokens.
	OutputTokens int64 `json:"output_tokens,required"`
	// The cost in credits for the tokens.
	TokenCost float64 `json:"token_cost,required"`
	// The total number of tokens.
	TotalTokens int64 `json:"total_tokens,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CostPerAgent respjson.Field
		InputTokens  respjson.Field
		OutputTokens respjson.Field
		TokenCost    respjson.Field
		TotalTokens  respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClientBatchedGridWorkflowCompleteWorkflowResponseUsage) RawJSON() string { return r.JSON.raw }
func (r *ClientBatchedGridWorkflowCompleteWorkflowResponseUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClientBatchedGridWorkflowCompleteWorkflowParams struct {
	// The description of the batched grid workflow.
	Description param.Opt[string] `json:"description,omitzero"`
	// The maximum number of loops to be completed by the batched grid workflow.
	MaxLoops param.Opt[int64] `json:"max_loops,omitzero"`
	// The name of the batched grid workflow.
	Name param.Opt[string] `json:"name,omitzero"`
	// The agent completions to be completed by the batched grid workflow.
	AgentCompletions []AgentSpecParam `json:"agent_completions,omitzero"`
	// The images to be used by the batched grid workflow.
	Imgs []string `json:"imgs,omitzero"`
	// The tasks to be completed by the batched grid workflow.
	Tasks []string `json:"tasks,omitzero"`
	paramObj
}

func (r ClientBatchedGridWorkflowCompleteWorkflowParams) MarshalJSON() (data []byte, err error) {
	type shadow ClientBatchedGridWorkflowCompleteWorkflowParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClientBatchedGridWorkflowCompleteWorkflowParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
