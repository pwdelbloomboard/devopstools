# About Kubernetes Ops

https://github.com/kubernetes/kops

## Kubernetes Operations - Kops

* This is kubectl for clusters.
* Helps you create, destroy, upgrade and maintain production grade, highly available K8s Cluster.
* Also provision necessary cloud infrastructure.

## Documentation

* [Documentation is Kept Here](https://kops.sigs.k8s.io/)

### Kops CLI

* [Kops CLI](https://kops.sigs.k8s.io/cli/kops/)

#### kops completion

* Generate auto-completion script for a specified shell.
* Basically, an auto-completeion script for a particular shell environment.

For example, if you run:

```
kops completion bash
```

* This will print out a script which can be used to run a bunch of kops functionality as a bash script.

#### kops create

* Create resource by command line, filename, or stdin.

##### Example

```
  # Create a cluster from the configuration specification in a YAML file.
  kops create -f my-cluster.yaml

  # Create secret from secret spec file.
  kops create -f secret.yaml

  # Create an instancegroup based on the YAML passed into stdin.
  cat instancegroup.yaml | kops create -f -
```

#### kops delete

#### kops distrust

#### kops edit

#### kops export

* Export a configuration

For example:

```
kops export kubecfg thing.domain.com --state s3://secret-state-store-bucket --admin
```
* The above exports a kubeconfig file with the cluster admin user (make sure you keep this user safe!)

* A kubeconfig file is a file used to configure access to Kubernetes when used in conjunction with the kubectl commandline tool (or other clients).  So this command would essentially give access to a cluster.

#### kops get

Get a cluster and its instance groups

```
kops get k8s-cluster.example.com
```
Example:

```
  assets         Display assets for cluster.
  clusters       Get one or many clusters.
  instancegroups Get one or many instance groups.
  instances      Display cluster instances.
  keypairs       Get one or many keypairs.
  secrets        Get one or many secrets.
  sshpublickeys  Get one or many secrets.
```

#### kops promote

#### kops replace

#### kops rolling-update

#### kops toolbox

#### kops trust

#### kops update

#### kops upgrade

#### kops validate