#1 /usr/bin/env sh

set -e

# wait for chaos-arcade
kubectl rollout status deployment/chaos-arcade --timeout=3m

# test chaos-arcade
helm test chaos-arcade
