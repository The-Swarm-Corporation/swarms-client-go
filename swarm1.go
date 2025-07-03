// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package swarms

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/swarms-go/internal/apijson"
	"github.com/stainless-sdks/swarms-go/internal/requestconfig"
	"github.com/stainless-sdks/swarms-go/option"
	"github.com/stainless-sdks/swarms-go/packages/param"
)

// SwarmService contains methods and other services that help with interacting with
// the swarms API.
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
func (r *SwarmService) CheckAvailable(ctx context.Context, query SwarmCheckAvailableParams, opts ...option.RequestOption) (res *SwarmCheckAvailableResponse, err error) {
	if !param.IsOmitted(query.XAPIKey) {
		opts = append(opts, option.WithHeader("x-api-key", fmt.Sprintf("%s", query.XAPIKey)))
	}
	opts = append(r.Options[:], opts...)
	path := "v1/swarms/available"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Get all API request logs for the user associated with the provided API key,
// excluding any logs that contain a client_ip field in their data.
func (r *SwarmService) GetLogs(ctx context.Context, query SwarmGetLogsParams, opts ...option.RequestOption) (res *SwarmGetLogsResponse, err error) {
	if !param.IsOmitted(query.XAPIKey) {
		opts = append(opts, option.WithHeader("x-api-key", fmt.Sprintf("%s", query.XAPIKey)))
	}
	opts = append(r.Options[:], opts...)
	path := "v1/swarm/logs"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Run a swarm with the specified task.
func (r *SwarmService) Run(ctx context.Context, params SwarmRunParams, opts ...option.RequestOption) (res *SwarmRunResponse, err error) {
	if !param.IsOmitted(params.XAPIKey) {
		opts = append(opts, option.WithHeader("x-api-key", fmt.Sprintf("%s", params.XAPIKey)))
	}
	opts = append(r.Options[:], opts...)
	path := "v1/swarm/completions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, params, &res, opts...)
	return
}

type SwarmSpecParam struct {
	// A comprehensive description of the swarm's objectives, capabilities, and
	// intended outcomes.
	Description param.Opt[string] `json:"description,omitzero"`
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
	// A flag indicating whether the swarm should return its execution history along
	// with the final output.
	ReturnHistory param.Opt[bool] `json:"return_history,omitzero"`
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
	// "DeepResearchSwarm", "CouncilAsAJudge", "InteractiveGroupChat".
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
)

type SwarmCheckAvailableResponse map[string]any

type SwarmGetLogsResponse map[string]any

type SwarmRunResponse map[string]any

type SwarmCheckAvailableParams struct {
	XAPIKey string `header:"x-api-key,required" json:"-"`
	paramObj
}

type SwarmGetLogsParams struct {
	XAPIKey string `header:"x-api-key,required" json:"-"`
	paramObj
}

type SwarmRunParams struct {
	SwarmSpec SwarmSpecParam
	XAPIKey   string `header:"x-api-key,required" json:"-"`
	paramObj
}

func (r SwarmRunParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.SwarmSpec)
}
func (r *SwarmRunParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.SwarmSpec)
}
