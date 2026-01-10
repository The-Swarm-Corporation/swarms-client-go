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

// ClientGraphWorkflowService contains methods and other services that help with
// interacting with the swarms-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClientGraphWorkflowService] method instead.
type ClientGraphWorkflowService struct {
	Options []option.RequestOption
}

// NewClientGraphWorkflowService generates a new service that applies the given
// options to each request. These options are applied after the parent client's
// options (if there is one), and before any request-specific options.
func NewClientGraphWorkflowService(opts ...option.RequestOption) (r ClientGraphWorkflowService) {
	r = ClientGraphWorkflowService{}
	r.Options = opts
	return
}

// Execute a graph workflow with directed agent nodes and edges. Enables complex
// multi-agent collaboration with parallel execution, automatic compilation, and
// comprehensive workflow orchestration.
func (r *ClientGraphWorkflowService) ExecuteWorkflow(ctx context.Context, body ClientGraphWorkflowExecuteWorkflowParams, opts ...option.RequestOption) (res *ClientGraphWorkflowExecuteWorkflowResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/graph-workflow/completions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Output schema for GraphWorkflow completion responses.
type ClientGraphWorkflowExecuteWorkflowResponse struct {
	// The job ID of the graph workflow.
	JobID string `json:"job_id,required"`
	// The outputs of the graph workflow.
	Outputs any `json:"outputs,required"`
	// The status of the graph workflow.
	Status string `json:"status,required"`
	// The timestamp of the graph workflow execution.
	Timestamp string `json:"timestamp,required"`
	// The usage statistics of the workflow.
	Usage ClientGraphWorkflowExecuteWorkflowResponseUsage `json:"usage,required"`
	// The description of the graph workflow.
	Description string `json:"description,nullable"`
	// The name of the graph workflow.
	Name string `json:"name,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		JobID       respjson.Field
		Outputs     respjson.Field
		Status      respjson.Field
		Timestamp   respjson.Field
		Usage       respjson.Field
		Description respjson.Field
		Name        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClientGraphWorkflowExecuteWorkflowResponse) RawJSON() string { return r.JSON.raw }
func (r *ClientGraphWorkflowExecuteWorkflowResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The usage statistics of the workflow.
type ClientGraphWorkflowExecuteWorkflowResponseUsage struct {
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
func (r ClientGraphWorkflowExecuteWorkflowResponseUsage) RawJSON() string { return r.JSON.raw }
func (r *ClientGraphWorkflowExecuteWorkflowResponseUsage) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type ClientGraphWorkflowExecuteWorkflowParams struct {
	// Whether to automatically compile the workflow for optimization.
	AutoCompile param.Opt[bool] `json:"auto_compile,omitzero"`
	// The description of the graph workflow.
	Description param.Opt[string] `json:"description,omitzero"`
	// Optional image path for vision-enabled agents.
	Img param.Opt[string] `json:"img,omitzero"`
	// The maximum number of execution loops for the workflow.
	MaxLoops param.Opt[int64] `json:"max_loops,omitzero"`
	// The name of the graph workflow.
	Name param.Opt[string] `json:"name,omitzero"`
	// The task to be executed by the workflow.
	Task param.Opt[string] `json:"task,omitzero"`
	// Whether to enable detailed logging.
	Verbose param.Opt[bool] `json:"verbose,omitzero"`
	// List of agent specifications to be used as nodes in the workflow graph.
	Agents []AgentSpecParam `json:"agents,omitzero"`
	// List of edges connecting nodes. Can be EdgeSpec objects or dictionaries with
	// 'source' and 'target' keys.
	Edges []ClientGraphWorkflowExecuteWorkflowParamsEdgeUnion `json:"edges,omitzero"`
	// List of node IDs that serve as ending points for the workflow.
	EndPoints []string `json:"end_points,omitzero"`
	// List of node IDs that serve as starting points for the workflow.
	EntryPoints []string `json:"entry_points,omitzero"`
	paramObj
}

func (r ClientGraphWorkflowExecuteWorkflowParams) MarshalJSON() (data []byte, err error) {
	type shadow ClientGraphWorkflowExecuteWorkflowParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClientGraphWorkflowExecuteWorkflowParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type ClientGraphWorkflowExecuteWorkflowParamsEdgeUnion struct {
	OfEdgeSpec *ClientGraphWorkflowExecuteWorkflowParamsEdgeEdgeSpec `json:",omitzero,inline"`
	OfAnyMap   map[string]any                                        `json:",omitzero,inline"`
	paramUnion
}

func (u ClientGraphWorkflowExecuteWorkflowParamsEdgeUnion) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfEdgeSpec, u.OfAnyMap)
}
func (u *ClientGraphWorkflowExecuteWorkflowParamsEdgeUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *ClientGraphWorkflowExecuteWorkflowParamsEdgeUnion) asAny() any {
	if !param.IsOmitted(u.OfEdgeSpec) {
		return u.OfEdgeSpec
	} else if !param.IsOmitted(u.OfAnyMap) {
		return &u.OfAnyMap
	}
	return nil
}

// Schema for defining an edge between nodes in the workflow graph.
//
// The properties Source, Target are required.
type ClientGraphWorkflowExecuteWorkflowParamsEdgeEdgeSpec struct {
	// The source node ID.
	Source string `json:"source,required"`
	// The target node ID.
	Target string `json:"target,required"`
	// Optional metadata for the edge.
	Metadata map[string]any `json:"metadata,omitzero"`
	paramObj
}

func (r ClientGraphWorkflowExecuteWorkflowParamsEdgeEdgeSpec) MarshalJSON() (data []byte, err error) {
	type shadow ClientGraphWorkflowExecuteWorkflowParamsEdgeEdgeSpec
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ClientGraphWorkflowExecuteWorkflowParamsEdgeEdgeSpec) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
