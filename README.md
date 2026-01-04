# Secure Ledger Service: A Cloud-Native Case Study

![Go](https://img.shields.io/badge/Language-Go_1.25-00ADD8)
![Kubernetes](https://img.shields.io/badge/Orchestrator-Kubernetes-326CE5)
![Security](https://img.shields.io/badge/Security-Zero_Trust-green)
![Status](https://img.shields.io/badge/Status-Production_Ready-success)

## 1. Executive Summary
The **Secure Ledger Service** is a high-performance, distributed transaction processing system designed to demonstrate **Zero Trust security principles** and **observability-first engineering**. Unlike typical CRUD applications, this project focuses on the *architectural rigors* of running financial software in a hostile cloud environment: resilience to failure, immunity to common attack vectors, and deep visibility into system behavior.

**Key Achievements**:
- **Zero Trust**: Implements immutable infrastructure with read-only filesystems and dropped capabilities.
- **Chaos Resilient**: Self-healing architecture that survives pod failures and network latency.
- **Auditable**: Infrastructure as Code (Terraform) with automated VPC Flow Logging.

---

## 2. Architecture Overview
The system follows a microservices architecture hosted on Kubernetes.

| Component | Responsibility | Design Choice |
| :--- | :--- | :--- |
| **Ledger API** | Processes transactions | **Go (Golang)** for concurrency and type safety. |
| **Orchestrator** | Manages lifecycle/scaling | **Kubernetes** for self-healing and topology awareness. |
| **Data Store** | State management | **In-Memory (Mutex)** for nanosecond latency (MVP). |
| **Observability** | Metrics & Monitoring | **Prometheus/Grafana** for P99 latency tracking. |
| **Network** | Cloud Infrastructure | **Terraform** for reproducible AWS VPCs. |

👉 **[View Full Architecture Diagram](docs/ARCHITECTURE.md)**

---

## 3. Security Model (Zero Trust)
We assume the network is compromised. Security is enforced at the pod level.

- **Identity**: No reliance on IP allowlists.
- **Least Privilege**: Application runs as `UID 1000` (non-root).
- **Attack Surface**: Distroless-style Alpine image with **Read-Only Root Filesystem**.

👉 **[View Threat Model](docs/THREAT_MODEL.md)** | 👉 **[Zero Trust Defense](docs/zero-trust.md)**

---

## 4. Observability & Failure Analysis
The system instrumented with **Golden Signal** metrics.

- **Metric**: `http_request_duration_seconds` (Histogram)
- **Aggregation**: P99 (99th Percentile) to catch tail latency.
- **Chaos Testing**: Confirmed resilience by killing pods and injecting 500ms latency.

👉 **[Failure Analysis Report](docs/failure-analysis.md)** | 👉 **[Observability Deep Dive](docs/observability.md)**

---

## 5. Quick Start (Demo)

Want to see it in action? Follow the **Golden Path** demo script.

```bash
# 1. Clone & Start
minikube start
kubectl apply -f k8s/

# 2. Test
./scripts/generate_traffic.sh
```

👉 **[Run the Full Demo Script](docs/demo.md)**

---

## 6. Known Limitations (Trade-offs)
- **Persistence**: Ledger is in-memory. Restarting pods clears balances. (Decision: Simplicity for Architecture demo).
- **Auth**: No JWT/OAuth implementation (Decision: Focus on Infrastructure/Security).

👉 **[Read Design Decisions](docs/design-decisions.md)**

---

## 7. Future Work
- [ ] Migrate storage to PostgreSQL (AWS RDS).
- [ ] Implement Mutual TLS (mTLS) with Istio.
- [ ] Add Structured Logging (ELK Stack).

---
*Created by Delwin*
