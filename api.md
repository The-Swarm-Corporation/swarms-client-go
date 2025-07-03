# swarms

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go">swarms</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go#GetRootResponse">GetRootResponse</a>

Methods:

- <code title="get /">client.<a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go#SwarmService.GetRoot">GetRoot</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go">swarms</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go#GetRootResponse">GetRootResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Health

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go">swarms</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go#HealthCheckResponse">HealthCheckResponse</a>

Methods:

- <code title="get /health">client.Health.<a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go#HealthService.Check">Check</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go">swarms</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go#HealthCheckResponse">HealthCheckResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Agent

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go">swarms</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go#AgentCompletionParam">AgentCompletionParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go">swarms</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go#AgentSpecParam">AgentSpecParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go">swarms</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go#AgentRunResponse">AgentRunResponse</a>

Methods:

- <code title="post /v1/agent/completions">client.Agent.<a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go#AgentService.Run">Run</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go">swarms</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go#AgentRunParams">AgentRunParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go">swarms</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go#AgentRunResponse">AgentRunResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Batch

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go">swarms</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go#AgentBatchRunResponse">AgentBatchRunResponse</a>

Methods:

- <code title="post /v1/agent/batch/completions">client.Agent.Batch.<a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go#AgentBatchService.Run">Run</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go">swarms</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go#AgentBatchRunParams">AgentBatchRunParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go">swarms</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go#AgentBatchRunResponse">AgentBatchRunResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Models

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go">swarms</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go#ModelListAvailableResponse">ModelListAvailableResponse</a>

Methods:

- <code title="get /v1/models/available">client.Models.<a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go#ModelService.ListAvailable">ListAvailable</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go">swarms</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go#ModelListAvailableResponse">ModelListAvailableResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Swarms

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go">swarms</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go#SwarmSpecParam">SwarmSpecParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go">swarms</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go#SwarmCheckAvailableResponse">SwarmCheckAvailableResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go">swarms</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go#SwarmGetLogsResponse">SwarmGetLogsResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go">swarms</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go#SwarmRunResponse">SwarmRunResponse</a>

Methods:

- <code title="get /v1/swarms/available">client.Swarms.<a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go#SwarmService.CheckAvailable">CheckAvailable</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go">swarms</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go#SwarmCheckAvailableResponse">SwarmCheckAvailableResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /v1/swarm/logs">client.Swarms.<a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go#SwarmService.GetLogs">GetLogs</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go">swarms</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go#SwarmGetLogsResponse">SwarmGetLogsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /v1/swarm/completions">client.Swarms.<a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go#SwarmService.Run">Run</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go">swarms</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go#SwarmRunParams">SwarmRunParams</a>) (<a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go">swarms</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go#SwarmRunResponse">SwarmRunResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Batch

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go">swarms</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go#SwarmBatchRunResponse">SwarmBatchRunResponse</a>

Methods:

- <code title="post /v1/swarm/batch/completions">client.Swarms.Batch.<a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go#SwarmBatchService.Run">Run</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go">swarms</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go#SwarmBatchRunParams">SwarmBatchRunParams</a>) ([]<a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go">swarms</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/swarms-go#SwarmBatchRunResponse">SwarmBatchRunResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
