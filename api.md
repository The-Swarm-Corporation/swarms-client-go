# swarms

Response Types:

- <a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#GetRootResponse">GetRootResponse</a>

Methods:

- <code title="get /">client.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#SwarmService.GetRoot">GetRoot</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#GetRootResponse">GetRootResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Health

Response Types:

- <a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#HealthCheckResponse">HealthCheckResponse</a>

Methods:

- <code title="get /health">client.Health.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#HealthService.Check">Check</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#HealthCheckResponse">HealthCheckResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Agent

Params Types:

- <a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#AgentCompletionParam">AgentCompletionParam</a>
- <a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#AgentSpecParam">AgentSpecParam</a>
- <a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#McpConnectionParam">McpConnectionParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#AgentListResponse">AgentListResponse</a>
- <a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#AgentRunResponse">AgentRunResponse</a>

Methods:

- <code title="get /v1/agents/list">client.Agent.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#AgentService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#AgentListResponse">AgentListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/agent/completions">client.Agent.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#AgentService.Run">Run</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#AgentRunParams">AgentRunParams</a>) (\*<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#AgentRunResponse">AgentRunResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Batch

Response Types:

- <a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#AgentBatchRunResponse">AgentBatchRunResponse</a>

Methods:

- <code title="post /v1/agent/batch/completions">client.Agent.Batch.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#AgentBatchService.Run">Run</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#AgentBatchRunParams">AgentBatchRunParams</a>) (\*<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#AgentBatchRunResponse">AgentBatchRunResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Models

Response Types:

- <a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#ModelListAvailableResponse">ModelListAvailableResponse</a>

Methods:

- <code title="get /v1/models/available">client.Models.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#ModelService.ListAvailable">ListAvailable</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#ModelListAvailableResponse">ModelListAvailableResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Swarms

Params Types:

- <a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#SwarmSpecParam">SwarmSpecParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#SwarmCheckAvailableResponse">SwarmCheckAvailableResponse</a>
- <a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#SwarmGetLogsResponse">SwarmGetLogsResponse</a>
- <a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#SwarmRunResponse">SwarmRunResponse</a>

Methods:

- <code title="get /v1/swarms/available">client.Swarms.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#SwarmService.CheckAvailable">CheckAvailable</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#SwarmCheckAvailableResponse">SwarmCheckAvailableResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/swarm/logs">client.Swarms.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#SwarmService.GetLogs">GetLogs</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#SwarmGetLogsResponse">SwarmGetLogsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/swarm/completions">client.Swarms.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#SwarmService.Run">Run</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#SwarmRunParams">SwarmRunParams</a>) (\*<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#SwarmRunResponse">SwarmRunResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Batch

Response Types:

- <a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#SwarmBatchRunResponse">SwarmBatchRunResponse</a>

Methods:

- <code title="post /v1/swarm/batch/completions">client.Swarms.Batch.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#SwarmBatchService.Run">Run</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#SwarmBatchRunParams">SwarmBatchRunParams</a>) (\*[]<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#SwarmBatchRunResponse">SwarmBatchRunResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# ReasoningAgents

Response Types:

- <a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#ReasoningAgentNewCompletionResponse">ReasoningAgentNewCompletionResponse</a>
- <a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#ReasoningAgentListTypesResponse">ReasoningAgentListTypesResponse</a>

Methods:

- <code title="post /v1/reasoning-agent/completions">client.ReasoningAgents.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#ReasoningAgentService.NewCompletion">NewCompletion</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#ReasoningAgentNewCompletionParams">ReasoningAgentNewCompletionParams</a>) (\*<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#ReasoningAgentNewCompletionResponse">ReasoningAgentNewCompletionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/reasoning-agent/types">client.ReasoningAgents.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#ReasoningAgentService.ListTypes">ListTypes</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#ReasoningAgentListTypesResponse">ReasoningAgentListTypesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Client

## Rate

Response Types:

- <a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#RateLimitWindow">RateLimitWindow</a>
- <a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#ClientRateGetLimitsResponse">ClientRateGetLimitsResponse</a>

Methods:

- <code title="get /v1/rate/limits">client.Client.Rate.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#ClientRateService.GetLimits">GetLimits</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#ClientRateGetLimitsResponse">ClientRateGetLimitsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## AutoSwarmBuilder

