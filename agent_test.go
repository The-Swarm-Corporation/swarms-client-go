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

func TestAgentRunWithOptionalParams(t *testing.T) {
	t.Skip("skipped: tests are disabled for the time being")
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
	_, err := client.Agent.Run(context.TODO(), swarms.AgentRunParams{
		AgentCompletion: swarms.AgentCompletionParam{
			AgentConfig: swarms.AgentSpecParam{
				AgentName:          swarms.String("agent_name"),
				AutoGeneratePrompt: swarms.Bool(true),
				Description:        swarms.String("description"),
				MaxLoops:           swarms.Int(0),
				MaxTokens:          swarms.Int(0),
				McpURL:             swarms.String("mcp_url"),
				ModelName:          swarms.String("model_name"),
				Role:               swarms.String("role"),
				StreamingOn:        swarms.Bool(true),
				SystemPrompt:       swarms.String("system_prompt"),
				Temperature:        swarms.Float(0),
				ToolsListDictionary: []map[string]any{{
					"foo": "bar",
				}},
			},
			History: swarms.AgentCompletionHistoryUnionParam{
				OfAnyMap: map[string]any{
					"foo": "bar",
				},
			},
			Img:    swarms.String("img"),
			Imgs:   []string{"string"},
			Stream: swarms.Bool(true),
			Task:   swarms.String("task"),
		},
	})
	if err != nil {
		var apierr *swarms.Error
		if errors.As(err, &apierr) {
			t.Log(string(apierr.DumpRequest(true)))
		}
		t.Fatalf("err should be nil: %s", err.Error())
	}
}
