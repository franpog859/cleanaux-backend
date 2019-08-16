#!/usr/bin/env bash

echo "Checking if resources are ready..."

READY=0
for i in {1..90}; do # 90 * 10 s = 15 min
    PODS="$( kubectl get po | grep 0/ )"
    INGRESS="$( kubectl get ing | grep \\. )"

    if [ "$PODS" == "" ] && [ "$INGRESS" != "" ]; then
        READY=1
        break
    fi

    echo "Waiting for resources..."
    sleep 10
done

if [ $READY == 0 ]; then
    echo "Something went wrong. It has been going on for over 15 minutes!"
    exit 1
fi

echo "Cleanaux Backend is ready!"
