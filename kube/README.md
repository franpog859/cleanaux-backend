# Apply Kubernetes resources on Minikube

To apply deployments and services run:

```bash
minikube start
kubectl apply -f entry-service/
kubectl apply -f auth-service/
kubectl apply -f content-service/
```

To apply ingress run:

```bash
minikube addons enable ingress
minikube addons disable addon-manager # see if everything works with: minikube addons list
kubectl apply -f ingress.yaml # wait a while
kubectl get ingress # remember {ADDRESS} and {HOSTS}
sudo vim /etc/hosts # add ingress {ADDRESS} {HOSTS} to the /etc/hosts file
```

To apply ingress on GCP:

```bash
kubectl create clusterrolebinding cluster-admin-binding --clusterrole cluster-admin --user $(gcloud config get-value account)
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/mandatory.yaml
minikube addons enable ingress # for minikube
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/provider/cloud-generic.yaml # for GKE
kubectl get pods --all-namespaces -l app.kubernetes.io/name=ingress-nginx --watch
```

To open MySQL client for content-service run:

```bash
kubectl run -it --rm --image=mysql:5.6 --restart=Never mysql-client -- mysql -h mysql -ppassword
```

Remember to manually delete volumes after work by running:

```bash
kubectl delete deployment,svc mysql
kubectl delete pvc mysql-pv-claim
kubectl delete pv mysql-pv-volume
```
