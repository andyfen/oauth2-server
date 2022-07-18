# oauth2 server



### Local
```
$ docker-compose up -d
```

```
$ go run main.go
```

#### Credentials
```
curl http://localhost:8080/credentials
```

```
{
    "client_id": "M2JkMjFhYjRiYWI3YzdmYmU3NGI3MWEwNDI1M2Q0NjQ3YTUxZGZlN2EzODA1YTgzZGNkNmQ3MDIyODkwZjgyMw==",
    "client_secret": "wNDI1M2Q0NjQ3YTUxZGZlN2EzODAM2JGNkNmQ3MDIyODkwZjgyMwrRkMjFhYjRiYWI3YzdmYmU3NGI3MWE1YTgzZ"
}
```

#### Access token

```
curl -X POST http://localhost:8080/oauth2/token \
 -d "grant_type=client_credentials&client_id=<CLIENT_ID>&client_secret=<CLIENT_SECRET>&scope=read"
```

Response
```
{
  "access_token": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJiNjhhYjA0Zi0yYTEzLTRjNDItYmU0ZC05ZmEwMTVmYmRmYzIiLCJleHAiOjE2MTE1NTM2MjZ9.NWh_mTk9XUgCgRiDrfhRF7X5GvptUwyFR77cgQdQIDzIa4t22gnO50EBejkrzsud6cHE8OtOZW454M9V05qllg",
  "expires_in": 7200,
  "scope": "read",
  "token_type": "Bearer"
}
```

#### Protected

Param
```
curl http://localhost:8080/api/protected?access_token=<ACCESS_TOKEN>
```

Header Authorization: Bearer
```
curl -v -X GET http://localhost:8080/api/protected \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer <ACCESS_TOKEN>"
```


#### Scopes (Roll-Based Access Control)

todo


