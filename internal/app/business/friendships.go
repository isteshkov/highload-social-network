package business

import (
	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/logging"
	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/repositories"
	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/services"
)

type FriendshipsService struct {
	l    logging.Logger
	repo repositories.Friendships
}

func NewFriendshipsService(l logging.Logger, r repositories.Friendships) *FriendshipsService {
	return &FriendshipsService{
		l:    l,
		repo: r,
	}
}

func (u FriendshipsService) WithLogger(l logging.Logger) services.Friendships {
	u.l = l
	u.repo.WithLogger(l)
	return &u
}
