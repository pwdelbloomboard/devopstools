// https://pkg.go.dev/html/template

package main

import (
	"html/template"
	"io/ioutil"
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {

	check := func(err error) {
		if err != nil {
			log.Fatal(err)
		}
	}

	gotmplContentsByteArray, err := ioutil.ReadFile("./helloworld.gotmpl.yaml")
	if err != nil {
		logrus.Fatal(err)
	}

	gotmplContentsStr := string(gotmplContentsByteArray)

	logrus.Info("gotmplContentsStr: ", gotmplContentsStr)

	t, err := template.New("testyaml").Parse(gotmplContentsStr)
	check(err)

	fillWithDataByteArray, err := ioutil.ReadFile("./other.vals.yaml")
	if err != nil {
		logrus.Fatal(err)
	}

	fillWithDataStr := string(fillWithDataByteArray)

	logrus.Info("fillWithDataStr: ", fillWithDataStr)

	data := struct {
		indent string
	}{fillWithDataStr}

	err = t.Execute(os.Stdout, data)
	check(err)

}
