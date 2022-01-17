package repositories

import (
	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/database"
	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/logging"
	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/models"
	"time"
)

type Users interface {
	WithLogger(l logging.Logger) Users

	ByUUID(UUID string) (user models.User, err error)
	Set(user models.User, withTx database.Transaction) (err error)
	SetDeleted(user models.User, deletedAt time.Time, withTx database.Transaction) (err error)
}
