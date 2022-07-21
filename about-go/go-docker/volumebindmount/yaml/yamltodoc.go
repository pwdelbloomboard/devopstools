package main

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

var data = `---
id: document1
---
id: document2
---
id: document3
`

// a type Doc is a bson formatted struct https://pkg.go.dev/go.mongodb.org/mongo-driver/x/bsonx#Doc
type Doc struct {
	test string
}

func main() {

	yamlByte, err := ioutil.ReadFile("test.yaml")
	if err != nil {
		logrus.Fatal(err)
	}

	// convert byte to string
	yamlStr := string(yamlByte)
	// display yaml
	logrus.Info(yamlStr)

	//

	// unmarshal only parses first segment of yaml
	decoder := yaml.NewDecoder(bytes.NewBufferString(yamlStr))
	for {
		var d Doc
		if err := decoder.Decode(&d); err != nil {
			if err == io.EOF {
				break
			}
			panic(fmt.Errorf("Document decode failed: %w", err))
		}
		fmt.Printf("%+v\n", d)
	}
	logrus.Info("All documents decoded")

}
