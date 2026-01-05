// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package swarms_test

import (
	"context"
	"errors"
	"os"
	"testing"

	"github.com/The-Swarm-Corporation/swarms-client-go"
	"github.com/The-Swarm-Corporation/swarms-client-go/internal/testutil"
	"github.com/The-Swarm-Corporation/swarms-client-go/option"
)

func TestClientAutoSwarmBuilderNewCompletionWithOptionalParams(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := swarms.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Client.AutoSwarmBuilder.NewCompletion(context.TODO(), swarms.ClientAutoSwarmBuilderNewCompletionParams{
		Description:   swarms.String("description"),
		ExecutionType: swarms.ClientAutoSwarmBuilderNewCompletionParamsExecutionTypeReturnAgents,
		MaxLoops:      swarms.Int(0),
		MaxTokens:     swarms.Int(0),
		ModelName:     swarms.String("model_name"),
		Name:          swarms.String("name"),
		Task:          swarms.String("task"),
	})
	if err != nil {
		var apierr *swarms.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}

func TestClientAutoSwarmBuilderListExecutionTypes(t *testing.T) {
	t.Skip("Prism tests are disabled")
	baseURL := "http://localhost:4010"
	if envURL, ok := os.LookupEnv("TEST_API_BASE_URL"); ok {
		baseURL = envURL
	}
	if !testutil.CheckTestServer(t, baseURL) {
		return
	}
	client := swarms.NewClient(
		option.WithBaseURL(baseURL),
		option.WithAPIKey("My API Key"),
	)
	_, err := client.Client.AutoSwarmBuilder.ListExecutionTypes(context.TODO())
	if err != nil {
		var apierr *swarms.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
