package stores

import (
	"context"

	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/database"
	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/logging"
	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/models"
	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/repositories"
)

type friendships struct {
	db database.Database
	l  logging.Logger
}

func NewFriendshipsStore(db database.Database, l logging.Logger) repositories.Friendships {
	return &friendships{
		db: db,
		l:  l,
	}
}

func (r friendships) WithLogger(l logging.Logger) repositories.Friendships {
	r.l = l
	r.db = r.db.WithLogger(l)
	return &r
}

func (r friendships) ByUserUUID(ctx context.Context, userUUID string) ([]models.Friendship, error) {
	panic("implement me")
}

func (r friendships) RequestsByUserUUID(ctx context.Context, userUUID string) ([]models.Friendship, error) {
	panic("implement me")
}
