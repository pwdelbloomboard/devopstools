# Starting Out

## Spinning Up an Debian Container and Manual Install

To start off with, we can simply create a Dockerfile which has Debian and do an attempted manual install:

```
wget -q -O - https://raw.githubusercontent.com/k3d-io/k3d/main/install.sh | bash
```

When we do this, there is no output, because -q is quiet mode.

```
wget -O - https://raw.githubusercontent.com/k3d-io/k3d/main/install.sh | bash
```
When we run the above, we get the error:

```
wget -O - https://raw.githubusercontent.com/k3d-io/k3d/main/install.sh | bash
--2022-05-03 17:41:25--  https://raw.githubusercontent.com/k3d-io/k3d/main/install.sh
Resolving raw.githubusercontent.com (raw.githubusercontent.com)... 185.199.109.133, 185.199.111.133, 185.199.108.133, ...
Connecting to raw.githubusercontent.com (raw.githubusercontent.com)|185.199.109.133|:443... connected.
ERROR: The certificate of 'raw.githubusercontent.com' is not trusted.
ERROR: The certificate of 'raw.githubusercontent.com' doesn't have a known issuer.
```
This was because due to how we were building our docker image, several important dependencies including the certificate file was not installing properly. After fixing this we were able to run:

```
wget -O - https://raw.githubusercontent.com/k3d-io/k3d/main/install.sh | bash
```
With results:

```
--2022-05-03 18:34:59--  https://raw.githubusercontent.com/k3d-io/k3d/main/install.sh
Resolving raw.githubusercontent.com (raw.githubusercontent.com)... 185.199.108.133, 185.199.111.133, 185.199.110.133, ...
Connecting to raw.githubusercontent.com (raw.githubusercontent.com)|185.199.108.133|:443... connected.
HTTP request sent, awaiting response... 200 OK
Length: 5282 (5.2K) [text/plain]
Saving to: 'STDOUT'

-                                                   100%[=================================================================================================================>]   5.16K  --.-KB/s    in 0.001s

2022-05-03 18:34:59 (5.36 MB/s) - written to stdout [5282/5282]

Preparing to install k3d into /usr/local/bin
k3d installed into /usr/local/bin/k3d
Run 'k3d --help' to see what you can do with it.
```
Checking out the version we have:

```
k3d version
k3d version v5.4.1
k3s version v1.22.7-k3s1 (default)
```

## Automatically Installing In Dockerfile

This was pretty easy, the following line just had to be added to the Dockerfile:

```
RUN wget -O - https://raw.githubusercontent.com/k3d-io/k3d/main/install.sh | bash
```

## Spinning Up a Test Cluster

So just to start out, we attmempted to spin up a test cluster with:

```
 k3d cluster create testcluster
```
Which lamentably resulted in:
```
ERRO[0000] Failed to get nodes for cluster 'testcluster': docker failed to get containers with labels 'map[k3d.cluster:testcluster]': failed to list containers: Cannot connect to the Docker daemon at unix:///var/run/docker.sock. Is the docker daemon running?
INFO[0000] Prep: Network
ERRO[0000] Failed Cluster Preparation: Failed Network Preparation: failed to create cluster network: failed to check for duplicate docker networks: docker failed to list networks: Cannot connect to the Docker daemon at unix:///var/run/docker.sock. Is the docker daemon running?
ERRO[0000] Failed to create cluster >>> Rolling Back
INFO[0000] Deleting cluster 'testcluster'
ERRO[0000] Failed to get nodes for cluster 'testcluster': docker failed to get containers with labels 'map[k3d.cluster:testcluster]': failed to list containers: Cannot connect to the Docker daemon at unix:///var/run/docker.sock. Is the docker daemon running?
ERRO[0000] failed to get cluster: No nodes found for given cluster
FATA[0000] Cluster creation FAILED, also FAILED to rollback changes!
```

Why are we getting these errors?  First off, we have to explore what docker.sock is:

### docker.sock and k3d

So first off, what is docker.sock?

* docker.sock is the UNIX socket that Docker daemon is listening to. It's the main entry point for Docker API. It also can be TCP socket but by default for security reasons Docker defaults to use UNIX socket.

