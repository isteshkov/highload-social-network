package business

import (
	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/logging"
	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/repositories"
	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/services"
)

type ProfilesService struct {
	l    logging.Logger
	repo repositories.Profiles
}

func NewProfilesService(l logging.Logger, r repositories.Profiles) *ProfilesService {
	return &ProfilesService{
		l:    l,
		repo: r,
	}
}

func (u ProfilesService) WithLogger(l logging.Logger) services.Profiles {
	u.l = l
	u.repo.WithLogger(l)
	return &u
}
