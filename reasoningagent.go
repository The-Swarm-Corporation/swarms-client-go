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
)

// ReasoningAgentService contains methods and other services that help with
// interacting with the swarms-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewReasoningAgentService] method instead.
type ReasoningAgentService struct {
	Options []option.RequestOption
}

// NewReasoningAgentService generates a new service that applies the given options
// to each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewReasoningAgentService(opts ...option.RequestOption) (r ReasoningAgentService) {
	r = ReasoningAgentService{}
	r.Options = opts
	return
}

// Run a reasoning agent with the specified task.
func (r *ReasoningAgentService) NewCompletion(ctx context.Context, body ReasoningAgentNewCompletionParams, opts ...option.RequestOption) (res *ReasoningAgentNewCompletionResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/reasoning-agent/completions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Get the types of reasoning agents available.
func (r *ReasoningAgentService) ListTypes(ctx context.Context, opts ...option.RequestOption) (res *ReasoningAgentListTypesResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/reasoning-agent/types"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ReasoningAgentNewCompletionResponse map[string]any

type ReasoningAgentListTypesResponse map[string]any

type ReasoningAgentNewCompletionParams struct {
	// The unique name assigned to the reasoning agent.
	AgentName param.Opt[string] `json:"agent_name,omitzero"`
	// A detailed explanation of the reasoning agent's purpose and capabilities.
	Description param.Opt[string] `json:"description,omitzero"`
	// The maximum number of times the reasoning agent is allowed to repeat its task.
	MaxLoops param.Opt[int64] `json:"max_loops,omitzero"`
	// The memory capacity for the reasoning agent.
	MemoryCapacity param.Opt[int64] `json:"memory_capacity,omitzero"`
	// The name of the AI model that the reasoning agent will utilize.
	ModelName param.Opt[string] `json:"model_name,omitzero"`
	// The number of knowledge items to use for the reasoning agent.
	NumKnowledgeItems param.Opt[int64] `json:"num_knowledge_items,omitzero"`
	// The number of samples to generate for the reasoning agent.
	NumSamples param.Opt[int64] `json:"num_samples,omitzero"`
	// The initial instruction or context provided to the reasoning agent.
	SystemPrompt param.Opt[string] `json:"system_prompt,omitzero"`
	// The task to be completed by the reasoning agent.
	Task param.Opt[string] `json:"task,omitzero"`
	// The type of output format for the reasoning agent.
	//
	// Any of "list", "dict", "dictionary", "string", "str", "final", "last", "json",
	// "all", "yaml", "xml", "dict-all-except-first", "str-all-except-first",
	// "basemodel", "dict-final", "list-final".
	OutputType ReasoningAgentNewCompletionParamsOutputType `json:"output_type,omitzero"`
	// The type of reasoning swarm to use (e.g., reasoning duo, self-consistency, IRE).
	//
	// Any of "reasoning-duo", "self-consistency", "ire", "reasoning-agent",
	// "consistency-agent", "ire-agent", "ReflexionAgent", "GKPAgent", "AgentJudge".
	SwarmType ReasoningAgentNewCompletionParamsSwarmType `json:"swarm_type,omitzero"`
	paramObj
}

func (r ReasoningAgentNewCompletionParams) MarshalJSON() (data []byte, err error) {
	type shadow ReasoningAgentNewCompletionParams
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *ReasoningAgentNewCompletionParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The type of output format for the reasoning agent.
type ReasoningAgentNewCompletionParamsOutputType string

const (
	ReasoningAgentNewCompletionParamsOutputTypeList               ReasoningAgentNewCompletionParamsOutputType = "list"
	ReasoningAgentNewCompletionParamsOutputTypeDict               ReasoningAgentNewCompletionParamsOutputType = "dict"
	ReasoningAgentNewCompletionParamsOutputTypeDictionary         ReasoningAgentNewCompletionParamsOutputType = "dictionary"
	ReasoningAgentNewCompletionParamsOutputTypeString             ReasoningAgentNewCompletionParamsOutputType = "string"
	ReasoningAgentNewCompletionParamsOutputTypeStr                ReasoningAgentNewCompletionParamsOutputType = "str"
	ReasoningAgentNewCompletionParamsOutputTypeFinal              ReasoningAgentNewCompletionParamsOutputType = "final"
	ReasoningAgentNewCompletionParamsOutputTypeLast               ReasoningAgentNewCompletionParamsOutputType = "last"
	ReasoningAgentNewCompletionParamsOutputTypeJson               ReasoningAgentNewCompletionParamsOutputType = "json"
	ReasoningAgentNewCompletionParamsOutputTypeAll                ReasoningAgentNewCompletionParamsOutputType = "all"
	ReasoningAgentNewCompletionParamsOutputTypeYaml               ReasoningAgentNewCompletionParamsOutputType = "yaml"
	ReasoningAgentNewCompletionParamsOutputTypeXml                ReasoningAgentNewCompletionParamsOutputType = "xml"
	ReasoningAgentNewCompletionParamsOutputTypeDictAllExceptFirst ReasoningAgentNewCompletionParamsOutputType = "dict-all-except-first"
	ReasoningAgentNewCompletionParamsOutputTypeStrAllExceptFirst  ReasoningAgentNewCompletionParamsOutputType = "str-all-except-first"
	ReasoningAgentNewCompletionParamsOutputTypeBasemodel          ReasoningAgentNewCompletionParamsOutputType = "basemodel"
	ReasoningAgentNewCompletionParamsOutputTypeDictFinal          ReasoningAgentNewCompletionParamsOutputType = "dict-final"
	ReasoningAgentNewCompletionParamsOutputTypeListFinal          ReasoningAgentNewCompletionParamsOutputType = "list-final"
)

// The type of reasoning swarm to use (e.g., reasoning duo, self-consistency, IRE).
type ReasoningAgentNewCompletionParamsSwarmType string

const (
	ReasoningAgentNewCompletionParamsSwarmTypeReasoningDuo     ReasoningAgentNewCompletionParamsSwarmType = "reasoning-duo"
	ReasoningAgentNewCompletionParamsSwarmTypeSelfConsistency  ReasoningAgentNewCompletionParamsSwarmType = "self-consistency"
	ReasoningAgentNewCompletionParamsSwarmTypeIre              ReasoningAgentNewCompletionParamsSwarmType = "ire"
	ReasoningAgentNewCompletionParamsSwarmTypeReasoningAgent   ReasoningAgentNewCompletionParamsSwarmType = "reasoning-agent"
	ReasoningAgentNewCompletionParamsSwarmTypeConsistencyAgent ReasoningAgentNewCompletionParamsSwarmType = "consistency-agent"
	ReasoningAgentNewCompletionParamsSwarmTypeIreAgent         ReasoningAgentNewCompletionParamsSwarmType = "ire-agent"
	ReasoningAgentNewCompletionParamsSwarmTypeReflexionAgent   ReasoningAgentNewCompletionParamsSwarmType = "ReflexionAgent"
	ReasoningAgentNewCompletionParamsSwarmTypeGkpAgent         ReasoningAgentNewCompletionParamsSwarmType = "GKPAgent"
	ReasoningAgentNewCompletionParamsSwarmTypeAgentJudge       ReasoningAgentNewCompletionParamsSwarmType = "AgentJudge"
)
