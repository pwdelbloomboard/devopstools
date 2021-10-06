# What Does kubectl config Do?

> Use kubeconfig files to organize information about clusters, users, namespaces, and authentication mechanisms. The kubectl command-line tool uses kubeconfig files to find the information it needs to choose a cluster and communicate with the API server of a cluster.

Runnning the following command:

```
kubectl config current-context
```

* ...Will show the name of the current context, for example: "dev.hellotest.com"

Whereas running:

```
kubectl config use-context CONTEXT_NAME [options]
```
Will select a context based upon CONTEXT_NAME.

Looking at all of the contexts will pull up the following (example shown):

```
kubectl config get-contexts

CURRENT   NAME                       CLUSTER                    AUTHINFO                   NAMESPACE
*         dev.hellotest.com          dev.hellotest.com          dev.hellotest.com          default
          prod.hellotest.com         prod.hellotest.com         prod.hellotest.com         default
```
Whereas, "viewing" the config will show information about the clusters, contexts, etc. in document format:

```
kubectl config view

apiVersion: v1
clusters:
- cluster:
    certificate-authority-data: DATA+OMITTED
    server: https://api.dev.example.com
  name: dev.hellotest.com
- cluster:
    certificate-authority-data: DATA+OMITTED
    server: https://api.prod.hellotest.com
  name: prod.hellotest.com
contexts:
- context:
    cluster: dev.hellotest.com
    namespace: default
    user: dev.hellotest.com
  name: dev.hellotest.com
- context:
    ~etc...additional information for additional cluster

kind: Config
preferences: {}
users:
- name: dev.hellotest.com
  user:
    exec:
      apiVersion: client.authentication.k8s.io/v1alpha1
      args:
      - token
      - -i
      - dev.hellotest.com
      command: /Users/USERNAME/.asdf/shims/aws-iam-authenticator
      env:
      - name: AWS_PROFILE
        value: default
      provideClusterInfo: false
- name: 
    ~ etc. -- additional user profile.

```

So the above, "document format," creates the context, in YAML.

* Where is this configuration file stored?
* On a local machine, this should be found at: HOME/.kube/config

### Simple Overview of Config File

* A configuration file describes clusters, users, and contexts. Below, the config-demo file has the framework to describe two clusters, two users, and three contexts.
* The config-demo file goes in the following folder structure: Create a directory named config-exercise. In your config-exercise directory, create a file named config-demo with this content:

```
apiVersion: v1
kind: Config
preferences: {}

clusters:
- cluster:
  name: development
- cluster:
  name: scratch

users:
- name: developer
- name: experimenter

contexts:
- context:
  name: dev-frontend
- context:
  name: dev-storage
- context:
  name: exp-scratch
```

Access to clusters can be configured by different methods and in different configurations, for example, you could have a setup where:

* Access to the development cluster requires authentication by certificate. 
* Access to the scratch cluster requires authentication by username and password.

### Default Kube Config Environmental Variable

The KUBECONFIG environment variable holds a list of kubeconfig files.

* If the KUBECONFIG environment variable doesn't exist, kubectl uses the default kubeconfig file, $HOME/.kube/config.
* If the KUBECONFIG environment variable does exist, kubectl uses an effective configuration that is the result of merging the files listed in the KUBECONFIG environment variable.

```
cat $HOME/.kube/config
```

## Configuring Across Multiple Clusters

*[Configuring Across Multiple Clusters](https://kubernetes.io/docs/tasks/access-application-cluster/configure-access-multiple-clusters/) goes into more information about how to do so.


1. Create a directory called [config-exercise](/about-kubernetes/config-exercise/)
2. Create a file in that directory called [config-demo](/about-kubernetes/config-exercise/config-demo) which has the above config-demo content structure.
3. config-demo file has the framework to describe two clusters, two users, and three contexts.
4. With our terminal, we can enter into the "config-exercise" directory and set the following details with the following commands:

Adding Clusters:

```
kubectl config --kubeconfig=config-demo set-cluster development --server=https://1.2.3.4 --certificate-authority=fake-ca-file
kubectl config --kubeconfig=config-demo set-cluster scratch --server=https://5.6.7.8 --insecure-skip-tls-verify
```

Adding Users:

```
kubectl config --kubeconfig=config-demo set-credentials developer --client-certificate=fake-cert-file --client-key=fake-key-seefile
kubectl config --kubeconfig=config-demo set-credentials experimenter --username=exp --password=some-password
```
After running these commands we should get responses, "User "developer" set." and "User "experimenter" set."

Then, entering into the config-demo file, we see the following has been changed:

```
users:
- name: developer
  user:
    client-certificate: fake-cert-file
    client-key: fake-key-seefile
- name: experimenter
  user:
    password: some-password
    username: exp
```

Note:

> To delete a user you can run kubectl --kubeconfig=config-demo config unset users.<name>
> To remove a cluster, you can run kubectl --kubeconfig=config-demo config unset clusters.<name>
> To remove a context, you can run kubectl --kubeconfig=config-demo config unset contexts.<name>

Then, we can run the following commands to modify the contexts:

```
kubectl config --kubeconfig=config-demo set-context dev-frontend --cluster=development --namespace=frontend --user=developer

Context "dev-frontend" modified.

kubectl config --kubeconfig=config-demo set-context dev-storage --cluster=development --namespace=storage --user=developer

Context "dev-storage" modified.

kubectl config --kubeconfig=config-demo set-context exp-scratch --cluster=scratch --namespace=default --user=experimenter

Context "exp-scratch" modified.

```
To view everything, we can do the following:

```
kubectl config --kubeconfig=config-demo view
```

The output of that shows two clusters, two users and three contexts.

> The fake-ca-file, fake-cert-file and fake-key-file above are the placeholders for the pathnames of the certificate files. You need to change these to the actual pathnames of certificate files in your environment.

> Sometimes you may want to use Base64-encoded data embedded here instead of separate certificate files; in that case you need to add the suffix -data to the keys, for example, certificate-authority-data, client-certificate-data, client-key-data.

* Each context is a triple (cluster, user, namespace). For example, the dev-frontend context says, "Use the credentials of the developer user to access the frontend namespace of the development cluster".

5. Set the context with:

```
kubectl config --kubeconfig=config-demo use-context dev-frontend
```

Before doing that, we can view our current context to be able to revert back:

```
kubectl config current-context

example.testexample.com
```
So in theory, we should be able to revert back to, "example.testexample.com"



# Resources

* [kubectl Context and Configuration](https://kubernetes.io/docs/reference/kubectl/cheatsheet/#kubectl-context-and-configuration)
* [configuring accessing multiple clusters](https://kubernetes.io/docs/tasks/access-application-cluster/configure-access-multiple-clusters/)
* [organizing cluster access with kubeconfig](https://kubernetes.io/docs/concepts/configuration/organize-cluster-access-kubeconfig/)