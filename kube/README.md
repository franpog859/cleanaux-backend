# Apply Kubernetes resources on Minikube

To apply deployments and services run:

```sh
minikube start
kubectl apply -f entry-service/
kubectl apply -f auth-service/
```

To apply ingress run:

```sh
minikube addons enable ingress
minikube addons disable addon-manager # see if everything works with: minikube addons list
kubectl apply -f ingress.yaml # wait a while
kubectl get ingress # remember {ADDRESS} and {HOSTS}
sudo vim /etc/hosts # add ingress {ADDRESS} {HOSTS} to the /etc/hosts file
```
