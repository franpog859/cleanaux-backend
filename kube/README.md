# Apply Kubernetes resources

To install Cleanaux Backend run scripts from `/scripts` or do:

## Minikube

To start minikube run:

```bash
minikube start --vm-driver={VM_DRIVER}
```

To apply ingress on minikube run:

```bash
minikube addons enable ingress
minikube addons disable addon-manager # see if everything works with: minikube addons list
kubectl apply -f ingress.yaml # wait a while
kubectl get ingress # remember {ADDRESS} and {HOSTS}
sudo vim /etc/hosts # add ingress {ADDRESS} {HOSTS} to the /etc/hosts file
```

Remember to manually delete volumes after work by running:

```bash
kubectl delete deployment,svc mysql
kubectl delete pvc mysql-pv-claim
kubectl delete pv mysql-pv-volume
```

## Google Cloud Platform

To connect to the cluster run the command provided by GCP:

```bash
gcloud container clusters get-credentials {CLUSTER_NAME} --zone {ZONE} --project {PROJECT}
```

To apply ingress on GCP run:

```bash
kubectl create clusterrolebinding cluster-admin-binding --clusterrole cluster-admin --user $(gcloud config get-value account)
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/mandatory.yaml
kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/master/deploy/provider/cloud-generic.yaml 
kubectl get pods --all-namespaces -l app.kubernetes.io/name=ingress-nginx --watch
kubectl apply -f ingress.yaml
```

## Both of them

To apply deployments and services run:

```bash
kubectl apply -f entry-service/
kubectl apply -f auth-service/
kubectl apply -f content-service/
```

To open MySQL client for content-service run:

```bash
kubectl run -it --rm --image=mysql:5.6 --restart=Never mysql-client -- mysql -h mysql -ppassword
```
