// https://pkg.go.dev/github.com/Masterminds/sprig#section-readme
// http://masterminds.github.io/sprig/

package main

import (
	"fmt"
	"os"
	"reflect"
	"text/template"

	sprig "github.com/Masterminds/sprig/v3"
	"github.com/sirupsen/logrus"
)

func main() {
	// Set up variables and template.
	vars := map[string]interface{}{"Name": "  John Jacob Jingleheimer Schmidt "}
	tpl := `Hello {{.Name | trim | lower | indent 4}}`

	// Get the Sprig function map.
	fmap := sprig.TxtFuncMap()
	t, err := template.New("test").Funcs(fmap).Parse(tpl)
	if err != nil {
		fmt.Printf("Error during template execution: %s", err)
		return
	}

	logrus.Info("vars type: ", reflect.TypeOf(vars))
	logrus.Info("tpl type: ", reflect.TypeOf(tpl))
	logrus.Info("t type: ", reflect.TypeOf(t))

	err = t.Execute(os.Stdout, vars)
	if err != nil {
		fmt.Printf("Error during template execution: %s", err)
		return
	}

}
