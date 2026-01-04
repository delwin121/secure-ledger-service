# Observability Model

"You cannot fix what you cannot see."

## Why Prometheus?
We chose **Prometheus** over traditional log parsing for metrics because:
1.  **Pull Model**: Prometheus scrapes our service, meaning our service doesn't need to know where the monitoring server protects it from being overloaded.
2.  **Dimensionality**: We can slice data by `status_code`, `method`, or `pod_name`.
3.  **Standard**: It is the cloud-native standard for K8s monitoring.

## Why P99 (99th Percentile) vs Average?
**Average latency is a lie.**
If 100 requests take:
- 99 requests: 10ms
- 1 request: 10,000ms (10s)

The **Average** is ~110ms. This looks "fine".
The **P99** is 10,000ms. This reveals the critical problem.

We monitor **P95** and **P99** to ensure the system performs well for the *slowest* queries, which are often the most important users (or the ones about to churn).

## The Stack
1.  **Metrics Injection**: `MetricsMiddleware` in Go wraps every handler, recording `start_time` and `duration` before the request finishes.
2.  **Exposition**: The `/metrics` endpoint exposes these in Prometheus format.
3.  **Collection**: A Prometheus container scrapes `http://host.docker.internal:8080/metrics` every 5s.
4.  **Visualization**: Grafana queries Prometheus to draw the graphs.
