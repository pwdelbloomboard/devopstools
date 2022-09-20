package main

import (
	"bytes"
	"fmt"
	"html/template"
	"io/fs"
	"os"

	"github.com/Masterminds/sprig/v3"
	"github.com/sirupsen/logrus"
)

func main() {

	templateFile := fmt.Sprintf("%s./config-map.gotmpl.yaml")
	templateContents, err := os.ReadFile(templateFile)
	errorcheck(err)

	templateInstance, err := template.New("config").Funcs(sprig.FuncMap()).Parse(string(templateContents))
	errorcheck(err)

	var processedTemplate bytes.Buffer

	err = templateInstance.ExecuteTemplate(
		&processedTemplate,
		"config",
		trimmedOutput,
	)
	errorcheck(err)

	outfile := fmt.Sprintf("%sconfig-map.yaml", viper.GetString("yosemite-path"))
	err = os.WriteFile(outfile, processedTemplate.Bytes(), fs.ModeAppend)
	errorcheck(err)

}

func errorcheck(err error) {
	if err != nil {
		logrus.Info(err)
	}
}
