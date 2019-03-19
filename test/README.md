# Test

To perform integration tests run:

```sh
bash integration-test.sh
```

To test Entry Service and Auth Service manually run:

```sh
kubectl expose deployment entry-service --type=LoadBalancer --name=entry-loadbalancer # if ingress does not work correctly
minikube service entry-loadbalancer --url
curl -d "username=user1&password=pass1" -X POST {URL}/login # should return token
```

To test Content Service manually:

- create `mysql` resource
- run mysql client and apply test `*.sql` files
- expose loadbalancer for `mysql-service`
- get url from the loadbalancer via `minikube service {LOADBALANCER} --url`
- change Content Service `databaseBase` to the url without `http://`
- build docker image and run it on port `8000`
- get content with `curl http://localhost:8000/user/content`
- update content with `curl -H 'Accept: application/json' -X PUT -d '{"id":2}' http://localhost:8000/user/content`
- see the content changes