Here is the concatstrings.go file:

```
package main

import (
	"fmt"
)

func main() {

	// our accountid
	accountid := `49653`

	// the metric name
	metricname := `host.disk.freeBytes`

	// first part of query string to concatenate
	firstpart := `
    mutation {
        nrqlDropRulesCreate(
            accountId: `

	// second part of query string to concatenate
	secondpart :=
		`
		rules: [
                    {
                        action: DROP_ATTRIBUTES
                        nrql: "FROM Metric SELECT `

	// third part of query string to concatenate
	thirdpart := `"
                        description: "Removes `

	// fourth part of query string to concatenate
	fourthpart := ` from Metric"
                    }
                ]
        ) {
            successes {
                id
                action
                nrql
                description
        }
            failures {
                submitted {
                    nrql
                }
            error {
                reason
                description
            }
        }
    }
    }`

	// print out
	/*
		fmt.Printf(firstpart)
		fmt.Printf(accountid)
		fmt.Printf(secondpart)
		fmt.Printf(metricname)
		fmt.Printf(thirdpart)
		fmt.Printf(metricname)
		fmt.Printf(fourthpart)
	*/

	// concatstrings
	finalstring := firstpart + accountid + secondpart + metricname + thirdpart + metricname + fourthpart
	// print out final version
	fmt.Printf(finalstring)
}
```