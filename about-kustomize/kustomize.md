## About Kustomize

* With Kubernetes, you have tons of yaml files and are applying them.
* However, if you're using nothing but yaml files, you are copying and pasting a lot of the same thing.
* At the same time, different environments need different levels of customizations.
* Similarly, you might have a different number of replicas in staging vs. production and so on and sofourth, the list goes on and on!
* You will almost certainly have variations across different environments, and you have to apply different changes against whatever third-party settings, while at the same time copy and pasting what are essentially, templates.

Helm solved this problem using, "templates."  However, this may not always be the best choice.

Essentially, it allows us to get the same results as what helm provides, but without templates, purely with yaml files.

It allows you to modify yaml files to accomplish all of the above, but in an easier way.

## Prerequisites for Kustomize

* Need a Kubernetes Cluster.
* Need an Nginx Ingress.
* Need the Kustomize CLI.

Basically, we used the repo we put together getting k3d up and running.

Installing Kustomize-CLI on this debian image involves the following:

```
apt update
apt install snapd
snap install core
```
Then we were able to install kustomize with:

```
snap install kustomize
```

When we attempt to do this, we get an, "error: cannot communicate with server:" 

So the first thing we can check is whether snap is working at all.


```
systemctl status snapd.service
System has not been booted with systemd as init system (PID 1). Can't operate.
Failed to connect to bus: Host is down
```

Long story short, we can't seem to use snapd on our container, because it appears that snap must be activated somehow, which cannot take place because we can't have access to systemctl (systemd) in a container.

We can't have access to systemd in a container because it mounts filesystems, controls kernel parameters, and has its own internal system for process output, sap space, POSIX message queues, inter-process message bus, terminal logins, etc. You would have to use this as, --privledged, which is typically a bad idea because it breaks Docker isolation.

### Installing Kustomize with Brew

We can install Brew in a linux container with:

```
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

This takes a while.  Once brew is installed, you can add it to the bash profile with:

```
    echo 'eval "$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)"' >> /root/.profile
    eval "$(/home/linuxbrew/.linuxbrew/bin/brew shellenv)"
```
Finally, we can use brew to install kustomize:

```
brew install kustomize
```
However, this might not work on our Docker container, since it's ARM rather than Intel, and the [installation page on Brew](https://formulae.brew.sh/formula/kustomize) does not show ARM compatibility for 64-bit linux.

But that being said, it installed anyway, and then we get:

```
Bash completion has been installed to:
  /home/linuxbrew/.linuxbrew/etc/bash_completion.d
```

### Versioning

* Brew:

```
brew --version
Homebrew 3.4.10
Homebrew/homebrew-core (git revision 874c68efb70; last commit 2022-05-04)
```
* Kustomize

```
kustomize version
{Version:kustomize/v4.5.4 GitCommit:cf3a452ddd6f83945d39d582243b8592ec627ae3 BuildDate:2022-03-28T23:06:20Z GoOs:linux GoArch:amd64}
```
Interestingly, Victor Farcic replied right away when asking about using Kustomize full vs. native Kustomize on kubectl.

> At the time I recorded that video, kustomize that was baked into kubectl was very old and there were some issues that prevented the team from upgrading it. That is now solved and now I use only the one in kubectl.

#### Alternate Installation Method

It is also possible to install Kustomize with:

```
curl -s "https://raw.githubusercontent.com/\
kubernetes-sigs/kustomize/master/hack/install_kustomize.sh"  | bash
```

## Using Kustomize

We started out by cloning:

* [ArgoCD Git Repo](git clone https://github.com/vfarcic/argocd-production.git

Then we move into the project folder and run:

```
ls -1
README.md
apps-kustomize.yaml
apps-manual.yaml
apps.yaml
argo-cd
argo-events
argo-workflows
argocd
codefresh
helm
kustomize
orig
project.yaml
sealed-secrets
```

Most of these have to do with argo, which we're not working with at the moment, we're just going to use argo-workflows.

Looking into the, "argo-workflows" - we see base and overlays, move into the base folder, then take a look at the kustomization.yaml.

```
ls -1 argo-workflows
base
overlays
ls -1 base/
cat kustomization.yaml
```

For now imagine that overlays doesn't exist, just remember that the directory structure can be whatever we want...the base contains the, "main stuff," whereas overlays is extra.

What's inside of base?

The majority is fairly self-explanatory, but we should look at kustomization.yaml and config.yaml.

From a kubernetes perspective, this is yet another, "kind" - another kubernetes resrouces, it specifies definitions that we want to apply to our cluster.

So what do we see here within kustomization.yaml?

* Basically, this is another kubernetes yaml with kind, "Kustomization".
* This is yet another kubernetes kind, pure yaml.
* Specifies the list of definitions that we want to apply to our cluster if we choose to apply this specific customization. Thesea re namespaces that are in the repositories.
* We don't want to copy and paste everything from, "/argo-workflows/", e.g. the resources: again and again, we just want to point to those repositories' manifests, e.g. manifests/base, manifest/cluster-install, manifest/cluster-install/argo-server-rbac, etc. over into our cluster.
* We also have a, "patchesStrategicMerge" which is kind of like a values/override file, adding a configmap patch rather than using values files as helm does, to override yaml files.

```
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
resources:
- namespace.yaml
- github.com/argoproj/argo-workflows/manifests/base
- github.com/argoproj/argo-workflows/manifests/cluster-install/workflow-controller-rbac
- github.com/argoproj/argo-workflows/manifests/cluster-install/argo-server-rbac
- ingress.yaml
patchesStrategicMerge:
- config.yaml
namespace: argo
```
* Looking a bit more closer at manifests/base/ ... you can look at the kustimization.yaml file. Each directory referenced by kustomize must have a kustomization.yaml file. This kustomization file shows whe resources for that directory and how they are used.  So if we actually go to: [github.com/argoproj/argo-workflows/manifests/base](https://github.com/argoproj/argo-workflows/tree/master/manifests/base) and look we see:
* kustomization.yaml

And this file contains:

```
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization

