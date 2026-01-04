# Threat Model: Secure Ledger Service

**Methodology**: STRIDE (Spoofing, Tampering, Repudiation, Information Disclosure, Denial of Service, Elevation of Privilege)

## 1. STRIDE Analysis

| Threat Category | Potential Risk | Mitigation in Place | Status |
| :--- | :--- | :--- | :--- |
| **Spoofing** | Attacker impersonates a legitimate client. | **Future Work**: No Authentication in MVP. Real system needs mTLS/OIDC. | ⚠️ Accepted Risk (MVP) |
| **Tampering** | Attacker modifies the ledger binary or config. | **Immutable Infrastructure**: Read-only filesystem. Non-root user. Checksummed Go modules. | ✅ Mitigated |
| **Repudiation** | User denies sending a transaction. | **Audit Logs**: Structured JSON logs record every request with timestamps. | ✅ Mitigated |
| **Information Disclosure** | Leakage of internal metrics or logs. | **Logging Hygiene**: No sensitive data (PII/Secrets) in logs. Minimal Docker image. | ✅ Mitigated |
| **Denial of Service** | Flooding the service to crash it. | **Scaling**: 3 Replicas + Topology Spread. **Resource Limits**: (Future: Pod Quotas). | ⚠️ Partial |
| **Elevation of Privilege** | Attacker escapes container to Host. | **Hardening**: `runAsNonRoot: true`, `readOnlyRootFilesystem: true`, `capabilities: drop: ["ALL"]`. | ✅ Mitigated |

## 2. Attack Vectors & Defenses

### Vector: Container Breakout
**Scenario**: An attacker finds a vulnerability in the Go runtime to execute code.
**Defense Layer**:
1.  **User Isolation**: Process runs as UID 1000, not Root. Cannot map to Host Root.
2.  **Capability Dropping**: Even if they escalate inside, they have no Linux capabilities (like `CAP_NET_ADMIN`).
3.  **Filesystem Lock**: They cannot download or write a exploit script because `/` is Read-Only.

### Vector: Supply Chain Attack
**Scenario**: A malicious dependency is added to `go.mod`.
**Defense Layer**:
1.  **Deterministic Builds**: `go.sum` ensures checksum of modules.
2.  **Image Scanning**: Trivy scans the final image for known CVEs in the pipeline.
3.  **Minimal Base**: Using `alpine` reduces the number of potential vulnerable packages hiding in the system.

### Vector: Insider Threat (Cloud)
**Scenario**: A rogue admin deletes logs to cover tracks.
**Defense Layer**:
1.  **VPC Flow Logs**: Network traffic is logged to CloudWatch (Immutable storage with retention policies).
2.  **IaC Review**: Terraform changes require code review (Pull Request) via the CI pipeline.

## 3. Residual Risks (MVP)

1.  **Lack of Authentication**: Any user who can reach the API can post transactions.
    - *Fix Plan*: Implement JWT Middleware.
2.  **In-Memory Data Loss**: If all pods crash, the ledger is empty.
    - *Fix Plan*: Connect to AWS RDS (PostgreSQL).
3.  **Unencrypted In-Transit**: HTTP is used between Ingress and Pods.
    - *Fix Plan*: Implement TLS (or Service Mesh mTLS).

---
*This model represents the security posture of Release v1.0.0*
