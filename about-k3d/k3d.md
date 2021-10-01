# About k3d

* [k3d.io](https://k3d.io/v4.4.8/) is a lightweight wrapper to run k3s, (which is a minimal Kubernetes distribution), - k3d allows one to run k3s in docker.

* [k3s](https://github.com/k3s-io/k3s) is lightweight kubernetes, which is production ready, half of the memory (presumably of kubernetes), within a binary of less than 100MB.

Some of the marketed features of k3s is that it could be used for Edge, IoT, CI, Development, ARM, embedding k8s, etc.

To get going with k3d, given a familiarity with k3s (discussed below), you can start off by running through the [k3d quickstart](https://k3d.io/#quick-start).

## About k3s

k3s is a production-ready Kubernetes distribution that is packaged as a single binary that wraps Kubernetes in a single, simple launcher. There are no OS dependencies, and eliminates the need to epose a port on K8s worker nodes.

* bundles several technologies for cohesive distributions for CNI, DNS, Metric Server, ingress, kube router, helm controller, etc.
* name comes from k8s, where 1+8+1=10, 10/2 = 5 (e.g. half the memory footprint of k8s), so then 8-5=3, therefore k3s = (k8s)/2
* The memory footprint was reduced by running many components inside of a single process, which eliminates significant overhead that would otherwise be duplicated for each component.
* binary is smaller by removing third-party storage drivers and cloud providders (upstream kubernetes).
- In-tree storage drivers
- In-tree cloud providers

### About In-Tree Storage Drivers

* [In-Tree Implementation Models](https://cloud-provider-vsphere.sigs.k8s.io/concepts/in_tree_vs_out_of_tree.html) were originally the default for Kubernetes. However as more infrastructure providers support Kubernetes, the old method becomes impractical. So basically it's a legacy system, bloatware that needs to stick around possibly for compatibility purposes, but [there is a pathway to migrate to out-of-tree methods](https://kubernetes.io/blog/2019/12/09/kubernetes-1-17-feature-csi-migration-beta/).

* [Container Storage Interface](https://github.com/container-storage-interface/spec/blob/master/spec.md) is a developing standard, "Container Storage Interface" that will enable storage providers to develop a plugin once and have it work across a number of different container orchetration systems, so it's like "USB."
* [Cloud Controller Manager](https://kubernetes.io/docs/tasks/administer-cluster/running-cloud-controller/) - is there to help interface with cloud providers vs. the Kubernetes project, sort of like an API between cloud providers and Kubernetes.

> Since cloud providers develop and release at a different pace compared to the Kubernetes project, abstracting the provider-specific code to the cloud-controller-manager binary allows cloud vendors to evolve independently from the core Kubernetes code.

## Usage of kubectl

* [kubctl](https://kubernetes.io/docs/tasks/tools/) is the Kubernetes command line tul, which allows running commands aginst k8s clusters.


## k3s vs Kind vs Minikube

* Minikube spawns a VM that is a single node k8s cluster. It is very beginner firendly tool.

* Kind or, "K in D" moves the cluster into docker containers, which results in a significantly faster startup time compared to spawning a VM.

* k3s is basically k8s with dispensible features removed, as mentioned above. The application is split into the k3s server and agent, the former acts as a manager while the latter is responsible for handling the workload. This allows you to deploy k8s manifests and helm charts by putting them in a specific directory.

# Launching a Cluster on Edge with k3d

To use k3d to start a cluster, you can simply run:

```
k3d cluster create mycluster
```
This will run through a list of steps ranging from preperation of the network to pulling the image from rancher, to some usage instructions.

```
$ k3d cluster create mycluster
INFO[0000] Prep: Network
INFO[0000] Created network 'k3d-mycluster' (8ee3b306e8a55df6e6cc2c8cfdb786877e0d4afafcb8aa74ff9d5a0b2ef27816)
INFO[0000] Created volume 'k3d-mycluster-images'
INFO[0001] Creating node 'k3d-mycluster-server-0'
INFO[0002] Pulling image 'docker.io/rancher/k3s:v1.21.1-k3s1'
INFO[0006] Creating LoadBalancer 'k3d-mycluster-serverlb'
INFO[0007] Pulling image 'docker.io/rancher/k3d-proxy:v4.4.6'
INFO[0009] Starting cluster 'mycluster'
INFO[0009] Starting servers...
INFO[0009] Starting Node 'k3d-mycluster-server-0'
INFO[0016] Starting agents...
INFO[0016] Starting helpers...
INFO[0016] Starting Node 'k3d-mycluster-serverlb'
INFO[0017] (Optional) Trying to get IP of the docker host and inject it into the cluster as 'host.k3d.internal' for easy access
INFO[0021] Successfully added host record to /etc/hosts in 2/2 nodes and to the CoreDNS ConfigMap
INFO[0021] Cluster 'mycluster' created successfully!
INFO[0021] --kubeconfig-update-default=false --> sets --kubeconfig-switch-context=false
INFO[0021] You can now use it like this:
kubectl config use-context k3d-mycluster
kubectl cluster-info
```

Then if you run docker ps -a to get a look at the processes you see:

```
$ docker ps -a
CONTAINER ID   IMAGE                      COMMAND                  CREATED          STATUS          PORTS                             NAMES
4862f3d0b744   rancher/k3d-proxy:v4.4.6   "/bin/sh -c nginx-pr…"   10 minutes ago   Up 9 minutes    80/tcp, 0.0.0.0:61395->6443/tcp   k3d-mycluster-serverlb
d58c09ecc571   rancher/k3s:v1.21.1-k3s1   "/bin/k3s server --t…"   10 minutes ago   Up 10 minutes                                     k3d-myc
```
Further, if you curl the IP address, you get the message:

* "Client sent an HTTP request to an HTTPS server."

* Note, in this context, local computing, 0.0.0.0 means, "all addresses on the local machine," whereas normally, the usage of, "127.0.0.1 ~ 127.255.255.254" is the [loopback address](https://www.sciencedirect.com/topics/computer-science/loopback-address) in ipv4, also known as, "localhost," which is the internal address that routes back to the local system.  In ipv6 this is 0:0:0:0:0:0:0:1 or ::1.
* If your local virtual machine has two IP addresses, 192.168.1.1 and 10.1.2.1, and you allow a webserver daemon like apache to listen on 0.0.0.0, it will be reachable at both of those IP addresses.
* Why TCP? TCP is for communicating between server and client, it offers guaranteed delivery of data packets in order, whereas UDP would not offer delivery in order or guaranteed delivery. HTTP uses port 80 while TCP uses no port.
* Why port 61395? [Port 61395](https://www.speedguide.net/port.php?port=61395) appears to be a private, dynamic port, randomly selected, without much reasoning other than it is dynamic and private. Since TCP uses no port this may be assigned for security purposes.

If we wanted to create a k3d cluster with 2 agent nodes, we could feed in command:

```
k3d cluster create -a 2
```
* What is an agent node?

To find this out, [in the context of k3d commands](https://k3d.io/v4.4.8/usage/commands/), we can look at the documentation by using the flags, "k3d cluster -h".

![](/img/k3dagentvsserver.png)

As can be seen in the image, there are two main constructs within k3d - Agent and Server.

#### k3s Server

* Servers contain a, "Process," which contains several important components:
- Database (via SQLite)
- API Server
- Tunnel Proxy (connects to the k3s Agent)
- Scheduler
- Controller Manager

These components are roughly analogous to a [k8s Control Plane](about-kubernetes/kubernetes.md#control-plane-components)

#### k3s Agent

Agents contain a, "Process" which contains:
- Tunnel Proxy (to connect to the k3s Server)
- Kube proxy
- [Flannel](https://github.com/flannel-io/flannel#flannel) - Flannel is an overlay network that works by assigning a range of subnet addresses (usually IPv4 with a /24 or /16 subnet mask). An overlay network is a computer network that is built on top of another network. Nodes in the overlay network can be thought of as being connected by virtual or logical links, each of which corresponds to a path, perhaps through many physical links, in the underlying network.
- Kublet, which has a ContainerID as well as various Pods.

These components are roughly analogous to a [k8s Node](about-kubernetes/kubernetes.md#node-components).

##### Flannel Illustrated

![](/img/flanneloverview.png)




# Creating a Setup Deployment

A setup, "hello world," deployment would include:

* Kubectl
* Docker
* k3d / k3s.


# Resources

* [k3d vs minikube vs kind](https://brennerm.github.io/posts/minikube-vs-kind-vs-k3s.html)
* [k3s + k3d = k8s: perfect match for dev and test](https://en.sokube.ch/post/k3s-k3d-k8s-a-new-perfect-match-for-dev-and-test-1)