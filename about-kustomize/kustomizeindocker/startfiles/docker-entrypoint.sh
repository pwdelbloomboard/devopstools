#!/bin/bash

# we are starting out in the directory /home/k3d-demo/ as specified in the dockerfile

# create a test cluster
k3d cluster create testcluster --api-port 6550 -p "8081:80@loadbalancer"

# once test cluster is up, apply pre-downloaded settings
kubectl apply --filename k8s/


# If a command was passed in, use it. Otherwise use /app/start.
# Needed to support passing flags to run script for local use.
if [[ -n "$*" ]]; then
    # shellcheck disable=SC2068
    exec $@
else
    # keep container running
    tail -f /dev/null
fi
