# Architectural Design Decisions

Every line of code represents a choice. Here is the defense of those choices.

## 1. Why Go?
We chose **Go (Golang)** over Python/Node.js for the Ledger Service.
- **Concurrency**: The `goroutine` model allows us to handle thousands of concurrent transactions with minimal memory footprint compared to OS threads.
- **Static Typing**: Financial systems require correctness. Compile-time checks prevent entire classes of type-related bugs.
- **Performance**: Go compiles to machine code. Lower latency means faster locking and unlocking of the ledger mutex.

## 2. Why Kubernetes (vs Docker Compose)?
- **Self-Healing**: Docker run stops if the process dies. K8s restarts it.
- **Rolling Updates**: K8s allows us to release new versions (v2) without dropping a single connection using `RollingUpdate` strategy.
- **Ecosystem**: K8s is the "Operating System of the Cloud". Using it demonstrates readiness for AWS EKS / Google GKE.

## 3. Trade-off: In-Memory Store (The "Elephant in the Room")
**Decision**: We built an in-memory `map[string]float64` instead of using PostgreSQL.
**Defense**:
- **MVP Focus**: The goal was to demonstrate *Cloud Engineering* (CI/CD, K8s, Security, Observability), not *Database Administration*.
- **Performance**: In-memory is the fastest possible storage (nanosecond access).
- **Honesty**: We acknowledge this is not "Durable". A production version would replace the `Store` interface implementation with `PostgresStore`.

## 4. Why Stateless?
The application logic contains no state. The state resides in the Store.
This allows us to scale from 1 pod to 100 pods instantly. If the app held state (e.g., sticky sessions), scaling would be exponentially harder.
