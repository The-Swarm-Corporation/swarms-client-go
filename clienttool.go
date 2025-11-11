// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package swarms

import (
	"context"
	"net/http"
	"slices"

	"github.com/The-Swarm-Corporation/swarms-client-go/internal/apijson"
	"github.com/The-Swarm-Corporation/swarms-client-go/internal/requestconfig"
	"github.com/The-Swarm-Corporation/swarms-client-go/option"
	"github.com/The-Swarm-Corporation/swarms-client-go/packages/respjson"
)

// ClientToolService contains methods and other services that help with interacting
// with the swarms-client API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewClientToolService] method instead.
type ClientToolService struct {
	Options []option.RequestOption
}

// NewClientToolService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewClientToolService(opts ...option.RequestOption) (r ClientToolService) {
	r = ClientToolService{}
	r.Options = opts
	return
}

// Retrieve comprehensive information about all available tools and capabilities
// supported by the Swarms API.
func (r *ClientToolService) ListAvailable(ctx context.Context, opts ...option.RequestOption) (res *ClientToolListAvailableResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "v1/tools/available"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ClientToolListAvailableResponse struct {
	// The status of the available tools.
	Status string `json:"status,nullable"`
	// The list of available tools.
	Tools []string `json:"tools,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Status      respjson.Field
		Tools       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClientToolListAvailableResponse) RawJSON() string { return r.JSON.raw }
func (r *ClientToolListAvailableResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