resources:
- crds
- workflow-controller
- argo-server
```

Which references other directories with kustomization.yaml's so you can further define and branch in as many directories as you want.  Looking into workflow-controller, the directory, there is a kustomization.yaml file, as well as other yaml files, such as workflow-controller-configmap.yaml:

```
apiVersion: v1
kind: ConfigMap
metadata:
  name: workflow-controller-configmap
```


And what about the config.yaml file in the original ?

```
cat config.yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: workflow-controller-configmap
data:
  config: |
    containerRuntimeExecutor: k8sapi
```

* It's not defining the whole configmap, it's going and saying, use "workflow-controller-configmap" - and use this, and find, "data : config : | " and then patch it with, "containerRuntimeExecutor: k8sapi."
* So basically, kustomize is a patching mechanism, leveraging configmaps to patch configurations and over-write through the configmap, whereas helm is more about templating and replacing values.

Now we attempt to apply kustomize by using, "apply -k base"

```
kubectl apply -k base
namespace/argo created
customresourcedefinition.apiextensions.k8s.io/clusterworkflowtemplates.argoproj.io created
customresourcedefinition.apiextensions.k8s.io/cronworkflows.argoproj.io created
customresourcedefinition.apiextensions.k8s.io/workfloweventbindings.argoproj.io created
customresourcedefinition.apiextensions.k8s.io/workflows.argoproj.io created
customresourcedefinition.apiextensions.k8s.io/workflowtaskresults.argoproj.io created
customresourcedefinition.apiextensions.k8s.io/workflowtasksets.argoproj.io created
customresourcedefinition.apiextensions.k8s.io/workflowtemplates.argoproj.io created
serviceaccount/argo created
serviceaccount/argo-server created
role.rbac.authorization.k8s.io/argo-role created
clusterrole.rbac.authorization.k8s.io/argo-aggregate-to-admin created
clusterrole.rbac.authorization.k8s.io/argo-aggregate-to-edit created
clusterrole.rbac.authorization.k8s.io/argo-aggregate-to-view created
clusterrole.rbac.authorization.k8s.io/argo-cluster-role created
clusterrole.rbac.authorization.k8s.io/argo-server-cluster-role created
rolebinding.rbac.authorization.k8s.io/argo-binding created
clusterrolebinding.rbac.authorization.k8s.io/argo-binding created
clusterrolebinding.rbac.authorization.k8s.io/argo-server-binding created
configmap/workflow-controller-configmap created
service/argo-server created
service/workflow-controller-metrics created
priorityclass.scheduling.k8s.io/workflow-controller created
deployment.apps/argo-server created
deployment.apps/workflow-controller created
ingress.networking.k8s.io/argo-server created
```
This doesn't really tell us too much, as much as how it's applying the agro workflows, which is not super important at the moment. 

If we wanted to output the, "combined patched version of all kustomized yaml files," we could just simply use, "kustomize build."

So going back into argo-workflows, we can apply the base file with:

```
kustomize build base
```
Which will output a super long yaml file, which represents our finished, patched yaml file combining all of the yaml files we referred to from base, and all of the tree yaml files connected to that (without applying it).

The primitive command to use, if you didn't use, "kubectl apply -k ." would be:

```
kustomize build base | kubectl appl --filename -
```
Now we can look at what has been applied by looking at the appropriate, new namespace whih was applied:

```
kubectl get namespaces
NAME              STATUS   AGE
kube-system       Active   43h
default           Active   43h
kube-public       Active   43h
kube-node-lease   Active   43h
argo              Active   18h
workflows         Active   3m8s
...

