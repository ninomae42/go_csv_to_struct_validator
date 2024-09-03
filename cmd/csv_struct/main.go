package main

import (
	"flag"
	"log"

	"github.com/ninomae42/csv_struct_validator/assert"
)

func main() {
	fileNameCSV := flag.String("csv", "", "CSV file path")
	fileNameStruct := flag.String("struct", "", "Struct file path")
	assertionMode := flag.String("mode", "all", "assertion mode")
	flag.Parse()

	mode := assert.AssertModeAll
	if *assertionMode == "print" {
		mode = assert.AssertModeAssertAndPrint
	}

	ok, err := assert.AssertCSVStruct(*fileNameCSV, *fileNameStruct, mode)
	if err != nil {
		log.Fatalln(err)
	}
	if !ok {
		log.Fatalln("assertion failed")
	} else {
		log.Println("assertion passed")
	}
}
