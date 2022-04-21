package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
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

	// invoke the readCsvFile function
	records := readCsvFile("valuesets.csv")

	// Get number of records
	recordslen := len(records)

	// go through all records
	for i := 0; i <= recordslen-1; i++ {
		// extract second value in array point
		setvalue, err := strconv.Atoi(records[i][1])
		// deal with forced error handling
		if err != nil {
			log.Fatal("Go Error from strconv.Atoi(records[i][1]): ", err)
		}
		// print if the setvalue of a particular record line is 1
		if setvalue == 1 {
			fmt.Println(records[i][0])
		}
	}

}
