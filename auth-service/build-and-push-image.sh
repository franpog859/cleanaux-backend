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

echo "Building auth-service image..."
IMAGE_ID=$(docker build -t auth-service . 2>/dev/null | awk '/Successfully built/{print $NF}')
if [[ "${IMAGE_ID}" == "" ]]; then
    echo "Failed to build the image!"
    exit 1
fi
echo "Pushing the auth-service image..."
docker tag ${IMAGE_ID} franpog859/auth-service:${TAG}
docker push franpog859/auth-service:${TAG}

echo "Successfully built and pushed the image with TAG: ${TAG}"
