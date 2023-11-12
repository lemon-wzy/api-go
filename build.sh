#!/bin/zsh
working_dir=$(cd $(dirname $0); pwd)
echo "working dir: $working_dir"
cd $working_dir
if test -e api-go-darwin-arm64 ; then
  rm api-go*
fi
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o api-go-linux main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o  api-go-windows main.go
CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o api-go-darwin-arm64 main.go
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o api-go-darwin-amd64 main.go
