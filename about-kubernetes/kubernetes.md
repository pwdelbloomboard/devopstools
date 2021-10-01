# About Kubernetes

## Background

Kubernetes is a microservice clustering application. Microservices are small, autonomous services that work together - the opposite of which would be called a monolith.

Large enterprises may think of microservices as extremely small, independent and managed by teams that don't even know each other exists. This is creating a backlash where some developers advocate returning to a monolith strucutre, which hypothetially would introduce a greater level of coding discipline, through the introduction of classes.

For smaller organizations who can actually graph out what their entire microservice infrastrucutre looks like, who are in growth mode...or otherwise put, "when you have a good reason to use microservices."

Other important reasoning:

* Zero downtime, independent deployability.
* Isolation of data and processing around said data
* Use of microservices to reflect organizational structure.

## Usage

* If there are 3 processes and 2 computers, and neither computer is running all 3 processes, you can approach the problem by deciding which of the 2 computers can run 2 processes and which runs 1...then network the machines together.
* You can also split up tasks based upon which computers have more resources, or based upon what is happening operationally - if you have available downtime for one process A, while B is in use you can flip A over to the weaker computer and B to the stronger computer.
* It's essentially a sum of bash scripts and best practices that system administrators would cobble together over time, presented as a single system behind declarative APIs.
* It's a system for automating deployment, scaling and management of containerized applications. It groups containers to make up an application into logical units for easy management and discovery.

## Usage with k3d

* [More background on k3d at k3d.md](/about-k3d/k3d.md)

### Clusters

* Clusters are groups of machines, nodes, which work together, in this case a k8s cluster. 
* These can be of any size, a single node cluster would consist of one machine that hosts the k8s ontrol-plane (exposing API and maintaining the cluster).
* These can also be expanded to up to 5000 nodes total, as of k8s v1.18.

### Starting a Cluster Locally with k3d

Starting a Cluster locally on an, "edge" machine such as a personal linux or MacOS involves the use of k3d.

Further discussion [can be found in the k3d.md documentation](/about-k3d/k3d.md#launching-a-cluster-on-edge-with-k3d).



# Kubernetes Components

When you deploy Kubernetes, you deploy a cluster. [Kubernetes clusters](https://kubernetes.io/docs/concepts/overview/components/) refer to a k8s cluster which includes:

* A set of worker machines, called nodes, that run containerized applications.
* Every cluster has at least one worker node.
* Worker nodes host "pods," which are the components of the application workload.
* The Control Plane - manages the worker nodes and the pods in the cluster, in production environments, the control plane usually runs across multiple computers.
* Cluster runs multiple nodes, providing fault-tolerance and high availability.

![](/about-kubernetes/img/components-of-kubernetes.svg)

## Control Pane Components

### kube-apiserver

### etcd

### kube-scheduler

### kube-controller-manager

### cloud-controller-manager

## Node Componenets

### kubelet

### kube-proxy

### Container Runtime

## Addons

### DNS

### Web UI

### Container Resource Monitoring

### Cluster Level Logging


# Kubernetes API




# Tools
### kubectl

kubectl is a command used to interact with clusters once it's up and running. There are various ways to install kubctl listed online depending upon the machine you're working with.

* [kubctl documentation](https://kubernetes.io/docs/reference/generated/kubectl/kubectl-commands)
* [Overall Reference Documentation](https://kubernetes.io/docs/reference/kubectl/)
* [kubctl cheat sheet](https://kubernetes.io/docs/reference/kubectl/cheatsheet/)


### kube prometheus

Includes various kubernetes monitoring components.

* [prometheus.md](/about-kubernetes/prometheus)

# Resources

* [Devops with Kubernetes - Introductory Tutorial](https://devopswithkubernetes.com/part-0) - utilizes k3d/k3s.
* [Deploying Simple Kubernetes Cluster](https://www.appvia.io/blog/tutorial-deploy-kubernetes-cluster) - interesting tutorial, but uses kind rather than k3d.
* [Run Kubernetes Locally with k3d - Youtube](https://www.youtube.com/watch?v=mCesuGk-Fks)
* [k3d vs minikube vs kind](https://brennerm.github.io/posts/minikube-vs-kind-vs-k3s.html)
* [Stackexchange Server Questions](https://serverfault.com/questions/78048/whats-the-difference-between-ip-address-0-0-0-0-and-127-0-0-1)
* [How Rancher Lab's K3d ... Kubernetes on the Edge](https://thenewstack.io/how-rancher-labs-k3s-makes-it-easy-to-run-kubernetes-at-the-edge/)