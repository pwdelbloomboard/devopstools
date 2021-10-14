# Helm Chart Development Guide

[Helm Chart Development Guide Source](https://helm.sh/docs/topics/charts/)

## Helm Charts vs Helm Templates

 > Helm Charts are source trees that contain a self-descriptor file, Chart.yaml, and one or more templates. 
 
 > Templates are Kubernetes manifest files that describe the resources you want to have on the cluster. Helm uses the Go templating engine by default.

## Charts In General

> Helm uses a packaging format called charts. A chart is a collection of files that describe a related set of Kubernetes resources. A single chart might be used to deploy something simple, like a memcached pod, or something complex, like a full web app stack with HTTP servers, databases, caches, and so on.

* Charts are created as files laid out in a directory trree.

> A chart is organized as a collection of files inside of a directory. The directory name is the name of the chart (without versioning information). 

```
wordpress/
  Chart.yaml          # A YAML file containing information about the chart
  LICENSE             # OPTIONAL: A plain text file containing the license for the chart
  README.md           # OPTIONAL: A human-readable README file
  values.yaml         # The default configuration values for this chart
  values.schema.json  # OPTIONAL: A JSON Schema for imposing a structure on the values.yaml file
  charts/             # A directory containing any charts upon which this chart depends.
  crds/               # Custom Resource Definitions
  templates/          # A directory of templates that, when combined with values,
                      # will generate valid Kubernetes manifest files.
  templates/NOTES.txt # OPTIONAL: A plain text file containing short usage notes
```
### The Chart YAML File

* [Chart YAML File](https://helm.sh/docs/topics/charts/#the-chartyaml-file)


#### Chart Types

> application and library. Application is the default type and it is the standard chart which can be operated on fully. The library chart provides utilities or functions for the chart builder. A library chart differs from an application chart because it is not installable and usually doesn't contain any resource objects.

#### Managing Dependencies Manually

[Managing Dependencies Manually via the Chat Directory](https://helm.sh/docs/topics/charts/#managing-dependencies-manually-via-the-charts-directory)

#### Template Files

[Template Files](https://helm.sh/docs/topics/charts/#template-files)

> Template files follow the standard conventions for writing Go templates (see the text/template Go package documentation for details). An example template file might look something like this:

```
apiVersion: v1
kind: ReplicationController
metadata:
  name: deis-database
  namespace: deis
  labels:
    app.kubernetes.io/managed-by: deis
spec:
  replicas: 1
  selector:
    app.kubernetes.io/name: deis-database
  template:
    metadata:
      labels:
        app.kubernetes.io/name: deis-database
    spec:
      serviceAccount: deis-database
      containers:
        - name: deis-database
          image: {{ .Values.imageRegistry }}/postgres:{{ .Values.dockerTag }}
          imagePullPolicy: {{ .Values.pullPolicy }}
          ports:
            - containerPort: 5432
          env:
            - name: DATABASE_STORAGE
              value: {{ default "minio" .Values.storage }}
```

Note above, the, "image: {{ .Values.imageRegistry }}/postgres:{{ .Values.dockerTag }}" is following the [Go Templating language](/about-go/gotemplates.md), which means that, "{{ .Values.imageRegistry }}" is a .FieldName action, which replaces the FieldName value of a given data structure (struct) on parse time.

So basically, the field name of that "image Registry" will be replaced on parse time, basically filling in items in a sort of for loop for each imageRegistry.

Other fields in the above example include:

> imageRegistry: The source registry for the Docker image.
> dockerTag: The tag for the docker image.
> pullPolicy: The Kubernetes pull policy.
> storage: The storage backend, whose default is set to "minio"

These values are defined by the template author, Helm does not dictate any parameters.


## Tips and Tricks

* [Charts Tips and Tricks](https://helm.sh/docs/howto/charts_tips_and_tricks/)

* Template Functions

> Helm uses [Go templates](/about-go/gotemplates.md) for templating your resource files.

[Go Templates](https://pkg.go.dev/text/template)

* there are also two special template files: include and required
* include function allows you to bring in another template, and then pass the results to other template functions.

* Quote Strings, Dont Quote Integers

> When you are working with string data, you are always safer quoting the strings than leaving them as bare words:

```
name: {{ .Values.MyName | quote }}
```



# Resources

* [Charts Tips and Tricks](https://helm.sh/docs/howto/charts_tips_and_tricks/)