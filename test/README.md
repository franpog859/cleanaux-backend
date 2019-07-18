# Test

To test the scenario manually:

- run mysql client and apply test `*.sql` files
- get ingress host with `kubectl get ingress`
- get JWT token with `curl -X POST {HOST}/login -v`
- get content with `curl -H "Authorization: {TOKEN}" {HOST}/user/content -v`
- update content with `curl -H 'Accept: application/json' -H 'Authorization: {TOKEN}' -X PUT -d '{"id":{CONTENT_ID}}' {HOST}/user/content -v`
- get content with `curl -H "Authorization: {TOKEN}" {HOST}/user/content -v`
- see the content changes
