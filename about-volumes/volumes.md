# About Kubernetes Volumes

* [About Volumes in the Context of Docker](/about-docker/docker-volumes.md)


"Volumes," in general is a confusing term because it has multiple specific definitions across different environments.

* You can mount a USB disk to your physical machine
* You can mount disk volume to your docker VM
* 

## Docker Volumes

Docker has a concept of volumes, though it is somewhat looser and less managed. A Docker volume is a directory on disk or in another container. Docker provides volume drivers, but the functionality is somewhat limited.

## Kubernetes Volumes

There are multiple different types of Kubernetes Volumes.

### configmap Volume

* [configmap Volume](https://kubernetes.io/docs/concepts/storage/volumes/#configmap)

A ConfigMap provides a way to inject configuration data into pods. The data stored in a ConfigMap can be referenced in a volume of type configMap and then consumed by containerized applications running in a pod.

```
apiVersion: v1
kind: Pod
metadata:
  name: configmap-pod
spec:
  containers:
    - name: test
      image: busybox
      volumeMounts:
        - name: config-vol
          mountPath: /etc/config
  volumes:
    - name: config-vol
      configMap:
        name: log-config
        items:
          - key: log_level
            path: log_level
```

### local Volume

[local volume](https://kubernetes.io/docs/concepts/storage/volumes/#local)

> A local volume represents a mounted local storage device such as a disk, partition or directory.

> Local volumes can only be used as a statically created PersistentVolume. Dynamic provisioning is not supported.

> Compared to hostPath volumes, local volumes are used in a durable and portable manner without manually scheduling pods to nodes. The system is aware of the volume's node constraints by looking at the node affinity on the PersistentVolume.

> However, local volumes are subject to the availability of the underlying node and are not suitable for all applications. If a node becomes unhealthy, then the local volume becomes inaccessible by the pod. The pod using this volume is unable to run. Applications using local volumes must be able to tolerate this reduced availability, as well as potential data loss, depending on the durability characteristics of the underlying disk.


### persistantvolumeclaim

> A persistentVolumeClaim volume is used to mount a PersistentVolume into a Pod. PersistentVolumeClaims are a way for users to "claim" durable storage (such as a GCE PersistentDisk or an iSCSI volume) without knowing the details of the particular cloud environment.

## Kubernetes Mount Propagation

[Mount Propagation](https://kubernetes.io/docs/concepts/storage/volumes/#mount-propagation)

Mount propagation allows for sharing volumes mounted by a container to other containers in the same pod, or even to other pods on the same node.





# Resources

* [Kubernetes Storage Volumes](https://kubernetes.io/docs/concepts/storage/volumes/)
* [How to Use Kubernetes Volumes - New Relic](https://newrelic.com/blog/how-to-relic/how-to-use-kubernetes-volumes)
* [TutorialsPoint - Kubernetes Volumes](https://www.tutorialspoint.com/kubernetes/kubernetes_volumes.htm)
* [Keeping the State of Apps - Volume Mounts](https://www.kubermatic.com/blog/keeping-the-state-of-apps-1-introduction-to-volume-and-volumemounts/)
* [Unofficial Kubernetes Documentation](https://unofficial-kubernetes.readthedocs.io/en/latest/concepts/storage/volumes/)
* [Kubernetes Storage an In-Depth Look](https://cloud.netapp.com/blog/cvo-blg-kubernetes-storage-an-in-depth-look)