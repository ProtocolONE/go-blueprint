#!/usr/bin/env sh

if [ -n "$1" ] && [ ${0:0:4} = "/bin" ]; then
  ROOT_DIR=$1/..
else
  ROOT_DIR="$( cd "$( dirname "$0" )" && pwd )/.."
fi

GO_IMAGE=p1hub/go
GO_IMAGE_TAG=1.12
DIND_IMAGE=p1hub/dind
DIND_IMAGE_TAG=latest
PROTOTOOL_IMAGE=p1hub/prototool
PROTOTOOL_IMAGE_TAG=latest
GO_PKG=github.com/ProtocolONE/go-blueprint
GOOS="linux"
GOARCH="amd64"
DOCKER_NETWORK="protocol-one-blueprint-default"
DOCKER_IMAGE=p1hub/protocol-one-blueprint
PROJECT_NAME="protocol-one-blueprint"