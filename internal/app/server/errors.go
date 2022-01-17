package application

import (
	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/errors"
)

const (
	ErrorTypeInvalidData   = "INVALID_DATA_ERROR"
	ErrorTypeGeneral       = "GENERAL_ERROR"
	ErrorTypeInfraService  = "INFRA_SERVICE_ERROR"
	ErrorTypeUnauthorized  = "UNAUTHORIZED"
	ErrorTypeForbidden     = "FORBIDDEN"
	ErrorTypeDoesNotExists = "ENTITY_DOES_NOT_EXISTS"
	ErrorCodeCantBeDeleted = "ENTITY_CANT_BE_DELETED"
)

var (
	ErrorProducerUnauthorized  = errors.NewProducer(ErrorTypeUnauthorized)
	ErrorProducerInvalidData   = errors.NewProducer(ErrorTypeInvalidData)
	ErrorProducerGeneral       = errors.NewProducer(ErrorTypeGeneral)
	ErrorProducerInfraService  = errors.NewProducer(ErrorTypeInfraService)
	ErrorProducerForbidden     = errors.NewProducer(ErrorTypeForbidden)
	ErrorProducerDoesNotExists = errors.NewProducer(ErrorTypeDoesNotExists)
	ErrorProducerCantBeDeleted = errors.NewProducer(ErrorCodeCantBeDeleted)

	ErrorsList = []*errors.ErrorProducer{
		ErrorProducerInvalidData,
		ErrorProducerGeneral,
		ErrorProducerInfraService,
		ErrorProducerForbidden,
		ErrorProducerDoesNotExists,
		ErrorProducerCantBeDeleted,
	}
)

func processError(errPtr *error) {
	if errPtr == nil || *errPtr == nil {
		return
	}

	err := *errPtr
	if errors.IsProducedBy(err, ErrorsList...) {
		return
	}

	*errPtr = ErrorProducerGeneral.Wrap(err)
	return
}
