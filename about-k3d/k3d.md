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

#### Merging Kubeconfig - kubeconfig merge

##### Useful Flags

--trace (super verbose)
--verbose

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
### Deployment

[Per this Tutorial](https://devopswithkubernetes.com/part-1/1-first-deploy#deployment)

To deploy an application we create a "Deployment," resource with the image.

```
kubectl create deployment {{METADATANAME}} --image={{IMAGE_NAME}}
  deployment.apps/{{METADATANAME}} created
```
So in our case, this will be:

```
kubectl create deployment buysellguess-dep --image=ghcr.io/pwdelbloomboard/ps-container
  deployment.apps/buysellguess-dep created
```
However, using the local machine that we're on, we might see an error like the following:

```
error: failed to create deployment: deployments.apps is forbidden: User "arn:aws:iam::846056206988:user/patrick" cannot create resource "deployments" in API group "apps" in the namespace "default"
```
[kubctl](https://kubernetes.io/docs/reference/generated/kubectl/kubectl-commands) is trying to make a deployment, but some of the settings already on our machine are set up to deploy to AWS via our user.

So, what we have to do in this situation is to create a new context, which is discussed further in [/about-kubernetes/kubectlconfig.md](/about-kubernetes/kubectlconfig.md)


# Creating a Setup Deployment

A setup, "hello world," deployment would include:

* Kubectl
* Docker
* k3d / k3s.


# Resources

* [k3d vs minikube vs kind](https://brennerm.github.io/posts/minikube-vs-kind-vs-k3s.html)
* [k3s + k3d = k8s: perfect match for dev and test](https://en.sokube.ch/post/k3s-k3d-k8s-a-new-perfect-match-for-dev-and-test-1)
* [Github Now has a Container Registry](https://dev.to/github/github-container-registry-better-than-docker-hub-1o9k)
* [Documentation on Creating a Github Personal Access Token](https://docs.github.com/en/authentication/keeping-your-account-and-data-secure/creating-a-personal-access-token)
* [Devops with Kubernetes Tutorial](https://devopswithkubernetes.com/part-1/1-first-deploy)
* [Play with Kubernetes](https://labs.play-with-k8s.com/)
