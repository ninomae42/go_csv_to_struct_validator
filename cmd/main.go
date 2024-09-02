package main

import (
	"flag"
	"log"

	"github.com/ninomae42/csv_struct_validator/assert"
)

func main() {
	fileNameCSV := flag.String("csv", "", "CSV file path")
	fileNameStruct := flag.String("struct", "", "Struct file path")
	flag.Parse()

	ok, err := assert.AssertCSVStruct(*fileNameCSV, *fileNameStruct)
	if err != nil {
		log.Fatalln(err)
	}
	if !ok {
		log.Fatalln("assertion failed")
	} else {
		log.Println("assertion passed")
	}
}
