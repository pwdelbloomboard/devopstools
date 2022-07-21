package main

import (
	"bytes"
	"fmt"
	"io"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v2"
)

var data = `---
id: document1
---
id: document2
`

type Doc struct {
	Id string
}

func main() {

	// unmarshal only parses first segment of yaml
	decoder := yaml.NewDecoder(bytes.NewBufferString(data))
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
