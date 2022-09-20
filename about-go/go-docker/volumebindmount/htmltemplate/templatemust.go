package main

import (
	"html/template"
	"os"
	"reflect"

	"github.com/sirupsen/logrus"
)

type Student struct {

	// declaring fields which are
	// exported and accessible
	// outside of package as they
	// begin with a capital letter
	Name  string
	Id    int
	Marks int
}

func main() {

	student_01 := Student{"Patrick", 12, 99}

	tmpl := template.Must(template.ParseFiles("layout.html"))
	logrus.Info("tmpl type: ", reflect.TypeOf(tmpl))
	logrus.Info("*tmpl type: ", reflect.TypeOf(*tmpl))

	t, err := template.ParseFiles("layout.html")
	errcheck(err)
	logrus.Info("t type: ", reflect.TypeOf(t))
	logrus.Info("*t type: ", reflect.TypeOf(*t))

	err = tmpl.Execute(os.Stdout, student_01)
	errcheck(err)

}

func errcheck(err error) {
	if err != nil {
		logrus.Info(err)
	}
}
