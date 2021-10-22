# About Storage Classes

> A StorageClass provides a way for administrators to describe the "classes" of storage they offer. Different classes might map to quality-of-service levels, or to backup policies, or to arbitrary policies determined by the cluster administrators. Kubernetes itself is unopinionated about what classes represent. This concept is sometimes called "profiles" in other storage systems.

## Storage Class Resource

* provisioner - an actual cloud services provider provisioner such as aws-ebs
* parameters
* reclaimPolicy
* mountOptions
* volumeBindingMode

```
apiVersion: storage.k8s.io/v1
kind: StorageClass
metadata:
  name: standard
provisioner: kubernetes.io/aws-ebs
parameters:
  type: gp2
reclaimPolicy: Retain
allowVolumeExpansion: true
mountOptions:
  - debug
volumeBindingMode: Immediate
```

# Resources

* [Storage Classes](https://kubernetes.io/docs/concepts/storage/storage-classes/)

