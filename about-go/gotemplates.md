# Go Templates

* [Go Templates](https://pkg.go.dev/text/template)

* [Helm](/about-helm/helm.md) uses Go templates for templating resource files.

> Go templates are a powerful method to customize output however you want, whether you’re creating a web page, sending an e-mail, working with Buffalo, Go-Hugo, or just using some CLI such as kubectl.

## Template Syntax

* Actions represents the data evaluations, functions or control loops. They’re delimited by {{ }}. Other, non delimited parts are left untouched.

```
{{ }}
```
* Data Evaluations

> Usually, when using templates, you’ll bind them to some data structure (e.g. struct) from which you’ll obtain data. To obtain data from a struct, you can use the {{ .FieldName }} action, which will replace it with FieldName value of given struct, on parse time.

```
{{ .FieldName }}
```
* Conditions

> You can also use if loops in templates. For example, you can check if FieldName non-empty, and if it is, print its value: 

```
{{if .FieldName}} Value of FieldName is {{ .FieldName }} {{end}}.
```

* Loops

* Functions, Pipelines, Variables

* Parsing Templates




# Resources

* [Go Templates](https://pkg.go.dev/text/template)
* [Go Templates - Gopher Academy](https://blog.gopheracademy.com/advent-2017/using-go-templates/)