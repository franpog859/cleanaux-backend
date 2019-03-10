echo "Applying resources..."
cd ../kube
kubectl apply -f entry-service/
kubectl apply -f auth-service/
kubectl apply -f content-service/

echo "Setting up ingress..."
kubectl create clusterrolebinding cluster-admin-binding --clusterrole cluster-admin --user $(gcloud config get-value account)
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/mandatory.yaml
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/provider/cloud-generic.yaml # for GKE
kubectl get pods --all-namespaces -l app.kubernetes.io/name=ingress-nginx

echo "Applying ingress..."
kubectl apply -f ingress.yaml

echo "Wait for ingress and other resources to start."
echo "For more information go to the /kube/README.md file!"