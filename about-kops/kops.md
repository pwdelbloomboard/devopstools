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

#### kops get

#### kops promote

#### kops replace

#### kops rolling-update

#### kops toolbox

#### kops trust

#### kops update

#### kops upgrade

#### kops validate