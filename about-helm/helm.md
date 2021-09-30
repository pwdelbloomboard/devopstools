# About Helm

[Helm](https://helm.sh/) is a package manager for Kubernetes.

Basically, it helps described the structure of an application through charts with helmfiles.

## What is the Definition a Kubernetes Package?

Within the context of Kubernetes, a package would essentially be a service or piece of software specifically built for the context of Kuberentes. [ArtifactHub](https://artifacthub.io/) is one website that hosts a collection of Kubernetes packages, for example, "MariaDB for Kubernetes."

## Overview of Helm

* Example - Ecommerce Application, includes 1. NodeJS Application, with two replicas. 2. MongoDB Database. 3. NodePort service to access the service.
* To deploy this type of application step, is to write a couple yaml files.

Key elements of YAML files:

* Values.YAML
* *deployment
- image: node/mongo1
- replicas: 2
* *service
- type: NodePort
- port: 8080
* Deployment.YAML
- image: {{values.deployment.image}}
- replicas: {{values.deployment.replicas}} // this looks up to the Values.YAML file.
* Service.YAML
- type: {{values.svc.type}}
- port: {{values.svc.port}}

Helm holds:
1. Configuration.
2. Chart (template)

* Chart consists of all of the files that you template.

How do you inject values into a template?

You can write values into a Values.YAML file, and then have Deployment.YAML and Service.YAML refer to those values within their structure.

Then, when you deploy, you run a CLI command:

```
helm install myApp
```
This will go into, "Teller" which maps services and deployments into a way that Kubernetes will understand. This can also configure the entire, "hardware setup," e.g. kubernetes setup in a way that scales up and down resources in a, "hot," manner, basically things get templated and sent over to Kubernetes dynamically. This is done as followS:

```
helm upgrade myApp
```

If there was a mistake, helm keeps a list of previous configurations, you can roll back to the last known configuration with:

```
helm rollback myApp
```

You can also work with repos, make use of repos, deploying a configuration by using:

```
helm package
```

This allows anyone in your organization to use the same package, or sets of packages for various scenarios under different operating conditions (for example, the holidays are over an a particular ECommerce store doesn't need as many resrouces anymore, you can set up a different helm package for, "low maintainence mode.")



## Resources

* [Introductory Tutorial on Helm](https://www.youtube.com/watch?v=fy8SHvNZGeE)