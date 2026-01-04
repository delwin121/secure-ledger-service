# Interview / Viva Preparation Guide

This document contains "Model Answers" for questions you might be asked during a viva or DevOps interview about this project.

## Q1: "Walk me through a transaction request."
**Answer**:
1.  "The user sends a JSON POST request to port 80."
2.  "The Kubernetes **Service** loads balances this request to one of the 3 available Pods."
3.  "The request hits the Go HTTP server. It first passes through the **Chaos Middleware** (checking for latency injection), then the **Metrics Middleware** (recording start time)."
4.  "The **Handler** validates the input (e.g., non-negative amount)."
5.  "A Mutex lock is acquired on the **Store** to ensure thread safety."
6.  "The balance is updated in memory, the lock is released, and a 200 OK is returned."
7.  "Finally, the Metrics Middleware calculates the total duration and updates the Prometheus counter."

## Q2: "What happens if a node goes down?"
**Answer**:
"I configured **Topology Spread Constraints** in the deployment. This tells Kubernetes to try to run the 3 replicas on different nodes. If one node fails, we lose only 1 pod. The K8s Controller Manager detects this and immediately schedules a replacement pod on a healthy node to return to the desired state of 3."

## Q3: "Why did you use 'readOnlyRootFilesystem'?"
**Answer**:
"It's a Zero Trust principle. If an attacker manages to exploit a vulnerability in my code, their first move is usually to download a script (like a reverse shell or crypto miner) and make it executable. By making the filesystem read-only, I physically prevent them from writing any payload to disk. It neutralizes the persistence of the attack."

## Q4: "How does your system detect performance issues?"
**Answer**:
"I don't rely on users complaining. I use **Prometheus** to scrape golden signals (Latency, Errors, Traffic) every 5 seconds. I specifically monitor the **99th Percentile (P99)** latency in Grafana. This alerted me immediately when I injected a 500ms delay during my Chaos Engineering tests, whereas an 'Average' metric might have hidden the spike."

## Q5: "What would you change for production?"
**Answer**:
1.  **Persistence**: Replace the in-memory map with a clustered PostgreSQL database (Amazon Aurora).
2.  **Authentication**: Add an OIDC / OAuth2 middleware to verify user tokens (JWTs).
3.  **Secrets**: Move configuration from plain ConfigMaps to AWS Secrets Manager or HashiCorp Vault.
