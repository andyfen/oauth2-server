package auth

import (
	"github.com/andyfen/oauth-server/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/go-oauth2/oauth2/v4/store"
	oredis "github.com/go-oauth2/redis/v4"
	"github.com/go-redis/redis/v8"
	"log"
)

// NewAuthManager ...
func NewAuthManager(config *config.Config, clientStore *store.ClientStore) *manage.Manager {
	manager := manage.NewDefaultManager()

	// use redis token store
	manager.MapTokenStorage(oredis.NewRedisStore(&redis.Options{
		Addr:     config.Redis.Addr,
		Password: "", // no password set
		DB:       0,  // use default DB
	}))

	manager.MapClientStorage(clientStore)

	// generate jwt access token
	manager.MapAccessGenerate(generates.NewJWTAccessGenerate("", []byte(config.Auth.JWTKey), jwt.SigningMethodHS512))

	return manager
}

// NewClientStore ...
func NewClientStore() *store.ClientStore {
	clientStore := store.NewClientStore()
	return clientStore
}

// NewAuthServer ...
func NewAuthServer(manager *manage.Manager) *server.Server {

	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)

	srv.SetClientInfoHandler(server.ClientFormHandler)

	manager.SetRefreshTokenCfg(manage.DefaultRefreshTokenCfg)

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	return srv
}
