#!/bin/bash
set -e

echo "Running $0"

GOLANGCI_LINT_VER=v1.39.0
DOCKER_CMD=${DOCKER_CMD:-docker}

if [ ! $(command -v ${DOCKER_CMD}) ]; then
    exit 0
fi

echo "Linting top-level module..."
${DOCKER_CMD} run --rm -v $(pwd):/opt/workspace -w /opt/workspace golangci/golangci-lint:$GOLANGCI_LINT_VER golangci-lint run
