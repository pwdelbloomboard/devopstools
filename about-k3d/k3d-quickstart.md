# k3d Quickstart Guide

## Foundational k3d Stuff

### k3d create registry

* [k3d registries guide](https://k3d.io/usage/guides/registries/)

```
k3d registry create registry.localhost --port 5000
```

* This creates a registry at localhost:5000
* This will show up as a docker container k3d-registry.localhost

The purpose of this local registry is to prevent having to send an image to a remote repo every time it is built.

### k3d local cluster

```
k3d cluster create local                                                                                \
            --registry-use k3d-registry.localhost:5000                                                          \
            --k3s-server-arg '--disable=traefik'                                                                \
            --k3s-server-arg '--disable=metrics-server'                                                         \
            --k3s-server-arg '--kubelet-arg=eviction-hard=imagefs.available<3%,nodefs.available<3%'             \
            --k3s-server-arg '--kubelet-arg=eviction-minimum-reclaim=imagefs.available=1%,nodefs.available=1%'  \
            --port 27017:27017@loadbalancer                                                                     \
            --port 5432:5432@loadbalancer                                                                       \
            --port 80:80@loadbalancer                                                                           
```

* We create a local registry
* We use the registry we created above.
* Use ports

- 27017
- 5432
- 80

## Creating a Cluster

* Create a Cluster by the name, "buysellguess" with ports at 8082, opening up to a loadbalancer on 8081.
* One agent.

![](/img/k8ssetupdiagram.png)

```
k3d cluster create buysellguess --port '8082:30080@agent[0]' -p 8081:80@loadbalancer --agents 1
```

After this point, we will have two clusters and a registry: 

* The (1) local registry
* The local cluster which is just a (2) server and a (3) load balancer
* The, "buysellguess" cluster which includes a (4) server, (5) load balancer and an (6) agent.

The above servers, load balancers registry and agent each are contained within their own container, for a total of 6 containers.

```
k3d cluster list
NAME           SERVERS   AGENTS   LOADBALANCER
buysellguess   1/1       1/1      true
local          1/1       0/0      true
```

![](/img/six-containers.png)

However, at this point if we attempt to access the application at the designated port, localhost:8081, we get a 404 error:

![](/img/404erroratlocalhost.png)

This is because we have not, "deployed" the application to our cluster.

## Creating the Deployment

Once we have completed the above, we can create a deployment.

### Designing the Manifests

We will need three manifests:

* deployment.yaml

Sets the name, template and image we draw off of to, "inject" into the agent.

* service.yaml

Connects the ingress and the agent together on TCP

* ingress.yaml

Allows the outside world to view the agent through the load balancer.

#### Visual Representation of Manifest

![](/img/k8ssetupdiagram.png)

### Creating the Deployment Using "kubectl apply"

Assuming our manifests to our existing cluster that we had set up in the steps above, by navigating into the appropriate folder within the terminal and running::

```
kubectl apply -f manifests/deployment.yaml

deployment.apps/buysellguess-dep created
```

#### Verifying Deployment is Up and Running

We can verify that the deployment is up and running with:

```
kubectl get deployments
NAME               READY   UP-TO-DATE   AVAILABLE   AGE
buysellguess-dep   1/1     1            1           22m
```

### Creating a Service, Verifying It's Running

We can then set up the service with a similar command:

```
kubectl apply -f manifests/service.yaml

kubectl get svc

NAME                                    TYPE           CLUSTER-IP      EXTERNAL-IP   PORT(S)                         AGE
kubernetes                              ClusterIP      10.43.0.1       <none>        443/TCP                         64m
...
buysellguess-svc                        ClusterIP      10.43.61.61     <none>        2345/TCP                        22s

```

Note that the service is listed on port 2345 as a TCP service.


### Creating an Ingress, Verifying It's Running

Finally we can create an ingress and verify that it is running:

```
kubectl apply -f manifests/ingress.yaml

kubectl get svc

Warning: extensions/v1beta1 Ingress is deprecated in v1.14+, unavailable in v1.22+; use networking.k8s.io/v1 Ingress
ingress.extensions/buysellguess-dep-ingress created

kubectl get ing

NAME                       CLASS    HOSTS   ADDRESS                 PORTS   AGE
buysellguess-dep-ingress   <none>   *       172.28.0.2,172.28.0.3   80      8s

```
Note that the ingress connects on port 80, which we had originaly specified above as the port that the loadbalancer would be connected to with the command:

```
k3d cluster create buysellguess --port '8082:30080@agent[0]' -p 8081:80@loadbalancer --agents 1
```

However, even though we have set up the deployment, ingress and the service above, after having created the cluster as mentioned, we still get a 404 not found error.

```
kubectl get pods
NAME                                READY   STATUS    RESTARTS   AGE
buysellguess-dep-6867c7cfdf-lphf7   1/1     Running   0          4m34s

kubectl get nodes
NAME                        STATUS   ROLES                  AGE   VERSION
k3d-buysellguess-server-0   Ready    control-plane,master   15h   v1.21.1+k3s1
k3d-buysellguess-agent-0    Ready    <none>                 15h   v1.21.1+k3s1

```
