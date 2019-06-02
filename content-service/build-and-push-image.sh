#!/usr/bin/env bash

echo "Preparing shell..."
set -o errexit ; set -o nounset
CURRENT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
cd ${CURRENT_DIR}

echo "Reading TAG name..."
if [ $# -ne 1 ]; then
    echo "You should provide the TAG name!"
    exit 1
fi
TAG=$1

echo "Building content-service image..."
IMAGE_ID=$(docker build -t content-service . 2>/dev/null | awk '/Successfully built/{print $NF}')
if [[ "${IMAGE_ID}" == "" ]]; then
    echo "Failed to build the image!"
    exit 1
fi
echo "Pushing the content-service image..."
docker tag ${IMAGE_ID} franpog859/content-service:${TAG}
docker push franpog859/content-service:${TAG}

echo "Successfully built and pushed the image with TAG: ${TAG}"
