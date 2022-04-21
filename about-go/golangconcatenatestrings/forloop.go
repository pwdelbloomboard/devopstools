package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func main() {

	records := readCsvFile("valuesets.csv")
	fmt.Println(records)

	// declare metricsarray string of size 2
	var metricsarray [2]string
	// set values for each part of metricsarray string
	metricsarray[0] = `host.disk.freeBytes`
	metricsarray[1] = `host.disk.inodesUsed`
	// loop through and print off each value of metricsarray string
	for i := 0; i <= 1; i++ {
		fmt.Println(metricsarray[i])
	}

}
