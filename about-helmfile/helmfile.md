# About HelmFile

# Helmfile

* [helmfile](https://github.com/roboll/helmfile)

> Helmfile is a declarative spec for deploying helm charts. It lets you...
> Keep a directory of chart value files and maintain changes in version control.
> Apply CI/CD to configuration changes.
> Periodically sync to avoid skew in environments.

### Documentation Note

> CAUTION: This documentation is for the development version of Helmfile. If you are looking for the documentation for any of releases, please switch to the corresponding release tag like v0.92.1.


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