#!/usr/bin/env bash

echo "Preparing shell..."
set -o errexit ; set -o nounset
CURRENT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
KUBERNETES_DIR="${CURRENT_DIR}/../kube"

echo "Setting up VM driver..."
if [[ "$OSTYPE" == "darwin"* ]]; then
    VM_DRIVER="hyperkit"
elif [[ "$OSTYPE" == "linux-gnu" ]]; then
    VM_DRIVER="kvm2" 
else
    echo "Unknown system. Unable to run minikube! Exitting..."
    exit 1
fi

echo "Starting minikube..."
minikube start --vm-driver=${VM_DRIVER}

echo "Applying databases..."
kubectl apply -f ${KUBERNETES_DIR}/mysql-database/persistent-volume-minikube.yaml
kubectl apply -f ${KUBERNETES_DIR}/mysql-database/deployment.yaml
kubectl apply -f ${KUBERNETES_DIR}/mysql-database/service.yaml
sleep 60

kubectl apply -f ${KUBERNETES_DIR}/mongo-database/persistent-volume-minikube.yaml
kubectl apply -f ${KUBERNETES_DIR}/mongo-database/deployment.yaml
kubectl apply -f ${KUBERNETES_DIR}/mongo-database/service.yaml
sleep 60

echo "Applying services..."
kubectl apply -f ${KUBERNETES_DIR}/auth-service/
kubectl apply -f ${KUBERNETES_DIR}/content-service/

echo "Setting up ingress..."
minikube addons enable ingress
minikube addons disable addon-manager

echo "Applying ingress..."
kubectl apply -f ${KUBERNETES_DIR}/ingress/ingress.yaml

echo "Waiting for ingress and other resources to start..."
bash ${CURRENT_DIR}/is-ready.sh

echo "For more information go to the /kube/README.md file!"
echo "Remember to cleanup minikube after your work with cleanup-minikube.sh script!"
