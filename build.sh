#!/usr/bin/env bash

# build for app
echo "GOPROXY: $(go env GOPROXY)"
echo $(date -u)
echo "start to build"

go build -ldflags="-w -s" -trimpath -o dirichlet