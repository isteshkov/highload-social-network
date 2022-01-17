package stores

import (
	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/database"
	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/errors"
)

const (
	ErrorTypeUnspecified   = "STORE_UNSPECIFIED_ERROR"
	ErrorTypeDoesNotExists = "STORE_ENTITY_DOES_NOT_EXISTS"
	ErrorTypeAlreadyExists = "STORE_ENTITY_ALREADY_EXISTS"
	ErrorTypeInconsistent  = "STORE_ENTITY_INCONSISTENT"
)

var (
	ErrorProducerUnspecified   = errors.NewProducer(ErrorTypeUnspecified)
	ErrorProducerDoesNotExists = errors.NewProducer(ErrorTypeDoesNotExists)
	ErrorProducerAlreadyExists = errors.NewProducer(ErrorTypeAlreadyExists)
	ErrorProducerInconsistent  = errors.NewProducer(ErrorTypeInconsistent)
)

var ErrorsList = []*errors.ErrorProducer{
	ErrorProducerUnspecified,
	ErrorProducerDoesNotExists,
	ErrorProducerAlreadyExists,
	ErrorProducerInconsistent,
}

func processError(errPtr *error) {
	if errPtr == nil || *errPtr == nil {
		return
	}

	err := *errPtr

	if errors.IsProducedBy(err, ErrorsList...) {
		return
	}

	if errors.IsProducedBy(err, database.ErrorsList...) {
		switch {
		case errors.IsProducedBy(err, database.ErrorProducerDoesNotExist):
			*errPtr = ErrorProducerDoesNotExists.Wrap(err)
		case errors.IsProducedBy(err, database.ErrorProducerAlreadyExist):
			*errPtr = ErrorProducerAlreadyExists.Wrap(err)
		default:
			*errPtr = ErrorProducerUnspecified.Wrap(err)
		}
		return
	}

	*errPtr = ErrorProducerUnspecified.Wrap(err)
}
