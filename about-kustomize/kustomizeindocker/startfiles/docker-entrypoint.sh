#!/bin/bash

# we are starting out in the directory /home/k3d-demo/ as specified in the dockerfile

# create a test cluster, append errors and results to results.txt
k3d cluster create testcluster --api-port 6550 -p "8081:80@loadbalancer" 2>&1 | tee -a results.txt

# once test cluster is up, apply pre-downloaded settings, append errors and results to results.txt
kubectl apply --filename k8s/ 2>&1 | tee -a results.txt

# add the current annotations to the results.txt
echo -e "\n ------------------------------------------" >> results.txt

echo -e "\n --- Annotations Prior to Kustomization ---" >> results.txt

echo -e "\n deployment \n " >> results.txt

kubectl describe deployment | grep -n Annotations | tee -a results.txt

echo -e "\n ------------------------------------------" >> results.txt

# once pre-downloaded settings are complete and within the cluster, apply kustomization, append errors and results to results.txt
kubectl apply -k k8s/ 2>&1 | tee -a results.txt

# add the current annotations to the results.txt
echo -e "\n ------------------------------------------" >> results.txt

echo -e "\n ----- Annotations After Kustomization ----" >> results.txt

echo -e "\n deployment \n " >> results.txt

kubectl describe deployment | grep -n Annotations | tee -a results.txt

echo -e "\n ------------------------------------------" >> results.txt


# If a command was passed in, use it. Otherwise use /app/start.
# Needed to support passing flags to run script for local use.
if [[ -n "$*" ]]; then
    # shellcheck disable=SC2068
    exec $@
else
    # keep container running
    tail -f /dev/null
fi
