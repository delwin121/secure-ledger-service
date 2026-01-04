#!/bin/bash
set -e

# 1. Start Prometheus
echo "Starting Prometheus..."
docker run -d \
  --name prometheus \
  -p 9090:9090 \
  -v $(pwd)/observability/prometheus.yaml:/etc/prometheus/prometheus.yml \
  --add-host host.docker.internal:host-gateway \
  prom/prometheus

# 2. Start Grafana
echo "Starting Grafana..."
docker run -d \
  --name grafana \
  -p 3000:3000 \
  --add-host host.docker.internal:host-gateway \
  grafana/grafana

echo "Observability Stack Started!"
echo "Prometheus: http://localhost:9090"
echo "Grafana:    http://localhost:3000"
