package database

import (
	"database/sql"
	stdErrors "errors"

	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/errors"

	"github.com/lib/pq"
)

var (
	ErrorProducerUnspecified = errors.NewProducer("unspecified error")

	ErrorProducerDoesNotExist = errors.NewProducer("entity does not exist")
	ErrorProducerAlreadyExist = errors.NewProducer("entity already exist")
)

var ErrorsList = []*errors.ErrorProducer{
	ErrorProducerUnspecified,
	ErrorProducerDoesNotExist,
	ErrorProducerAlreadyExist,
}

//nolint:errorlint
func processError(errPtr *error) {
	if errPtr == nil || *errPtr == nil {
		return
	}

	err := *errPtr

	if errors.IsProducedBy(err, ErrorsList...) {
		return
	}

	if pqErr, ok := err.(*pq.Error); ok {
		// pq errors code corresponding duplicate primary key
		if pqErr.Code == "23505" {
			*errPtr = ErrorProducerAlreadyExist.Wrap(err, 1)
			return
		}
	}

	if stdErrors.Is(err, sql.ErrNoRows) {
		*errPtr = ErrorProducerDoesNotExist.Wrap(err, 1)
		return
	}

	*errPtr = ErrorProducerUnspecified.Wrap(err, 1)
}
