# Test

To run unit tests:

- run `./../scripts/check-code.sh` script

To test the scenario manually:

- run mysql client with `kubectl run -it --rm --image=mysql:5.6 --restart=Never mysql-client -- mysql -h mysql-database-internal -ppassword` and apply test `*.sql` files
- run mongo interactively with `kubectl exec -it $(kubectl get po | grep -oE "\b(mongo)([a-zA-Z0-9-])+\b") /usr/bin/mongo` and apply test files
- get ingress host with `kubectl get ingress | grep -oE "\b([0-9]{1,3}\.){3}[0-9]{1,3}\b"`
- get JWT token with `curl -H 'Authorization: Basic dXNlcjE6cGFzczE=' -X POST {HOST}/login -v`
- get content with `curl -H 'Authorization: Bearer {TOKEN}' {HOST}/content -v`
- update content with `curl -H 'Authorization: Bearer {TOKEN}' -H 'Accept: application/json' -X PUT -d '{"id":{CONTENT_ID}}' {HOST}/content -v`
- get content with `curl -H 'Authorization: Bearer {TOKEN}' {HOST}/content -v`
- see the content changes
