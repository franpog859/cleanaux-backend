echo "Preparing shell..."
set -o errexit ; set -o nounset
CURRENT_DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"
KUBERNETES_DIR="${CURRENT_DIR}/../kube"

minikube delete
minikube start

echo "Applying resources..."
kubectl apply -f ${KUBERNETES_DIR}/entry-service/
kubectl apply -f ${KUBERNETES_DIR}/auth-service/
kubectl apply -f ${KUBERNETES_DIR}/content-service/
kubectl apply -f ${KUBERNETES_DIR}/mysql-database/

echo "Setting up ingress..."
minikube addons enable ingress
minikube addons disable addon-manager

echo "Applying ingress..."
kubectl apply -f ${KUBERNETES_DIR}/ingress.yaml

echo "Wait for ingress and other resources to start."
echo "For more information go to the /kube/README.md file!"
