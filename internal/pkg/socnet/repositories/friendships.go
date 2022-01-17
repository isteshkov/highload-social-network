package repositories

import (
	"context"

	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/logging"
	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/models"
)

type Friendships interface {
	WithLogger(logger logging.Logger) Friendships

	ByUserUUID(ctx context.Context, userUUID string) ([]models.Friendship, error)
	RequestsByUserUUID(ctx context.Context, userUUID string) ([]models.Friendship, error)
}
