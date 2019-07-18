#!/usr/bin/env bash

echo "Preparing shell..."
set -o errexit ; set -o nounset
CURRENT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd ${CURRENT_DIR}

echo "Checking auth-service code..."
./../auth-service/check-code.sh

echo "Checking content-service code..."
./../content-service/check-code.sh
