# Dynamic Provisioning (DP)

## General About Dynamic Provisioning

* [Dynamic Provisioning Kubernetes Documentation](https://kubernetes.io/docs/concepts/storage/dynamic-provisioning/)

* Allows storage volumes to be created on-demand.
* Without DP, cluster admins have to manually make calls to their cloud or storage provider to create new storage volumes, and create PV objects to represent them in K8s.

* To enable dynamic provisioning, a cluster administrator needs to pre-create one or more StorageClass objects for users. StorageClass objects define which provisioner should be used and what parameters should be passed to that provisioner when dynamic provisioning is invoked. The name of a StorageClass object must be a valid DNS subdomain name.

The following manifest creates a storage class "slow" which provisions standard disk-like persistent disks.

```
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: slow
provisioner: kubernetes.io/gce-pd
parameters:
  type: pd-standard
```

* The following manifest creates a storage class "fast" which provisions SSD-like persistent disks.

```
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: fast
provisioner: kubernetes.io/gce-pd
parameters:
  type: pd-ssd
```

## About Storage Classes

* [About Storage Classes](https://kubernetes.io/docs/concepts/storage/storage-classes/)

Familiarity with volumes and persistent volumes is suggested.

> A StorageClass provides a way for administrators to describe the "classes" of storage they offer. Different classes might map to quality-of-service levels, or to backup policies, or to arbitrary policies determined by the cluster administrators. Kubernetes itself is unopinionated about what classes represent. This concept is sometimes called "profiles" in other storage systems.

* So to look at a StorageClass we can do:

```
$ k get StorageClass --all-namespaces
NAME                   PROVISIONER             RECLAIMPOLICY   VOLUMEBINDINGMODE
local-path (default)   rancher.io/local-path   Delete          WaitForFirstConsumer
```

* We can inspect the StorageClass with (used on Rancher):

```
$ k get StorageClass -o yaml
apiVersion: v1
items:
- apiVersion: storage.k8s.io/v1
  kind: StorageClass
  metadata:
    annotations:

      objectset.rio.cattle.io/id: ""
      objectset.rio.cattle.io/owner-gvk: k3s.cattle.io/v1, Kind=Addon
      objectset.rio.cattle.io/owner-name: local-storage
      objectset.rio.cattle.io/owner-namespace: kube-system
      storageclass.kubernetes.io/is-default-class: "true"
    creationTimestamp: "2022-10-28T18:28:36Z"
    labels:
      objectset.rio.cattle.io/hash: XXX
    name: local-path
    resourceVersion: "246"
    uid: XXX
  provisioner: rancher.io/local-path
  reclaimPolicy: Delete
  volumeBindingMode: WaitForFirstConsumer
kind: List
metadata:
  resourceVersion: ""
```


## Volumes

https://kubernetes.io/docs/concepts/storage/volumes/

> On-disk files in a container are ephemeral, which presents some problems for non-trivial applications when running in containers. One problem is the loss of files when a container crashes. The kubelet restarts the container but with a clean state. A second problem occurs when sharing files between containers running together in a Pod. The Kubernetes volume abstraction solves both of these problems.

In short, similar to a Docker volume.

### Re: Shared Volumes

* "Storage backend with support for shared volumes"

* What a shared volume is does not seem to be explicitly spelled out in the K8s documentation, but it seems to be implied that it is simply a volume that is shared between pods.  E.g. no special resource type is needed, it’s just that more than one pod shares a volume.
https://kubernetes.io/docs/concepts/storage/volumes/#using-subpath 

## Persistent Volumes

https://kubernetes.io/docs/concepts/storage/persistent-volumes/

* Whereas a volume is something that is attached, "exterior to" a cluster, a PersistentVolume is a storage class which gets provisioned within a cluster.

* They have lifecycles independent of pods, but are dependant upon the cluster.

> A PersistentVolume (PV) is a piece of storage in the cluster that has been provisioned by an administrator or dynamically provisioned using Storage Classes. It is a resource in the cluster just like a node is a cluster resource. PVs are volume plugins like Volumes, but have a lifecycle independent of any individual Pod that uses the PV. This API object captures the details of the implementation of the storage, be that NFS, iSCSI, or a cloud-provider-specific storage system.

* So basically if a pod gets deleted, the Persistent Volume still exists, but if the cluster gets deleted, the Persistent Volume gets deleted.

* A, "claim," is an allocation or request for storage by a user (which can be a particular application).