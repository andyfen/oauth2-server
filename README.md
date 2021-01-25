# OAuth2 Server

### Run
```
$ docker-compose up -d
```

### Credentials
```
curl http://localhost:8080/credentials
```

```
{
  "client_id": "3dbb92f9-4e08-4e12-b277-b7ce8a108edb",
  "client_secret": "70b963f8-7d9b-4d19-845c-44a20239dd20"
}
```

### Access token
GET
```
curl http://localhost:8080/oauth2/token?grant_type=client_credentials&client_id=b68ab04f-2a13-4c42-be4d-9fa015fbdfc2&client_secret=b93755cb-25f8-4d72-8755-6b05bc1ed169&scope=read
```

POST
```
curl -v -X POST http://localhost:8080/oauth2/token \
 -d "grant_type=client_credentials&client_id=3dbb92f9-4e08-4e12-b277-b7ce8a108edb&client_secret=70b963f8-7d9b-4d19-845c-44a20239dd20&scope=read"
```


```
{
  "access_token": "eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJiNjhhYjA0Zi0yYTEzLTRjNDItYmU0ZC05ZmEwMTVmYmRmYzIiLCJleHAiOjE2MTE1NTM2MjZ9.NWh_mTk9XUgCgRiDrfhRF7X5GvptUwyFR77cgQdQIDzIa4t22gnO50EBejkrzsud6cHE8OtOZW454M9V05qllg",
  "expires_in": 7200,
  "scope": "read",
  "token_type": "Bearer"
}
```

### Protected
Query Param
```
curl http://localhost:8080/api/protected?access_token=eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJiNjhhYjA0Zi0yYTEzLTRjNDItYmU0ZC05ZmEwMTVmYmRmYzIiLCJleHAiOjE2MTE1NTM3MDV9.3oor4USEWfJVNWxIZGLkaMx7afi3nGqykrb7Q8lY7bmMyGalDAWOUW7f5747324cTxPnJujT9wmYJzA4pxvPyg
```

Authorization Header
```
curl -v -X GET http://localhost:8080/api/protected \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJiNjhhYjA0Zi0yYTEzLTRjNDItYmU0ZC05ZmEwMTVmYmRmYzIiLCJleHAiOjE2MTE1NTM3MDV9.3oor4USEWfJVNWxIZGLkaMx7afi3nGqykrb7Q8lY7bmMyGalDAWOUW7f5747324cTxPnJujT9wmYJzA4pxvPyg"
```


#### Dev 
```
$ make dev
```

```
$ make stop
```

