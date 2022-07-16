package token

import (
	"context"

	"github.com/go-oauth2/oauth2/v4"
)

var (
	TokenCtxKey = &contextKey{"Token"}
	ErrorCtxKey = &contextKey{"Error"}
)

type contextKey struct {
	name string
}

func NewContext(ctx context.Context, t oauth2.TokenInfo, err error) context.Context {
	ctx = context.WithValue(ctx, TokenCtxKey, t)
	ctx = context.WithValue(ctx, ErrorCtxKey, err)
	return ctx
}

func FromContext(ctx context.Context) (oauth2.TokenInfo, error) {
	tokenInfo, _ := ctx.Value(TokenCtxKey).(oauth2.TokenInfo)

	err, _ := ctx.Value(ErrorCtxKey).(error)

	return tokenInfo, err
}
