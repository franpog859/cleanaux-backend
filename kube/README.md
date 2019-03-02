# Apply Kubernetes resources on Minikube

To apply deployments and services run:

```sh
minikube start
kubectl apply -f entry-service/
kubectl apply -f auth-service/
kubectl apply -f content-service/
```

To apply ingress run:

```sh
minikube addons enable ingress
minikube addons disable addon-manager # see if everything works with: minikube addons list
kubectl apply -f ingress.yaml # wait a while
kubectl get ingress # remember {ADDRESS} and {HOSTS}
sudo vim /etc/hosts # add ingress {ADDRESS} {HOSTS} to the /etc/hosts file
```

To open MySQL client for content-service run:

```sh
kubectl run -it --rm --image=mysql:5.6 --restart=Never mysql-client -- mysql -h mysql -ppassword
```

Remember to manually delete volumes after work by running:

```sh
kubectl delete deployment,svc mysql
kubectl delete pvc mysql-pv-claim
kubectl delete pv mysql-pv-volume
```
