#!/usr/bin/env bash

function log() {
  echo "===== $(date) ==> ${*}"
}

log "Running go vet"
if ! go vet ./...; then
  echo "go vet failed"
  exit 1
fi

D=${PWD##*/}

name=${1:-$D}
dist=dist
mkdir -p "${dist}"

if [[ -e version ]]; then
  version=$(tr -d '\n' <./version)
else
  version=1
fi
oses=(linux darwin)
archs=(amd64 arm64)

log "building version ${version}..."
echo "Building "

for GOOS in "${oses[@]}"; do
  printf "  - %s " "${GOOS}"
  for GOARCH in "${archs[@]}"; do
    printf "%s " "${GOARCH}"
    dst="${dist}"
    mkdir -p "${dst}"
    echo "${version}" >"${dst}/version"
    # shellcheck disable=SC2097,SC2098
    GOOS=${GOOS} GOARCH=${GOARCH} CGO_ENABLED=0 \
      go build -ldflags "-s -w -X 'main.Version=${version}'" \
      -o "${dst}/${version}/${name}-${GOOS}-${GOARCH}" .
  done
  echo
done

version=$((version + 1))
echo ${version} >version
echo
