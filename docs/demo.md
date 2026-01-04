# Golden Path Demo Script

This document outlines the repeatable demo sequence to validte the Ledger Service's core functionality, resilience, and observability.

**Pre-requisites**:
- Running Kubernetes Cluster (Minikube).
- `kubectl` configured.
- `docker` installed.

## 1. Start the Cluster & Deploy
```bash
# 1. Start Minikube
minikube start

# 2. Build & Load Image (If not already done)
docker build -t ledger-service:latest -f docker/Dockerfile app/
minikube image load ledger-service:latest

# 3. Deploy Manifests
kubectl apply -f k8s/configmap.yaml
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml
```

## 2. Establish Observability (The "Eyes")
We need to see what's happening.
```bash
# Start Prometheus & Grafana (in separate terminal)
./observability/start.sh

# Forward Port for App (in another terminal)
# NOTE: "--address 0.0.0.0" is CRITICAL for Docker to see Minikube
sudo kubectl port-forward --address 0.0.0.0 svc/ledger-service 8080:80
```
- Open Grafana: http://localhost:3000 (admin/admin).

## 3. The "Happy Path" (Transactions)
Generate normal traffic to show the system works.
```bash
# Send 10 transactions
for i in {1..10}; do
  curl -X POST http://localhost:8080/transaction \
  -d '{"account_id":"user_demo","amount":100}'
done
```
**Observation**:
- Check Grafana "Request Rate". It should show a small bump.
- Check "95th Percentile Latency". It should be low (<10ms).

## 4. The "Chaos" (Kill a Pod)
Demonstrate Self-Healing.
```bash
# 1. List pods
kubectl get pods

# 2. Delete one
kubectl delete pod <pod-name-from-above>

# 3. Watch it happen
kubectl get pods -w
```
**Observation**:
- `kubectl` immediately shows the pod `Terminating`.
- Almost instantly, a `ContainerCreating` appears (ReplicaSet controller at work).
- **Traffic**: If you run the curl loop *during* the delete, you might see 0 errors because K8s reroutes to the other 2 healthy pods.

## 5. The "Performance Crisis" (Latency Injection)
Demonstrate Observability.
```bash
# 1. Update ConfigMap to inject 500ms latency
# Edit k8s/configmap.yaml -> Change SIMULATE_DELAY_MS to "500"
kubectl apply -f k8s/configmap.yaml

# 2. Restart deployment to pick up config (Zero Downtime)
kubectl rollout restart deploy/ledger-app

# 3. Generate Traffic
for i in {1..20}; do
  curl -X POST http://localhost:8080/transaction \
  -d '{"account_id":"load","amount":1}'
done
```
**Observation**:
- **Grafana**: The "95th Percentile Latency" graph will spike to **0.5s (500ms)**.
- This proves we can detect "slowness" not just "crashes".

## 6. Show Security Controls
Prove the system is hardened.
```bash
# Try to write a file to the root system (Should FAIL)
kubectl exec -it deploy/ledger-app -- touch /root/malware.sh
# Result: "touch: /root/malware.sh: Read-only file system"
```
