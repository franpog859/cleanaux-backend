#!/usr/bin/env bash

echo "Preparing shell..."
set -o errexit ; set -o nounset # TODO: Make sure that errexit is ok to be here. 
CURRENT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
KUBERNETES_DIR="${CURRENT_DIR}/../kube"

echo "Deleting databases..."
kubectl delete -f ${KUBERNETES_DIR}/mysql-database/deployment.yaml
kubectl delete -f ${KUBERNETES_DIR}/mysql-database/service.yaml
kubectl delete -f ${KUBERNETES_DIR}/mysql-database/persistent-volume-minikube.yaml
kubectl delete -f ${KUBERNETES_DIR}/mongo-database/deployment.yaml
kubectl delete -f ${KUBERNETES_DIR}/mongo-database/service.yaml
kubectl delete -f ${KUBERNETES_DIR}/mongo-database/persistent-volume-minikube.yaml

echo "Deleting services..."
kubectl delete -f ${KUBERNETES_DIR}/auth-service/
kubectl delete -f ${KUBERNETES_DIR}/content-service/

echo "Cleaning up minikube..."
minikube addons disable ingress
minikube addons enable addon-manager

echo "Deleting minikube..."
minikube delete

echo "Minikube cleaned successfully!"