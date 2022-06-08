# Kubernetes Documentation and Basics

* The most basic thing to understand about kubernetes is the deployment, service and ingress paradigm.  It's important to understand how these three items fit together, and that there is a rough breakdown of how the different of how deployment, service and ingress configurations all hook together through metadata tags and selectors, e.g.:

![k8smeta](about-kubernetes/img/k8s_meta.png)

* The so-called, "Kubernetes Documentation" is a starting point, a sort of collection of tutorials which can be used to understand the general workflow of how everything fits together. This can be found at [Kubernetes.io](https://kubernetes.io/docs/concepts/).

* The Kubernetes API is considerably more complex, but goes into absolute detail on what every possible option within every possible field of every possible API version. There is a [Kubernetes Documentation Reference API link](https://kubernetes.io/docs/reference/kubernetes-api/) as well as [individual links to individual API versions, with this case being v1.21.0](https://kubernetes.io/docs/reference/generated/kubernetes-api/v1.21/).

### Technique

To work through and build-out Kubernetes, the important thing is to understand the concepts with the, "Kubernetes Concepts Documentation," first and then drill down into the API itself to build something specific.