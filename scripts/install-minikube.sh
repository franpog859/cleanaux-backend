#!/usr/bin/env bash

echo "Preparing shell..."
set -o errexit ; set -o nounset
CURRENT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
KUBERNETES_DIR="${CURRENT_DIR}/../kube"

echo "Setting up VM driver..."
if [[ "$OSTYPE" == "darwin"* ]]; then
    VM_DRIVER="hyperkit"
elif [[ "$OSTYPE" == "linux-gnu" ]]; then
    VM_DRIVER="kvm2" # TODO: Check if it works correctly.
else
    echo "Unknown system. Unable to run minikube! Exitting..."
    exit 1
fi

echo "Starting minikube..."
minikube start --vm-driver=${VM_DRIVER}

echo "Applying services..."
kubectl apply -f ${KUBERNETES_DIR}/entry-service/
kubectl apply -f ${KUBERNETES_DIR}/auth-service/
kubectl apply -f ${KUBERNETES_DIR}/content-service/

echo "Applying databases..."
kubectl apply -f ${KUBERNETES_DIR}/mysql-database/deployment.yaml
kubectl apply -f ${KUBERNETES_DIR}/mysql-database/service.yaml
kubectl apply -f ${KUBERNETES_DIR}/mysql-database/persistent-volume-minikube.yaml
kubectl apply -f ${KUBERNETES_DIR}/mongo-database/deployment.yaml
kubectl apply -f ${KUBERNETES_DIR}/mongo-database/service.yaml
kubectl apply -f ${KUBERNETES_DIR}/mongo-database/persistent-volume-minikube.yaml

echo "Setting up ingress..."
minikube addons enable ingress
minikube addons disable addon-manager

echo "Applying ingress..."
kubectl apply -f ${KUBERNETES_DIR}/ingress.yaml

echo "Wait for ingress and other resources to start."
echo "For more information go to the /kube/README.md file!"
