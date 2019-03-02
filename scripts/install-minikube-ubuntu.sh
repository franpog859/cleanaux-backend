minikube delete
minikube start

cd ../kube
kubectl apply -f entry-service/
kubectl apply -f auth-service/
kubectl apply -f content-service/

minikube addons enable ingress
minikube addons disable addon-manager
kubectl apply -f ingress.yaml

echo "For more informations go to /kube/README.md file!"