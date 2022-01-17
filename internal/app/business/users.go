package business

import (
	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/logging"
	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/repositories"
	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/services"
)

type UsersService struct {
	l    logging.Logger
	repo repositories.Users
}

func NewUsersService(l logging.Logger, r repositories.Users) *UsersService {
	return &UsersService{
		l:    l,
		repo: r,
	}
}

func (u UsersService) WithLogger(l logging.Logger) services.Users {
	u.l = l
	u.repo.WithLogger(l)
	return &u
}
