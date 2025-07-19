// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package swarms

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/The-Swarm-Corporation/swarms-client-go/internal/apijson"
	"github.com/The-Swarm-Corporation/swarms-client-go/internal/requestconfig"
	"github.com/The-Swarm-Corporation/swarms-client-go/option"
	"github.com/The-Swarm-Corporation/swarms-client-go/packages/param"
)

// AgentService contains methods and other services that help with interacting with
// the swarms-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentService] method instead.
type AgentService struct {
	Options []option.RequestOption
	Batch   AgentBatchService
}

// NewAgentService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAgentService(opts ...option.RequestOption) (r AgentService) {
	r = AgentService{}
	r.Options = opts
	r.Batch = NewAgentBatchService(opts...)
	return
}

// Run an agent with the specified task.
func (r *AgentService) Run(ctx context.Context, body AgentRunParams, opts ...option.RequestOption) (res *AgentRunResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/agent/completions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AgentCompletionParam struct {
	// An optional image URL that may be associated with the agent's task or
	// representation.
	Img param.Opt[string] `json:"img,omitzero"`
	// A flag indicating whether the agent should stream its output.
	Stream param.Opt[bool] `json:"stream,omitzero"`
	// The task to be completed by the agent.
	Task param.Opt[string] `json:"task,omitzero"`
	// The history of the agent's previous tasks and responses. Can be either a
	// dictionary or a list of message objects.
	History AgentCompletionHistoryUnionParam `json:"history,omitzero"`
	// A list of image URLs that may be associated with the agent's task or
	// representation.
	Imgs []string `json:"imgs,omitzero"`
	// The configuration of the agent to be completed.
	AgentConfig AgentSpecParam `json:"agent_config,omitzero"`
	paramObj
}

func (r AgentCompletionParam) MarshalJSON() (data []byte, err error) {
	type shadow AgentCompletionParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentCompletionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type AgentCompletionHistoryUnionParam struct {
	OfAnyMap         map[string]any      `json:",omitzero,inline"`
	OfMapOfStringMap []map[string]string `json:",omitzero,inline"`
	paramUnion
}

func (u AgentCompletionHistoryUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAnyMap, u.OfMapOfStringMap)
}
func (u *AgentCompletionHistoryUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *AgentCompletionHistoryUnionParam) asAny() any {
	if !param.IsOmitted(u.OfAnyMap) {
		return &u.OfAnyMap
	} else if !param.IsOmitted(u.OfMapOfStringMap) {
		return &u.OfMapOfStringMap
	}
	return nil
}

// The property AgentName is required.
type AgentSpecParam struct {
	// The unique name assigned to the agent, which identifies its role and
	// functionality within the swarm.
	AgentName param.Opt[string] `json:"agent_name,omitzero,required"`
	// A flag indicating whether the agent should automatically create prompts based on
	// the task requirements.
	AutoGeneratePrompt param.Opt[bool] `json:"auto_generate_prompt,omitzero"`
	// A detailed explanation of the agent's purpose, capabilities, and any specific
	// tasks it is designed to perform.
	Description param.Opt[string] `json:"description,omitzero"`
	// The maximum number of times the agent is allowed to repeat its task, enabling
	// iterative processing if necessary.
	MaxLoops param.Opt[int64] `json:"max_loops,omitzero"`
	// The maximum number of tokens that the agent is allowed to generate in its
	// responses, limiting output length.
	MaxTokens param.Opt[int64] `json:"max_tokens,omitzero"`
	// The URL of the MCP server that the agent can use to complete its task.
	McpURL param.Opt[string] `json:"mcp_url,omitzero"`
	// The name of the AI model that the agent will utilize for processing tasks and
	// generating outputs. For example: gpt-4o, gpt-4o-mini, openai/o3-mini
	ModelName param.Opt[string] `json:"model_name,omitzero"`
	// The designated role of the agent within the swarm, which influences its behavior
	// and interaction with other agents.
	Role param.Opt[string] `json:"role,omitzero"`
	// A flag indicating whether the agent should stream its output.
	StreamingOn param.Opt[bool] `json:"streaming_on,omitzero"`
	// The initial instruction or context provided to the agent, guiding its behavior
	// and responses during execution.
	SystemPrompt param.Opt[string] `json:"system_prompt,omitzero"`
	// A parameter that controls the randomness of the agent's output; lower values
	// result in more deterministic responses.
	Temperature param.Opt[float64] `json:"temperature,omitzero"`
	// A dictionary of tools that the agent can use to complete its task.
	ToolsListDictionary []map[string]any `json:"tools_list_dictionary,omitzero"`
	paramObj
}

func (r AgentSpecParam) MarshalJSON() (data []byte, err error) {
	type shadow AgentSpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentSpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRunResponse map[string]any

type AgentRunParams struct {
	AgentCompletion AgentCompletionParam
	paramObj
}

func (r AgentRunParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.AgentCompletion)
}
func (r *AgentRunParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.AgentCompletion)
}
