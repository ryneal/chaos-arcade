# chaos-arcade

Chaos arcade is a sample service which is designed to demonstrate concepts of chaos engineering.

This project is hugely influenced by [KubeInvaders](https://github.com/lucky-sideburn/KubeInvaders)

Arcade games included so far include:
* Space invaders
* Snake

Specifications:

* Arcade games
* Kill random pod API
* Health checks (readiness and liveness)

### Install

Helm:

```bash
helm repo add chaos-arcade https://ryneal.github.io/chaos-arcade

helm upgrade --install --wait frontend \
--namespace chaos \
--set replicaCount=2 \
--set allowedNamespaces=test \
chaos-arcade/chaos-arcade

```