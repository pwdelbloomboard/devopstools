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

## Starting off with Simple Commands

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

* What is an Agent?  Basically it's analogous to a, "Node" within Kubernetes.

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

Flannel is analogous to the k8s kube-proxy.

![](/img/flanneloverview.png)

## Continuing Simple Commands

In the above section we created a cluster and looked at the anaologies to the terminologies between k3d/k3s and k8s.

There are other fundamental commands that can be run using, "cluster," which we can observe via "k3d cluster -h" .

```
$ k3d cluster list
NAME        SERVERS   AGENTS   LOADBALANCER
mycluster   1/1       0/0      true
```

IF we run the, "stop" cluster command, we see a slight delay, then listing, we see:

```
$ k3d cluster stop mycluster
INFO[0000] Stopping cluster 'mycluster'

$ k3d cluster list
NAME        SERVERS   AGENTS   LOADBALANCER
mycluster   0/1       0/0      true
```
Which shows us that the Servers are now down.  Using Docker to inspect the containers:

```
$ docker ps -a
CONTAINER ID   IMAGE                      COMMAND                  CREATED        STATUS                       PORTS     NAMES
4862f3d0b744   rancher/k3d-proxy:v4.4.6   "/bin/sh -c nginx-pr…"   21 hours ago   Exited (137) 2 minutes ago             k3d-mycluster-serverlb
d58c09ecc571   rancher/k3s:v1.21.1-k3s1   "/bin/k3s server --t…"   21 hours ago   Exited (1) 2 minutes ago               k3d-mycluster-server-0
```
We see that the statuses are now shown as, "Exited."

We can of course start these, "Exited" clusters back up and re-observe them with docker once again:

```
$ k3d cluster start mycluster
INFO[0000] Starting cluster 'mycluster'
INFO[0000] Starting servers...
INFO[0000] Starting Node 'k3d-mycluster-server-0'
INFO[0006] Starting agents...
INFO[0006] Starting helpers...
INFO[0006] Starting Node 'k3d-mycluster-serverlb'

$ docker ps -a
CONTAINER ID   IMAGE                      COMMAND                  CREATED        STATUS          PORTS                             NAMES
4862f3d0b744   rancher/k3d-proxy:v4.4.6   "/bin/sh -c nginx-pr…"   21 hours ago   Up 11 seconds   80/tcp, 0.0.0.0:61395->6443/tcp   k3d-mycluster-serverlb
d58c09ecc571   rancher/k3s:v1.21.1-k3s1   "/bin/k3s server --t…"   21 hours ago   Up 18 seconds                                     k3d-mycluster-server-0
```
We can also stop and delete the cluster with:

```
$ k3d cluster delete mycluster
INFO[0000] Deleting cluster 'mycluster'
INFO[0000] Deleted k3d-mycluster-serverlb
INFO[0001] Deleted k3d-mycluster-server-0
INFO[0001] Deleting cluster network 'k3d-mycluster'
INFO[0001] Deleting image volume 'k3d-mycluster-images'
INFO[0001] Removing cluster details from default kubeconfig...
INFO[0001] Removing standalone kubeconfig file (if there is one)...
INFO[0001] Successfully deleted cluster mycluster!
```
Inspecting mycluster with docker ps -a will yield a blank result.

So of course, any cluster with the capability to do something should have at least one agent, so let's create a cluster with two agents as a demonstration.

```
$ k3d cluster create -a 2
INFO[0000] Prep: Network
INFO[0000] Created network 'k3d-k3s-default' (a0c553abd851f9fd91f1dbb35c816aaf3102dd2b0af304ec0ca7614a5a7ee12b)
INFO[0000] Created volume 'k3d-k3s-default-images'
INFO[0001] Creating node 'k3d-k3s-default-server-0'
INFO[0001] Creating node 'k3d-k3s-default-agent-0'
INFO[0001] Creating node 'k3d-k3s-default-agent-1'
INFO[0001] Creating LoadBalancer 'k3d-k3s-default-serverlb'
INFO[0001] Starting cluster 'k3s-default'
INFO[0001] Starting servers...
INFO[0001] Starting Node 'k3d-k3s-default-server-0'
INFO[0008] Starting agents...
INFO[0008] Starting Node 'k3d-k3s-default-agent-0'
INFO[0020] Starting Node 'k3d-k3s-default-agent-1'
INFO[0028] Starting helpers...
INFO[0028] Starting Node 'k3d-k3s-default-serverlb'
INFO[0029] (Optional) Trying to get IP of the docker host and inject it into the cluster as 'host.k3d.internal' for easy access
INFO[0031] Successfully added host record to /etc/hosts in 4/4 nodes and to the CoreDNS ConfigMap
INFO[0031] Cluster 'k3s-default' created successfully!
INFO[0031] --kubeconfig-update-default=false --> sets --kubeconfig-switch-context=false
INFO[0031] You can now use it like this:
kubectl config use-context k3d-k3s-default
kubectl cluster-info
```
Running through our series of inspection commands:

```
$ k3d cluster list
NAME          SERVERS   AGENTS   LOADBALANCER
k3s-default   1/1       2/2      true

$ docker ps -a
CONTAINER ID   IMAGE                      COMMAND                  CREATED          STATUS          PORTS                             NAMES
3b38c8105626   rancher/k3d-proxy:v4.4.6   "/bin/sh -c nginx-pr…"   26 minutes ago   Up 26 minutes   80/tcp, 0.0.0.0:63568->6443/tcp   k3d-k3s-default-serverlb
75ebee0ba262   rancher/k3s:v1.21.1-k3s1   "/bin/k3s agent"         26 minutes ago   Up 26 minutes                                     k3d-k3s-default-agent-1
f5c294a91cb1   rancher/k3s:v1.21.1-k3s1   "/bin/k3s agent"         26 minutes ago   Up 26 minutes                                     k3d-k3s-default-agent-0
2655afa00fa3   rancher/k3s:v1.21.1-k3s1   "/bin/k3s server --t…"   26 minutes ago   Up 26 minutes                                     k3d-k3s-default-server-0

```
We can see that we now have 2 agents, 1 server and 4 docker containers with the respective servers, agents and a load balancer. Neat!

If we run the, "kubectl cluster-info" command, we see the following:

```
$ kubectl cluster-info
Kubernetes control plane is running at https://0.0.0.0:63568
CoreDNS is running at https://0.0.0.0:63568/api/v1/namespaces/kube-system/services/kube-dns:dns/proxy
Metrics-server is running at https://0.0.0.0:63568/api/v1/namespaces/kube-system/services/https:metrics-server:/proxy
```
This of course tells us similar info to our, "docker ps -a" command in a different format and higher resolution, showing the control pane accessible at the proxy listed at the same address shown in docker ps -a, DNS showing that address and 

We could optionally set up a cluster with two agents and no load balancer via:

```
k3d cluster create -a 2 --no-lb
```
There is also a config file for the cluster, kubeconfig.  We can inquire on this with the name of the cluster and the command listed in the -h list, "get":

