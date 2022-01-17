package business

import (
	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/context"
	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/logging"
	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/repositories"
	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/services"
)

type AuthService struct {
	l    logging.Logger
	repo repositories.Sessions
}

func NewAuthService(l logging.Logger, r repositories.Sessions) *AuthService {
	return &AuthService{
		l:    l,
		repo: r,
	}
}

func (a AuthService) WithLogger(l logging.Logger) services.Auth {
	a.l = l
	a.repo.WithLogger(l)
	return &a
}

func (a AuthService) ValidateAccessToken(token string, ctx *context.Context) (userUid string, err error) {
	panic("implement me")
}
