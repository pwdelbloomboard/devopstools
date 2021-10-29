# Ingresses
## Getting Ingress DNS's

The easiest way to get the ingress, which is a configuration that exposes an application (via a service)

This presupposes knowledge about, "k get pods" or "kubectl get pods" and the capability to obtain a podname.

```
kubectl get ingress | grep [part-of-podname]
```

The abbreviated, alias version of this command would be, "kgi | grep [part-of-podname]"
### What's going on with this command here?

Here's a breakdown of the above command.

* kubectl get ingress ... "pulls down" all of the ingresses
* | or pipe takes the output of that first command and gives it to the next command, which is...
* grep (which means filter given the search term)
* [part-of-podname] ... which literally represents the part of the podname
### Example Use of kubectl get ingress

* So for example, if you want to find the DNS for your application, which is sitting on a pod, and you found that the pod name is, "example123" -- you can run the command:

```
kubectl get ingress | grep example
```
This would show all of the pod-ingresses which have the word, "example" in it, so you would be able to find, "example123" in that listing, which would be formatted as follows:

```
NAME        CLASS   HOSTS                           ADDRESS     PORTS   AGE
example123  <none>  example123.local.domain.com     172.23.0.3  80      1h23m
```

The alias (shortened) format of the, "kubectl get ingress," command is:

```
kgi | grep [part-of-podname]
```