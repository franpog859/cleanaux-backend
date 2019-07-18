#!/usr/bin/env bash

echo "Preparing shell..."
set -o errexit ; set -o nounset
CURRENT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
ROOT_DIR="${CURRENT_DIR}/.."

echo "Reading TAG name..."
if [ $# -ne 1 ]; then
    echo "You should provide the TAG name! Run:"
    echo "${0} {YOUR_TAG_NAME}"
    exit 1
fi
TAG=$1

echo "Building and pushing auth-service image..."
cd ${ROOT_DIR}/auth-service
bash build-and-push-image.sh ${TAG}

echo "Building and pushing content-service image..."
cd ${ROOT_DIR}/content-service
bash build-and-push-image.sh ${TAG}

echo "Successfully built images with TAG: ${TAG}"
