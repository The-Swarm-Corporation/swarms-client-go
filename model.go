// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package swarms

import (
	"context"
	"fmt"
	"net/http"

	"github.com/stainless-sdks/swarms-go/internal/requestconfig"
	"github.com/stainless-sdks/swarms-go/option"
)

// ModelService contains methods and other services that help with interacting with
// the swarms API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewModelService] method instead.
type ModelService struct {
	Options []option.RequestOption
}

// NewModelService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewModelService(opts ...option.RequestOption) (r ModelService) {
	r = ModelService{}
	r.Options = opts
	return
}

// Get all available models.
func (r *ModelService) ListAvailable(ctx context.Context, query ModelListAvailableParams, opts ...option.RequestOption) (res *ModelListAvailableResponse, err error) {
	if !param.IsOmitted(query.XAPIKey) {
		opts = append(opts, option.WithHeader("x-api-key", fmt.Sprintf("%s", query.XAPIKey)))
	}
	opts = append(r.Options[:], opts...)
	path := "v1/models/available"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ModelListAvailableResponse map[string]any

type ModelListAvailableParams struct {
	XAPIKey string `header:"x-api-key,required" json:"-"`
	paramObj
}
