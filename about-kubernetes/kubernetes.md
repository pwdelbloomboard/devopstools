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

### Starting a Cluster Locally with k3d

Starting a Cluster locally on an, "edge" machine such as a personal linux or MacOS involves the use of k3d.

Further discussion [can be found in the k3d.md documentation](/about-k3d/k3d.md#launching-a-cluster-on-edge-with-k3d).



# Kubernetes Components

### Pods

Pods are the fundamental, atomic component of k8s, they represent a group of one or more application containers, (such as a Docker Container), as well as "resource containers," (such as a Docker Volume) or Networking information, and Information on how to run the container, such as versioning and ports.

A single Pod might contain a container with a React app, and then a different container with a Flask app, with the two containers feeding data back and fourth to each other. The containers in a Pod share an IP Address Space and Port Space, are always co-located and co-scheduled, and run in a shared context in a, "k8s node," (discussed in the next section).

![](/img/pods-overview.svg)

As can be seen above, Pods are quite simply, collections of, "stuff," Containers which hold Applications, and each Pod has its own IP address/space. The stuff might be one or more Containerized app, Volumes or other resources.

### Diagrams of Pods and Nodes

The below diagram shows a very abstracted view of the, "worker node," and its internals. Worker Nodes are basically a collection of things that allow Pods to run (Kubelet, Runtime, Proxy, etc.), as well as the Pods themselves. At this point we can see there is something called a, "Master Node," which does something as well.

![](/img/k8s-model-workernode.png)

The below diagram shows an overall view of k8s including a Master Node and individual Worker Nodes. This Master Node has Controllers and Schedulers, discussed further in detail below, but essentially they are, "loops," which are constantly running and checking to make sure everything is in a healthy, running and expected state - like a smart thermostat making sure your house is at the right temperature.

![](/img/k8s-model-full.png)

The below diagram shows a model including a load balancer. An, ["External Load Balancer"](https://kubernetes.io/docs/tasks/access-application-cluster/create-external-load-balancer/) is an optional strategy that helps balance traffic from many users in the outside world between different Worker Nodes to keep everything even and steady and prevent any individual Worker from getting too overloaded causing a crash.

![](/img/k8s-diagram.jpg)
### Clusters

* Clusters are groups of machines, or, "Groups of Nodes," using our above terminology, which work together, in this case a k8s cluster (vs. a k3d cluster, which would be a lightweight version of k8s for edge or local machines and laptops). 
* These can be of any size, a single node cluster would consist of one machine that hosts the k8s ontrol-plane (exposing API and maintaining the cluster).
* These can also be expanded to up to 5000 nodes total, as of k8s v1.18.

When you deploy Kubernetes, you deploy a cluster, even if it's just that one Master Node, it's a cluster of one. [Kubernetes clusters](https://kubernetes.io/docs/concepts/overview/components/) refer to a k8s cluster which includes:

* This may also include as mentioned above, a set of worker machines, called Nodes, that run containerized applications.
* Every cluster should probably have at least one worker node.
* Worker nodes host "Pods," which are the components of the application workload.
* The Control Plane - manages the worker nodes and the pods in the cluster, in production environments, the control plane usually runs across multiple computers.
* Clusters can run multiple Nodes, providing fault-tolerance and high availability, or orchestrate an app in a particular configuration for some other useful reason not discussed in this document.

![](/img/components-of-kubernetes.svg)

## Control Plane Components

[Documentation on Control Plane Components](https://kubernetes.io/docs/concepts/overview/components/#control-plane-components)

![](/img/k8s-controlplane.jpg)

### kube-apiserver

* Exposes the [Kubernetes API](https://kubernetes.io/docs/concepts/overview/kubernetes-api/).
* Scales by deploying more instances.
### etcd

* Highly available key-value store used as backing store for cluster data.
* There should be a backup plan if using this.
* [etcd documentation](https://etcd.io/docs/)
### kube-scheduler

* watches for newly created pods with no assigned node and selects node for them to run on.
* factors taken into account include, individual and collective resource requirements, hardware / software policy constraints, affinity and anti-affinity specifications, data locality, interworkload interference, and deadlines.

### kube-controller-manager

[Controller Documentation](https://kubernetes.io/docs/concepts/architecture/controller/)

* runs controller processes - a control loop is a non-terminating loop that regulates the state of a system, arbitrating between a "current state" and a "desired state."
* "Controllers" in Kubernetes are control loops that watch the state of your cluster, then request changes where needed.
* The "Controller Manager," run various controller processes. Each controller is a seperate process, but they are compiled into a single binary and run as a single process.
* Examples: Node Controller...responsible for noticing and responding when nodes go down. Job Controller...watches for job objects that represent one-off-tasks, then creates pods to run these tasks to completion Endpoints Controller - populates endpoints (joins Service & Pods).  Etc.
### cloud-controller-manager

* Lets you link your cluster to your cloud provider's API, seperates out components that interact with that cloud platform from components that only interact with your cluster.
* This is one of the items that seems to be deleted within k3d/k3s.
* Similar to the kube-controller-manager, this is a control loop, and it's specific to what cloud provider you use (for now, but in the future everything seems to be moving toward a universal standard).
## Node Componenets

![](/img/k8s-clusternodes.jpg)

### kubelet

An agent that runs on each Node in the cluster - makes sure that containers are running in a Pod.

The kubelet takes what are called, "PodSpecs," provided through various mechanims and ensures containers described in those PodSpecs.

(note: the kublet does not touch containers not managed by Kubernetes)

### kube-proxy

* Network proxy that runs on each node in a cluster, maintaining network rules on nodes.
### Container Runtime

* Just the software responsible for running containers, such as Docker, containerd, CRI-O, etc.
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
* [Kubernetes in 3 Diagrams](https://tsuyoshiushio.medium.com/kubernetes-in-three-diagrams-6aba8432541c)
* [K8s Enterprise Architecture](https://platform9.com/blog/kubernetes-enterprise-chapter-2-kubernetes-architecture-concepts/)
* [k8s Basics](https://kubernetes.io/docs/tutorials/kubernetes-basics/explore/explore-intro/)