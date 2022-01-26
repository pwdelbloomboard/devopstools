# helm install

Helm install is used within the context of a cluster to install a chart (package), either from a repo or a set of files.

In some cases there are dependency issues, or crash loops or other reasons why something might not install.  This guide goes through the procedure for troubleshooting an install.

## helm install -h

The following is the helm install -h file.

This command installs a chart archive.

The install argument must be a chart reference, a path to a packaged chart,
a path to an unpacked chart directory or a URL.

To override values in a chart, use either the '--values' flag and pass in a file
or use the '--set' flag and pass configuration from the command line, to force
a string value use '--set-string'. In case a value is large and therefore
you want not to use neither '--values' nor '--set', use '--set-file' to read the
single large value from file.

    $ helm install -f myvalues.yaml myredis ./redis

or

    $ helm install --set name=prod myredis ./redis

or

    $ helm install --set-string long_int=1234567890 myredis ./redis

or

    $ helm install --set-file my_script=dothings.sh myredis ./redis

You can specify the '--values'/'-f' flag multiple times. The priority will be given to the
last (right-most) file specified. For example, if both myvalues.yaml and override.yaml
contained a key called 'Test', the value set in override.yaml would take precedence:

    $ helm install -f myvalues.yaml -f override.yaml  myredis ./redis

You can specify the '--set' flag multiple times. The priority will be given to the
last (right-most) set specified. For example, if both 'bar' and 'newbar' values are
set for a key called 'foo', the 'newbar' value would take precedence:

    $ helm install --set foo=bar --set foo=newbar  myredis ./redis


To check the generated manifests of a release without installing the chart,
the '--debug' and '--dry-run' flags can be combined.

If --verify is set, the chart MUST have a provenance file, and the provenance
file MUST pass all verification steps.

There are five different ways you can express the chart you want to install:

1. By chart reference: helm install mymaria example/mariadb
2. By path to a packaged chart: helm install mynginx ./nginx-1.2.3.tgz
3. By path to an unpacked chart directory: helm install mynginx ./nginx
4. By absolute URL: helm install mynginx https://example.com/charts/nginx-1.2.3.tgz
5. By chart reference and repo url: helm install --repo https://example.com/charts/ mynginx nginx

CHART REFERENCES

A chart reference is a convenient way of referencing a chart in a chart repository.

When you use a chart reference with a repo prefix ('example/mariadb'), Helm will look in the local
configuration for a chart repository named 'example', and will then look for a
chart in that repository whose name is 'mariadb'. It will install the latest stable version of that chart
until you specify '--devel' flag to also include development version (alpha, beta, and release candidate releases), or
supply a version number with the '--version' flag.

To see the list of chart repositories, use 'helm repo list'. To search for
charts in a repository, use 'helm search'.

Usage:
  helm install [NAME] [CHART] [flags]

Flags:
      --atomic                       if set, the installation process deletes the installation on failure. The --wait flag will be set automatically if --atomic is used
      --ca-file string               verify certificates of HTTPS-enabled servers using this CA bundle
      --cert-file string             identify HTTPS client using this SSL certificate file
      --create-namespace             create the release namespace if not present
      --dependency-update            run helm dependency update before installing the chart
      --description string           add a custom description
      --devel                        use development versions, too. Equivalent to version '>0.0.0-0'. If --version is set, this is ignored
      --disable-openapi-validation   if set, the installation process will not validate rendered templates against the Kubernetes OpenAPI Schema
      --dry-run                      simulate an install
  -g, --generate-name                generate the name (and omit the NAME parameter)
  -h, --help                         help for install
      --insecure-skip-tls-verify     skip tls certificate checks for the chart download
      --key-file string              identify HTTPS client using this SSL key file
      --keyring string               location of public keys used for verification (default "/Users/patrick.delaneybloomboard.com/.gnupg/pubring.gpg")
      --name-template string         specify template used to name the release
      --no-hooks                     prevent hooks from running during install
  -o, --output format                prints the output in the specified format. Allowed values: table, json, yaml (default table)
      --password string              chart repository password where to locate the requested chart
      --post-renderer postrenderer   the path to an executable to be used for post rendering. If it exists in $PATH, the binary will be used, otherwise it will try to look for the executable at the given path (default exec)
      --render-subchart-notes        if set, render subchart notes along with the parent
      --replace                      re-use the given name, only if that name is a deleted release which remains in the history. This is unsafe in production
      --repo string                  chart repository url where to locate the requested chart
      --set stringArray              set values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2)
      --set-file stringArray         set values from respective files specified via the command line (can specify multiple or separate values with commas: key1=path1,key2=path2)
      --set-string stringArray       set STRING values on the command line (can specify multiple or separate values with commas: key1=val1,key2=val2)
      --skip-crds                    if set, no CRDs will be installed. By default, CRDs are installed if not already present
      --timeout duration             time to wait for any individual Kubernetes operation (like Jobs for hooks) (default 5m0s)
      --username string              chart repository username where to locate the requested chart
  -f, --values strings               specify values in a YAML file or a URL (can specify multiple)
      --verify                       verify the package before using it
      --version string               specify the exact chart version to use. If this is not specified, the latest version is used
      --wait                         if set, will wait until all Pods, PVCs, Services, and minimum number of Pods of a Deployment, StatefulSet, or ReplicaSet are in a ready state before marking the release as successful. It will wait for as long as --timeout
      --wait-for-jobs                if set and --wait enabled, will wait until all Jobs have been completed before marking the release as successful. It will wait for as long as --timeout

Global Flags:
      --debug                       enable verbose output
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

## --debug flag

* Use the debug flag to see what is going wrong with an installation while it attempts to install.
* Push the output to an output file to be able to read later if needed.  SomeCommand &> SomeFile.txt  

### Viewing the Debug Output

When you look at the debug output you might see something like this:

```
client.go:558: [debug] newrelic-nri-metadata-injection-admission-create: Jobs active: 1, jobs failed: 0, jobs succeeded: 0
```

Followed by a timeout error.

```
Error: failed pre-install: timed out waiting for the condition
helm.go:81: [debug] failed pre-install: timed out waiting for the condition
```

This essentially means that the last Job is the one that failed.

### Using a --dry-run to Create Output File

Adding the flag:

```
--dry-run
```

Will allow an output file to be created.

When you look at the dry run, you can look at all of the sub-dependency charts and try to see if you can find the name linked to the specific sub-dependency chart, if there is one.

Using that sub-dependency, you can zoom in on that one dependency to try to diagnose problems.


#### Zooming In on Sub-Dependency

To zoom in on the sub-deopendency itself, just the chart specified can be used rather than the entire bundle.


### Adding Parameters to Skip Dependencies

TBD


