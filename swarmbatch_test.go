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

func TestSwarmBatchRun(t *testing.T) {
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
	_, err := client.Swarms.Batch.Run(context.TODO(), swarms.SwarmBatchRunParams{
		Body: []swarms.SwarmSpecParam{{
			Agents: []swarms.AgentSpecParam{{
				AgentName:                 swarms.String("agent_name"),
				AutoGeneratePrompt:        swarms.Bool(true),
				Description:               swarms.String("description"),
				DynamicTemperatureEnabled: swarms.Bool(true),
				LlmArgs: map[string]any{
					"foo": "bar",
				},
				MaxLoops:     swarms.Int(0),
				MaxTokens:    swarms.Int(0),
				McpURL:       swarms.String("mcp_url"),
				ModelName:    swarms.String("model_name"),
				Role:         swarms.String("role"),
				StreamingOn:  swarms.Bool(true),
				SystemPrompt: swarms.String("system_prompt"),
				Temperature:  swarms.Float(0),
				ToolsListDictionary: []map[string]any{{
					"foo": "bar",
				}},
			}},
			Description:                      swarms.String("description"),
			HeavySwarmLoopsPerAgent:          swarms.Int(0),
			HeavySwarmQuestionAgentModelName: swarms.String("heavy_swarm_question_agent_model_name"),
			HeavySwarmWorkerModelName:        swarms.String("heavy_swarm_worker_model_name"),
			Img:                              swarms.String("img"),
			MaxLoops:                         swarms.Int(0),
			Messages: swarms.SwarmSpecMessagesUnionParam{
				OfMapOfAnyMap: []map[string]any{{
					"foo": "bar",
				}},
			},
			Name:          swarms.String("name"),
			RearrangeFlow: swarms.String("rearrange_flow"),
			Rules:         swarms.String("rules"),
			ServiceTier:   swarms.String("service_tier"),
			Stream:        swarms.Bool(true),
			SwarmType:     swarms.SwarmSpecSwarmTypeAgentRearrange,
			Task:          swarms.String("task"),
			Tasks:         []string{"string"},
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
