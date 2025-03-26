#!/usr/bin/env bash

function log() {
  echo "===== $(date) ==> ${*}"
}

log "Cleaning..."

rm -rfv version dist/*

log "Done"
