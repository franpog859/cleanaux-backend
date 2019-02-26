# Test

To perform integration tests run:

```sh
bash integration-test.sh
```

To test it manually run:

```sh
kubectl expose deployment entry-service --type=LoadBalancer --name=entry-test # if ingress does not work correctly
minikube service entry-test --url
curl -d "username=user1&password=pass1" -X POST {URL}/login # should return token
```