#!/usr/bin/env bash

echo "Preparing shell..."
set -o errexit ; set -o nounset # TODO: Make sure that errexit is ok to be here. 
CURRENT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
KUBERNETES_DIR="${CURRENT_DIR}/../kube"

echo "Deleting databases..."
kubectl delete -f ${KUBERNETES_DIR}/mysql-database/
kubectl delete -f ${KUBERNETES_DIR}/mongo-database/

echo "Deleting services..."
kubectl delete -f ${KUBERNETES_DIR}/entry-service/
kubectl delete -f ${KUBERNETES_DIR}/auth-service/
kubectl delete -f ${KUBERNETES_DIR}/content-service/

echo "Deleting minikube..."
minikube delete

echo "Minikube cleaned successfully!"