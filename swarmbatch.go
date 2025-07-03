// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package swarms

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/stainless-sdks/swarms-go/internal/requestconfig"
	"github.com/stainless-sdks/swarms-go/option"
)

// SwarmBatchService contains methods and other services that help with interacting
// with the swarms-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewSwarmBatchService] method instead.
type SwarmBatchService struct {
	Options []option.RequestOption
}

// NewSwarmBatchService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewSwarmBatchService(opts ...option.RequestOption) (r SwarmBatchService) {
	r = SwarmBatchService{}
	r.Options = opts
	return
}

// Run a batch of swarms with the specified tasks using a thread pool.
func (r *SwarmBatchService) Run(ctx context.Context, body SwarmBatchRunParams, opts ...option.RequestOption) (res *[]SwarmBatchRunResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/swarm/batch/completions"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

type SwarmBatchRunResponse map[string]any

type SwarmBatchRunParams struct {
	Body []SwarmSpecParam
	paramObj
}

func (r SwarmBatchRunParams) MarshalJSON() (data []byte, err error) {
	return json.Marshal(r.Body)
}
func (r *SwarmBatchRunParams) UnmarshalJSON(data []byte) error {
	return json.Unmarshal(data, &r.Body)
}