```
$ k3d cluster list
NAME          SERVERS   AGENTS   LOADBALANCER
k3s-default   1/1       2/2      true

$ k3d kubeconfig get k3s-default
---
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: LS0tLS...
    server: https://0.0.0.0:63568
  name: k3d-k3s-default
contexts:
- context:
    cluster: k3d-k3s-default
    user: admin@k3d-k3s-default
  name: k3d-k3s-default
current-context: k3d-k3s-default
kind: Config
preferences: {}
users:
- name: admin@k3d-k3s-default
  user:
    client-certificate-data: LS0tLS1

```
The above shows the cluster with certificates, server url and port, name, context(s) which includes the user, name and any other users and preferences.

If we want to delete the cluster finally, we use the command, "stop" and then, "delete."

### More Sophisticated Way of Running a k3d Cluster Without a Deployment

Typically, Clusters are not run without Deployments.  Deployments provide declarative updates for Pods and ReplicaSets. More information about [Kubernetes Deployments](https://github.com/pwdelbloomboard/devopstools/blob/main/about-kubernetes/kubernetesdeployment.md#deployment-vs-no-deployment) 

```
k3d cluster create mycluster --api-port 127.0.0.1:6445 --servers 3 --agents 2 --volume '/Users/{{username}}/Projects/dockerreactjs-yarn/app:/code@agent[*]' --port '8080:80@loadbalancer'
```
* 1 load balancer   [--port '8080:80@loadbalancer']
* 3 servers (control-plane nodes)   [--servers 3]
* 2 agents (formerly worker nodes)   [--agents 2]
* --api-port 127.0.0.1:6445, tells the k3d to map the Kubernetes API Port (6443 internally) to 127.0.0.1/localhost’s port 6445.
* --volume /Users/{{username}}/Projects/dockerreactjs-yarn/app:/code@agent[*] bind mounts your local directory ../app to the path /code inside all ([*] of your agent nodes). Replace * with an index (here: 0 or 1) to only mount it into one of them.
* --port '8080:80@loadbalancer ... maps your local host’s port 8080 to port 80 on the load balancer (serverlb), which can be used to forward HTTP ingress traffic to your cluster. For example, you can now deploy a web app into the cluster (Deployment), which is exposed (Service) externally via an Ingress such as myapp.k3d.localhost.

> Note: You have to have some mechanism set up to route to resolve myapp.k3d.localhost to your local host IP (127.0.0.1). The most common way is using entries of the form 127.0.0.1 myapp.k3d.localhost in your /etc/hosts file

If we don't have a way to resolve the localhost IP, or an ingress, or port forwarding, then we won't be able to view the app, even if it is running.  On top of this, the command we used above, points to a file on the local machine, rather than a volume with a running, "container" that could be a pod, so we don't have a way of running the app itself.  We will get one of the following messages:

![](/img/nginx_notfound.png)

Then (provided that everything is set up to resolve that domain to your local host IP), you can point your browser to http://myapp.k3d.localhost:8080 to access your app (assuming ingresses are set up correctly). Traffic then flows from your host through the Docker bridge interface to the load balancer. From there, it’s proxied to the cluster, where it passes via Ingress and Service to your application Pod.

So using the, "whoami" and "pwd" command, we can get an idea of where we are and what the username is. With that information, we can set up a command which creates a, "non-deployment cluster, as shown below:

```
k3d cluster create nondeploymentcluster --api-port 127.0.0.1:6445 --servers 1 --agents 1 --volume '/Users/patrick.delaneybloomboard.com/Projects/dockerreactjs-yarn/app:/code@agent[*]' --port '8080:80@loadbalancer'
```
This should spit out the following:

```
INFO[0000] Prep: Network                                
INFO[0000] Created network 'k3d-nondeploymentcluster' (61e6e44237ba7f28aecc083fe471430406bc8e81d9b0ccf5c8f12ca69fa5eed5) 
INFO[0000] Created volume 'k3d-nondeploymentcluster-images' 
INFO[0001] Creating node 'k3d-nondeploymentcluster-server-0' 
INFO[0001] Creating node 'k3d-nondeploymentcluster-agent-0' 
INFO[0001] Creating LoadBalancer 'k3d-nondeploymentcluster-serverlb' 
INFO[0001] Starting cluster 'nondeploymentcluster'      
INFO[0001] Starting servers...                          
INFO[0001] Starting Node 'k3d-nondeploymentcluster-server-0' 
INFO[0011] Starting agents...                           
INFO[0011] Starting Node 'k3d-nondeploymentcluster-agent-0' 
INFO[0024] Starting helpers...                          
INFO[0024] Starting Node 'k3d-nondeploymentcluster-serverlb' 
INFO[0026] (Optional) Trying to get IP of the docker host and inject it into the cluster as 'host.k3d.internal' for easy access 
INFO[0033] Successfully added host record to /etc/hosts in 3/3 nodes and to the CoreDNS ConfigMap 
INFO[0033] Cluster 'nondeploymentcluster' created successfully! 
INFO[0033] --kubeconfig-update-default=false --> sets --kubeconfig-switch-context=false 
INFO[0033] You can now use it like this:                
kubectl config use-context k3d-nondeploymentcluster
kubectl cluster-info
```
Running kubctl cluster-info shows the following information:

```
Kubernetes control plane is running at https://127.0.0.1:6445
CoreDNS is running at https://127.0.0.1:6445/api/v1/namespaces/kube-system/services/kube-dns:dns/proxy
Metrics-server is running at https://127.0.0.1:6445/api/v1/namespaces/kube-system/services/https:metrics-server:/proxy
```
If we run, "kubectl get pods", we get:
```
kubectl get pods
No resources found in default namespace.
```
But if we run, "kubectl get nodes" we get:
```
$ kubectl get nodes
NAME                                STATUS   ROLES                  AGE   VERSION
k3d-nondeploymentcluster-agent-0    Ready    <none>                 23m   v1.21.1+k3s1
k3d-nondeploymentcluster-server-0   Ready    control-plane,master   23m   v1.21.1+k3s1
```
This seems to be showing that the pods are not exposed, while the nodes are. The question is - can we set up port forwarding on these nodes to allow the pod to get exposed so that we can see the application?

```
kubectl port-forward k3d-nondeploymentcluster-agent-0 3003:3000
Error from server (NotFound): pods "k3d-nondeploymentcluster-agent-0" not found
```
So basically, since k3d-nondeploymentcluster-agent-0 is a node and not a pod, we can't seem to forward it.  How can we have nodes but without pods?  This goes back to the [defnition of nodes and pods](/about-kubernetes/kubernetes.md#diagrams-of-pods-and-nodes) - basically pods are application containers, while nodes are operating system containers. The pods run on top of the nodes, so you could have an empty node running, "no application," because there is no pod inside of it.

IF we look at the resources that our nodes and pods are using, we can see that our nodes take up space, while our pods do not:

```
kubectl top node
W1007 14:31:41.403250   76548 top_node.go:119] Using json format to get metrics. Next release will switch to protocol-buffers, switch early by passing --use-protocol-buffers flag
NAME                                CPU(cores)   CPU%   MEMORY(bytes)   MEMORY%   
k3d-nondeploymentcluster-agent-0    45m          2%     225Mi           1%        
k3d-nondeploymentcluster-server-0   172m         8%     636Mi           3%    

kubectl top pod
W1007 14:36:50.185384   77172 top_pod.go:140] 

No resources found in default namespace.
```
So basically what is happening here is not that port forwarding is not being set up correctly, it's that the directory, "/Users/patrick.delaneybloomboard.com/Projects/dockerreactjs-yarn/app" does not point to a volume, but instead to the source code, unbuilt, of the ReactJS app as well as the Dockerfile and everything else, which presumably does nothing for us in terms of building any type of app, because it's not even containerized, it's just source code.

If we want to see which clusters we have available we can use the following handy commands, which are available within k3d:

```
k3d cluster list
NAME                   SERVERS   AGENTS   LOADBALANCER
buysellguess           1/1       0/0      true
local                  1/1       0/0      true
nondeploymentcluster   1/1       1/1      true
```

We can stop, list, and delete / list these clusters with:

```
k3d cluster stop nondeploymentcluster
INFO[0000] Stopping cluster 'nondeploymentcluster'      

k3d cluster list
NAME                   SERVERS   AGENTS   LOADBALANCER
buysellguess           1/1       0/0      true
local                  1/1       0/0      true
nondeploymentcluster   0/1       0/1      true

k3d cluster delete

k3d cluster delete nondeploymentcluster

INFO[0000] Deleting cluster 'nondeploymentcluster'      
INFO[0000] Deleted k3d-nondeploymentcluster-serverlb    
INFO[0000] Deleted k3d-nondeploymentcluster-agent-0     
INFO[0000] Deleted k3d-nondeploymentcluster-server-0    
INFO[0000] Deleting cluster network 'k3d-nondeploymentcluster' 
INFO[0001] Deleting image volume 'k3d-nondeploymentcluster-images' 
INFO[0001] Removing cluster details from default kubeconfig... 
INFO[0001] Removing standalone kubeconfig file (if there is one)... 
INFO[0001] Successfully deleted cluster nondeploymentcluster!

k3d cluster list

NAME           SERVERS   AGENTS   LOADBALANCER
buysellguess   1/1       0/0      true
local          1/1       0/0      true

```

In order to properly set up a working application using the, "non-deployment," method we need to understand more about [volumes](/about-volumes/volumes.md) as well as [registries](/about-registires/registries.md).

### Creating K3D Cluster and Importing a Docker Image

* [Rancher Blog - Setup k3d high Availability](https://rancher.com/blog/2020/set-up-k3s-high-availability-using-k3d)

There is another way to create a k3d cluster, simply using an image.

```
k3d cluster create --servers 1 --agents 1 -–image ghcr.io/pwdelbloomboard/ps-container --port '8080:80@loadbalancer' --api-port 127.0.0.1:6445
```
However, when we run this we get the error, "FATA[0000] unknown shorthand flag: 'â' in -–image "

* Is this because the, "--image" needs to be a Kubernetes-ready image?
* Is this because the, "--image" must pull from a registry?  [discussion on registries](/about-registries/registries.md)
* It may also be because, "--image" is not a flag that is compatible with MacOS yet on k3d

However, it is probably just because it's a docker image, rather than a k3d-compatible image, whatever that may be.

There is another command [k3d image import](https://k3d.io/v5.0.0/usage/commands/k3d_image_import/) which essentially, "imports" a docker image into a k3d cluster.

```
Flags:
  -c, --cluster stringArray   Select clusters to load the image to. (default [k3s-default])
  -h, --help                  help for import
  -k, --keep-tarball          Do not delete the tarball containing the saved images from the shared volume

Global Flags:
      --timestamps   Enable Log timestamps
      --trace        Enable super verbose output (trace logging)
      --verbose      Enable verbose output (debug logging)
```

So using this strategy, we first create the cluster, then import our docker image with:

```
k3d cluster create nondeploymentcluster --api-port 127.0.0.1:6445 --servers 1 --agents 1 --port '8080:80@loadbalancer'

... (info on cluster deployment)

k3d image import ghcr.io/pwdelbloomboard/ps-container:latest -c nondeploymentcluster --trace

INFO[0000] Importing image(s) into cluster 'nondeploymentcluster' 
INFO[0000] Starting k3d-tools node...                   
INFO[0001] Pulling image 'docker.io/rancher/k3d-tools:v4.4.6' 
INFO[0002] Starting Node 'k3d-nondeploymentcluster-tools' 
INFO[0003] Saving 1 image(s) from runtime...            
INFO[0050] Importing images into nodes...               
INFO[0050] Importing images from tarball '/k3d/images/k3d-nondeploymentcluster-images-20211008102755.tar' into node 'k3d-nondeploymentcluster-server-0'... 
INFO[0050] Importing images from tarball '/k3d/images/k3d-nondeploymentcluster-images-20211008102755.tar' into node 'k3d-nondeploymentcluster-agent-0'... 
INFO[0114] Removing the tarball(s) from image volume... 
INFO[0115] Removing k3d-tools node...                   
INFO[0116] Deleted k3d-nondeploymentcluster-tools       
INFO[0116] Successfully imported image(s)               
INFO[0116] Successfully imported 1 image(s) into 1 cluster(s)
```
Note that although the documentation says, :latest is assumed, we had to explicitly put, :latest as the tag in order for this to work.

So once the Docker image is successfully imported into the cluster, we should in theory be able to view the pod. However if we run the kubectl get pods command, we see:

```
kubectl get pods
No resources found in default namespace.
```
* So, perhaps our Pods were not set up within the, "default" namespace. 
* Or perhaps our Pods were not set up at all within these Nodes.

Of course, the cluster itself is running:

```
$ k3d cluster list
NAME                   SERVERS   AGENTS   LOADBALANCER
buysellguess           1/1       0/0      true
local                  1/1       0/0      true
nondeploymentcluster   1/1       1/1      true
```
So if we run a wider scope view of Pods across all Namespaces, we see the following:

```
kubectl get pods --all-namespaces --output wide
NAMESPACE     NAME                                      READY   STATUS      RESTARTS   AGE   IP          NODE                                NOMINATED NODE   READINESS GATES
kube-system   coredns-7448499f4d-d5g2v                  1/1     Running     0          13h   10.42.0.3   k3d-nondeploymentcluster-server-0   <none>           <none>
kube-system   metrics-server-86cbb8457f-bjs9r           1/1     Running     0          13h   10.42.1.2   k3d-nondeploymentcluster-agent-0    <none>           <none>
kube-system   helm-install-traefik-crd-25xxg            0/1     Completed   0          13h   10.42.1.3   k3d-nondeploymentcluster-agent-0    <none>           <none>
kube-system   helm-install-traefik-bfmlw                0/1     Completed   2          13h   10.42.0.2   k3d-nondeploymentcluster-server-0   <none>           <none>
kube-system   svclb-traefik-4gh6q                       2/2     Running     0          13h   10.42.1.5   k3d-nondeploymentcluster-agent-0    <none>           <none>
kube-system   svclb-traefik-986lc                       2/2     Running     0          13h   10.42.0.5   k3d-nondeploymentcluster-server-0   <none>           <none>
kube-system   traefik-97b44b794-xznjc                   1/1     Running     0          13h   10.42.1.4   k3d-nondeploymentcluster-agent-0    <none>           <none>
kube-system   local-path-provisioner-5ff76fc89d-t4xsx   1/1     Running     1          13h   10.42.0.4   k3d-nondeploymentcluster-server-0   <none>           <none>
```

Above we can see the various pods on nondeploymentcluster server-0, and agent-0. However, where is our imported application pod?  What other information can we get from our available k3d commands?

* k3d kubeconfig get nondeploymentcluster ... gets the certificate of authority information, context, user, and ip information.

```
k3d kubeconfig get nondeploymentcluster
---
apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: LS0t  ... (very long certificate)
    server: https://127.0.0.1:6445
  name: k3d-nondeploymentcluster
contexts:
- context:
    cluster: k3d-nondeploymentcluster
    user: admin@k3d-nondeploymentcluster
  name: k3d-nondeploymentcluster
current-context: k3d-nondeploymentcluster
kind: Config
preferences: {}
users:
- name: admin@k3d-nondeploymentcluster
  user:
    client-certificate-data: LS0t ... (very long certificate)
    client-key-data: LS
```
* k3d node list ... lists out the nodes
```
k3d node list
NAME                                ROLE           CLUSTER                STATUS
k3d-buysellguess-server-0           server         buysellguess           running
k3d-buysellguess-serverlb           loadbalancer   buysellguess           running
k3d-local-server-0                  server         local                  running
k3d-local-serverlb                  loadbalancer   local                  running
k3d-nondeploymentcluster-agent-0    agent          nondeploymentcluster   running
k3d-nondeploymentcluster-server-0   server         nondeploymentcluster   running
k3d-nondeploymentcluster-serverlb   loadbalancer   nondeploymentcluster   running
k3d-registry.localhost              registry                              running
```
* We see above that there are three nodes specific to the, "nondeploymentcluster," - e.g. the agent, server0, and serverlb (load balancer)  Note that there is also a, "registry" node.

From a comment on a [k3d issue 642](https://github.com/rancher/k3d/issues/642), 

> Since you imported the image and didn't push it to a registry, the image "is just there" in the nodes, so you cannot pull it. Make sure that your imagePullPolicy to ifNotPresent, and not Always, so it doesn't try to pull.  You can use the image name just like you used it for the import command.

In other words, if we were to use the, "import image" command, the image would be just sitting in the Node, rather than being, "pulled," from a registry, where it would presumably go into a pod. The imagePullPolicy default appears to be, "Always" try to pull from a registry, so when it attempts to pull, if there is nothing there, it will not pull into a Pod.  However, this could be in the instance where we are using a .yaml file, and not the command line method.

If we attempt to run the command again, with --trace flag enabled, we get the following:

```
TRAC[0001] Exec process '[./k3d-tools save-image -d /k3d/images/k3d-nondeploymentcluster-images-20211008120948.tar ghcr.io/pwdelbloomboard/ps-container:latest]' still running in node 'k3d-nondeploymentcluster-tools'.. sleeping for 1 second... 
```

What might be happening is that we're pulling a, ["k3d-tools," image](https://github.com/rancher/k3d/tree/main/tools) from Docker.io, and not the, "ghcr.io," image because when we look at the output from one of the above "import" commands, we see:

```
$ k3d image import ghcr.io/pwdelbloomboard/ps-container:latest -c nondeploymentcluster
INFO[0000] Importing image(s) into cluster 'nondeploymentcluster' 
INFO[0000] Starting k3d-tools node...                   
INFO[0001] Pulling image 'docker.io/rancher/k3d-tools:v4.4.6' 
INFO[0002] Starting Node 'k3d-nondeploymentcluster-tools' 
INFO[0003] Saving 1 image(s) from runtime...            
INFO[0050] Importing images into nodes...               
INFO[0050] Importing images from tarball '/k3d/images/k3d-nondeploymentcluster-images-20211008102755.tar' into node 'k3d-nondeploymentcluster-server-0'... 
INFO[0050] Importing images from tarball '/k3d/images/k3d-nondeploymentcluster-images-20211008102755.tar' into node 'k3d-nondeploymentcluster-agent-0'... 
INFO[0114] Removing the tarball(s) from image volume... 
INFO[0115] Removing k3d-tools node...                   
INFO[0116] Deleted k3d-nondeploymentcluster-tools       
INFO[0116] Successfully imported image(s)               
INFO[0116] Successfully imported 1 image(s) into 1 cluster(s) 
```
#### Final Thoughts on Importing Images

* Using the, "import image," command line function is different than using the, "Deployment," command-line function.
* Importing images through the command line with "k3d import" is like importing an operating system, it's better to just stick with the stock k3d off-the-shelf images provided.
* Using, "Deployments," is more similar to creating an Application, or a version of an application, installing a, "Deployment," on an existing operating system.
* When using, "Deployments," you can use a YAML file and specify the, "image" as an independent Docker file on your machine.

## Deploying with k3d

In order to deploy, we need an application to deploy.

* As a sample application, we can use [this simple ReactJS](https://github.com/pwdel/dockerreactjs) application that the author built previously and deployed to Docker successfully.

* To start off with, we briefly updated that application so that it uses yarn to deploy rather than npm (as yarn is more secure) and set it up under github.com/pwdelbloomboard.

Per the tutorial we're working with:

> To deploy we need the cluster to have an access to the image. By default, Kubernetes is intended to be used with a registry. K3d offers import-images command, but since that won't work when we go to non-k3d solutions we'll use the now possibly very familiar registry Docker Hub, we used that in DevOps with Docker.

Github [now has a container registry](https://dev.to/github/github-container-registry-better-than-docker-hub-1o9k), so rather than signing up with Dockerhub, we can use our already-established Github account to get registered and push images diretly to Github.

### Github Container Registry

> authenticate using your GitHub Username and a PAT (Personal Access Token) with the write packages scope (watch this video to see how to create a PAT in GitHub), for example with Docker Login, and push the container as you would normally do.

> You just have to tag the image with the format ghcr.io/OWNER/IMAGE_NAME:version, where OWNER is the name of your user or the organization.

> And if you are doing it in GitHub Actions it's even easier.

#### Creating a Github Personal Access Token (PAT)

* [Documentation on Creating a Github Personal Access Token](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token)

You can create a PAT at this direct link [here](https://github.com/settings/tokens).

* The scope needed is to, "read, write and delete packages."  If you click that while creating a personal access token, then, "repo" will automatically be selected as well.

So then to push the image to a registry, you first tag it using the format:

```
ghcr.io/OWNER/IMAGE_NAME:version
```
So in our case, it will be:

```
docker build -t ghcr.io/pwdel/ps-container:latest .
```
...where we arbitrarily selected version "latest" for now. In order to use and push this tagged image to the registry, we can follow Github's documentation.

#### Working with Github Packages

* [Github Packages - (Github Container Registry Equivalent) ](https://docs.github.com/en/packages/working-with-a-github-packages-registry/working-with-the-container-registry)

##### Steps to Work with Github Packages

The background guide to how to Publish a Github Package can be found [here](https://docs.github.com/en/actions/publishing-packages/publishing-docker-images).

The idea is - you want to create a workflow that performs a Docker build, and then publishes Docker images to Docker Hub or GitHub Packages.  In our case, we will use Github packages.

Some prerequisites to review ahead of time:

* [Encrypted Secrets](https://docs.github.com/en/actions/security-guides/encrypted-secrets) - these are essentially encrypted environmental variables that we create in an organization, repo or repo environment. These are secrets that will be available to use in Github Actions workflows.
* [Automated Token Authorization](https://docs.github.com/en/actions/security-guides/automatic-token-authentication) - When you enable GitHub Actions, GitHub installs a GitHub App on your repository. [Github Actions are enabled by setting up an Action within a repo](https://github.com/pwdelbloomboard/dockerreactjs-yarn/actions/new).
* [Working with Github Container Registry](https://docs.github.com/en/packages/working-with-a-github-packages-registry/working-with-the-container-registry) is regular, run-of-the-mill tagging and pushing to a particular Container Registry, such as the Docker registry, with a command such as, "docker push ghcr.io/OWNER/IMAGE_NAME:latest" assuming an image has been pre-named/tagged with ghcr.io/OWNER/IMAGE_NAME:latest ," as an example.

*Some Quick Definitions for Github Actions*

"Github Actions," is the service in general, but it is also a standalone, one-line command, the smallest building block of the entire workflow.

* The workflow is an automated procedure that you add to your repository. Workflows are made up of one or more jobs and can be scheduled or triggered by an event. The workflow can be used to build, test, package, release, or deploy a project on GitHub. You can reference a workflow within another workflow, see "Reusing workflows."
* Events - An event is a specific activity that triggers a workflow. For example, activity can originate from GitHub when someone pushes a commit to a repository or when an issue or pull request is created. You can also use the repository dispatch webhook to trigger a workflow when an external event occurs. For a complete list of events that can be used to trigger workflows, see Events that trigger workflows.
* A job is a set of steps that execute on the same runner. By default, a workflow with multiple jobs will run those jobs in parallel. You can also configure a workflow to run jobs sequentially. For example, a workflow can have two sequential jobs that build and test code, where the test job is dependent on the status of the build job. If the build job fails, the test job will not run.
* A step is an individual task that can run commands in a job. A step can be either an action or a shell command. Each step in a job executes on the same runner, allowing the actions in that job to share data with each other.
* A, "runner" is the server on which all of the above run, it can be either through Github or we can host our own.


The format we use for builds can be found at the [docker-hub/builds](https://docs.docker.com/docker-hub/builds/) documentation.

So basically what happens is Github Actions is triggering a workflow, which uses the token (as specified in ${{ secrets.GITHUB_TOKEN }}) to connect to an outside app, which is a docker build runner.

The build image we need to make is essentially a set of command line commands running of [our Dockerfile](https://github.com/pwdelbloomboard/dockerreactjs-yarn/blob/main/app/Dockerfile), and pushing to the Github container registry, with each, "step" being an action within our Dockerfile.

```
buildImage:
  name: Build Docker Image
  runs-on: node:12-alpine

  steps:
    - name: Build Container Image
      run: docker build . --file Dockerfile --tag $IMAGE_NAME

    - name: Log into GitHub Container Registry
      run: echo "${{ secrets.GITHUB_TOKEN }}" | docker login https://ghcr.io -u ${{ github.actor }} --password-stdin

    - name: Push image to GitHub Container Registry
      run: |
        IMAGE_ID=ghcr.io/${{ github.repository_owner }}/MyBeautifulContainer:123        
        docker push $IMAGE_ID

```
A more throurough rundown of the above yaml file [may be found here at, "Publishing Docker Images to Github Packages](https://docs.github.com/en/actions/publishing-packages/publishing-docker-images#publishing-images-to-github-packages).


One the package is published, you can [connect a repository to a package](https://docs.github.com/en/packages/learn-github-packages/connecting-a-repository-to-a-package) within Github.


1. [Authenticate to Github Packages Container Registry](https://docs.github.com/en/packages/learn-github-packages/introduction-to-github-packages#authenticating-to-github-packages)

You can login using the command:

```
docker login ghcr.io 
```
With ghcr.io being the github container registry. When it prompts you for your username/password, use the access token for both. In theory the token username and password should be different, with the username being the actual username, but in this instance it worked with both plugged in.

You can set up your token as an environmental variable on your machine, like so:

```
export CR_PAT=YOUR_TOKEN
```
Then, the following command will log you in to the ghcr.io using your token.
```
echo $CR_PAT | docker login ghcr.io -u USERNAME --password-stdin
```
2. Push the container to the container registry [according to this documentation, "Working with a Github Packages Registry."](https://docs.github.com/en/packages/working-with-a-github-packages-registry/working-with-the-container-registry):

We use the following format:

```
docker push ghcr.io/OWNER/IMAGE-NAME:TAG
```
Which in our situation is:

```
docker push ghcr.io/pwdel/ps-container:latest
```
In this case you may get a, "permission deinied," error.

If that is the case, you may need to [Configure Access for the Member](https://github.community/t/unauthorized-error-when-pushing-images/131217) which is discussed within the [Configuring Package Access Control and Visibility Section](https://docs.github.com/en/packages/learn-github-packages/configuring-a-packages-access-control-and-visibility).

Basically, under the free plan, Access Permissions must be set to Public for Images. Public images allow anonymous access and can be pulled without authentication or signing in via the CLI.  Access Permissions seem to be something that is set for an organization, which makes sense because typically, "packages," and releases are something that is built within the context of a team.

So, we set up a seperate organization, [githubpackagetest](https://github.com/orgs/githubpackagestest/repositories) and transfered the repo, [dockerreactjs-yarn](https://github.com/githubpackagestest/dockerreactjs-yarn) over to this organization so that it could hypothetically not have a, "Permission Denied," error for the pwdelbloomboard user.

Notice that when the image was originally tagged, it was tagged with the username, "pwdel," which is not our username, it is actually, "pwdelbloomboard."  So to fix this, we first have to tag the build image correctly with:

```
docker build -t ghcr.io/pwdelbloomboard/ps-container:latest .
```
Then after verifying the build, try the proper push with:
```
docker push ghcr.io/pwdelbloomboard/ps-container:latest
```
Once this has been properly pushed, you will see a success message with:
```
The push refers to repository [ghcr.io/pwdelbloomboard/ps-container]
3e796ce6091b: Pushed 
32fc229d0177: Pushed 
4848bc580068: Pushed 
62595f19d440: Pushed 
446ec7c50f08: Pushed 
b8f0e895f520: Pushed 
f8700d3a252f: Pushed 
39982b2a789a: Pushed 
latest: digest: sha256:8f995...baf20 size: 1998
```
Now that we have successfully pushed to a registry, we see that the package is visible under our package menu:

![](/img/packagesview.png)

And clicking further into this specific package we see that we can link it to a repository:

![](/img/linktorepo.png)

Of course, since we had assigned our repo over to the organization above, since we thought that all of this needed to be under an organization, we no longer see the repo available as one we can assign a package to. So, after all of the above, we transferred this back over to [/pwdelbloomboard](https://github.com/pwdelbloomboard/dockerreactjs-yarn).

To connect it to the package, we have the following interface:

![](/img/connecttopackage.png)

To make this package usable and buildable, we're going to need to make its settings visible to the world / non-private. This can be set under [/user/packages/container/NAME/settings](https://github.com/users/pwdelbloomboard/packages/container/ps-container/settings).  At the bottom, go to the, "danger zone," and set to public.


3. [Connect a Repo to an Image](https://docs.github.com/en/packages/learn-github-packages/connecting-a-repository-to-a-package) Within Github, go to, "Profile" and then ["Packages."](https://github.com/pwdelbloomboard?tab=packages).  


You can use [Github Actions](https://docs.github.com/en/actions) to make this process automated, as shown below:

```
- name: Log into GitHub Container Registry
  run: echo "${{ secrets.GITHUB_TOKEN }}" | docker login https://ghcr.io -u ${{ github.actor }} --password-stdin

- name: Push image to GitHub Container Registry
  run: |
    IMAGE_ID=ghcr.io/${{ github.repository_owner }}/MyBeautifulContainer:123        
    docker push $IMAGE_ID
```
### Deployment Strategy

[Per this Tutorial](https://devopswithkubernetes.com/part-1/1-first-deploy#deployment)

To deploy an application we create a "Deployment," resource with the image.

```
kubectl create deployment {{METADATANAME}} --image={{IMAGE_NAME}}
  deployment.apps/{{METADATANAME}} created
```
So in our case, this will be:

```
kubectl create deployment buysellguess-dep --image=ghcr.io/pwdelbloomboard/ps-container
```
However, using the local machine that we're on, we might see an error like the following:

```
error: failed to create deployment: deployments.apps is forbidden: User "arn:aws:iam::846056206988:user/patrick" cannot create resource "deployments" in API group "apps" in the namespace "default"
```
[kubctl](https://kubernetes.io/docs/reference/generated/kubectl/kubectl-commands) is trying to make a deployment, but some of the settings already on our machine are set up to deploy to AWS via our user.

The idea behind creating a [deployment](/about-kubernetes/kubernetesdeployment.md) is that you're creating a completely seperate, walled off automation infrastructure, you are specifying how Nodes will be set up and used within a cluster. This is typical for K8s usage. Setting up a context is what walls off the deployment, and this is discussed further in [/about-kubernetes/kubectlconfig.md](/about-kubernetes/kubectlconfig.md)

That being said, if you already have a deployment setup, or if you wanted to just set things up as a demo, you could start off by setting up one Node at a time and manually killing it, for a simple application.

#### Running Off of Existing Deployment - Imperative, Terminal-Based Approach

So if we want to set up an app based upon our existing deployment, we can run the following command, which uses the previously registered image, "ghcr.io/pwdelbloomboard/ps-container"

```
kubectl create deployment buysellguess-dep --image=ghcr.io/pwdelbloomboard/ps-container
```
So once this is running, if we check our Docker desktop, we see that additional containers are running with the k3d label:

![](/img/k3d_initialdeploy.png)

Now if we run the "kubectl get deployments" and, "kubectl get pods," command, we see the following:

```
kubectl get deployments

NAME               READY   UP-TO-DATE   AVAILABLE   AGE
buysellguess-dep   1/1     1            1           34m

kubectl get pods

NAME                                READY   STATUS    RESTARTS   AGE
buysellguess-dep-7c476c64cd-trsl8   1/1     Running   0          35m
```

The above commands show that the deployments and pods are running as expected.

This deployment can be taken down with the following command:

```
kubectl delete deployment buysellguess-dep

deployment.apps "buysellguess-dep" deleted

```
Running either, "kubectl get deployments" or "kubectl get pods" should lead to a, "no resources found," message.

#### Declaratively Running Deployment with YAML File

We can create a deployment declaratively rather than through the command line.  Previously all deployments were run iteratively, through a correct series of commands, however that approach is cumbersome and may lead to mistakes, so the new way to do it is to use a yaml file and declaratively show everything that a deployment needs.

Within our [ReactJS Project](https://github.com/pwdelbloomboard/dockerreactjs-yarn/tree/main/app) in the application file, the same file where the Dockerfile resides, we may create a folder called, "manifests" and a file, /app/manifests/deployment.yaml.

To fill out this deployment, we need a container name, which means we first have to run the docker app off of our pre-build image:

```
sudo docker run -it --rm \
-v ${PWD}:/app \
-v /app/node_modules \
-p 3001:3000 \
-e CHOKIDAR_USEPOLLING=true \
--name buysellguessapp \
ghcr.io/pwdelbloomboard/ps-container
```

* From this, the, "name" of the container was set as, "buysellguessapp" - so this can be set within the deployment.yaml as, "containers: -name:"
* For the image itself, the id used is the full, sha256 id number which can be identified with:

```
docker inspect ghcr.io/pwdelbloomboard/ps-container
```

* More on what this number is is described at [docker.md](/about-docker/docker.md)

So in essense our manifests file should look like the following:

```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: buysellguess-dep
spec:
  replicas: 1
  selector:
    matchLabels:
      app: buysellguess
  template:
    metadata:
      labels:
        app: buysellguess
    spec:
      containers:
        - name: buysellguessapp
          image: 23f9401d2b7d275acf143295324d8f0bf3870988c3ea8432ce1db6b0e11bd36a
```
We then navigate to the application folder within terminal and use, "kubectl apply" to apply the manifest deployment.yaml settings.

```
kubectl apply -f manifests/deployment.yaml

deployment.apps/buysellguess-dep created
```
Now however if we are to run, "kubectl get pods" or "kubectl get deployments" we get the following results:

```
deployments...

NAME               READY   UP-TO-DATE   AVAILABLE   AGE
buysellguess-dep   0/1     1            0           38s

pods...

NAME                                READY   STATUS             RESTARTS   AGE
buysellguess-dep-7bf8648964-jts2k   0/1     InvalidImageName   0          47s
```
Which is showing that the deployments are not ready, due to the reason that the imagename is invalid.

This may have been because we set the image id of the container "buysellguessapp" as the original Docker source image id.

* Sha-256 Image ID (source image)

So instead of using that long string of ID number, we use the actual Built Docker image to run the application in the YAML file, as shown below.

```
      containers:
        - name: buysellguessapp
          image: ghcr.io/pwdelbloomboard/ps-container
```

After we switched this, we run "kubectl apply -f manifests/deployment.yaml" again and then ran "get pods," and "get deployments" and we find that everything is running. Whew!

Brief note on, "Apply" vs "Create" - in the context of using a .yaml file to create a deployment, "Create" and "Apply" seem to do the same thing.

#### Viewing the App with Port Forwarding

> port-forward command is mainly for debugging, also grants you temporary access to this application on your local machine. If you need to connect to the application under a specific port and you do not want to expose it using the Ingress object, then you should use the port-forward command.

So looking at our pod name with the, "get pods" command, we see the following:

```
kubectl get pods

NAME                                READY   STATUS    RESTARTS   AGE
buysellguess-dep-6867c7cfdf-4nlp5   1/1     Running   0          18m
```

So given the name that has been assigned by the deployment above, we can forward this pod using the, "port forwarding," command, like so:

```
kubectl port-forward buysellguess-dep-6867c7cfdf-4nlp5 8080:8080
```
However, when we try to reach this page at localhost:8080, we get the following error:

```
Forwarding from [::1]:8080 -> 8080
Handling connection for 8080
Handling connection for 8080
E1006 15:19:53.606951   41216 portforward.go:400] an error occurred forwarding 8080 -> 8080: error forwarding port 8080 to pod 41552e66a350a52b86a5235aa4fea4afb9e85a02ad671db8a89d049472c1c48d, uid : failed to execute portforward in network namespace "/var/run/netns/cni-703a8f70-f3bc-f4f4-7827-b3dc63c26756": failed to dial 8080: dial tcp4 127.0.0.1:8080: connect: connection refused
```
It could be that 8080 is just being used as a port already. Using port 3003, we do:

```
kubectl port-forward buysellguess-dep-6867c7cfdf-4nlp5 3003:3000
```
And then visiting http://localhost:3003 we see the page up and running.

```
$ kubectl port-forward buysellguess-dep-6867c7cfdf-4nlp5 3003:3000
Forwarding from 127.0.0.1:3003 -> 3000
Forwarding from [::1]:3003 -> 3000
Handling connection for 3003
Handling connection for 3003
Handling connection for 3003
Handling connection for 3003
Handling connection for 3003
```
Also notable, if we kill the Docker process for this app which had been running on localhost:3001, the deployment pod still works. This is because kubernetes is not drawing anything from the docker process and, "forwarding" it into kubernetes in any way, it is literally spinning up a new node, which contains the docker image (with the app) embedded inside of it, so in a sense, k8s not, "leveraging docker," k8s is running its own new containerized layer inside of a completely seperate system (e.g. within the node).

More of what the port forwarding process looks like is described under [about-networking](/about-networking/networking.md)

Port forwarding might be helpful for debugging, but it will be easier to understand if you set up services and ingress.

#### Set Up a Service and an Ingress

* [Introduction to Networking](https://devopswithkubernetes.com/part-1/3-introduction-to-networking)

#### Connecting Ports - 8081 to Server and 8082 to Agent port 30080

The objective of this step is to grant access through port 8081 to our server node (actually all nodes) and 8082 to one of our agent nodes port 30080. They will be used to showcase different methods of communicating with the servers.

> K3d has helpfully prepared us a port to access the API in 6443 and, in addition, has opened a port to 80. All requests to the load balancer here will be proxied to the same ports of all server nodes of the cluster. However, for testing purposes, we'll want an individual port open for a single node. Let's delete our old cluster and create a new one with port some ports open.

> K3d documentation tells us how the ports are opened, we'll open local 8081 to 80 in k3d-k3s-default-serverlb and local 8082 to 30080 in k3d-k3s-default-agent-0. The 30080 is chosen almost completely randomly, but needs to be a value between 30000-32767 for the next step:

```
k3d cluster create --port '8082:30080@agent[0]' -p 8081:80@loadbalancer --agents 2

INFO[0000] Prep: Network                                
INFO[0000] Re-using existing network 'k3d-k3s-default' (9a13490949d1c8f61699605d5e440e33e01cc18184e16ada6b8ea5c8b58cb380) 
INFO[0000] Created volume 'k3d-k3s-default-images'      
INFO[0001] Creating node 'k3d-k3s-default-server-0'     
INFO[0001] Creating node 'k3d-k3s-default-agent-0'      
INFO[0001] Creating node 'k3d-k3s-default-agent-1'      
INFO[0001] Creating LoadBalancer 'k3d-k3s-default-serverlb' 
INFO[0001] Starting cluster 'k3s-default'               
INFO[0001] Starting servers...                          
INFO[0001] Starting Node 'k3d-k3s-default-server-0'     
INFO[0013] Starting agents...                           
INFO[0013] Starting Node 'k3d-k3s-default-agent-0'      
INFO[0027] Starting Node 'k3d-k3s-default-agent-1'      
INFO[0036] Starting helpers...                          
INFO[0036] Starting Node 'k3d-k3s-default-serverlb'     
INFO[0037] (Optional) Trying to get IP of the docker host and inject it into the cluster as 'host.k3d.internal' for easy access 
INFO[0047] Successfully added host record to /etc/hosts in 4/4 nodes and to the CoreDNS ConfigMap 
INFO[0047] Cluster 'k3s-default' created successfully!  
INFO[0047] --kubeconfig-update-default=false --> sets --kubeconfig-switch-context=false 
INFO[0047] You can now use it like this:                
kubectl config use-context k3d-k3s-default
kubectl cluster-info

```

We can apply the settings from our deployment again with:

```
kubectl apply -f manifests/deployment.yaml

deployment.apps/buysellguess-dep created

```
> Now we have access through port 8081 to our server node (actually all nodes) and 8082 to one of our agent nodes port 30080. They will be used to showcase different methods of communicating with the servers.

#### What is a Service?

As Deployment resources took care of deployments for us. Service resource will take care of serving the application to connections from outside of the cluster.

![](/img/whatsaservice.png)

To take care of services, we can create a service.yaml file in the [manifest folder](https://github.com/pwdelbloomboard/dockerreactjs-yarn/tree/main/app/manifests) to declare the following:

* Declare that we want a Service
* Declare which port to listen to
* Declare the application where the request should be directed to
* Declare the port where the request should be directed to

```
apiVersion: v1
kind: Service
metadata:
  name: buysellguess-dep
spec:
  type: NodePort
  selector:
    app: buysellguess # This is the app as declared in the deployment.
  ports:
    - name: http
      nodePort: 30080 # This is the port that is available outside. Value for nodePort can be between 30000-32767
      protocol: TCP
      port: 1234 # This is a port that is available to the cluster, in this case it can be ~ anything
      targetPort: 3000 # This is the target port
```

The comments above explain what the declarations mean, including the 1. Matching the name: and the app: to our deployment.yaml, 2. Using port to port to the outside world, with protocol TCP.

Once we have added this file, we can use the, "kubectl apply" command:

```
kubectl apply -f manifests/service.yaml
  service/hashresponse-svc created
```
> As we've published 8082 as 30080 we can access it now via http://localhost:8082.

![](/img/buysellguessapp_connectedviaservice8082.png)


> We've now defined a nodeport with type: NodePort. NodePorts simply ports that are opened by Kubernetes to all of the nodes and the service will handle requests in that port. NodePorts are not flexible and require you to assign a different port for every application. As such NodePorts are not used in production but are helpful to know about.

> What we'd want to use instead of NodePort would be a LoadBalancer type service but this "only" works with cloud providers as it configures a, possibly costly, load balancer for it. We'll get to know them in part 3.

#### What is an Ingress?

![](/img/whatsaningress.png)

Whereas a Service is on Layer 4 of the OSI model, an Ingress is on Layer 7.

![](/img/osi7layer.png)

A more detailed breakdown shows that an ingress is more like the following:

![](/img/ingress_detail.png)

To set up an ingress, first we have to delete our previous service.

```
kubectl delete -f manifests/service.yaml
service "buysellguess-dep" deleted
```
Instead of the above service which defined a NodePort and Published on 30080, we are setting up on port 2345.

```
apiVersion: v1
kind: Service
metadata:
  name: buysellguess-svc
spec:
  type: ClusterIP
  selector:
    app: buysellguess
  ports:
    - port: 2345
      protocol: TCP
      targetPort: 3000
```
So once we have deleted the above previous service and replaced service.yaml with the simplier version, showing up on port 2345 with taretport 3000, 

We now add, "ingress.yaml" - 

```
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: buysellguess-dep-ingress
spec:
  rules:
  - http:
      paths:
      - path: /
        backend:
          serviceName: buysellguess-dep
          servicePort: 2345
```

We apply these new service.yaml and ingress.yaml and get:

```
kubectl apply -f manifests/service.yaml
service/buysellguess-dep created

kubectl apply -f manifests/ingress.yaml
Warning: extensions/v1beta1 Ingress is deprecated in v1.14+, unavailable in v1.22+; use networking.k8s.io/v1 Ingress
ingress.extensions/buysellguess-dep-ingress created
```
Verify that these exist with:

```
kubectl get svc

NAME               TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)    AGE
kubernetes         ClusterIP   10.43.0.1      <none>        443/TCP    3m37s
buysellguess-svc   ClusterIP   10.43.41.146   <none>        2345/TCP   4s

kubectl get ing
NAME                       CLASS    HOSTS   ADDRESS                            PORTS   AGE
buysellguess-dep-ingress   <none>   *       172.24.0.2,172.24.0.3,172.24.0.4   80      3m17s

```

> We can see that the ingress is listening on port 80. As we already opened port there we can access the application on http://localhost:8081.

As shown below:

![](/img/buysellguessapp_connectedviaingress8082.png)

![](/img/serviceingresstaxonomy.png)
#### Networking Between Pods

> Kubernetes includes a DNS service so communication between pods and containers in Kubernetes is as much of a challenge as it was with containers in docker-compose. Containers in a pod share the network. As such every other container inside a pod is accessible from localhost. 

> For communication between Pods a Service is used as they expose the Pods as a network service.

So in short:

* Containers in Pods share a network. Every container inside a pod is accessible from localhost.
* To communicate **between** Pods, we have to set up a Service and expose the Pods as a Network Service.

> The following creates a cluster-internal IP which will enable other pods in the cluster to access the port 8080 of "example" application from http://example-service. ClusterIP is the default type for a Service.

```
apiVersion: v1
kind: Service
metadata:
  name: example-service
spec:
  type: ClusterIP
  selector:
    app: example
  ports:
    - name: http
      protocol: TCP
      port: 80
      targetPort: 8080
```

You can go into a pod and make a request rather than spinning up a new pod:

```
kubectl exec -it example_pod
```

#### Namespaces

Namespaces are useful to organize highly complex K8s systems.

```
k get namespaces
NAME              STATUS   AGE
default           Active   69m
kube-system       Active   69m
kube-public       Active   69m
kube-node-lease   Active   69m
```

You can tag resources with namespaces.

To get resources under a particular namespace:

```
kubectl get pods -n default    
NAME                                READY   STATUS    RESTARTS   AGE
buysellguess-dep-6867c7cfdf-8h7rg   1/1     Running   0          69m
```

#### Configuring Applications


#### StatefulSets and Jobs



#### Monitoring



### One Node at a Time Strategy - No Deployment

* Running one Node at a time is not recommended.
#### Debugging

Information on debugging has been placed in a couple different locations:

* [about lens](/about-lens/lens.md)
* [about-kubernetes ](/about-kubernetes/kubernetes.md#)

# Questions Outstanding

* N/A at this time.

# Resources

* [k3d vs minikube vs kind](https://brennerm.github.io/posts/minikube-vs-kind-vs-k3s.html)
* [k3s + k3d = k8s: perfect match for dev and test](https://en.sokube.ch/post/k3s-k3d-k8s-a-new-perfect-match-for-dev-and-test-1)
* [Github Now has a Container Registry](https://dev.to/github/github-container-registry-better-than-docker-hub-1o9k)
* [Documentation on Creating a Github Personal Access Token](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token)
* [Devops with Kubernetes Tutorial](https://devopswithkubernetes.com/part-1/1-first-deploy)
* [Play with Kubernetes](https://labs.play-with-k8s.com/)
* [Create Multi-Node K8s Cluster with K3d](https://mohitgoyal.co/2021/05/28/create-multi-node-kubernetes-cluster-with-k3d/)
* [YouTube - Cheap Quick Kubernetes Development Cluster Using k3d/k3s](https://www.youtube.com/watch?v=jUPL4ZOlJ0E)
* [Run Your First App with Kubernetes](https://medium.com/@m.sedrowski/run-your-first-application-on-kubernetes-e54d5194e84b)
* [Introduction to Networking](https://devopswithkubernetes.com/part-1/3-introduction-to-networking)
* [Introduction to k3d on k3s](https://www.suse.com/c/introduction-k3d-run-k3s-docker-src/)
* [Rancher Blog - Setup k3d high Availability](https://rancher.com/blog/2020/set-up-k3s-high-availability-using-k3d)
* [k3d issue 642](https://github.com/rancher/k3d/issues/642)
* [Setup k3d for Local Testing and Development](https://thoughtexpo.com/setup-k3d-cluster-for-local-testing-or-development/)