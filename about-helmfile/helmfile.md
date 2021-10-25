# About HelmFile

# Helmfile

* [helmfile](https://github.com/roboll/helmfile)

> Helmfile is a declarative spec for deploying helm charts. It lets you...
> Keep a directory of chart value files and maintain changes in version control.
> Apply CI/CD to configuration changes.
> Periodically sync to avoid skew in environments.

### Documentation Note

> CAUTION: This documentation is for the development version of Helmfile. If you are looking for the documentation for any of releases, please switch to the corresponding release tag like v0.92.1.

### How Helmfile is Used

* Helmfile can be used in production and within testing to verify prior to moving something into production.

The output of a test script will, "write" gotmpl files which can be reviewed to ensure correct ouput.

![](/img/helmfile_test_output.png)


### What Helmfile References

* Helmfile references an overall apps.yaml file which lists releases for each application within an overall kubernetes cluster.

* Each release will have a named set which includes an individual application name, along with various added variables such as the, "branchSlug" which helps keep a particular release seperate from any past or future release.

* Templates are also referenced - note in this below image the template logic specifically states, "if environment is prod then FALSE, else TRUE." (kind of the reverse of what you might expect)

![](/img/applicationwithinhelmfilesappsyaml.png)

* Labels are also applied as release names. [Labels within the context of kubernetes](https://kubernetes.io/docs/concepts/overview/working-with-objects/labels/) are more analogous to, "tags" in that they allow users to define a structure and can be applied ad-hoc. Labels contrast to namespaces, which are designed more to partition spaces of an cluster, set resource quotas and globally apply actions to a set of resources.

### Helmfile in Use

#### --environment

# The list of environments managed by helmfile.
#
# The default is `environments: {"default": {}}` which implies:
#
# - `{{ .Environment.Name }}` evaluates to "default"
# - `{{ .Values }}` being empty
environments:
  # The "default" environment is available and used when `helmfile` is run without `--environment NAME`.
  default:
    # Everything from the values.yaml is available via `{{ .Values.KEY }}`.
    # Suppose `{"foo": {"bar": 1}}` contained in the values.yaml below,
    # `{{ .Values.foo.bar }}` is evaluated to `1`.
    values:
    - environments/default/values.yaml
    # Each entry in values can be either a file path or inline values.
    # The below is an example of inline values, which is merged to the `.Values`
    - myChartVer: 1.0.0-dev
  # Any environment other than `default` is used only when `helmfile` is run with `--environment NAME`.
  # That is, the "production" env below is used when and only when it is run like `helmfile --environment production sync`.





# Resources

* [helmfile](https://github.com/roboll/helmfile)