# Deployment vs. No Deployment

![](/img/k8sdeployment.png)

A, "Deployment," is a way to tell Kubernetes what container you want, how they should be running and how many of them should be running, in an automatic way. The opposite of a deployment would be, "running pods directly."

Use cases for deployments include:

* Running multiple instances of an application
* Scaling the number of instances of an application up or down
* Updating every running instance of an application
* Rolling back all running instances of an application to another version

A ReplicaSet resource is a way to tell how many replicas of a Pod you want. It will delete or create Pods until the number of Pods you wanted are running. ReplicaSets are managed by Deployments and you should not have to manually define or modify them.

Deployments will create pods that have a naming convention like so:

> mongo-8b6c997f-mgv7l

This is because they are being automatically created.

* Deployments do not by default decide where applications run on a cluster. Deployments can be configured to show where applications run in the cluster. A good example use case for that would be if you had jobs you wanted to run on nodes with GPUs for data crunching.
* Nodes with GPUs are more expensive, so instead of making all of the nodes in your cluster have GPUs you spin up some nodes with GPUs and tag those nodes as such. Then when creating your deployment you can tell it to run only on nodes (or not to run on nodes) with that specific tag.

### Services

* Deployments and Services are often used in tandem: Deployments working to define the desired state of the application and Services working to make sure communication between almost any kind of resource and the rest of the cluster is stable and adaptable. It is highly recommended that most workloads use both, but in some cases that may not make sense depending on application behavior. Here are some overviews of what would happen if you chose not to run a deployment or a service for your application.

* Without a deployment - Running a pod without a deployment can be done, but it is generally not recommended. For very simple testing this may be an effective method to increase velocity but for anything of importance this approach has a number of flaws.

* Without a deployment, Pods can still be created and run through unmanaged ReplicaSets. While you will still be able to scale your application you lose out on a lot of base functionality deployments provide and drastically increase your maintenance burden. Kubernetes now recommends running almost all Pods in Deployments instead of using custom ReplicaSets.

### daemonsets



### statefulset

* statefulsets are like deployments that give the pods fixed hostnames

### running pods directly



### ingress

services