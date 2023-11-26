
go build -o greet.so -buildmode=plugin plugin.go
go build -o hotfix.so -buildmode=plugin dispatch.go