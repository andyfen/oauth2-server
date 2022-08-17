# oauth2 authorization server



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

curl http://localhost:8080/credentials


curl -X POST http://localhost:8080/oauth2/token \
 -d "grant_type=client_credentials&client_id=YjAwY2Y3Nzc5NTg4NTBhOTNkZjhiZjg2ZTg5YWVkOTY1ODYwMGI3MzlhNTQ0MDU1ZDgwNGQ3OTdlYzkwN2Q0ZQ==&client_secret=NGZhMjk5NGIwZjA3OWU3MDM1MjNhNjIyMjZkNDg0YmE3OWZjZDE5ODI2ZWQwZGQ3N2FjMmY3YWNhYzYyMzg2ZQ==&scope=read"

curl http://localhost:8080/api/protected?access_token=eyJhbGciOiJIUzUxMiIsInR5cCI6IkpXVCJ9.eyJhdWQiOiJZakF3WTJZM056YzVOVGc0TlRCaE9UTmtaamhpWmpnMlpUZzVZV1ZrT1RZMU9EWXdNR0kzTXpsaE5UUTBNRFUxWkRnd05HUTNPVGRsWXprd04yUTBaUT09IiwiZXhwIjoxNjYwNDk4NzMwfQ.CBWeTGuUwcmMqiS-fYXqcdGm6_lbJOhjA7QuH_CKdpaqry8omS76u16fu-k3LBWzPNFTziJ1FTrAkP6aenqaxA