* Docker cli client uses this socket to execute docker commands by default. You can override these settings as well.

* There might be different reasons why you may need to mount Docker socket inside a container. Like launching new containers from within another container. Or for auto service discovery and Logging purposes. This increases attack surface so you should be careful if you mount docker socket inside a container there are trusted codes running inside that container otherwise you can simply compromise your host that is running docker daemon, since Docker by default launches all containers as root.

Basically, what k3d is doing is spinning up new docker containers. Typically k3d is installed on a machine such as a laptop, but in this case we're actually installing it within a Docker container, so this is like installing a docker container within a docker container.  In order to do that we have to mount the docker socket as a volume by adding the following to our docker-compose.yaml:

```
    volumes:
      - /var/run/docker.sock:/var/run/docker.sock
```

Essentially what this is doing, is creating a bind mount between our local machine's Unix socket that the Docker daemon is listening to at /var/run/docker.sock and connecting this to the actual Docker container's socket at /var/run/docker.sock, essentially connecting our physical machine disk to the container. This is not recommended for production as it is insecure.

After running through and re-building and re-installing everything, we get:

```
k3d cluster create testcluster
INFO[0000] Prep: Network
INFO[0000] Created network 'k3d-testcluster'
INFO[0000] Created image volume k3d-testcluster-images
INFO[0000] Starting new tools node...
INFO[0001] Pulling image 'ghcr.io/k3d-io/k3d-tools:5.4.1'
INFO[0001] Creating node 'k3d-testcluster-server-0'
INFO[0002] Pulling image 'docker.io/rancher/k3s:v1.22.7-k3s1'
INFO[0002] Starting Node 'k3d-testcluster-tools'
INFO[0007] Creating LoadBalancer 'k3d-testcluster-serverlb'
INFO[0007] Pulling image 'ghcr.io/k3d-io/k3d-proxy:5.4.1'
INFO[0009] Using the k3d-tools node to gather environment information
INFO[0009] Starting new tools node...
INFO[0010] Starting Node 'k3d-testcluster-tools'
INFO[0011] Starting cluster 'testcluster'
INFO[0011] Starting servers...
INFO[0011] Starting Node 'k3d-testcluster-server-0'
INFO[0019] All agents already running.
INFO[0019] Starting helpers...
INFO[0020] Starting Node 'k3d-testcluster-serverlb'
INFO[0027] Injecting records for hostAliases (incl. host.k3d.internal) and for 3 network members into CoreDNS configmap...
INFO[0029] Cluster 'testcluster' created successfully!
INFO[0029] You can now use it like this:
kubectl cluster-info
```

Which shows that the cluster can be succesffully installed.  Note that:

* Load Balancer was created.
* server-0 was created.
* testcluster-tools was created.

Due to the nature of how docker works and our bind mount, we were able to actually view the containers on our local machine desktop, so there was no reason to further install Docker on the container itself to inspect the containers...everything is viewable by the API across the socket.


### Docker Socket Running Containers on Local

So because we connected the 

## Installing kubectl

After doing an install for kubctl with curl, this was added to the Dockerfile, along with an alias, "k" for kubectl.

From here, we can perform several functions:

```
k cluster-info
Kubernetes control plane is running at https://host.docker.internal:44889
CoreDNS is running at https://host.docker.internal:44889/api/v1/namespaces/kube-system/services/kube-dns:dns/proxy
Metrics-server is running at https://host.docker.internal:44889/api/v1/namespaces/kube-system/services/https:metrics-server:https/proxy

To further debug and diagnose cluster problems, use 'kubectl cluster-info dump'.
root@c7294bbadc15:/# k get all
NAME                 TYPE        CLUSTER-IP   EXTERNAL-IP   PORT(S)   AGE
service/kubernetes   ClusterIP   10.43.0.1    <none>        443/TCP   26m
root@c7294bbadc15:/# k get nodes
NAME                       STATUS   ROLES                  AGE   VERSION
k3d-testcluster-server-0   Ready    control-plane,master   26m   v1.22.7+k3s1
```
### Checking Out Docker Containers

