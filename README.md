# OAuth2 Server

### Run
```
$ docker-compose up -d
```

commands use [jq](https://stedolan.github.io/jq/)

### Credentials
```
curl http://localhost:8080/credentials |\
jq "."
```

```
{
  "client_id":"000000",
  "client_secret":"999999"
}
```

### Access token
via GET
```
curl http://localhost:8080/oauth2/token?grant_type=client_credentials&client_id=000000&client_secret=999999&scope=read | \
jq "."
```

via POST
```
curl -v -X POST http://localhost:8080/oauth2/token \
 -d "grant_type=client_credentials&client_id=000000&client_secret=999999&scope=read" | \
jq "."
```


```
{
  "access_token": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiIwMDAwMDAiLCJleHAiOjE2MTE1Mzg4ODB9.-DKQhKotHv3hHDJJK0TChntHBaUMdwdceRp2U9fhx8INu4bqkvRF9OBJyVxHDKOIxC7g2Zdn9ukeOO7a-iIrBw",
  "expires_in": 7200,
  "scope": "read",
  "token_type": "Bearer"
}
```

### Protected
via Query Param
```
curl http://localhost:8080/api/protected?access_token=eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiIwMDAwMDAiLCJleHAiOjE2MTE1NDkzNTN9.vtzYhjRwI7N-xDsbvvVtq17Bl5HQv6vABhdlXFV6DaeAwen7xpbrfKrZnN7vl-oxeQpW95qGWXeN4rsH_t3xSQ | \
jq "."
```

via Header
```
curl -v -X GET http://localhost:8080/api/protected \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiIwMDAwMDAiLCJleHAiOjE2MTE1NDkzNTN9.vtzYhjRwI7N-xDsbvvVtq17Bl5HQv6vABhdlXFV6DaeAwen7xpbrfKrZnN7vl-oxeQpW95qGWXeN4rsH_t3xSQ" | \
jq "."
```


#### Dev 
```
$ make dev 
```

```
$ make dev-down
```

