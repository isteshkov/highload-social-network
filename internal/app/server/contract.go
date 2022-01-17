package application

import (
	"net/http"

	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/common"
	error_codes "github.com/isteshkov/highload-social-network/internal/pkg/socnet/errors"
	DTO "github.com/isteshkov/highload-social-network/internal/app/server/dto"

	"github.com/gin-gonic/gin"
)

// ProcessEndpointResult
//send result of endpoint according our protocol(error codes e.t.c.)
func (a *Application) ProcessEndpointResult(c *gin.Context, response *DTO.Response, errPtr *error, endpointTag string) {
	if errPtr == nil || *errPtr == nil {
		if response != nil {
			if response.Status == "" {
				if response.ErrorCode == "" {
					response.Status = DTO.ResponseStatusSuccess
				} else {
					response.Status = DTO.ResponseStatusError
				}
			}
		}

		a.sendResponse(c, response, http.StatusOK)
		return
	}
	processError(errPtr)

	response.Description = (*errPtr).Error()
	response.Status = DTO.ResponseStatusError

	if err, ok := (*errPtr).(error_codes.HasType); ok {
		if !ok {
			a.sendResponse(c, response, http.StatusInternalServerError)
			a.l.ErrorF(*errPtr, "error in %s endpoint", endpointTag)
		}
		if withMessage, ok := (*errPtr).(error_codes.HasMsg); ok {
			response.Description = withMessage.ErrorMessage()
		}

		switch err.Type() {
		case ErrorTypeForbidden:
			softWriteResponseErrorCode(response, error_codes.ErrorCodeForbidden)
			a.sendResponse(c, response, http.StatusForbidden)
		case ErrorTypeInvalidData:
			softWriteResponseErrorCode(response, error_codes.ErrorCodeInvalidData)
			a.sendResponse(c, response, http.StatusBadRequest)
		case ErrorTypeUnauthorized:
			softWriteResponseErrorCode(response, error_codes.ErrorCodeUnAuthorized)
			a.sendResponse(c, response, http.StatusUnauthorized)
		default:
			if withCode, ok := err.(error_codes.HasCode); ok {
				if withCode.Code() != "" {
					softWriteResponseErrorCode(response, withCode.Code())
				} else {
					softWriteResponseErrorCode(response, err.Type())
				}
			} else {
				softWriteResponseErrorCode(response, err.Type())
			}
			a.sendResponse(c, response, http.StatusOK)
			a.l.ErrorF(*errPtr, "error in %s endpoint", endpointTag)
		}
		return
	}

	c.AbortWithStatus(http.StatusInternalServerError)
	a.l.ErrorF(*errPtr, "error in %s endpoint", endpointTag)
	return
}

func softWriteResponseErrorCode(response *DTO.Response, code string) {
	if response.ErrorCode == "" {
		response.ErrorCode = code
	}
}

func (a *Application) sendResponse(c *gin.Context, response *DTO.Response, statusCode int) {
	common.ResponseToGin(c, statusCode, response)
	return
}
