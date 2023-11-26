
go build -o greet.so -buildmode=plugin plugin.go
# go build -o hotfix.so.1 -buildmode=plugin dispatch.go
go build -o hotfix.so.2 -buildmode=plugin dispatch.go