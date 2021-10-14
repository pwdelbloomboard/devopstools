# Helm Template Debugging

The reference documentation for debugging helm templates can be found at:

* [Helm Template Debugging](https://helm.sh/docs/chart_template_guide/debugging/)

> Debugging templates can be tricky because the rendered templates are sent to the Kubernetes API server, which may reject the YAML files for reasons other than formatting.

## Helm Charts vs Helm Templates

 > Helm Charts are source trees that contain a self-descriptor file, Chart.yaml, and one or more templates. 
 
 > Templates are Kubernetes manifest files that describe the resources you want to have on the cluster. Helm uses the Go templating engine by default.

## Helm Prerequisites

> The following prerequisites are required for a successful and properly secured use of Helm.

> 1. A Kubernetes cluster
> 2. Deciding what security configurations to apply to your installation if any 
> 3. Installing and configuring Helm

### Cluster Prerequisite

* Based upon the above, it's important to double check to ensure there is a cluster running in order to run helm.

So if we have an application, "mysampleapplication" on which we expect a cluster to be present, and assuming we are using k3d, we should be able to view the cluster with the following:

```
k3d cluster list

NAME          SERVERS   AGENTS   LOADBALANCER
k3s-default   1/1       2/2      true
local         1/1       0/0      true
```

However, we don't see an, "application" cluster listed, only a, "k3s-default" cluster. This likely means we're good to go, but if a specific, named application cluster is needed, this can be revisited.

### Security Prerequisites

### Helm Installed and Configured

## Helm Repositories

* [Chart Repository](https://helm.sh/docs/topics/chart_repository/)

* Charts are published to a registry or repo and have release versions.However, Charts can also be locally stored and used and don't have to pull from the cloud.

### Helm Commands

#### helm search hub [SEARCHTERM]

This command searches the Artifact hub and returns options for pre-published charts.

```
helm search hub nginx

URL                                               	CHART VERSION	APP VERSION         	DESCRIPTION                                       
https://artifacthub.io/packages/helm/dysnix/nginx 	7.1.8        	1.19.4              	Chart for the nginx server                        
https://artifacthub.io/packages/helm/shubhamtat...	0.1.12       	1.19.6              	Nginx Helm chart for Kubernetes        

...

```
#### helm install

> To install a new package, use the helm install command. At its simplest, it takes two arguments: A release name that you pick, and the name of the chart you want to install.

> Note that installing a chart creates a new release object. 

> During installation, the helm client will print useful information about which resources were created, what the state of the release is, and also whether there are additional configuration steps you can or should take.

> Helm does not wait until all of the resources are running before it exits. Many charts require Docker images that are over 600M in size, and may take a long time to install into the cluster.


##### Customizing Charts

* [Customizing Charts Before Installing](https://helm.sh/docs/intro/using_helm/#customizing-the-chart-before-installing)


#### Other Installation Methods

[Other Helm Chart Installation Methods](https://helm.sh/docs/intro/using_helm/#more-installation-methods)

* A chart repository (as we've seen above)
* A local chart archive (helm install foo foo-0.1.1.tgz)
* An unpacked chart directory (helm install foo path/to/foo)
* A full URL (helm install foo https://example.com/charts/foo-1.2.3.tgz)


#### helm upgrade and helm rollback

* [helm upgrade and rollback](https://helm.sh/docs/intro/using_helm/#helm-upgrade-and-helm-rollback-upgrading-a-release-and-recovering-on-failure)


## Looking at a Template

As an example, we can look at our application, the [buysellguessapp](https://github.com/pwdelbloomboard/dockerreactjs-yarn/tree/main/app/helm-buysellguessapp).

We have 


## 