kubectl --namespace argo get ingresses

...

NAME          CLASS    HOSTS      ADDRESS   PORTS   AGE
argo-server   <none>   acme.com             80      18h
```
* So what if we wanted to keep everything the same but just change one elemnt, such as change the host? This might be similar to using, .gotmpl files in helm.
* We can use overlays, as were found at "/home/argocd-production/argo-workflows/overlays" -- specifically, we can look at the production/ folder.

```
ingress_patch.json  kustomization.yaml
```

* What this kustomization.yaml file is saying is, "patch only the ingress_patch.json" on everything under ../../base 
* We could provide more patches, but here we're only applying this one patch.
* target is the target of what we are patching, and the *path* will be applied to that target.

```
apiVersion: kustomize.config.k8s.io/v1beta1
kind: Kustomization
resources:
- ../../base
- ../workflows
patches:
- path: ingress_patch.json
  target:
    group: networking.k8s.io
    version: v1
    kind: Ingress
    name: argo-server
```
Looking at that patch, this is showing that the path, /spec/rules/host will be overwritten by the value, argo-workflows.~:

```
cat ingress_patch.json
[
  {
    "op": "replace",
    "path": "/spec/rules/0/host",
    "value": "argo-workflows.192.168.64.2.nip.io"
  }
]
```
To apply this patch we can do the following from back in the /argo-workflows/ folder:

```
kustomize build overlays/production | kubectl apply --filename -
```
Which will do all of the same stuff as the kustomize build ~ | kubectl apply --filename - command above, but it is replacing that ingress.

```
kubectl --namespace argo get ingresses
...

NAME          CLASS    HOSTS                            
argo-server   <none>   argo-workflows.192.168.64.2.nip.io

```
* We can view to see what kind of image is being used by running:

```
kubectl --namespace argo get deployment argo-server --output yaml

...

    spec:
      containers:
      - args:
        - server
        image: quay.io/argoproj/argocli:latest

```
* So we're using that image from quay.io/argoproj/argocli:latest -- now, if we wanted to overlay that image, we could do that by applying another patch..

There are shorthand ways to apply these patches, rather than hard-coding the patch, we can move faster by using a shell command, such as (within /overlays/production, since this is a patch):

```
kustomize edit set image argoproj/argocli=argoproj/argocli:v2.12.4
```

* Now if we look at the kustomization yaml, you can see that the patch configuration gets applied to update the version.

* Basically, you can create as many versions of the base yaml, with as many overlays as you want, and it really depends upon how many environments and variations that you might be using.

* It's a very convenient way to customize, "kustomize" variations of kubernetes settings. It doesn't do the same thing, there are differences between kustomize and helm, but the end goal is very similar, to provide means of providing variations of our manifests, upgrade applications. 

### Is it better/worse/same as helm?



## Resources

* [Kustomize - Hwo to Simplify Kubernetes Configuration Management](https://www.youtube.com/watch?v=Twtbg6LFnAg)
* [Commands Used in Above](https://gist.github.com/vfarcic/07b0b4642b5694d0239ee7c1629173ce)
* [Running Systemd on Docker Container](https://stackoverflow.com/questions/51979553/is-it-recommended-to-run-systemd-inside-docker-container)
* [More About Docker and Systemd](https://stackoverflow.com/questions/59466250/docker-system-has-not-been-booted-with-systemd-as-init-system)
* [Helm vs. Kustomize](https://www.youtube.com/watch?v=ZMFYSm0ldQ0)