package helper

import (
	"go/ast"
	"go/parser"
	"go/token"
	"reflect"
	"strings"
)

type StructInfo struct {
	Name   string
	Fields []StructRecord
}

type StructRecord struct {
	FieldName string
	CSVTag    string
}

func GetStructInfo(fileName string) (StructInfo, error) {
	fset := token.NewFileSet()

	node, err := parser.ParseFile(fset, fileName, nil, parser.AllErrors)
	if err != nil {
		return StructInfo{}, err
	}

	infoList := make([]StructInfo, 0)
	for _, decl := range node.Decls {
		genDecl, ok := decl.(*ast.GenDecl)
		if !ok {
			continue
		}

		for _, spec := range genDecl.Specs {
			typeSpec, ok := spec.(*ast.TypeSpec)
			if !ok {
				continue
			}

			structType, ok := typeSpec.Type.(*ast.StructType)
			if !ok {
				continue
			}

			records := make([]StructRecord, 0)
			for _, field := range structType.Fields.List {
				if field.Tag != nil {
					tag := reflect.StructTag(strings.Trim(field.Tag.Value, "`"))
					csvTag := tag.Get("csv")

					if csvTag != "" {
						r := StructRecord{
							FieldName: field.Names[0].Name,
							CSVTag:    csvTag,
						}
						records = append(records, r)
					}
				}
			}
			i := StructInfo{
				Name:   typeSpec.Name.Name,
				Fields: records,
			}
			infoList = append(infoList, i)
		}

	}
	return infoList[0], nil
}
