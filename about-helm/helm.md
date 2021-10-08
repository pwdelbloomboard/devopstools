# About Helm

[Helm](https://helm.sh/) is a package manager for Kubernetes.

Kubernetes is managed and set up through a series of .yaml objects.  These .yaml objects are used to set up:

* Services
* Pods
* Config Maps
* Secrets
* Deployments
* Etc.

Helm allows you to manage all of these automatically.

Allows you to template your yaml file, so rather than having to re-deploy a new app needing to set up new yaml scripts, it applies the template and then you're ready to go.

Rather than writing a script, we tell kubernetes exactly what we want, and kubernetes makes it happen. Basically it does this by making sure that the actual state matches the desired states. The desired states are represented by the yaml files.


## Example

An application may have:

* deployment.yaml
- what container to deploy


## What is the Definition a Kubernetes Package?

* Within the context of Kubernetes, a package would essentially be "state settings" for virtual machines that are designed to help operate a service or piece of software specifically built for the context of Kubernetes, and possibly specifically customized for certain conditions (such as high useage or low usage). 
* [ArtifactHub](https://artifacthub.io/) is one website that hosts a collection of Kubernetes packages, for example, "MariaDB for Kubernetes."

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

## How to Use Helm, Helmcharts and Tiller

* You can think of Helm as, "apt" or "homebrew" for Kubernetes.

Kubernetes includes:

* pods
* services

Package Yaml files to control the following, bundled together is known as, "helm charts.":

* Stateful set
* Configmap
* k8s User Permissions
* Secret
* Services

Well known services all have existing configurations that can be found. The capability to share Helm charts is part of why helm became so popular, through public and private registries.

### Helm Templating Engine

* A kuberentes may have a collection of microservices. 
* The difference between yaml files between each of these microservices may be a few different values.

A template file is a sort of, "template," that defines configuration values within {{variables}} for a microservice - these variables can be re-set based upon which microservice you are configuring, rather than having to write out many different template files.

In essence, it's:

1) Common blueprint.
2) Dynamic values replaced by placeholders.

This may come in to use if you have different levels of configuration between development, staging, or production, where different values are needed for the various microservices.

### Example Helm Start Structure

* mychart
- chart.yml // main info about the chart
- values.yaml // the values stored
- /charts/ // where you store charts
- /templates/ // where you store templates

#### values.yaml

imageName: myapp
port: 8080
version: 1.0.0

You can also create overrides with versioning.

### Release Management

* Client is helm CLI
* Server is "Tiller," 

Whenever you create a new chart execution, the changes are recorded by Tiller, similar to the git paradigm - basically chart execution history.

Downsides of Tiller - there is a security issue in that it has too much power inside the k8s cluster, but beyond Helm version  

# Helm Chart Creation




## Resources

* [Collection of Helm Tutorials](https://jfrog.com/blog/10-helm-tutorials-to-start-your-kubernetes-journey/)
* [Introductory Tutorial on Helm](https://www.youtube.com/watch?v=fy8SHvNZGeE)
* [Helm in Kubernetes](https://www.youtube.com/watch?v=-ykwb1d0DXU)