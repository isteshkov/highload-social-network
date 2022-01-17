package services

import (
	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/context"
	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/logging"
)

type Users interface {
	WithLogger(l logging.Logger) Users
}

type Auth interface {
	WithLogger(l logging.Logger) Auth

	ValidateAccessToken(token string, ctx *context.Context) (userUid string, err error)
}

type Friendships interface {
	WithLogger(l logging.Logger) Friendships
}