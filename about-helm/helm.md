# Helm

## About Helm

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

# Helm Conceptual Overview

Helm is a, "package manager," for k8s, and its packages are, "charts," which are essentially sets of yaml files and tempaltes which instruct helm how to send, "releases."  These releases get sent to, "tiller" which operates, installs the chart configuration and deletes old releases in an actual k8s deployment.

![](/img/helm-architecture.png)

Where Helm fits in a general workflow is at the, "release and operate," stage, wherein a deployment of a set of images is being pushed into operation, (perhaps this could be development, staging or production), and then specifying how a cluster or set of images will behave within operation.

![](/img/helm-workflow.png)

Reiterating the Helm architecture, the charts may be stored in various locations, one strategy is to put them in a completely seperate repo from the main, "code / project" repo for the application in question. However they may also reside right within the application itself.

![](/img/helm_repo_description.png)

A developer, perhaps working through a local machine or through a remote server, sends helm commands to a remote cloud server which runs tiller, which then follows through on actions which make the release and delete possible within the k8s deployment. Tiller itself sits on the k8s cluster as well.

![](/img/helmclient.png)

As far as how a developer actually uses Helm, they physically work on building and writing templates which make up a chart, those charts then get pushed to a repo as a, "chart," and then that repo gets deployed (via either a local machine or remote service such as Rundeck), which then goes to tiller and updates or operates the release on the k8s cluster.

So essentially, developers are:

* Working on the templates
* Working on pushing commands using the helm CLI which pushes the templates to a particular release.

![](/img/howdevelopersusehelm.png)

## Example

An application may have:

* deployment.yaml
- what container to deploy
* secret.yaml
* persistent_storage.yaml

If you wanted to create a new application, you would have to copy-paste all of these files, tons of yaml files and duplication. 

> Helm groups all of these together in templates. It brings all of these yaml files together in what is called a "Chart," which has a, "templates folder," that holds a file which passes all of the values. Helm will inject values and parameters to make everything happen, using a, "Values File."

* We can inject values and parameters, allowing us to re-use charts among many applications and micro-services.
* Re-use it with different applications with different names.
* We can pass the values as a, "file" rather than a series of command line commands.

Helm will take the, "Values Files," into the, "Helm Chart," makes the chart generic.

### Secondary Overview of Helm

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

> You can write values into a Values.YAML file, and then have Deployment.YAML and Service.YAML refer to those values within their structure.

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


## What is the Definition a "Kubernetes Package?"

When talking about Helm, the term, "package manager," gets thrown around a lot. What does this really mean in the context of helm?

