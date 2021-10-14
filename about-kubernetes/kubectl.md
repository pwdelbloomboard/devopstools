# The API Resources

On the command line, api resources can be found with:

```
kubectl api-resources
```

## Clusters

The way to look at and work with clusters is through the [Kubernetes API](https://kubernetes.io/docs/tasks/administer-cluster/access-cluster-api/).

The way to check the location and credentials that kubectl knows about with this command:

```
kubectl config view
```

kubctl is just a way of conecting to the Kubernetes API, there is no way to turn on and off clusters or resources with kubctl, as with k3d.

## Getting All Resources Under a Namespace

```
kubectl get all
```

## Getting Pods and Deployments


```
kubectl get deployments
NAME               READY   UP-TO-DATE   AVAILABLE   AGE
buysellguess-dep   1/1     1            1           34m
```


```
kubectl get pods

NAME                                READY   STATUS    RESTARTS   AGE
buysellguess-dep-7c476c64cd-trsl8   1/1     Running   0          35m
```

## Debugging

### kubectl describe deployment

will describe details about the deployment.

```
kubectl describe deployment buysellguess-dep

Name:                   buysellguess-dep
Namespace:              default
CreationTimestamp:      Wed, 06 Oct 2021 14:58:13 -0500
Labels:                 <none>
Annotations:            deployment.kubernetes.io/revision: 1
Selector:               app=buysellguess
Replicas:               1 desired | 1 updated | 1 total | 1 available | 0 unavailable
StrategyType:           RollingUpdate
MinReadySeconds:        0
RollingUpdateStrategy:  25% max unavailable, 25% max surge
Pod Template:
  Labels:  app=buysellguess
  Containers:
   buysellguessapp:
    Image:        ghcr.io/pwdelbloomboard/ps-container
    Port:         <none>
    Host Port:    <none>
    Environment:  <none>
    Mounts:       <none>
  Volumes:        <none>
Conditions:
  Type           Status  Reason
  ----           ------  ------
  Available      True    MinimumReplicasAvailable
  Progressing    True    NewReplicaSetAvailable
OldReplicaSets:  <none>
NewReplicaSet:   buysellguess-dep-6867c7cfdf (1/1 replicas created)
Events:
  Type    Reason             Age   From                   Message
  ----    ------             ----  ----                   -------
  Normal  ScalingReplicaSet  38m   deployment-controller  Scaled up replica set buysellguess-dep-6867c7cfdf to 1

```
### kubectl describe pod

```
kubectl describe pod buysellguess-dep-6867c7cfdf-4nlp5
Name:         buysellguess-dep-6867c7cfdf-4nlp5
Namespace:    default
Priority:     0
Node:         k3d-buysellguess-server-0/172.20.0.2
Start Time:   Wed, 06 Oct 2021 14:58:13 -0500
Labels:       app=buysellguess
              pod-template-hash=6867c7cfdf
Annotations:  <none>
Status:       Running
IP:           10.42.0.12
IPs:
  IP:           10.42.0.12
Controlled By:  ReplicaSet/buysellguess-dep-6867c7cfdf
Containers:
  buysellguessapp:
    Container ID:   containerd://9f4e2dabe308dca0e1ea63fb066faf4ec3abaf23b30f0e255bbe4bcd37a25a8e
    Image:          ghcr.io/pwdelbloomboard/ps-container
    Image ID:       ghcr.io/pwdelbloomboard/ps-container@sha256:8f9959b48435e7689fa8a091fac9140766e061510daa4ed0db263c05ca4baf20
    Port:           <none>
    Host Port:      <none>
    State:          Running
      Started:      Wed, 06 Oct 2021 14:58:14 -0500
    Ready:          True
    Restart Count:  0
    Environment:    <none>
    Mounts:
      /var/run/secrets/kubernetes.io/serviceaccount from kube-api-access-hns5x (ro)
Conditions:
  Type              Status
  Initialized       True 
  Ready             True 
  ContainersReady   True 
  PodScheduled      True 
Volumes:
  kube-api-access-hns5x:
    Type:                    Projected (a volume that contains injected data from multiple sources)
    TokenExpirationSeconds:  3607
    ConfigMapName:           kube-root-ca.crt
    ConfigMapOptional:       <nil>
    DownwardAPI:             true
QoS Class:                   BestEffort
Node-Selectors:              <none>
Tolerations:                 node.kubernetes.io/not-ready:NoExecute op=Exists for 300s
                             node.kubernetes.io/unreachable:NoExecute op=Exists for 300s
Events:
  Type    Reason     Age   From               Message
  ----    ------     ----  ----               -------
  Normal  Scheduled  39m   default-scheduler  Successfully assigned default/buysellguess-dep-6867c7cfdf-4nlp5 to k3d-buysellguess-server-0
  Normal  Pulling    39m   kubelet            Pulling image "ghcr.io/pwdelbloomboard/ps-container"
  Normal  Pulled     39m   kubelet            Successfully pulled image "ghcr.io/pwdelbloomboard/ps-container" in 615.6774ms
  Normal  Created    39m   kubelet            Created container buysellguessapp
  Normal  Started    39m   kubelet            Started container buysellguessapp

```

### kubectl get events -w

* 

### kubectl get pv

* Getting persistent volumes.


### kubectl logs

Gets the logs of a pod.

```
kubectl logs buysellguess-dep-6867c7cfdf-4nlp5

Example output:

yarn run v1.22.5
$ react-scripts start
ℹ ｢wds｣: Project is running at http://10.42.0.12/
ℹ ｢wds｣: webpack output is served from 
ℹ ｢wds｣: Content not from webpack is served from /public
ℹ ｢wds｣: 404s will fallback to /
Starting the development server...

Compiled with warnings.

src/App.js
  Line 35:18:  The href attribute is required for an anchor to be keyboard accessible. Provide a valid, navigable address as the href value. If you cannot provide an href, but still need the element to resemble a link, use a button and change it with appropriate styles. Learn more: https://github.com/evcohen/eslint-plugin-jsx-a11y/blob/master/docs/rules/anchor-is-valid.md  jsx-a11y/anchor-is-valid

Search for the keywords to learn more about each warning.
To ignore, add // eslint-disable-next-line to the line before.

```

### kubctl cluster-info

"kubectl cluster-info" gives us information about the cluster itself, and the control pane.

```
Kubernetes control plane is running at https://127.0.0.1:6445
CoreDNS is running at https://127.0.0.1:6445/api/v1/namespaces/kube-system/services/kube-dns:dns/proxy
Metrics-server is running at https://127.0.0.1:6445/api/v1/namespaces/kube-system/services/https:metrics-server:/proxy
```

### kubectl get pods


### kubectl logs [PODNAME] -f

### kubectl get logs -l

Gets the logs of a particular pod.

### kubectl get pods --namespace [default]

### kubectl get pods [namespace]


### kubctl describe pod [ingress]


### kubectl get ingress

### kubctl logs [INGRESSNAME] -f

> This one does not seem to work

# Resources

* [Intro to Debugging](https://devopswithkubernetes.com/part-1/2-introduction-to-debugging)