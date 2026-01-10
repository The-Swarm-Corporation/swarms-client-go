// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package swarms

import (
	"context"
	"encoding/json"
	"net/http"
	"slices"

	"github.com/The-Swarm-Corporation/swarms-client-go/internal/apijson"
	shimjson "github.com/The-Swarm-Corporation/swarms-client-go/internal/encoding/json"
	"github.com/The-Swarm-Corporation/swarms-client-go/internal/requestconfig"
	"github.com/The-Swarm-Corporation/swarms-client-go/option"
	"github.com/The-Swarm-Corporation/swarms-client-go/packages/param"
	"github.com/The-Swarm-Corporation/swarms-client-go/packages/respjson"
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

// Get all unique agent configurations that the user has created or used, without
// task details. Allows users to reuse agent configs with new tasks.
func (r *AgentService) List(ctx context.Context, opts ...option.RequestOption) (res *AgentListResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/agents/list"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Run an agent with the specified task. Supports streaming when stream=True.
func (r *AgentService) Run(ctx context.Context, body AgentRunParams, opts ...option.RequestOption) (res *AgentRunResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/agent/completions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AgentCompletionParam struct {
	// An optional image URL that may be associated with the agent's task or
	// representation.
	Img param.Opt[string] `json:"img,omitzero"`
	// The task to be completed by the agent.
	Task param.Opt[string] `json:"task,omitzero"`
	// The history of the agent's previous tasks and responses. Can be either a
	// dictionary or a list of message objects.
	History AgentCompletionHistoryUnionParam `json:"history,omitzero"`
	// A list of image URLs that may be associated with the agent's task or
	// representation.
	Imgs []string `json:"imgs,omitzero"`
	// A list of tools that the agent should use to complete its task.
	ToolsEnabled []string `json:"tools_enabled,omitzero"`
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
	// A flag indicating whether the agent should dynamically adjust its temperature
	// based on the task.
	DynamicTemperatureEnabled param.Opt[bool] `json:"dynamic_temperature_enabled,omitzero"`
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
	// The effort to put into reasoning.
	ReasoningEffort param.Opt[string] `json:"reasoning_effort,omitzero"`
	// A parameter enabling an agent to use reasoning.
	ReasoningEnabled param.Opt[bool] `json:"reasoning_enabled,omitzero"`
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
	// The number of tokens to use for thinking.
	ThinkingTokens param.Opt[int64] `json:"thinking_tokens,omitzero"`
	// A parameter enabling an agent to summarize tool calls.
	ToolCallSummary param.Opt[bool] `json:"tool_call_summary,omitzero"`
	// Additional arguments to pass to the LLM such as top_p, frequency_penalty,
	// presence_penalty, etc.
	LlmArgs map[string]any `json:"llm_args,omitzero"`
	// The MCP connections to use for the agent. This is a list of MCP connections.
	// Includes multiple MCP connections.
	McpConfigs AgentSpecMcpConfigsParam `json:"mcp_configs,omitzero"`
	// A dictionary of tools that the agent can use to complete its task.
	ToolsListDictionary []map[string]any `json:"tools_list_dictionary,omitzero"`
	// The MCP connection to use for the agent.
	McpConfig McpConnectionParam `json:"mcp_config,omitzero"`
	paramObj
}

func (r AgentSpecParam) MarshalJSON() (data []byte, err error) {
	type shadow AgentSpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentSpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The MCP connections to use for the agent. This is a list of MCP connections.
// Includes multiple MCP connections.
//
// The property Connections is required.
type AgentSpecMcpConfigsParam struct {
	// List of MCP connections
	Connections []McpConnectionParam `json:"connections,omitzero,required"`
	paramObj
}

func (r AgentSpecMcpConfigsParam) MarshalJSON() (data []byte, err error) {
	type shadow AgentSpecMcpConfigsParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AgentSpecMcpConfigsParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type McpConnectionParam struct {
	// Authentication token for accessing the MCP server
	AuthorizationToken param.Opt[string] `json:"authorization_token,omitzero"`
	// Timeout for the MCP server
	Timeout param.Opt[int64] `json:"timeout,omitzero"`
	// The transport protocol to use for the MCP server
	Transport param.Opt[string] `json:"transport,omitzero"`
	// The type of connection, defaults to 'mcp'
	Type param.Opt[string] `json:"type,omitzero"`
	// The URL endpoint for the MCP server
	URL param.Opt[string] `json:"url,omitzero"`
	// Headers to send to the MCP server
	Headers map[string]string `json:"headers,omitzero"`
	// Dictionary containing configuration settings for MCP tools
	ToolConfigurations map[string]any `json:"tool_configurations,omitzero"`
	ExtraFields        map[string]any `json:"-"`
	paramObj
}

func (r McpConnectionParam) MarshalJSON() (data []byte, err error) {
	type shadow McpConnectionParam
	return param.MarshalWithExtras(r, (*shadow)(&r), r.ExtraFields)
}
func (r *McpConnectionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentListResponse map[string]any

type AgentRunResponse struct {
	// A description of the agent or completion.
	Description string `json:"description,nullable"`
	// The unique identifier for the agent completion.
	JobID string `json:"job_id,nullable"`
	// The name of the agent.
	Name string `json:"name,nullable"`
	// The outputs generated by the agent.
	Outputs any `json:"outputs"`
	// Indicates whether the agent completion was successful.
	Success bool `json:"success,nullable"`
	// The temperature setting used for the agent's response generation.
	Temperature float64 `json:"temperature,nullable"`
	// The timestamp when the agent completion was created.
	Timestamp string `json:"timestamp,nullable"`
	// Usage statistics or metadata for the agent completion.
	Usage map[string]any `json:"usage,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description respjson.Field
		JobID       respjson.Field
		Name        respjson.Field
		Outputs     respjson.Field
		Success     respjson.Field
		Temperature respjson.Field
		Timestamp   respjson.Field
		Usage       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AgentRunResponse) RawJSON() string { return r.JSON.raw }
func (r *AgentRunResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AgentRunParams struct {
	AgentCompletion AgentCompletionParam
	paramObj
}

func (r AgentRunParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.AgentCompletion)
}
func (r *AgentRunParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.AgentCompletion)
}
