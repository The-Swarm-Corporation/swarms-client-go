// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package swarms

import (
	"context"
	"net/http"

	"github.com/The-Swarm-Corporation/swarms-client-go/internal/requestconfig"
	"github.com/The-Swarm-Corporation/swarms-client-go/option"
)

// ClientRateService contains methods and other services that help with interacting
// with the swarms-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClientRateService] method instead.
type ClientRateService struct {
	Options []option.RequestOption
}

// NewClientRateService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewClientRateService(opts ...option.RequestOption) (r ClientRateService) {
	r = ClientRateService{}
	r.Options = opts
	return
}

// Get the rate limits and current usage for the user associated with the provided
// API key.
func (r *ClientRateService) GetLimits(ctx context.Context, opts ...option.RequestOption) (res *ClientRateGetLimitsResponse, err error) {
	opts = append(r.Options[:], opts...)
	path := "v1/rate/limits"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ClientRateGetLimitsResponse map[string]any
