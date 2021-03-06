#!/usr/bin/env bash

# build for app
echo "GOPROXY: $(go env GOPROXY)"
echo "GOOS: $(go env GOOS)"
echo "GOARCH: $(go env GOARCH)"

echo $(date -u)
echo "build swagger docs"
swag init -g ./apollo.go

echo "start to build"
go build -ldflags='-w -s' -trimpath -o apollo ./
