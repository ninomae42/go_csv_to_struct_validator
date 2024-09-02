package helper

import (
	"bytes"
	"encoding/csv"
	"io"
	"os"
)

type CSVInfo struct {
	FileName string
	Header   []string
}

func GetCsvInfo(fileName string) (CSVInfo, error) {
	r, err := NewReaderFromFile(fileName)
	if err != nil {
		return CSVInfo{}, err
	}

	header, err := GetCsvHeader(r)
	if err != nil {
		return CSVInfo{}, err
	}

	return CSVInfo{
		FileName: fileName,
		Header:   header,
	}, nil
}

func NewReaderFromFile(fileName string) (io.Reader, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, f); err != nil {
		return nil, err
	}

	return buf, nil
}

func GetCsvHeader(r io.Reader) ([]string, error) {
	reader := csv.NewReader(r)
	return reader.Read()
}
