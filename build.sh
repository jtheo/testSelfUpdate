#!/usr/bin/env bash
#

if [[ -z ${version} ]]; then
  echo "version var not defined... "
  exit 1
fi

echo "Building version ${version}"
go build -ldflags "-s -w -X main.Version=v${version}"

scp ${PWD##*/} fr2.hako.us:/store/www.default/
