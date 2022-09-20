// https://pkg.go.dev/github.com/Masterminds/sprig#section-readme
// http://masterminds.github.io/sprig/

package main

import (
	"fmt"
	"html/template"
	"os"
	"reflect"

	sprig "github.com/Masterminds/sprig/v3"
	"github.com/sirupsen/logrus"
)

func main() {
	// Set up variables and template.
	dataToInsertIntoTemplate := `All work and no play make jack a dull boy. `

	// get the template string
	templateByte, err := os.ReadFile("./ourtemplatefile.yaml")
	if err != nil {
		fmt.Printf("Error during template execution: %s", err)
		return
	}

	templateStr := string(templateByte)

	// Get the Sprig html function map.
	t, err := template.New("nameoftemplate").Funcs(sprig.HtmlFuncMap()).Parse(templateStr)
	if err != nil {
		fmt.Printf("Error during template execution: %s", err)
		return
	}

	logrus.Info("t type: ", reflect.TypeOf(t))

	err = t.ExecuteTemplate(
		os.Stdout,
		"nameoftemplate",
		dataToInsertIntoTemplate,
	)
	if err != nil {
		fmt.Printf("Error during template execution: %s", err)
		return
	}

}
