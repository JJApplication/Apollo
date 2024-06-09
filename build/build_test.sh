#!/usr/bin/env bash

# build for app
echo "GOPROXY: $(go env GOPROXY)"
echo "GOOS: $(go env GOOS)"
echo "GOARCH: $(go env GOARCH)"

date -u

echo "start to build"
build_date=$(date +%Y-%m-%d)
GOOS=$(go env GOOS)
GOARCH=$(go env GOARCH)
GOVersion=$(go env GOVERSION)
GitCommit=$(git rev-parse HEAD)
# deb mode with symbols
go build -mod=mod -o apollo ./

echo "done"
