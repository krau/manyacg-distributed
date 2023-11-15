go build -v -x -o dist/collector collector/main.go
go build -v -x -o dist/core core/main.go
go build -v -x -o dist/storage storage/main.go
