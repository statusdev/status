
Start instances with `./cmd/api/api --log-level=debug --http-addr=6660`

Subscribe instance B to instance A

````
curl -X POST -H 'Content-Type: application/json' -i http://localhost:6661/subscriptions --data '{
  "url": "http://localhost:6660"
}'
````

Upload new media

```
curl -i -X POST -H "Content-Type: multipart/form-data" -F "type=image/jpg" -F "data=@/home/chris/Downloads/users.png" http://localhost:6660/status
```