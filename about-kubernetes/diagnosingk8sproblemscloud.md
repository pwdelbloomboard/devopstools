# Diagnosing K8s Problems in the Cloud

### Introduction

Much of the procedures used in [Diagnosing K8s Problems Locally](/about-kubernetes/diagnosingk8sproblemslocal.md) can be used in the cloud. The main difference is that the setup must be run differently.  Likely anything cloud-related should be set up under a different context than local, so the starting point would be to switch contexts.

### How to Switch Contexts

The command, [kubectl config](/about-kubernetes/kubectlconfig.md) has a lot of options covering context.

Getting the current context:

```
kubectl config current-context
```
If k3d is setup, and you're on local, this should show k3d-local or something similar.

```
kubectl config get-contexts
```

Should get a list of contexts.  You can switch to the proper context with:

```
kubectl config use-context [CONTEXT_NAME]

Switched to context "[CONTEXT_NAME]"

```
Once you have switched, you can check and confirm your context again:

```
kubectl config current-context
```
### Running KGP on Current Context

Once you have set the context, you can attempt to check the pods, however you might get:

```
Error from server (Forbidden): pods is forbidden: User "arn:aws:iam::[NUMBER]:[USER]" cannot list resource "pods" in API group "" in the namespace "default"
```
Basically, this is a permissions error. If you get this kind of error, your user is likely not granted permissions to access the API within this area.

Once you have this type of error resolved, you can do:

```
kubectl get pods --namespace=default | grep [appname]
```
* This attempts to get all pods within a particular namespace, and then grep's (filters) by the appname that you place on the second part of the command.

From here, you might see an error, such as:

```
[APPNAME-RELEASENAME]   0/1     CrashLoopBackOff   212        18h
```
This indicates that the application is crashing and restarting again and again, so there is something happening underneath preventing successful deployment.

### 