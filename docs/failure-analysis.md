# Failure Analysis

We test against failure because **hope is not a strategy**. This document records the behavior of the system under various stress conditions.

## Recovery Scenarios

| Scenario | Action | Expected Outcome | Observed Result | Severity |
| :--- | :--- | :--- | :--- | :--- |
| **Kill One Pod** | `kubectl delete pod <name>` | Kubernetes starts a replacement immediately. Service availability > 99%. | ✅ Traffic rerouted to remaining 2 pods. Zero user-facing errors. | Low |
| **Kill Multiple Pods** | `kubectl delete pod <name1> <name2>` | Throughput capacity drops, but service remains up if 1 pod survives. | ✅ Latency increased slightly. New pods ready in <5s. | Medium |
| **Config Drift** | Change env var without rollout | App ignores change until restart. | ✅ ConfigMap updated, but Pods kept old config until `rollout restart`. | Low |

## Resilience Mechanisms

### 1. ReplicaSets
We run **3 replicas** of the application. The Kubernetes ReplicaSet Controller ensures that if the current state (e.g., 2 pods) doesn't match the desired state (3 pods), it immediately schedules a new one.

### 2. Topology Spread Constraints
```yaml
topologySpreadConstraints:
  - maxSkew: 1
    topologyKey: kubernetes.io/hostname
```
This ensures that if we were on a multi-node cluster (like AWS EKS), our pods would be spread across different physical machines (Nodes). If one Node dies, we only lose 1/3 of our capacity, not 100%.

### 3. Statelessness
The application holds no state on disk. This means any pod can be killed and replaced without data corruption or recovery procedures. *Note: For MVP, data is in memory, so it IS lost, but in a real DB architecture, this holds true.*
