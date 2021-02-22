# chaos-arcade

[![e2e](https://github.com/ryneal/chaos-arcade/workflows/e2e/badge.svg)](https://github.com/ryneal/chaos-arcade/blob/master/.github/workflows/e2e.yml)
[![test](https://github.com/ryneal/chaos-arcade/workflows/test/badge.svg)](https://github.com/ryneal/chaos-arcade/blob/master/.github/workflows/test.yml)
[![cve-scan](https://github.com/ryneal/chaos-arcade/workflows/cve-scan/badge.svg)](https://github.com/ryneal/chaos-arcade/blob/master/.github/workflows/cve-scan.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/ryneal/chaos-arcade)](https://goreportcard.com/report/github.com/ryneal/chaos-arcade)
[![Docker Pulls](https://img.shields.io/docker/pulls/ryneal/chaos-arcade)](https://hub.docker.com/r/ryneal/chaos-arcade)

chaos-arcade is a tiny web application made with Go that showcases best practices of running microservices in Kubernetes.

Specifications:

* Health checks (readiness and liveness)
* Graceful shutdown on interrupt signals
* File watcher for secrets and configmaps
* Instrumented with Prometheus
* Tracing with Istio and Jaeger
* Linkerd service profile
* Structured logging with zap 
* 12-factor app with viper
* Fault injection (random errors and latency)
* Swagger docs
* Helm and Kustomize installers
* End-to-End testing with Kubernetes Kind and Helm
* Kustomize testing with GitHub Actions and Open Policy Agent
* Multi-arch container image with Docker buildx and Github Actions
* CVE scanning with trivy

Web API:

* `GET /` prints runtime information
* `GET /version` prints chaos-arcade version and git commit hash 
* `GET /metrics` return HTTP requests duration and Go runtime metrics
* `GET /healthz` used by Kubernetes liveness probe
* `GET /readyz` used by Kubernetes readiness probe

gRPC API:

* `/grpc.health.v1.Health/Check` health checking

Web UI:

![chaos-arcade-ui](https://raw.githubusercontent.com/ryneal/chaos-arcade/gh-pages/screens/chaos-arcade-ui-v3.png)

To access the Swagger UI open `<chaos-arcade-host>/swagger/index.html` in a browser.

### Guides

* [GitOps Progressive Deliver with Flagger, Helm v3 and Linkerd](https://helm.workshop.flagger.dev/intro/)
* [GitOps Progressive Deliver on EKS with Flagger and AppMesh](https://eks.handson.flagger.dev/prerequisites/)
* [Automated canary deployments with Flagger and Istio](https://medium.com/google-cloud/automated-canary-deployments-with-flagger-and-istio-ac747827f9d1)
* [Kubernetes autoscaling with Istio metrics](https://medium.com/google-cloud/kubernetes-autoscaling-with-istio-metrics-76442253a45a)
* [Autoscaling EKS on Fargate with custom metrics](https://aws.amazon.com/blogs/containers/autoscaling-eks-on-fargate-with-custom-metrics/)
* [Managing Helm releases the GitOps way](https://medium.com/google-cloud/managing-helm-releases-the-gitops-way-207a6ac6ff0e)
* [Securing EKS Ingress With Contour And Let’s Encrypt The GitOps Way](https://aws.amazon.com/blogs/containers/securing-eks-ingress-contour-lets-encrypt-gitops/)

### Install

Helm:

```bash
helm repo add chaos-arcade https://ryneal.github.io/chaos-arcade

helm upgrade --install --wait frontend \
--namespace test \
--set replicaCount=2 \
--set backend=http://backend-chaos-arcade:9898/echo \
chaos-arcade/chaos-arcade

# Test pods have hook-delete-policy: hook-succeeded
helm test frontend

helm upgrade --install --wait backend \
--namespace test \
--set hpa.enabled=true \
chaos-arcade/chaos-arcade
```

Kustomize:

```bash
kubectl apply -k github.com/ryneal/chaos-arcade//kustomize
```

Docker:

```bash
docker run -dp 9898:9898 ryneal/chaos-arcade
```