As mentioned above, the Docker containers could be inspected simply on the local machine, for example with, "docker ps -a" which resulted in:

```
CONTAINER ID   IMAGE                              COMMAND                  CREATED              STATUS              PORTS                             NAMES
f0f3eb807094   ghcr.io/k3d-io/k3d-tools:5.4.1     "/app/k3d-tools noop"    About a minute ago   Up About a minute                                     k3d-testcluster-tools
687cb1ca5c63   ghcr.io/k3d-io/k3d-proxy:5.4.1     "/bin/sh -c nginx-pr…"   About a minute ago   Up About a minute   80/tcp, 0.0.0.0:43087->6443/tcp   k3d-testcluster-serverlb
28f4ad83ae18   rancher/k3s:v1.22.7-k3s1           "/bin/k3s server --t…"   About a minute ago   Up About a minute                                     k3d-testcluster-server-0
c980d4ace2f4   k3dpurposed_debian:bullseye-slim   "bash"                   2 minutes ago        Up 2 minutes                                          playwithk3d_container
```

# Notes and Observations

### Main Ingress

The ingress is already baked in with ghcr.io/k3d-io/k3d-proxy:5.4.1

...which means port forwarding is not needed.

### Pods

These are basically all of the pods necessary to run the cluster.

```
kubectl get pods -A
NAMESPACE     NAME                                      READY   STATUS      RESTARTS   AGE
kube-system   local-path-provisioner-84bb864455-4s2mb   1/1     Running     0          6m23s
kube-system   coredns-96cc4f57d-5s9gr                   1/1     Running     0          6m23s
kube-system   helm-install-traefik-crd--1-bd642         0/1     Completed   0          6m23s
kube-system   svclb-traefik-ltnsv                       2/2     Running     0          6m2s
kube-system   helm-install-traefik--1-bdxmq             0/1     Completed   1          6m23s
kube-system   metrics-server-ff9dbcb6c-ksdqc            1/1     Running     0          6m23s
kube-system   traefik-56c4b88c4b-x9bvd                  1/1     Running     0          6m2s
```

### Nodes

There is only one node per cluster by default with k3d.

```
k get nodes
NAME                       STATUS   ROLES                  AGE     VERSION
k3d-testcluster-server-0   Ready    control-plane,master   3m34s   v1.22.7+k3s1
```

### Specifying Version

We could specify a version of a cluster with the, "--image" flag to mimic something closer to what is in production, so basically, an older cluster.

```
k3d cluster create another-cluster --image rancher/k3s:v1.20.4-k3s1
```

### Deleting the Cluster

You can list/delete clusters with:

```
k3d cluster list
k3d cluster delete
```

### Working with Configuration Files and k3d

