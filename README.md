# Go CSV to Struct Validator

## Usage

Command

```shell
go run ./cmd/main.go --csv ./testdata/test.csv --struct ./testdata/struct.go
```


Output(success)
```shell
❯ go run ./cmd/main.go --csv ./testdata/test.csv --struct ./testdata/struct.go
2024/09/03 06:07:04 assertion passed
```

Output(failure)
```shell
❯ go run ./cmd/main.go --csv ./testdata/test.csv --struct ./testdata/struct.go
2024/09/03 06:07:04 assertion failed
```
