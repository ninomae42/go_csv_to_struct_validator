package assert

import (
	"fmt"

	"github.com/ninomae42/csv_struct_validator/helper"
)

var (
	ErrorHeaderLengthMismatch = fmt.Errorf("assert: header length mismatch")
	ErrorHeaderStringMismatch = fmt.Errorf("assert: header string mismatch")
)

func AssertCSVStruct(fileNameCSV string, fileNameStruct string) (bool, error) {
	csvInfo, err := helper.GetCsvInfo(fileNameCSV)
	if err != nil {
		return false, err
	}

	structInfo, err := helper.GetStructInfo(fileNameStruct)
	if err != nil {
		return false, err
	}

	if len(csvInfo.Header) != len(structInfo.Fields) {
		return false, ErrorHeaderLengthMismatch
	}

	for i, csvHeader := range csvInfo.Header {
		if csvHeader != structInfo.Fields[i].CSVTag {
			return false, ErrorHeaderStringMismatch
		}
	}

	return true, nil
}
