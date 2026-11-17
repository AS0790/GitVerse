# Kubernetes Deployment Architecture

GitVerse can be deployed to Kubernetes clusters as a stateless microservice.

## Architecture
* Deployment replica count: 2+
* Resource limits: 128Mi Memory, 200m CPU
* Health checks: HTTP GET /healthz
