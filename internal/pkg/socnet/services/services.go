package services

import "github.com/isteshkov/highload-social-network/internal/pkg/socnet/logging"

type Services struct {
	l logging.Logger
	Users
	Auth
	Friendships
}

func (s Services) WithLogger(l logging.Logger) *Services {
	s.l = l
	s.Auth = s.Auth.WithLogger(l)
	s.Users = s.Users.WithLogger(l)
	s.Friendships = s.Friendships.WithLogger(l)
	return &s
}

func NewServices(
	l logging.Logger,
	auth Auth,
	users Users,
	friendships Friendships,
) *Services {
	return &Services{
		l:           l,
		Users:       users,
		Auth:        auth,
		Friendships: friendships,
	}
}
