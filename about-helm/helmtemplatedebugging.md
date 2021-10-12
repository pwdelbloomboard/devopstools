# Helm Template Debugging

The reference documentation for debugging helm templates can be found at:

* [Helm Template Debugging](https://helm.sh/docs/chart_template_guide/debugging/)


> Debugging templates can be tricky because the rendered templates are sent to the Kubernetes API server, which may reject the YAML files for reasons other than formatting.

## Helm Prerequisites

> The following prerequisites are required for a successful and properly secured use of Helm.

> 1. A Kubernetes cluster
> 2. Deciding what security configurations to apply to your installation if any 
> 3. Installing and configuring Helm

### Cluster Prerequisite

* Based upon the above, it's important to double check to ensure there is a cluster running in order to run helm.

So if we have an application, "mysampleapplication" on which we expect a cluster to be present, and assuming we are using k3d, we should be able to view the cluster with the following:

```
k3d cluster list

NAME          SERVERS   AGENTS   LOADBALANCER
k3s-default   1/1       2/2      true
local         1/1       0/0      true
```

However, we don't see an, "application" cluster listed, only a, "k3s-default" cluster. This likely means we're good to go, but if a specific, named application cluster is needed, this can be revisited.

### Security Prerequisites

### Helm Installed and Configured

## 

