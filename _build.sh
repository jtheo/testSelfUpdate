#!/bin/bash

echo "Running go vet"
if ! go vet ./...; then
    echo "go vet failed"
    exit 1
fi

echo "Running go test"
if ! go test -v; then
    echo "go test failed"
    exit 1
fi

if type -p govulncheck; then
    echo "Running govulncheck"
    govulncheck ./...
fi

if type -p golangci-lint; then
    echo "Running golangci-lint"
    golangci-lint run -E revive -E errcheck -E nilerr -E gosec
fi

D=$(basename "${PWD}")

name=${1:-$D}
mkdir -p bin

read -r h m s <<< "$(date "+%H %M %S")"
minor=$(( (${h##0} + 1) * (${m##0} + 1) * (${s##0} + 1) ))
major=$(date +%Y%m%d)
version="${major}.${minor}"

oses=( linux darwin )
archs=( amd64 arm64 )

for GOOS in "${oses[@]}";do
    for GOARCH in "${archs[@]}"; do
        echo "Building ${GOARCH} for ${GOOS}..."
        # shellcheck disable=SC2097,SC2098
        GOOS=${GOOS} GOARCH=${GOARCH} CGO_ENABLED=0 \
            go build -ldflags "-s -w -X 'main.Version=v${version}'" \
              -o "bin/${name}-${GOOS}-${GOARCH}" .
    done
done
