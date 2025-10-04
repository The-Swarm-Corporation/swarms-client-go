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
	opts = slices.Concat(r.Options, opts)
	path := "v1/rate/limits"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

type ClientRateGetLimitsResponse struct {
	// The configured rate limits based on the user's subscription tier.
	Limits ClientRateGetLimitsResponseLimits `json:"limits,required"`
	// Current rate limit usage information for different time windows.
	RateLimits ClientRateGetLimitsResponseRateLimits `json:"rate_limits,required"`
	// The user's current subscription tier (free or premium).
	Tier string `json:"tier,required"`
	// ISO timestamp when the rate limits information was retrieved.
	Timestamp string `json:"timestamp,required"`
	// Indicates whether the rate limits request was successful.
	Success bool `json:"success,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Limits      respjson.Field
		RateLimits  respjson.Field
		Tier        respjson.Field
		Timestamp   respjson.Field
		Success     respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClientRateGetLimitsResponse) RawJSON() string { return r.JSON.raw }
func (r *ClientRateGetLimitsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The configured rate limits based on the user's subscription tier.
type ClientRateGetLimitsResponseLimits struct {
	// The maximum number of requests allowed per day.
	MaximumRequestsPerDay int64 `json:"maximum_requests_per_day,required"`
	// The maximum number of requests allowed per hour.
	MaximumRequestsPerHour int64 `json:"maximum_requests_per_hour,required"`
	// The maximum number of requests allowed per minute.
	MaximumRequestsPerMinute int64 `json:"maximum_requests_per_minute,required"`
	// The maximum number of tokens allowed per agent.
	TokensPerAgent int64 `json:"tokens_per_agent,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		MaximumRequestsPerDay    respjson.Field
		MaximumRequestsPerHour   respjson.Field
		MaximumRequestsPerMinute respjson.Field
		TokensPerAgent           respjson.Field
		ExtraFields              map[string]respjson.Field
		raw                      string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClientRateGetLimitsResponseLimits) RawJSON() string { return r.JSON.raw }
func (r *ClientRateGetLimitsResponseLimits) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Current rate limit usage information for different time windows.
type ClientRateGetLimitsResponseRateLimits struct {
	// Rate limit information for the last day.
	Day ClientRateGetLimitsResponseRateLimitsDay `json:"day,required"`
	// Rate limit information for the last hour.
	Hour ClientRateGetLimitsResponseRateLimitsHour `json:"hour,required"`
	// Rate limit information for the last minute.
	Minute ClientRateGetLimitsResponseRateLimitsMinute `json:"minute,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Day         respjson.Field
		Hour        respjson.Field
		Minute      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClientRateGetLimitsResponseRateLimits) RawJSON() string { return r.JSON.raw }
func (r *ClientRateGetLimitsResponseRateLimits) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Rate limit information for the last day.
type ClientRateGetLimitsResponseRateLimitsDay struct {
	// The number of requests made in this time window.
	Count int64 `json:"count,required"`
	// Whether the rate limit has been exceeded for this time window.
	Exceeded bool `json:"exceeded,required"`
	// The maximum number of requests allowed in this time window.
	Limit int64 `json:"limit,required"`
	// The number of requests remaining before hitting the limit.
	Remaining int64 `json:"remaining,required"`
	// ISO timestamp when the rate limit will reset.
	ResetTime string `json:"reset_time,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Exceeded    respjson.Field
		Limit       respjson.Field
		Remaining   respjson.Field
		ResetTime   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClientRateGetLimitsResponseRateLimitsDay) RawJSON() string { return r.JSON.raw }
func (r *ClientRateGetLimitsResponseRateLimitsDay) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Rate limit information for the last hour.
type ClientRateGetLimitsResponseRateLimitsHour struct {
	// The number of requests made in this time window.
	Count int64 `json:"count,required"`
	// Whether the rate limit has been exceeded for this time window.
	Exceeded bool `json:"exceeded,required"`
	// The maximum number of requests allowed in this time window.
	Limit int64 `json:"limit,required"`
	// The number of requests remaining before hitting the limit.
	Remaining int64 `json:"remaining,required"`
	// ISO timestamp when the rate limit will reset.
	ResetTime string `json:"reset_time,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Exceeded    respjson.Field
		Limit       respjson.Field
		Remaining   respjson.Field
		ResetTime   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClientRateGetLimitsResponseRateLimitsHour) RawJSON() string { return r.JSON.raw }
func (r *ClientRateGetLimitsResponseRateLimitsHour) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Rate limit information for the last minute.
type ClientRateGetLimitsResponseRateLimitsMinute struct {
	// The number of requests made in this time window.
	Count int64 `json:"count,required"`
	// Whether the rate limit has been exceeded for this time window.
	Exceeded bool `json:"exceeded,required"`
	// The maximum number of requests allowed in this time window.
	Limit int64 `json:"limit,required"`
	// The number of requests remaining before hitting the limit.
	Remaining int64 `json:"remaining,required"`
	// ISO timestamp when the rate limit will reset.
	ResetTime string `json:"reset_time,required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Count       respjson.Field
		Exceeded    respjson.Field
		Limit       respjson.Field
		Remaining   respjson.Field
		ResetTime   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r ClientRateGetLimitsResponseRateLimitsMinute) RawJSON() string { return r.JSON.raw }
func (r *ClientRateGetLimitsResponseRateLimitsMinute) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}
