#!/usr/bin/env bash

# build for app
echo "GOPROXY: $(go env GOPROXY)"
echo "GOOS: $(go env GOOS)"
echo "GOARCH: $(go env GOARCH)"

date -u
echo "build swagger docs"
swag init -g ./apollo.go

echo "start to build"
# deb mode with symbols
go build -mod=mod -trimpath -o apollo ./
# prod
#go build -mod=mod -ldflags='-w -s' -trimpath -o apollo ./
echo "done"
