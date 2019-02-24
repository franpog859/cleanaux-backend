# Entry Service

```sh
docker build -t entry-service .
docker tag {IMAGE_ID} franpog859/entry-service:{TAG}
docker push franpog859/entry-service
```

```sh
minikube start
kubectl apply -f deployment.yaml
kubectl apply -f service.yaml
```

To get URL to the service run:
```sh
minikube service entry-service --url
```

Apply Ingress
```sh
minikube addons list
minikube addons enable ingress
minikube addons disable addon-manager
kubectl apply -f ingress.yaml # wait a while
kubectl get ingress # add {ADDRESS} {HOSTS} to the /etc/hosts file
```