# Zero Trust Security Model

**"Never Trust, Always Verify."**

In traditional security, things "inside" the firewall are trusted. In our design, **nothing is trusted**.

## 1. Identity over IP
We do not rely on "Internal IP" allowlists.
- **Why?** In Kubernetes, IPs are ephemeral. An attacker pod could spin up and grab a "trusted" IP.
- **Our Approach**: We treat the internal network as hostile. (Future work: mTLS would cryptographically verify pod identity).

## 2. Least Privilege (IAM & Linux Capabilities)
We operate on the principle that the container **will** be compromised eventually. The goal is to limit the blast radius.
- **Non-Root**: We run as UID 1000. If an attacker breaks out, they are a nobody on the host.
- **Drop `ALL` Capabilities**: We explicitly drop `CAP_NET_ADMIN`, `CAP_SYS_ADMIN`, etc. The process cannot modify network routes or mount file systems even if it wanted to.

## 3. Immutable Infrastructure
We set `readOnlyRootFilesystem: true`.
- **Attack Scenario**: Attacker sends a remote code execution (RCE) payload to download `crypto_miner.sh` and run it.
- **Defense**: The `wget` or `curl` might succeed to memory, but they **cannot write the file to disk** to make it executable. `chmod +x` fails. The attack is neutralized.

## 4. Layered Defense (Defense in Depth)
1.  **Layer 1 (Code)**: Go memory safety.
2.  **Layer 2 (Container)**: Distroless/Alpine image (minimal attack surface).
3.  **Layer 3 (Runtime)**: Non-root, Read-only K8s context.
4.  **Layer 4 (Network)**: VPC Flow Logs audit all traffic.
