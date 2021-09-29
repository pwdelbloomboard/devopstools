# Interpreting a Debug Log

## How is a Debug Log Generated

Refer to [debuglog.md](/understanding-scripts/debuglog)

## What are the Main Sections of debug.txt?

* DNS - is the ping time unresaonbly long?
* docker info 
    - is the memory being used dangerously close to the system memory?
    - are the number of CPUs expected correct?

* docker-mac-info
    - are the number of CPU's correct?
    - is the memory appropriate?

* docker ps -a
    - are any docker processes interfering with each other?
    
* docker system df
    - are any docker images interfering with each other

* k3d cluster list
    - shows the servers actually working, or at least the number of servers

* kubectl --context=k3d-local version
    - gives info on what version of k3d is actually being used

* kubectl --context=k3d-local cluster-info 
    - 

* helm --kube-context=k3d-local ls --all
    - shows all local k3d clusters as well as their status, whether deployed or not

* kubectl --context=k3d-local get nodes -o wide 
    - shows k3d-local-server status, roles, version

* kubectl --context=k3d-local get all --all-namespaces
    - shows status of all kubernetes namespaces

Note - if the namespace for objects created by the Kubernetes system keeps crashing and restarting, you will see a "CrashloopBackOff" error - which, could be caused by either the application inside the container keeps crashing, some type of parameters of the pod or container are not being set up correctly, or some error has been made when deploying kubernetes, or something else.

    - also shows the CLUSTER-IP, EXTERNAL-IP and PORTS of the various services.
    - also shows daemonset of various apps which may include mongo, redis, nginx, postgres, etc.

* kubectl --context=k3d-local get ingress --all-namespaces

    - shows local ingress endpoints for various services, such as nginx, prometheus, or various apps.

* kubectl --context=k3d-local get persistentvolumes --all-namespaces 

    - shows volumes, such as postgres data volumes, mongo data volumes, aws or any seperate applications.

* kubectl --context=k3d-local get persistentvolumeclaims --all-namespaces 

    - A PersistentVolume (PV) is a piece of storage in the cluster that has been provisioned by an administrator or dynamically provisioned using Storage Classes. It is a resource in the cluster just like a node is a cluster resource. PVs are volume plugins like Volumes, but have a lifecycle independent of any individual Pod that uses the PV. This API object captures the details of the implementation of the storage, be that NFS, iSCSI, or a cloud-provider-specific storage system.
    - A PersistentVolumeClaim (PVC) is a request for storage by a user. It is similar to a Pod. Pods consume node resources and PVCs consume PV resources. Pods can request specific levels of resources (CPU and Memory). Claims can request specific size and access modes (e.g., they can be mounted ReadWriteOnce, ReadOnlyMany or ReadWriteMany, see AccessModes).

* kubectl --context=k3d-local describe nodes 

    - describes k3d server nodes in details, with labels, annotations, conditions, addresses, capacity, allocatable capability, system info, 
    - conditions describe various conditions of the network, memory, disk, PID and overall readiness.
    - non-terminated pods show a list of all default pods as well as their age, CPU Requests, CPU Limits, Memory Requests, Memory Limits, etc.
    - total resource limits are also shown - including cpu, memory, ephemeral storage and sofourth. This is a good place to look at whether memory or cpu is being overloaded.