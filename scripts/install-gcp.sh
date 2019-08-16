#!/usr/bin/env bash

echo "Preparing shell..."
set -o errexit ; set -o nounset
CURRENT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
KUBERNETES_DIR="${CURRENT_DIR}/../kube"

echo "Applying databases..."
kubectl apply -f ${KUBERNETES_DIR}/mysql-database/persistent-volume-gcp.yaml
kubectl apply -f ${KUBERNETES_DIR}/mysql-database/deployment.yaml
kubectl apply -f ${KUBERNETES_DIR}/mysql-database/service.yaml
kubectl apply -f ${KUBERNETES_DIR}/mongo-database/persistent-volume-gcp.yaml
kubectl apply -f ${KUBERNETES_DIR}/mongo-database/deployment.yaml
kubectl apply -f ${KUBERNETES_DIR}/mongo-database/service.yaml

echo "Applying services..."
kubectl apply -f ${KUBERNETES_DIR}/auth-service/
kubectl apply -f ${KUBERNETES_DIR}/content-service/

echo "Setting up ingress..."
kubectl create clusterrolebinding cluster-admin-binding --clusterrole cluster-admin --user $(gcloud config get-value account)
kubectl create namespace ingress-nginx
kubectl apply --kustomize ${KUBERNETES_DIR}/ingress/

echo "Applying ingress..."
kubectl apply -f ${KUBERNETES_DIR}/ingress/ingress.yaml

echo "Waiting for ingress and other resources to start..."
bash ${CURRENT_DIR}/is-ready.sh

echo "For more information go to the /kube/README.md file!"
