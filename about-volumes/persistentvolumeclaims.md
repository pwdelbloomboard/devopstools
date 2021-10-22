# Persistent Volume Claims

* [Kubernetes Persistent Volumes](https://kubernetes.io/docs/concepts/storage/persistent-volumes/)

> A PersistentVolume (PV) is a piece of storage in the cluster that has been provisioned by an administrator or dynamically provisioned using Storage Classes. It is a resource in the cluster just like a node is a cluster resource. PVs are volume plugins like Volumes, but have a lifecycle independent of any individual Pod that uses the PV. This API object captures the details of the implementation of the storage, be that NFS, iSCSI, or a cloud-provider-specific storage system.

> A PersistentVolumeClaim (PVC) is a request for storage by a user. It is similar to a Pod. Pods consume node resources and PVCs consume PV resources. Pods can request specific levels of resources (CPU and Memory). Claims can request specific size and access modes (e.g., they can be mounted ReadWriteOnce, ReadOnlyMany or ReadWriteMany, see AccessModes).

### Types of PersistentVolumes

* awsElasticBlockStore - AWS Elastic Block Store (EBS)
* azureDisk - Azure Disk
* azureFile - Azure File
* cephfs - CephFS volume
* csi - Container Storage Interface (CSI)
* fc - Fibre Channel (FC) storage
* flexVolume - FlexVolume
gcePersistentDisk - GCE Persistent Disk
glusterfs - Glusterfs volume
hostPath - HostPath volume (for single node testing only; WILL NOT WORK in a multi-node cluster; consider using local volume instead)
iscsi - iSCSI (SCSI over IP) storage
local - local storage devices mounted on nodes.
nfs - Network File System (NFS) storage
portworxVolume - Portworx volume
rbd - Rados Block Device (RBD) volume
vsphereVolume - vSphere VMDK volume

# Resources

* [Pod Has Unbound Persistent Volume Claim](https://stackoverflow.com/questions/52668938/pod-has-unbound-persistentvolumeclaims)