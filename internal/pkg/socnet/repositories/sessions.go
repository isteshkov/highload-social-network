package repositories

import (
	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/database"
	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/logging"
	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/models"
)

type Sessions interface {
	WithLogger(l logging.Logger) Sessions

	ByUUID(UUID string) (session models.Session, err error)
	ByUserUUID(UUID string) (session models.Session, err error)
	Set(session models.Session, withTx database.Transaction) (err error)
	Delete(userUUID string, withTx database.Transaction) (err error)
}
