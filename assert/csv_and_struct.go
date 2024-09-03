package assert

import (
	"fmt"
	"log"

	"github.com/ninomae42/csv_struct_validator/helper"
)

var (
	ErrorHeaderLengthMismatch = fmt.Errorf("assert: header length mismatch")
	ErrorHeaderStringMismatch = fmt.Errorf("assert: header string mismatch")
)

type AssertMode uint8

const (
	AssertModeAll AssertMode = iota
	AssertModeAssertAndPrint
)

func AssertCSVStruct(fileNameCSV string, fileNameStruct string, am AssertMode) (bool, error) {
	csvInfo, err := helper.GetCsvInfo(fileNameCSV)
	if err != nil {
		return false, err
	}

	structInfo, err := helper.GetStructInfo(fileNameStruct)
	if err != nil {
		return false, err
	}

	switch am {
	case AssertModeAll:
		if len(csvInfo.Header) != len(structInfo.Fields) {
			return false, ErrorHeaderLengthMismatch
		}

		for i, csvHeader := range csvInfo.Header {
			if csvHeader != structInfo.Fields[i].CSVTag {
				return false, ErrorHeaderStringMismatch
			}
		}
	case AssertModeAssertAndPrint:
		var lengthMismatch bool
		var stringMismatch bool
		if len(csvInfo.Header) != len(structInfo.Fields) {
			log.Printf("header length mismatch: %d != %d\n", len(csvInfo.Header), len(structInfo.Fields))
			lengthMismatch = true
		}

		for i, csvHeader := range csvInfo.Header {
			if len(structInfo.Fields) <= i {
				log.Printf("[%v] header missing in struct: %s\n", i, csvHeader)
				continue
			}
			if csvHeader != structInfo.Fields[i].CSVTag {
				log.Printf("[%v] header string mismatch: %s != %s\n", i, csvHeader, structInfo.Fields[i].CSVTag)
				stringMismatch = true
			}
		}
		switch {
		case lengthMismatch:
			return false, ErrorHeaderLengthMismatch
		case stringMismatch:
			return false, ErrorHeaderStringMismatch
		default:
		}
	}

	return true, nil
}