Response Types:

- <a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#ClientAutoSwarmBuilderNewCompletionResponse">ClientAutoSwarmBuilderNewCompletionResponse</a>

Methods:

- <code title="post /v1/auto-swarm-builder/completions">client.Client.AutoSwarmBuilder.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#ClientAutoSwarmBuilderService.NewCompletion">NewCompletion</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#ClientAutoSwarmBuilderNewCompletionParams">ClientAutoSwarmBuilderNewCompletionParams</a>) (\*<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#ClientAutoSwarmBuilderNewCompletionResponse">ClientAutoSwarmBuilderNewCompletionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/auto-swarm-builder/execution-types">client.Client.AutoSwarmBuilder.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#ClientAutoSwarmBuilderService.ListExecutionTypes">ListExecutionTypes</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*[]<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## AdvancedResearch

Response Types:

- <a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#ClientAdvancedResearchNewCompletionResponse">ClientAdvancedResearchNewCompletionResponse</a>

Methods:

- <code title="post /v1/advanced-research/completions">client.Client.AdvancedResearch.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#ClientAdvancedResearchService.NewCompletion">NewCompletion</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#ClientAdvancedResearchNewCompletionParams">ClientAdvancedResearchNewCompletionParams</a>) (\*<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#ClientAdvancedResearchNewCompletionResponse">ClientAdvancedResearchNewCompletionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

### Batch

Response Types:

- <a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#ClientAdvancedResearchBatchNewCompletionResponse">ClientAdvancedResearchBatchNewCompletionResponse</a>

Methods:

- <code title="post /v1/advanced-research/batch/completions">client.Client.AdvancedResearch.Batch.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#ClientAdvancedResearchBatchService.NewCompletion">NewCompletion</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#ClientAdvancedResearchBatchNewCompletionParams">ClientAdvancedResearchBatchNewCompletionParams</a>) (\*[]<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#ClientAdvancedResearchBatchNewCompletionResponse">ClientAdvancedResearchBatchNewCompletionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Tools

Response Types:

- <a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#ClientToolListAvailableResponse">ClientToolListAvailableResponse</a>

Methods:

- <code title="get /v1/tools/available">client.Client.Tools.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#ClientToolService.ListAvailable">ListAvailable</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#ClientToolListAvailableResponse">ClientToolListAvailableResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Marketplace

Response Types:

- <a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#ClientMarketplaceNewAgentResponse">ClientMarketplaceNewAgentResponse</a>

Methods:

- <code title="post /v1/marketplace/agents">client.Client.Marketplace.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#ClientMarketplaceService.NewAgent">NewAgent</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#ClientMarketplaceNewAgentParams">ClientMarketplaceNewAgentParams</a>) (\*<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#ClientMarketplaceNewAgentResponse">ClientMarketplaceNewAgentResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## BatchedGridWorkflow

Response Types:

- <a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#ClientBatchedGridWorkflowCompleteWorkflowResponse">ClientBatchedGridWorkflowCompleteWorkflowResponse</a>

Methods:

- <code title="post /v1/batched-grid-workflow/completions">client.Client.BatchedGridWorkflow.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#ClientBatchedGridWorkflowService.CompleteWorkflow">CompleteWorkflow</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#ClientBatchedGridWorkflowCompleteWorkflowParams">ClientBatchedGridWorkflowCompleteWorkflowParams</a>) (\*<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#ClientBatchedGridWorkflowCompleteWorkflowResponse">ClientBatchedGridWorkflowCompleteWorkflowResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## GraphWorkflow

Response Types:

- <a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#ClientGraphWorkflowExecuteWorkflowResponse">ClientGraphWorkflowExecuteWorkflowResponse</a>

Methods:

- <code title="post /v1/graph-workflow/completions">client.Client.GraphWorkflow.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#ClientGraphWorkflowService.ExecuteWorkflow">ExecuteWorkflow</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#ClientGraphWorkflowExecuteWorkflowParams">ClientGraphWorkflowExecuteWorkflowParams</a>) (\*<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go">swarms</a>.<a href="https://pkg.go.dev/github.com/The-Swarm-Corporation/swarms-client-go#ClientGraphWorkflowExecuteWorkflowResponse">ClientGraphWorkflowExecuteWorkflowResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
