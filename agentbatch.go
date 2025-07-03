// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package swarms

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/stainless-sdks/swarms-go/internal/requestconfig"
	"github.com/stainless-sdks/swarms-go/option"
)

// AgentBatchService contains methods and other services that help with interacting
// with the swarms-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAgentBatchService] method instead.
type AgentBatchService struct {
	Options []option.RequestOption
}

// NewAgentBatchService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewAgentBatchService(opts ...option.RequestOption) (r AgentBatchService) {
	r = AgentBatchService{}
	r.Options = opts
	return
}

// Run a batch of agents with the specified tasks using a thread pool.
//
// Args: agent_completions: List of agent completion tasks to process x_api_key:
// API key for authentication
//
// Returns: List[Dict[str, Any]]: List of results from completed agent tasks
//
// Raises: HTTPException: If there's an error processing the batch
func (r *AgentBatchService) Run(ctx context.Context, body AgentBatchRunParams, opts ...option.RequestOption) (res *AgentBatchRunResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/agent/batch/completions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type AgentBatchRunResponse map[string]any

type AgentBatchRunParams struct {
	Body []AgentCompletionParam
	paramObj
}

func (r AgentBatchRunParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}
func (r *AgentBatchRunParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}
