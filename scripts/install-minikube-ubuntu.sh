minikube delete
minikube start

echo "Applying resources..."
cd ../kube
kubectl apply -f entry-service/
kubectl apply -f auth-service/
kubectl apply -f content-service/

echo "Setting up ingress..."
minikube addons enable ingress
minikube addons disable addon-manager

echo "Applying ingress..."
kubectl apply -f ingress.yaml

echo "Wait for ingress and other resources to start."
echo "For more information go to the /kube/README.md file!"