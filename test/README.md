# Test

To run unit tests:

- run `./../scripts/check-code.sh`

To test the scenario manually:

- run mysql client and apply test `*.sql` files
- run mongo interactively and apply test files
- get ingress host with `kubectl get ingress`
- get JWT token with `curl -H 'Authorization: Basic dXNlcjE6cGFzczE=' -X POST {HOST}/login -v`
- get content with `curl -H 'Authorization: Bearer {TOKEN}' {HOST}/user/content -v`
- update content with `curl -H 'Authorization: Bearer {TOKEN}' -H 'Accept: application/json' -X PUT -d '{"id":{CONTENT_ID}}' {HOST}/user/content -v`
- get content with `curl -H 'Authorization: Bearer {TOKEN}' {HOST}/user/content -v`
- see the content changes
