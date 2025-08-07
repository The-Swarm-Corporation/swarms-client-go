// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package swarms

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/The-Swarm-Corporation/swarms-client-go/internal/apijson"
	shimjson "github.com/The-Swarm-Corporation/swarms-client-go/internal/encoding/json"
	"github.com/The-Swarm-Corporation/swarms-client-go/internal/requestconfig"
	"github.com/The-Swarm-Corporation/swarms-client-go/option"
	"github.com/The-Swarm-Corporation/swarms-client-go/packages/param"
	"github.com/The-Swarm-Corporation/swarms-client-go/packages/respjson"
)

// SwarmService contains methods and other services that help with interacting with
// the swarms-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSwarmService] method instead.
type SwarmService struct {
	Options []option.RequestOption
	Batch   SwarmBatchService
}

// NewSwarmService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewSwarmService(opts ...option.RequestOption) (r SwarmService) {
	r = SwarmService{}
	r.Options = opts
	r.Batch = NewSwarmBatchService(opts...)
	return
}

// Check the available swarm types.
func (r *SwarmService) CheckAvailable(ctx context.Context, opts ...option.RequestOption) (res *SwarmCheckAvailableResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/swarms/available"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get all API request logs for all API keys associated with the user identified by
// the provided API key, excluding any logs that contain a client_ip field in their
// data.
func (r *SwarmService) GetLogs(ctx context.Context, opts ...option.RequestOption) (res *SwarmGetLogsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/swarm/logs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Run a swarm with the specified task.
func (r *SwarmService) Run(ctx context.Context, body SwarmRunParams, opts ...option.RequestOption) (res *SwarmRunResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/swarm/completions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SwarmSpecParam struct {
	// A comprehensive description of the swarm's objectives, capabilities, and
	// intended outcomes.
	Description param.Opt[string] `json:"description,omitzero"`
	// The number of loops to run per agent in the heavy swarm.
	HeavySwarmLoopsPerAgent param.Opt[int64] `json:"heavy_swarm_loops_per_agent,omitzero"`
	// The model name to use for the question agent in the heavy swarm.
	HeavySwarmQuestionAgentModelName param.Opt[string] `json:"heavy_swarm_question_agent_model_name,omitzero"`
	// The model name to use for the worker agent in the heavy swarm.
	HeavySwarmWorkerModelName param.Opt[string] `json:"heavy_swarm_worker_model_name,omitzero"`
	// An optional image URL that may be associated with the swarm's task or
	// representation.
	Img param.Opt[string] `json:"img,omitzero"`
	// The maximum number of execution loops allowed for the swarm, enabling repeated
	// processing if needed.
	MaxLoops param.Opt[int64] `json:"max_loops,omitzero"`
	// The name of the swarm, which serves as an identifier for the group of agents and
	// their collective task.
	Name param.Opt[string] `json:"name,omitzero"`
	// Instructions on how to rearrange the flow of tasks among agents, if applicable.
	RearrangeFlow param.Opt[string] `json:"rearrange_flow,omitzero"`
	// Guidelines or constraints that govern the behavior and interactions of the
	// agents within the swarm.
	Rules param.Opt[string] `json:"rules,omitzero"`
	// The service tier to use for processing. Options: 'standard' (default) or 'flex'
	// for lower cost but slower processing.
	ServiceTier param.Opt[string] `json:"service_tier,omitzero"`
	// A flag indicating whether the swarm should stream its output.
	Stream param.Opt[bool] `json:"stream,omitzero"`
	// The specific task or objective that the swarm is designed to accomplish.
	Task param.Opt[string] `json:"task,omitzero"`
	// A list of agents or specifications that define the agents participating in the
	// swarm.
	Agents []AgentSpecParam `json:"agents,omitzero"`
	// A list of messages that the swarm should complete.
	Messages SwarmSpecMessagesUnionParam `json:"messages,omitzero"`
	// The classification of the swarm, indicating its operational style and
	// methodology.
	//
	// Any of "AgentRearrange", "MixtureOfAgents", "SpreadSheetSwarm",
	// "SequentialWorkflow", "ConcurrentWorkflow", "GroupChat", "MultiAgentRouter",
	// "AutoSwarmBuilder", "HiearchicalSwarm", "auto", "MajorityVoting", "MALT",
	// "DeepResearchSwarm", "CouncilAsAJudge", "InteractiveGroupChat", "HeavySwarm".
	SwarmType SwarmSpecSwarmType `json:"swarm_type,omitzero"`
	// A list of tasks that the swarm should complete.
	Tasks []string `json:"tasks,omitzero"`
	paramObj
}

func (r SwarmSpecParam) MarshalJSON() (data []byte, err error) {
	type shadow SwarmSpecParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *SwarmSpecParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Only one field can be non-zero.
//
// Use [param.IsOmitted] to confirm if a field is set.
type SwarmSpecMessagesUnionParam struct {
	OfMapOfAnyMap []map[string]any `json:",omitzero,inline"`
	OfAnyMap      map[string]any   `json:",omitzero,inline"`
	paramUnion
}

func (u SwarmSpecMessagesUnionParam) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfMapOfAnyMap, u.OfAnyMap)
}
func (u *SwarmSpecMessagesUnionParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, u)
}

func (u *SwarmSpecMessagesUnionParam) asAny() any {
	if !param.IsOmitted(u.OfMapOfAnyMap) {
		return &u.OfMapOfAnyMap
	} else if !param.IsOmitted(u.OfAnyMap) {
		return &u.OfAnyMap
	}
	return nil
}

// The classification of the swarm, indicating its operational style and
// methodology.
type SwarmSpecSwarmType string

const (
	SwarmSpecSwarmTypeAgentRearrange       SwarmSpecSwarmType = "AgentRearrange"
	SwarmSpecSwarmTypeMixtureOfAgents      SwarmSpecSwarmType = "MixtureOfAgents"
	SwarmSpecSwarmTypeSpreadSheetSwarm     SwarmSpecSwarmType = "SpreadSheetSwarm"
	SwarmSpecSwarmTypeSequentialWorkflow   SwarmSpecSwarmType = "SequentialWorkflow"
	SwarmSpecSwarmTypeConcurrentWorkflow   SwarmSpecSwarmType = "ConcurrentWorkflow"
	SwarmSpecSwarmTypeGroupChat            SwarmSpecSwarmType = "GroupChat"
	SwarmSpecSwarmTypeMultiAgentRouter     SwarmSpecSwarmType = "MultiAgentRouter"
	SwarmSpecSwarmTypeAutoSwarmBuilder     SwarmSpecSwarmType = "AutoSwarmBuilder"
	SwarmSpecSwarmTypeHiearchicalSwarm     SwarmSpecSwarmType = "HiearchicalSwarm"
	SwarmSpecSwarmTypeAuto                 SwarmSpecSwarmType = "auto"
	SwarmSpecSwarmTypeMajorityVoting       SwarmSpecSwarmType = "MajorityVoting"
	SwarmSpecSwarmTypeMalt                 SwarmSpecSwarmType = "MALT"
	SwarmSpecSwarmTypeDeepResearchSwarm    SwarmSpecSwarmType = "DeepResearchSwarm"
	SwarmSpecSwarmTypeCouncilAsAJudge      SwarmSpecSwarmType = "CouncilAsAJudge"
	SwarmSpecSwarmTypeInteractiveGroupChat SwarmSpecSwarmType = "InteractiveGroupChat"
	SwarmSpecSwarmTypeHeavySwarm           SwarmSpecSwarmType = "HeavySwarm"
)

type SwarmCheckAvailableResponse struct {
	Success    bool `json:"success,nullable"`
	SwarmTypes any  `json:"swarm_types"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Success     respjson.Field
		SwarmTypes  respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SwarmCheckAvailableResponse) RawJSON() string { return r.JSON.raw }
func (r *SwarmCheckAvailableResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SwarmGetLogsResponse struct {
	Count     int64  `json:"count,nullable"`
	Logs      any    `json:"logs"`
	Status    string `json:"status,nullable"`
	Timestamp string `json:"timestamp,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Logs        respjson.Field
		Status      respjson.Field
		Timestamp   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SwarmGetLogsResponse) RawJSON() string { return r.JSON.raw }
func (r *SwarmGetLogsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SwarmRunResponse struct {
	// The description of the swarm.
	Description string `json:"description,required"`
	// The execution time of the swarm.
	ExecutionTime float64 `json:"execution_time,required"`
	// The unique identifier for the swarm completion.
	JobID string `json:"job_id,required"`
	// The number of agents in the swarm.
	NumberOfAgents int64 `json:"number_of_agents,required"`
	// The output of the swarm.
	Output any `json:"output,required"`
	// The service tier of the swarm.
	ServiceTier string `json:"service_tier,required"`
	// The status of the swarm completion.
	Status string `json:"status,required"`
	// The name of the swarm.
	SwarmName string `json:"swarm_name,required"`
	// The type of the swarm.
	SwarmType string `json:"swarm_type,required"`
	// The usage of the swarm.
	Usage map[string]any `json:"usage,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Description    respjson.Field
		ExecutionTime  respjson.Field
		JobID          respjson.Field
		NumberOfAgents respjson.Field
		Output         respjson.Field
		ServiceTier    respjson.Field
		Status         respjson.Field
		SwarmName      respjson.Field
		SwarmType      respjson.Field
		Usage          respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r SwarmRunResponse) RawJSON() string { return r.JSON.raw }
func (r *SwarmRunResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type SwarmRunParams struct {
	SwarmSpec SwarmSpecParam
	paramObj
}

func (r SwarmRunParams) MarshalJSON() (data []byte, err error) {
	return shimjson.Marshal(r.SwarmSpec)
}
func (r *SwarmRunParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.SwarmSpec)
}
