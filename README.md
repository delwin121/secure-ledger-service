# Secure Ledger Service: DevOps & Security Case Study

> A production-grade, secure-by-default Go microservice demonstrating modern Cloud-Native engineering practices.

![Go](https://img.shields.io/badge/Go-1.25-blue)
![Kubernetes](https://img.shields.io/badge/Kubernetes-Ready-blue)
![Terraform](https://img.shields.io/badge/Terraform-Infrastructure-purple)
![Security](https://img.shields.io/badge/Security-Hardened-green)

## 📖 Overview
This project is more than just a Ledger Application. It is a reference architecture for:
- **Zero Trust Security** (Non-root, Read-only FS, No capabilities).
- **Observability** (Prometheus Metrics, Structured Logging).
- **Chaos Engineering** (Latency Injection, Pod Failure Recovery).
- **Infrastructure as Code** (Terraform, Kubernetes Manifests).

## 🚀 Quick Start (Run it in 5 Minutes)

### Prerequisites
- Docker
- Minikube
- Kubectl

### Step 1: Clone & Build
```bash
git clone <repo-url>
cd cloud_project

# Build the secure Docker image
docker build -t ledger-service:latest -f docker/Dockerfile app/
```

### Step 2: Deploy to Minikube
```bash
# Start Minikube
minikube start

# Load image into Minikube
minikube image load ledger-service:latest

# Deploy Configuration and App
kubectl apply -f k8s/configmap.yaml
kubectl apply -f k8s/deployment.yaml
kubectl apply -f k8s/service.yaml
```

### Step 3: Verify It Works
```bash
# Check if pods are running
kubectl get pods

# Forward port to access locally
kubectl port-forward svc/ledger-service 8080:80
```
*In another terminal:*
```bash
# Create a transaction
curl -X POST http://localhost:8080/transaction \
  -d '{"account_id":"user123","amount":500.00}'

# You should see: {"status":"success"}
```

## 🛠️ Engineering Highlights

### 1. Security First
We utilize a **distroless-style** approach with Alpine, running as a non-privileged user (UID 1000).
- **View Policy**: [docs/THREAT_MODEL.md](docs/THREAT_MODEL.md)
- **View Dockerfile**: [docker/Dockerfile](docker/Dockerfile)

### 2. Failure is Expected
The system is designed to survive failure.
- **Try it**: Delete a pod with `kubectl delete pod <name>`. Kubernetes resurrects it instantly.
- **Latency Testing**: The system has a built-in Chaos Middleware to simulate network lag.

### 3. Observable by Design
Metrics are exposed at `/metrics` for Prometheus.
- **View Dashboard**: [docs/ARCHITECTURE.md](docs/ARCHITECTURE.md)

## 📂 Project Structure
```
├── app/            # Go Application Code
├── docker/         # Dockerfile (Multi-stage, Secure)
├── k8s/            # Kubernetes Manifests (Deployment, Service, ConfigMap)
├── terraform/      # AWS Infrastructure (VPC, Flow Logs)
├── docs/           # Architecture & Security Documentation
└── .github/        # CI/CD Pipelines (Trivy Security Gates)
```

## 🤝 Contributing
This project enforces **Policy-as-Code**. All PRs must pass:
1.  Trivy IaC Scan (Terraform/K8s)
2.  Trivy Image Scan (No Critical CVEs)
