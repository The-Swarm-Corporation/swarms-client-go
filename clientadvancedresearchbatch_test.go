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

func TestClientAdvancedResearchBatchNewCompletion(t *testing.T) {
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
	_, err := client.Client.AdvancedResearch.Batch.NewCompletion(context.TODO(), swarms.ClientAdvancedResearchBatchNewCompletionParams{
		InputSchemas: []swarms.ClientAdvancedResearchBatchNewCompletionParamsInputSchema{{
			Config: swarms.ClientAdvancedResearchBatchNewCompletionParamsInputSchemaConfig{
				Description:            swarms.String("description"),
				DirectorAgentName:      swarms.String("director_agent_name"),
				DirectorMaxLoops:       swarms.Int(0),
				DirectorMaxTokens:      swarms.Int(0),
				DirectorModelName:      swarms.String("director_model_name"),
				ExaSearchMaxCharacters: swarms.Int(0),
				ExaSearchNumResults:    swarms.Int(0),
				MaxLoops:               swarms.Int(0),
				Name:                   swarms.String("name"),
				WorkerModelName:        swarms.String("worker_model_name"),
			},
			Task: swarms.String("task"),
			Img:  swarms.String("img"),
		}},
	})
	if err != nil {
		var apierr *swarms.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
