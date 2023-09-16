#!/usr/bin/env bash

# build for app
echo "GOPROXY: $(go env GOPROXY)"
echo "GOOS: $(go env GOOS)"
echo "GOARCH: $(go env GOARCH)"

date -u
echo "build swagger docs"
swag init -g ./apollo.go

echo "start to build"
build_date=$(date +%Y-%m-%d)
GOOS=$(go env GOOS)
GOARCH=$(go env GOARCH)
GOVersion=$(go env GOVERSION)
GitCommit=$(git rev-parse HEAD)
# deb mode with symbols
go build -mod=mod -trimpath -ldflags \
 "\
 -X 'github.com/JJApplication/Apollo/vars.BuildDate=${build_date}' \
 -X 'github.com/JJApplication/Apollo/vars.GOOS=${GOOS}' \
 -X 'github.com/JJApplication/Apollo/vars.GOARCH=${GOARCH}' \
 -X 'github.com/JJApplication/Apollo/vars.GOVersion=${GOVersion}' \
 -X 'github.com/JJApplication/Apollo/vars.GitCommit=${GitCommit}' \
 " \
 -o apollo ./
# prod
#go build -mod=mod -ldflags='-w -s' -trimpath -o apollo ./
echo "done"
