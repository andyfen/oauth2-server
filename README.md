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
    "client_id": "35c6261e85960ccdde15d4bd6ef6032c4d77b434217b411807242edbe2edafe7",
    "client_secret": "c2f7efda2b999e8acd565c793b31feab7d9cf2cb4cb1660d8bcfeae3f99b3f4d"
}
```

### Access token

```
curl -v -X POST http://localhost:8080/oauth2/token \
 -d "grant_type=client_credentials&client_id=35c6261e85960ccdde15d4bd6ef6032c4d77b434217b411807242edbe2edafe7&client_secret=c2f7efda2b999e8acd565c793b31feab7d9cf2cb4cb1660d8bcfeae3f99b3f4d&scope=read"
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

### Protected

Param
```
curl http://localhost:8080/api/protected?access_token=eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiIzNWM2MjYxZTg1OTYwY2NkZGUxNWQ0YmQ2ZWY2MDMyYzRkNzdiNDM0MjE3YjQxMTgwNzI0MmVkYmUyZWRhZmU3IiwiZXhwIjoxNjE0NzU3NTI2fQ.LJqHtXn_Y0waZRHHOz4CY7GRx3vGJjybAP4E2CrzVmdSLVCzWnJ1HBstzR4m-Y6XVXPqKR9ieh7vWwapLQemmw
```

Header Authorization: Bearer
```
curl -v -X GET http://localhost:8080/api/protected \
  -H "Content-Type: application/json" \
  -H "Authorization: Bearer eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiIzNWM2MjYxZTg1OTYwY2NkZGUxNWQ0YmQ2ZWY2MDMyYzRkNzdiNDM0MjE3YjQxMTgwNzI0MmVkYmUyZWRhZmU3IiwiZXhwIjoxNjE0NzU3NTI2fQ.LJqHtXn_Y0waZRHHOz4CY7GRx3vGJjybAP4E2CrzVmdSLVCzWnJ1HBstzR4m-Y6XVXPqKR9ieh7vWwapLQemmw"
```