* Within the context of Kubernetes, a package would essentially be "state settings" for virtual machines that are designed to help operate a service or piece of software specifically built for the context of Kubernetes, and possibly specifically customized for certain conditions (such as high useage or low usage). 
* [ArtifactHub](https://artifacthub.io/) is one website that hosts a collection of Kubernetes packages, for example, "MariaDB for Kubernetes."


### How to Use Helm, Helmcharts and Tiller

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

The below is a basic overview tutorial on how to create a Helm Chart.
### Create a Kubernetes Cluster.

To do this, we created a cluster using k3d, which is described more in the [tutorial on k3d](/about-k3d/k3d.md)

For this project, we have the following k3d setup going on our local machine, including all of the servers shown:

![](/img/k8s-setupforhelm.png)

Once you have this ready to go, you can enter into the Helm CLI.

### Installing the Helm CLI

* Next, a normal recommended process is to actually set up a small alpine linux container where you can install and play with helm, just to keep things clean and isolated away from your main machine that you might be working on, and to keep things consistent.
* You should also make sure kubectl is installed.
* For using Helm, Make sure you have the latest version of Helm installed by going to Github, looking at the releases, and downloading the most recent binary.

Once you have helm installed and apprpriately isolated, you can access Helm on the command line with:

```
helm

The Kubernetes package manager

Common actions for Helm:

- helm search:    search for charts
- helm pull:      download a chart to your local directory to view
- helm install:   upload the chart to Kubernetes
- helm list:      list releases of charts

Environment variables:

| Name                               | Description                                                                       |
|------------------------------------|-----------------------------------------------------------------------------------|
| $HELM_CACHE_HOME                   | set an alternative location for storing cached files.                             |
| $HELM_CONFIG_HOME                  | set an alternative location for storing Helm configuration.                       |
| $HELM_DATA_HOME                    | set an alternative location for storing Helm data.                                |
| $HELM_DEBUG                        | indicate whether or not Helm is running in Debug mode                             |
| $HELM_DRIVER                       | set the backend storage driver. Values are: configmap, secret, memory, postgres   |
| $HELM_DRIVER_SQL_CONNECTION_STRING | set the connection string the SQL storage driver should use.                      |
| $HELM_MAX_HISTORY                  | set the maximum number of helm release history.                                   |
| $HELM_NAMESPACE                    | set the namespace used for the helm operations.                                   |
| $HELM_NO_PLUGINS                   | disable plugins. Set HELM_NO_PLUGINS=1 to disable plugins.                        |
| $HELM_PLUGINS                      | set the path to the plugins directory                                             |
| $HELM_REGISTRY_CONFIG              | set the path to the registry config file.                                         |
| $HELM_REPOSITORY_CACHE             | set the path to the repository cache directory                                    |
| $HELM_REPOSITORY_CONFIG            | set the path to the repositories file.                                            |
| $KUBECONFIG                        | set an alternative Kubernetes configuration file (default "~/.kube/config")       |
| $HELM_KUBEAPISERVER                | set the Kubernetes API Server Endpoint for authentication                         |
| $HELM_KUBECAFILE                   | set the Kubernetes certificate authority file.                                    |
| $HELM_KUBEASGROUPS                 | set the Groups to use for impersonation using a comma-separated list.             |
| $HELM_KUBEASUSER                   | set the Username to impersonate for the operation.                                |
| $HELM_KUBECONTEXT                  | set the name of the kubeconfig context.                                           |
| $HELM_KUBETOKEN                    | set the Bearer KubeToken used for authentication.                                 |

Helm stores cache, configuration, and data based on the following configuration order:

- If a HELM_*_HOME environment variable is set, it will be used
- Otherwise, on systems supporting the XDG base directory specification, the XDG variables will be used
- When no other location is set a default location will be used based on the operating system

By default, the default directories depend on the Operating System. The defaults are listed below:

| Operating System | Cache Path                | Configuration Path             | Data Path               |
|------------------|---------------------------|--------------------------------|-------------------------|
| Linux            | $HOME/.cache/helm         | $HOME/.config/helm             | $HOME/.local/share/helm |
| macOS            | $HOME/Library/Caches/helm | $HOME/Library/Preferences/helm | $HOME/Library/helm      |
| Windows          | %TEMP%\helm               | %APPDATA%\helm                 | %APPDATA%\helm          |

Usage:
  helm [command]

Available Commands:
  completion  generate autocompletion scripts for the specified shell
  create      create a new chart with the given name
  dependency  manage a chart's dependencies
  diff        Preview helm upgrade changes as a diff
  env         helm client environment information
  get         download extended information of a named release
  help        Help about any command
  history     fetch release history
  install     install a chart
  lint        examine a chart for possible issues
  list        list releases
  package     package a chart directory into a chart archive
  plugin      install, list, or uninstall Helm plugins
  pull        download a chart from a repository and (optionally) unpack it in local directory
  repo        add, list, remove, update, and index chart repositories
  rollback    roll back a release to a previous revision
  search      search for a keyword in charts
  show        show information of a chart
  status      display the status of the named release
  template    locally render templates
  test        run tests for a release
  uninstall   uninstall a release
  upgrade     upgrade a release
  verify      verify that a chart at the given path has been signed and is valid
  version     print the client version information

Flags:
      --debug                       enable verbose output
  -h, --help                        help for helm
      --kube-apiserver string       the address and the port for the Kubernetes API server
      --kube-as-group stringArray   group to impersonate for the operation, this flag can be repeated to specify multiple groups.
      --kube-as-user string         username to impersonate for the operation
      --kube-ca-file string         the certificate authority file for the Kubernetes API server connection
      --kube-context string         name of the kubeconfig context to use
      --kube-token string           bearer token used for authentication
      --kubeconfig string           path to the kubeconfig file
  -n, --namespace string            namespace scope for this request
      --registry-config string      path to the registry config file (default "/Users/patrick.delaneybloomboard.com/Library/Preferences/helm/registry.json")
      --repository-cache string     path to the file containing cached repository indexes (default "/Users/patrick.delaneybloomboard.com/Library/Caches/helm/repository")
      --repository-config string    path to the file containing repository names and URLs (default "/Users/patrick.delaneybloomboard.com/Library/Preferences/helm/repositories.yaml")

Use "helm [command] --help" for more information about a command.

```

### About Charts and Starting to Create a Helm Chart

Charts have:

* Name
* Template Folder
* Values File

1. Using our [example application](https://github.com/pwdelbloomboard/dockerreactjs-yarn), we create a temp folder where we can create a chart.
2. Navigate into the folder for the app itself, we're going to create all of the files we need in the, "Option3" approach discussed below - wherein we maintain a service-specific chart in the same repo as the service itself. So we navigate to the, /app folder.
2. "helm create" - will create charts in the appropriate file.

```
helm create helm-buysellguessapp
Creating buysellguessapp

ls
Chart.yaml  charts  templates  values.yaml

```
We are now presented with a default file strucutre which helm created on our behalf within the structure.


3. Chart File Structure

*Folders*

* charts - you can nest charts in this folder. Embed more charts

* templates - all the yaml you would like to bundle up to form a chart. can be existing yaml files.

```
NOTES.txt  _helpers.tpl  deployment.yaml  hpa.yaml  ingress.yaml  service.yaml	serviceaccount.yaml  tests
```

* tests - you can craft your own chart by 

*Files*

* Chart.yaml - containing information about our chart, version, description, etc.

```
apiVersion: v2
name: buysellguessapp
description: A Helm chart for Kubernetes

# A chart can be either an 'application' or a 'library' chart.
#
# Application charts are a collection of templates that can be packaged into versioned archives
# to be deployed.
#
# Library charts provide useful utilities or functions for the chart developer. They're included as
# a dependency of application charts to inject those utilities and functions into the rendering
# pipeline. Library charts do not define any templates and therefore cannot be deployed.
type: application

# This is the chart version. This version number should be incremented each time you make changes
# to the chart and its templates, including the app version.
# Versions are expected to follow Semantic Versioning (https://semver.org/)
version: 0.1.0

# This is the version number of the application being deployed. This version number should be
# incremented each time you make changes to the application. Versions are not expected to
# follow Semantic Versioning. They should reflect the version the application is using.
# It is recommended to use it with quotes.
appVersion: "1.16.0"
```

* values.yaml - default configuration values, we can store stuff like override image tag, override our chart make our chart generic.
- Values can be made per file per service. values.dev.yaml or values.prod.yaml, for example.

```
# Default values for buysellguessapp.
# This is a YAML-formatted file.
# Declare variables to be passed into your templates.

replicaCount: 1

image:
  repository: nginx
  pullPolicy: IfNotPresent
  # Overrides the image tag whose default is the chart appVersion.
  tag: ""

imagePullSecrets: []
nameOverride: ""
fullnameOverride: ""

serviceAccount:
  # Specifies whether a service account should be created
  create: true
  # Annotations to add to the service account
  annotations: {}
  # The name of the service account to use.
  # If not set and create is true, a name is generated using the fullname template
  name: ""

podAnnotations: {}

podSecurityContext: {}
  # fsGroup: 2000

securityContext: {}
  # capabilities:
  #   drop:
  #   - ALL
  # readOnlyRootFilesystem: true
  # runAsNonRoot: true
  # runAsUser: 1000

service:
  type: ClusterIP
  port: 80

ingress:
  enabled: false
  annotations: {}
    # kubernetes.io/ingress.class: nginx
    # kubernetes.io/tls-acme: "true"
  hosts:
    - host: chart-example.local
      paths:
      - path: /
        backend:
          serviceName: chart-example.local
          servicePort: 80
  tls: []
  #  - secretName: chart-example-tls
  #    hosts:
  #      - chart-example.local

resources: {}
  # We usually recommend not to specify default resources and to leave this as a conscious
  # choice for the user. This also increases chances charts run on environments with little
  # resources, such as Minikube. If you do want to specify resources, uncomment the following
  # lines, adjust them as necessary, and remove the curly braces after 'resources:'.
  # limits:
  #   cpu: 100m
  #   memory: 128Mi
  # requests:
  #   cpu: 100m
  #   memory: 128Mi

autoscaling:
  enabled: false
  minReplicas: 1
  maxReplicas: 100
  targetCPUUtilizationPercentage: 80
  # targetMemoryUtilizationPercentage: 80

nodeSelector: {}

tolerations: []

affinity: {}

```

You can explore the values file to see what kind of examples are in there.

#### Convert a Default Chart to be Useable Chart

Doing this is fairly straightforward. Basically all we need to do is copy our existing k8s yaml into the empty helm yaml (in the template files) that we just created and use it as-is.

* delete everything you don't need.

Look at charts.yaml. We can set up the version.

* don't delete the helper file. 

* don't change the name of the file or folder after creation - this may change.

Basically, do the following:

1. Replace deployments.yaml, ingress.yaml and service.yaml within the templates folder.  Same for any secrets and config map if those exist.

We now have a helm chart without having done any changes, and it should work 100% according to what we had already rendered on k8s.

2. To test this out, we run the, "helm template" command.

Which has the format: helm template [NAME] [CHART] [flags]

* Meaning, [NAME] of the new template we want to create.
* [CHART] - is the name of our existing helm folder we just created above.

```
cd helm-buysellguessapp

helm template -h

helm template buysellguessapp-helmtemplate helm-buysellguessapp

---
# Source: helm_buysellguessapp/templates/serviceaccount.yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: buysellguessapp_helmtemplate-helm_buysellguessapp
  labels:
    helm.sh/chart: helm_buysellguessapp-0.1.0
    app.kubernetes.io/name: helm_buysellguessapp
    app.kubernetes.io/instance: buysellguessapp_helmtemplate
    app.kubernetes.io/version: "1.16.0"
    app.kubernetes.io/managed-by: Helm
---
# Source: helm_buysellguessapp/templates/service.yaml
apiVersion: v1
kind: Service
metadata:
  name: buysellguess-svc
spec:
  type: ClusterIP
  selector:
    app: buysellguess
  ports:
    - port: 2345
      protocol: TCP
      targetPort: 3000
---
# Source: helm_buysellguessapp/templates/deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: buysellguess-dep
spec:
  replicas: 1
  selector:
    matchLabels:
      app: buysellguess
  template:
    metadata:
      labels:
        app: buysellguess
    spec:
      containers:
        - name: buysellguessapp
          image: ghcr.io/pwdelbloomboard/ps-container
---
# Source: helm_buysellguessapp/templates/ingress.yaml
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: buysellguess-dep-ingress
spec:
  rules:
  - http:
      paths:
      - path: /
        backend:
          serviceName: buysellguess-dep
          servicePort: 2345
---
# Source: helm_buysellguessapp/templates/tests/test-connection.yaml
apiVersion: v1
kind: Pod
metadata:
  name: "buysellguessapp_helmtemplate-helm_buysellguessapp-test-connection"
  labels:
    helm.sh/chart: helm_buysellguessapp-0.1.0
    app.kubernetes.io/name: helm_buysellguessapp
    app.kubernetes.io/instance: buysellguessapp_helmtemplate
    app.kubernetes.io/version: "1.16.0"
    app.kubernetes.io/managed-by: Helm
  annotations:
    "helm.sh/hook": test
spec:
  containers:
    - name: wget
      image: busybox
      command: ['wget']
      args: ['buysellguessapp_helmtemplate-helm_buysellguessapp:80']
  restartPolicy: Never

```
So we now have a very basic helm chart displayed above. Note that this template does not store anywhere unless we tell it to, it is simply a command-line display of what our total template, completely concatenated together, will look like. Now we can, "install" the helm chart.

If you wanted to save the terminal output to a file, for example then you could use the following including an output command:

```
helm template buysellguessapp-helmtemplate helm-buysellguessapp > helmtemplate-output.txt
```

3. helm install has the format:

helm install [NAME] [CHART] [flags]

* with, "name" being the name we want to create
* [CHART] being the name of the chart we set up above (not during the previous step, that was a template.)

```
helm install buysellguessapp-helminstalled helm-buysellguessapp

Error: create: failed to create: Secret "sh.helm.release.v1.buysellguessapp_helminstalled.v1" is invalid: metadata.name: Invalid value: "sh.helm.release.v1.buysellguessapp_helminstalled.v1": a lowercase RFC 1123 subdomain must consist of lower case alphanumeric characters, '-' or '.', and must start and end with an alphanumeric character (e.g. 'example.com', regex used for validation is '[a-z0-9]([-a-z0-9]*[a-z0-9])?(\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*')
```
What happened? We got an error basically where the Secret was unable to create. Why?

"a lowercase RFC 1123 subdomain must consist of lower case alphanumeric characters, '-' or '.', and must start and end with an alphanumeric character"

So, we fix this and try again:

```
helm install buysellguessapp-helminstalled helm-buysellguessapp

Error: rendered manifests contain a resource that already exists. Unable to continue with install: Service "buysellguess-svc" in namespace "default" exists and cannot be imported into the current release: invalid ownership metadata; label validation error: missing key "app.kubernetes.io/managed-by": must be set to "Helm"; annotation validation error: missing key "meta.helm.sh/release-name": must be set to "buysellguessapp-helminstalled"; annotation validation error: missing key "meta.helm.sh/release-namespace": must be set to "default"
```

What does this mean?  There are several errors:

* Since we already set up, "buysellguess-svc" as a service within the namespace, "default," we are getting the message:

> Service "buysellguess-svc" in namespace "default" exists and cannot be imported into the current release

Just for the sake of sanity, we can try shutting down (stopping) the existing k8s servers that we have, to see if we get a different error message and to move towards investigating whether the, "svc" that the error is referring to is not one listed in another container.

When we do this, we get the errror:

> Error: Kubernetes cluster unreachable: Get "https://0.0.0.0:52387/version?timeout=32s": dial tcp 0.0.0.0:52387: connect: connection refused

According to [this Stackoverflow discussion](https://stackoverflow.com/questions/62309223/rendered-manifests-contain-a-resource-that-already-exists-could-not-get-informa) we can diagnose the issue by using, "helm template" and check what is rendered compared to what we have in our cluster. There may be a name defined with a particular value but not have the value defined in our values.yaml.

After reviewing more about helm in general, as well as [helm template debugging](/about-helm/helmtemplatedebugging.md), we can establish the following procedure:

So to diagnose this issue further, we can follow the following procedure:

3.1. Run kubectl get services to ensure that we see, "buysellguess-svc" as a service, which we do:

```
kubectl get services
NAME               TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)    AGE
kubernetes         ClusterIP   10.43.0.1      <none>        443/TCP    4d1h
buysellguess-dep   ClusterIP   10.43.97.120   <none>        2345/TCP   4d1h
buysellguess-svc   ClusterIP   10.43.41.146   <none>        2345/TCP   4d1h
```
3.2. Look at our "helmtemplate-output.txt" file which was generated by using the output ">" function.  Within that file, look for the, "buysellguess-svc" name.

Here we see that within service.yaml, we have:

```
# Source: helm-buysellguessapp/templates/service.yaml
apiVersion: v1
kind: Service
metadata:
  name: buysellguess-svc
spec:
  type: ClusterIP
  selector:
    app: buysellguess
  ports:
    - port: 2345
      protocol: TCP
      targetPort: 3000
```


3.3 Change the "metadata: name: X" to avoid a conflict.

Directly within the service.yaml file, we just change the name to include, "-helm" and save it to see if this clears the error.

```
# Source: helm-buysellguessapp/templates/service.yaml
apiVersion: v1
kind: Service
metadata:
  name: buysellguess-svc-helm
spec:
  type: ClusterIP
  selector:
    app: buysellguess
  ports:
    - port: 2345
      protocol: TCP
      targetPort: 3000
```
Once this has been completed, the error is cleared.

However, we get a similar error again, which we can clear by using the same above procedure:

```
Error: rendered manifests contain a resource that already exists. Unable to continue with install: Deployment "buysellguess-dep" in namespace "default" exists and cannot be imported into the current release: invalid ownership metadata;
...

```
This time, we create the template output file again, do a search and see that the, "buysellguess-dep" name appears to be on the Ingress.

```
# Source: helm-buysellguessapp/templates/ingress.yaml
apiVersion: extensions/v1beta1
kind: Ingress
metadata:
  name: buysellguess-dep-ingress
spec:
  rules:
  - http:
      paths:
      - path: /
        backend:
          serviceName: buysellguess-dep
          servicePort: 2345
```
However this name was also used in, "deployment.yaml."

```
# Source: helm-buysellguessapp/templates/deployment.yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: buysellguess-dep
spec:
  replicas: 1
  selector:
    matchLabels:
      app: buysellguess
  template:
    metadata:
      labels:
        app: buysellguess
    spec:
      containers:
        - name: buysellguessapp
          image: ghcr.io/pwdelbloomboard/ps-container
```
So, first we change the name within ingress.yaml to, "buysellgess-dep-ingress-backend" see what happens, then if it still catches an error, we take out out of deployment.yaml as well.

* After changing the ingress.yaml file, we still get an error, so we change the deployment.yaml file as well to, "buysellguess-dep-helm"
* Of course, even after changing this, we still get naming conflicts, however this time we also get warnings.

```
W1012 17:13:26.308475   89028 warnings.go:70] extensions/v1beta1 Ingress is deprecated in v1.14+, unavailable in v1.22+; use networking.k8s.io/v1 Ingress
Error: rendered manifests contain a resource that already exists. Unable to continue with install: Ingress "buysellguess-dep-ingress" in namespace "default" exists and cannot be imported into the current release: invalid ownership metadata;
```
* So, we basically have to change every possible name conflict file and add, "helm" onto it to ensure there are absolutely no conflicts.

Looking through the file we have the following evidently which need to change to clear the errors:

| template file   | original                                        | changed to                       |
|-----------------|-------------------------------------------------|----------------------------------|
| deployment.yaml | metadata:   name: buysellguess-dep              | buysellguess-dep-helm            |
| ingress.yaml    | metadata:   name: buysellguess-dep-ingress      | buysellguess-dep-ingress-helm    |
| ingress.yaml    | backend:          serviceName: buysellguess-dep | buysellguess-dep-ingress-backend |
| service.yaml    | buysellguess-svc                                | buysellguess-svc-helm            |

Once all of these names have been changed, we get the following:

```
helm install buysellguessapp-helminstalled helm-buysellguessapp

W1012 17:23:07.606557   60123 warnings.go:70] extensions/v1beta1 Ingress is deprecated in v1.14+, unavailable in v1.22+; use networking.k8s.io/v1 Ingress
W1012 17:23:07.733783   60123 warnings.go:70] extensions/v1beta1 Ingress is deprecated in v1.14+, unavailable in v1.22+; use networking.k8s.io/v1 

Ingress
NAME: buysellguessapp-helminstalled
LAST DEPLOYED: Tue Oct 12 17:23:07 2021
NAMESPACE: default
STATUS: deployed
REVISION: 1
NOTES:
1. Get the application URL by running these commands:
  export POD_NAME=$(kubectl get pods --namespace default -l "app.kubernetes.io/name=helm-buysellguessapp,app.kubernetes.io/instance=buysellguessapp-helminstalled" -o jsonpath="{.items[0].metadata.name}")
  export CONTAINER_PORT=$(kubectl get pod --namespace default $POD_NAME -o jsonpath="{.spec.containers[0].ports[0].containerPort}")
  echo "Visit http://127.0.0.1:8080 to use your application"
  kubectl --namespace default port-forward $POD_NAME 8080:$CONTAINER_PORT
```
To understand fully whether we really need to change the, "backend" name on the ingress, we can do another installation after changing the backend name back to, "serviceName: buysellguess-dep" - however we now have to change the helm installation name.

```
helm install buysellguessapp-helminstalled-simple helm-buysellguessapp
```
Unfortunately when we try this, we get another message:

```
Error: rendered manifests contain a resource that already exists. Unable to continue with install: Service "buysellguess-svc-helm" in namespace "default" exists and cannot be imported into the current release: invalid ownership metadata; annotation validation error: key "meta.helm.sh/release-name" must equal "buysellguessapp-helminstalled-simple": current value is "buysellguessapp-helminstalled"
```

4. We now get to the point where we can, "list" our helmchart with:

```
helm list

NAME                         	NAMESPACE	REVISION	UPDATED                             	STATUS  	CHART                     	APP VERSION
buysellguessapp-helminstalled	default  	1       	2021-10-12 17:23:07.473547 -0500 CDT	deployed	helm-buysellguessapp-0.1.0	1.16.0     

```
5. So to delete this chart so that we can re-install it, or re-release it, we do:

```
helm uninstall 

W1012 17:36:10.013328   96944 warnings.go:70] extensions/v1beta1 Ingress is deprecated in v1.14+, unavailable in v1.22+; use networking.k8s.io/v1 Ingress
release "buysellguessapp-helminstalled" uninstalled
```
Then once you do, "helm list" you should get an empty namespace.

We can then re-install helm with the above attempt we had tried in step 3.4 above, to see if it is indeed the meta-namespaces which conflict by running once again:

```
helm install buysellguessapp-helminstalled-simple helm-buysellguessapp

W1012 17:37:03.681615   13676 warnings.go:70] extensions/v1beta1 Ingress is deprecated in v1.14+, unavailable in v1.22+; use networking.k8s.io/v1 Ingress
W1012 17:37:03.761390   13676 warnings.go:70] extensions/v1beta1 Ingress is deprecated in v1.14+, unavailable in v1.22+; use networking.k8s.io/v1 

Ingress
NAME: buysellguessapp-helminstalled-simple
LAST DEPLOYED: Tue Oct 12 17:37:03 2021
NAMESPACE: default
STATUS: deployed
REVISION: 1
NOTES:
1. Get the application URL by running these commands:
  export POD_NAME=$(kubectl get pods --namespace default -l "app.kubernetes.io/name=helm-buysellguessapp,app.kubernetes.io/instance=buysellguessapp-helminstalled-simple" -o jsonpath="{.items[0].metadata.name}")
  export CONTAINER_PORT=$(kubectl get pod --namespace default $POD_NAME -o jsonpath="{.spec.containers[0].ports[0].containerPort}")
  echo "Visit http://127.0.0.1:8080 to use your application"
  kubectl --namespace default port-forward $POD_NAME 8080:$CONTAINER_PORT
```

#### Running a Test

```
helm test buysellguessapp-helminstalled-simple

$ helm test buysellguessapp-helminstalled-simple
NAME: buysellguessapp-helminstalled-simple
LAST DEPLOYED: Tue Oct 12 17:37:03 2021
NAMESPACE: default
STATUS: deployed
REVISION: 1
TEST SUITE:     buysellguessapp-helminstalled-simple-helm-buysellguessapp-test-connection
Last Started:   Tue Oct 12 17:42:37 2021
Last Completed: Tue Oct 12 17:42:40 2021
Phase:          Failed
NOTES:
1. Get the application URL by running these commands:
  export POD_NAME=$(kubectl get pods --namespace default -l "app.kubernetes.io/name=helm-buysellguessapp,app.kubernetes.io/instance=buysellguessapp-helminstalled-simple" -o jsonpath="{.items[0].metadata.name}")
  export CONTAINER_PORT=$(kubectl get pod --namespace default $POD_NAME -o jsonpath="{.spec.containers[0].ports[0].containerPort}")
  echo "Visit http://127.0.0.1:8080 to use your application"
  kubectl --namespace default port-forward $POD_NAME 8080:$CONTAINER_PORT
Error: pod buysellguessapp-helminstalled-simple-helm-buysellguessapp-test-connection failed

```





#### Running the Suggested Commands

```
 export POD_NAME=$(kubectl get pods --namespace default -l "app.kubernetes.io/name=helm-buysellguessapp,app.kubernetes.io/instance=buysellguessapp-helminstalled-simple" -o jsonpath="{.items[0].metadata.name}")

error: error executing jsonpath "{.items[0].metadata.name}": Error executing template: array index out of bounds: index 0, length 0. Printing more information for debugging the template:
	template was:
		{.items[0].metadata.name}
	object given to jsonpath engine was:
		map[string]interface {}{"apiVersion":"v1", "items":[]interface {}{}, "kind":"List", "metadata":map[string]interface {}{"resourceVersion":"", "selfLink":""}}

```




### Additional Errors

The following errors came from 

* The app.kubernetes.io/managed-by must be set to, "Helm" which is why we are getting:

> invalid ownership metadata; label validation error: missing key "app.kubernetes.io/managed-by": must be set to "Helm"

* meta.helm.sh/release-name must be set to the spefified.

> annotation validation error: missing key "meta.helm.sh/release-name": must be set to "buysellguessapp-helminstalled"

* meta.helm.sh/release-namespace must be set to, "default." which is why we are getting:

> annotation validation error: missing key "meta.helm.sh/release-namespace": must be set to "default"

```
helm ls --all-namespaces
NAME       	NAMESPACE  	REVISION	UPDATED                              	STATUS  	CHART             	APP VERSION
traefik    	kube-system	1       	2021-10-08 19:48:38.9196674 +0000 UTC	deployed	traefik-9.18.2    	2.4.8      
traefik-crd	kube-system	1       	2021-10-08 19:48:36.7649404 +0000 UTC	deployed	traefik-crd-9.18.2
```



## Best Practices for Locating Charts

### Option1 - Maintain one Big Shared Chart in a Chart Repo

* Create a seperate repo, "Chartmuseum," and maintain that seperately, have the person in charge of deployment infrastructure run that.

Shared charts can save a lot of hassle if services are similar in nature.

### Option2 - Maintain Several Service-Specific Charts in a Chart Repo

* You can make a change to one service without worrying about breaking something for another service, but they can cause duplicated work.

For example, you could maintain charts for 15 different microservices in a central repo. This makes it easier to update all of them in one place rather than submitting pull requests to 15 different repos.

### Option3- Maintain a Service-Specific Chart in the Same Repo as the Service Itself

* This is good when the microservices have significant differences - the pattern work better when you keep each chart in the same repo as the service code. If you store the Helm chart in the service repo, it's easier to continoulsy deploy the service independently from other projects.

However, this breaks down as you maintain more and more microservices.


## Resources

* [Collection of Helm Tutorials](https://jfrog.com/blog/10-helm-tutorials-to-start-your-kubernetes-journey/)
* [Introductory Tutorial on Helm](https://www.youtube.com/watch?v=fy8SHvNZGeE)
* [Helm in Kubernetes](https://www.youtube.com/watch?v=-ykwb1d0DXU)
* [Using Helm Github](https://github.com/marcel-dempers/docker-development-youtube-series/tree/master/kubernetes/helm)
* [Best Way to Manage Helm](https://insights.project-a.com/whats-the-best-way-to-manage-helm-charts-1cbf2614ec40)
* [helm 3 install for resource that already exists](https://stackoverflow.com/questions/59443834/helm-3-install-for-resouces-that-exist)
* [Rendered Manifests Contain Resource that Already Exists - Stackoverflow](https://stackoverflow.com/questions/62309223/rendered-manifests-contain-a-resource-that-already-exists-could-not-get-informa)
* [The Helm Docs has a section on, "debugging templates."](https://helm.sh/docs/chart_template_guide/debugging/)
