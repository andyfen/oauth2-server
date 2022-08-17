package auth

import (
	"log"
	"time"

	oauth2gorm "github.com/andyfen/oauth-server/server/auth/oauth2gorm"
	"github.com/andyfen/oauth-server/server/config"
	"github.com/dgrijalva/jwt-go"
	"github.com/go-oauth2/oauth2/v4/errors"
	"github.com/go-oauth2/oauth2/v4/generates"
	"github.com/go-oauth2/oauth2/v4/manage"
	"github.com/go-oauth2/oauth2/v4/server"
)

func NewAuthManager(config *config.Config, clientStore *oauth2gorm.ClientStore) *manage.Manager {
	manager := manage.NewDefaultManager()

	// use mysql token store
	store := oauth2gorm.NewTokenStore(
		oauth2gorm.NewConfig("postgres://root:secret@localhost:5432/mydb", ""),
		0,
	)

	defer store.Close()

	manager.MapTokenStorage(store)

	// ----
	manager.MapClientStorage(clientStore)

	// generate jwt access token
	manager.MapAccessGenerate(
		generates.NewJWTAccessGenerate("", []byte(config.JWTKey), jwt.SigningMethodHS512),
	)

	return manager
}

func NewClientStore() *oauth2gorm.ClientStore {
	return oauth2gorm.NewClientStore(oauth2gorm.NewConfig("postgres://root:secret@localhost:5432/mydb", ""))
}

func NewAuthServer(manager *manage.Manager) *server.Server {

	srv := server.NewDefaultServer(manager)
	srv.SetAllowGetAccessRequest(true)

	srv.SetClientInfoHandler(server.ClientFormHandler)

	manager.SetRefreshTokenCfg(manage.DefaultRefreshTokenCfg)

	// set the client grant token config
	manager.SetClientTokenCfg(&manage.Config{
		AccessTokenExp:    time.Duration(60) * time.Second, // FIXME
		RefreshTokenExp:   time.Duration(24) * time.Hour,
		IsGenerateRefresh: true,
	})

	srv.SetInternalErrorHandler(func(err error) (re *errors.Response) {
		log.Println("Internal Error:", err.Error())
		return
	})

	srv.SetResponseErrorHandler(func(re *errors.Response) {
		log.Println("Response Error:", re.Error.Error())
	})

	return srv
}
