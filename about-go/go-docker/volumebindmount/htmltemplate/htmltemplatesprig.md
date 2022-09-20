# Common Functions from html/template

### Must

```
func Must(t *Template, err error) *Template
```

* Helper that wraps a call to a function returning (*Template, error) and panics if the error is non-nil. It is intended for use in variable initializations such as:

```
var t = template.Must(template.New("name").Parse("html"))
```

### https://pkg.go.dev/html/template#Template.New

```
func (t *Template) New(name string) *Template
```

* New allocates a new HTML template associated with the given one and with the same delimiters. The association, which is transitive, allows one template to invoke another with a {{template}} action.
* If a template with the given name already exists, the new HTML template will replace it. The existing template will be reset and disassociated with t.

```
template.New("test")
```

* Inputs *Template
* Outputs *Template

### Template

> Template is a specialized Template from "text/template" that produces a safe HTML document fragment.

```
type Template struct {

	// The underlying template's parse tree, updated to be HTML-safe.
	Tree *parse.Tree
	// contains filtered or unexported fields
}
```
### template.Whatever()

> This lowercase version of template simply invokes the template library itself.


### template.ParseFiles("whatever.html")

* Outputs a *template.Template, basically a Template, which is the output result of a *.html templating file which uses special curley brackets {} to signify different templating functions, with the appropriate data filled in from a struct.
* Inputs an .html file.

### Parse

```
func Parse(name, text, leftDelim, rightDelim string, funcs ...map[string]any) (map[string]*Tree, error)
```
* Parse returns a map from template name to parse.Tree, created by parsing the templates described in the argument string. The top-level template will be given the specified name. If an error is encountered, parsing stops and an empty map is returned with the error.

* Inputs name, text, leftDelim, rightDelim string, funcs ...map[string]any
* Outputs map[string]*Tree

### Tree

> A tree is a representation of a single parsed template.

```
type Tree struct {
	Name      string    // name of the template represented by the tree.
	ParseName string    // name of the top-level template during parsing, for error messages.
	Root      *ListNode // top-level root of the tree.
	Mode      Mode      // parsing mode.
	// contains filtered or unexported fields
}
```

#### function *Template.Execute

* Inputs a *template.Template
* Executes on an io.Writer, with data
* Outputs an error and does a thing

```
func (t *Template) Execute(wr io.Writer, data any) error
```
> Execute applies a parsed template to the specified data object, writing the output to wr. If an error occurs executing the template or writing its output, execution stops, but partial results may already have been written to the output writer. A template may be executed safely in parallel, although if parallel executions share a Writer the output may be interleaved.

### function *Template.ExecuteTemplate()

```
func (t *Template) ExecuteTemplate(wr io.Writer, name string, data any) error
```
* ExecuteTemplate applies the template associated with t that has the given name to the specified data object and writes the output to wr. If an error occurs executing the template or writing its output, execution stops, but partial results may already have been written to the output writer. A template may be executed safely in parallel, although if parallel executions share a Writer the output may be interleaved.



### *Template.Funcs(funcMap)

* Funcs adds the elements of the argument map to the template's function map. It must be called before the template is parsed. It panics if a value in the map is not a function with appropriate return type. However, it is legal to overwrite elements of the map. The return value is the template, so calls can be chained.

```
func (t *Template) Funcs(funcMap FuncMap) *Template
```

### sprig ... func TxtFuncMap

* TxtFuncMap returns a 'text/template'.FuncMap

```
func TxtFuncMap() ttemplate.FuncMap
```

* The input can be empty, it just initializes an empty text/template FuncMap.

### text/template FuncMap

```
type FuncMap map[string]any
```

* FuncMap is the type of the map defining the mapping from names to functions. Each function must have either a single return value, or two return values of which the second has type error. In that case, if the second (error) return value evaluates to non-nil during execution, execution terminates and Execute returns that error.
* Errors returned by Execute wrap the underlying error; call errors.As to uncover them.
* When template execution invokes a function with an argument list, that list must be assignable to the function's parameter types. Functions meant to apply to arguments of arbitrary type can use parameters of type interface{} or of type reflect.Value. Similarly, functions meant to return a result of arbitrary type can return interface{} or reflect.Value.

### sprig ... func HtmlFuncMap

* Analogous to sprig TxtFuncMap
* HtmlFuncMap returns an 'html/template'.Funcmap