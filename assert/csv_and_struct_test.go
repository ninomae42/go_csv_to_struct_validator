package assert_test

import (
	"errors"
	"testing"

	"github.com/ninomae42/csv_struct_validator/assert"
)

func TestAssertCSVStruct(t *testing.T) {
	type (
		give struct {
			fileNameCSV    string
			fileNameStruct string
		}
		want struct {
			ok  bool
			err error
		}
	)
	tests := []struct {
		name string
		give give
		want want
	}{
		{
			name: "[OK] 有効なCSVファイルと構造体ファイルを指定した場合",
			give: give{
				fileNameCSV:    "./testdata/ok1.csv",
				fileNameStruct: "./testdata/ok1.go",
			},
			want: want{
				ok:  true,
				err: nil,
			},
		},
		{
			name: "[NG] CSVファイルのヘッダーが構造体のフィールド数と異なる場合",
			give: give{
				fileNameCSV:    "./testdata/ng1.csv",
				fileNameStruct: "./testdata/ok1.go",
			},
			want: want{
				ok:  false,
				err: assert.ErrorHeaderLengthMismatch,
			},
		},
		{
			name: "[NG] CSVファイルのヘッダーが構造体のフィールド名と異なる場合",
			give: give{
				fileNameCSV:    "./testdata/ng2.csv",
				fileNameStruct: "./testdata/ok1.go",
			},
			want: want{
				ok:  false,
				err: assert.ErrorHeaderStringMismatch,
			},
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			ok, err := assert.AssertCSVStruct(tt.give.fileNameCSV, tt.give.fileNameStruct)
			if ok != tt.want.ok {
				t.Errorf("got %v, want %v", ok, tt.want.ok)
			}
			if !errors.Is(err, tt.want.err) {
				t.Errorf("got %v, want %v", err, tt.want.err)
			}
		})
	}
}