Looking at: [Victor Farcic's Example](https://github.com/vfarcic/k3d-demo):

We can declaratively set up a cluster with the following options:

```
kind: Simple
apiVersion: k3d.io/v1alpha2
name: my-cluster
image: rancher/k3s:v1.20.4-k3s1
servers: 3
agents: 3
ports:
- port: 80:80
  nodeFilters:
  - loadbalancer
# options:
#   k3s:
#     extraServerArgs:
#     - --no-deploy=traefik
```
Basically, 3 servers, 3 agents, serving on port 80:80, has a loadbalancer, etc.

So if you wanted to use that example you could do:

```
git clone https://github.com/vfarcic/k3d-demo.git
```
To copy that above file, and then declaratively set up your cluster based upon that yaml file, which then can be used to create an actual cluster with:

```
k3d cluster create --config k3d.yaml
```
After this was created, we got some failures which may have been inherent in creating multiple-node clusters on a docker image, rather than doing it on a local machine.

### Attempting kubectl apply

* Kubectl apply is what is used to actually install the package into the cluster.

So using our above git repository, we attempt to apply with:

```
# kubectl apply --filename k8s/
deployment.apps/devops-toolkit created
ingress.networking.k8s.io/devops-toolkit created
service/devops-toolkit created
```
However, to be able to view this application we have to have a port available, which we did not set originally within our docker-compose.

So within docker-compose we can certainly set the port to 80 or whatever, although it's probably better to use a less used port.

We may also need to activate the, "extra_hosts" command:

```
    extra_hosts:
      - "host.docker.internal:host-gateway"
```
Further, from k3d, we may need to bind the host in order to expose the service.

From [k3d documentation on exposing services](https://k3d.io/v5.4.1/usage/exposing_services/):

```
k3d cluster create --api-port 6550 -p "8081:80@loadbalancer"
```
Of course we have to think carefully about what the different ports are doing.

When we run the following:

```
kubectl apply --filename k8s/
```
Within our k8s/ folder we are setting the port as 80:

```
spec:
  type: ClusterIP
  ports:
  - port: 80
    targetPort: 80
```
Therefore if we look at the actual ingress with kubectl get ingress:

```
NAME             CLASS    HOSTS   ADDRESS      PORTS   AGE
devops-toolkit   <none>   *       172.24.0.3   80      106s
```
Basically, this should work on all hosts.

Within our docker-compose.yaml file we have:

```
    ports: 
      - 80:80
    extra_hosts:
      - "host.docker.internal:host-gateway"      
```

So hypothetically we should be connecting port 80 within the container (all hosts) to port 80 on local, so that just regular, "localhost/" should work.

However, we don't see anything on localhost. But, we do see something on:


```
localhost:8081
```

So what happened here with the networking?  Let's take a closer look:

```
k3d cluster create --api-port 6550 -p "8081:80@loadbalancer"
```
Basically, the flag, "--api-port" specifies the Kubernetes API server port exposed on the load balancer. In other words, we are connecting the API port 6550 to the Host (8081) at the loadbalancer, 80.

Meanwhile, the --port flag maps ports from the node containers via the server load balancer to the host according to a specified format. So basically, our container port, 80, which was set in our various yaml files is connected to the host (node) port, 8081.

Since we connect Docker to our desktop/dev machine at 80 via docker-compose.yaml under, "-port," then the, "host" (meaning the k3d node) which is visible at 8081 is connected to that localhost port, allowing us to access our app at 8081.

```
--api-port [HOST:]HOSTPORT                                       

Specify the Kubernetes API server port exposed on the LoadBalancer (Format: [HOST:]HOSTPORT)
- Example: `k3d cluster create --api-port 0.0.0.0:6550`

 -p, --port [HOST:][HOSTPORT:]CONTAINERPORT[/PROTOCOL][@NODEFILTER]   
 Map ports from the node containers (via the serverlb) to the host (Format: [HOST:][HOSTPORT:]CONTAINERPORT[/PROTOCOL][@NODEFILTER])

```
The actual app comes from the image that we pull down, which is specified in our deployment.yaml:

```
apiVersion: apps/v1
kind: Deployment
metadata:
  name: devops-toolkit
  labels:
    app: devops-toolkit
spec:
  selector:
    matchLabels:
      app: devops-toolkit
  template:
    metadata:
      labels:
        app: devops-toolkit
    spec:
      containers:
      - name: devops-toolkit
        image: vfarcic/devops-toolkit-series
```

Basically the image, "vfarcic/devops-toolkit-series" which defaults to the [image on Dockerhub](https://hub.docker.com/r/vfarcic/devops-toolkit-series/tags), which is the web's main Docker image registry. 


## Resources

* [Installing k3d in Docker Containers](https://techviewleo.com/run-rancher-k3s-kubernetes-in-docker-containers/)
* [k3d locally using Rancher K3D - YouTube, DevOps Toolkit](https://www.youtube.com/watch?v=mCesuGk-Fks)
* [Running Docker Container in a Docker Container](https://devopscube.com/run-docker-in-docker/)
* [Docker Socket](https://stackoverflow.com/questions/35110146/can-anyone-explain-docker-sock)
* [Setting Up K3d - Linkedin](https://www.linkedin.com/pulse/install-k3d-linux-kubernetes-installation-guide-prayag-sangode?trk=articles_directory)
* [Victor Farcic's Example](https://github.com/vfarcic/k3d-demo)
* [k3d documentation on exposing services](https://k3d.io/v5.4.1/usage/exposing_services/)