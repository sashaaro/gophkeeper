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
createRelease() {
  curl -L -sf \
    -X POST \
    -H "Accept: application/vnd.github+json" \
    -H "Authorization: Bearer ${GITHUB_TOKEN}" \
    -H "X-GitHub-Api-Version: 2022-11-28" \
    https://api.github.com/repos/sashaaro/gophkeeper/releases \
    -d "{\"tag_name\":\"${APP_VERSION}\",\"target_commitish\":\"main\",\"name\":\"${APP_VERSION}\",\"body\":\"Description of the release\",\"draft\":false,\"prerelease\":false,\"generate_release_notes\":false}"
}

uploadAsset() {
  local ID=$1
  local FILENAME=$2
  local url=https://uploads.github.com/repos/sashaaro/gophkeeper/releases/${ID}/assets?name=${FILENAME}
  echo "Upload asset ${FILENAME} to release ${ID}"
  echo "POST ${url}"
  curl -sf \
    -X POST \
    -H "Accept: application/vnd.github+json" \
    -H "Authorization: Bearer ${GITHUB_TOKEN}" \
    -H "X-GitHub-Api-Version: 2022-11-28" \
    -H "Content-Type: application/octet-stream" \
    "${url}" \
    --data-binary "@${FILENAME}"
}

echo "Create Release"
RELEASE_ID=`createRelease | grep -m1 -oP '(?<="id": )([^,]*)'`
echo "Created RELEASE_ID: ${RELEASE_ID}"
cd ./build
ls -la
uploadAsset "${RELEASE_ID}" "client-linux-amd64.${APP_VERSION}.tar.gz"
uploadAsset "${RELEASE_ID}" "client-darwin-amd64.${APP_VERSION}.tar.gz"
uploadAsset "${RELEASE_ID}" "client-windows-amd64.${APP_VERSION}.tar.gz"