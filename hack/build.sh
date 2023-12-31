#!/bin/bash
set -ex

GIT_VERSION=$(git describe --always --tags || true)
VERSION=${CI_VERSION:-${GIT_VERSION}}
GIT_COMMIT=$(git rev-list -1 HEAD || true)
COMMIT=${CI_COMMIT:-${GIT_COMMIT}}
BUILD_DATE=$(date --utc -Iseconds)

mkdir -p bin

LDFLAGS="-s -w "
LDFLAGS+="-X github.com/mshitrit/customized-user-remediation/version.Version=${VERSION} "
LDFLAGS+="-X github.com/mshitrit/customized-user-remediation/version.GitCommit=${COMMIT} "
LDFLAGS+="-X github.com/mshitrit/customized-user-remediation/version.BuildDate=${BUILD_DATE} "
GOFLAGS=-mod=vendor CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="${LDFLAGS}" -o manager main.go
