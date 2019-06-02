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

echo "Building auth-service image..."
cd ${ROOT_DIR}/auth-service
IMAGE_ID=$(docker build -t auth-service . 2>/dev/null | awk '/Successfully built/{print $NF}')
if [[ "${IMAGE_ID}" == "" ]]; then
    echo "Failed to build the image!"
    exit 1
fi
echo "Pushing the auth-service image..."
docker tag ${IMAGE_ID} franpog859/auth-service:${TAG}
docker push franpog859/auth-service:${TAG}

echo "Building entry-service image..."
cd ${ROOT_DIR}/entry-service
IMAGE_ID=$(docker build -t entry-service . 2>/dev/null | awk '/Successfully built/{print $NF}')
if [[ "${IMAGE_ID}" == "" ]]; then
    echo "Failed to build the image!"
    exit 1
fi
echo "Pushing the entry-service image..."
docker tag ${IMAGE_ID} franpog859/entry-service:${TAG}
docker push franpog859/entry-service:${TAG}

echo "Building content-service image..."
cd ${ROOT_DIR}/content-service
IMAGE_ID=$(docker build -t content-service . 2>/dev/null | awk '/Successfully built/{print $NF}')
if [[ "${IMAGE_ID}" == "" ]]; then
    echo "Failed to build the image!"
    exit 1
fi
echo "Pushing the content-service image..."
docker tag ${IMAGE_ID} franpog859/content-service:${TAG}
docker push franpog859/content-service:${TAG}

echo "Successfully built images with TAG: ${TAG}"
