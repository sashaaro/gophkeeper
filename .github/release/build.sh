#!/bin/bash

set -e

export APP_VERSION=$(./.github/release/version.sh)
rm -rf ./build
mkdir ./build
export CGO_ENABLED=0
echo "Build linux"
GOOS=linux GOARCH=amd64 go build -ldflags="-X main.appVersion=${APP_VERSION}" -o ./build/client-linux-amd64 -v ./cmd/client
echo "Build Mac"
GOOS=darwin GOARCH=amd64 go build -ldflags="-X main.appVersion=${APP_VERSION}" -o ./build/client-darwin-amd64 -v ./cmd/client
echo "Build Windows"
GOOS=windows GOARCH=amd64 go build -ldflags="-X main.appVersion=${APP_VERSION}" -o ./build/client-windows-amd64.exe -v ./cmd/client
echo "Gzip binaries"
tar -czf ./build/client-linux-amd64.${APP_VERSION}.tar.gz ./build/client-linux-amd64
tar -czf ./build/client-darwin-amd64.${APP_VERSION}.tar.gz ./build/client-darwin-amd64
tar -czf ./build/client-windows-amd64.${APP_VERSION}.tar.gz ./build/client-windows-amd64.exe

# TODO Нужно опубликовать релиз с этой версией и бинарями. См. https://docs.github.com/en/rest/releases/assets?apiVersion=2022-11-28#upload-a-release-asset"