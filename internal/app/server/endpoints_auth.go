package application

import (
	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/context"
	DTO "github.com/isteshkov/highload-social-network/internal/app/server/dto"

	"github.com/gin-gonic/gin"
)

// HandlerSignIn
// @Summary sign in to account
// @Description returns profile data
// @Param auth body DTO.RequestSignInBody true "auth data"
// @Success 200 {object} DTO.ResponsePayloadSignIn
// @Failure 400 {object} DTO.Response
// @Failure 401 {object} DTO.Response
// @Failure 403 {object} DTO.Response
// @Failure 405 {object} DTO.Response
// @Failure 500 {object} DTO.Response
// @Router /api/v1/auth/signin [post]
func (a Application) HandlerSignIn(c *gin.Context) {
	var response DTO.Response
	var err error
	a.SetLogger(a.l.WithFields(context.LogFieldsFromContext(c)))
	defer a.ProcessEndpointResult(c, &response, &err, "auth_sign_in")

}

// HandlerSignUp
// @Summary sign up a new account
// @Description returns profile data
// @Param auth body DTO.RequestSignUpBody true "auth data"
// @Success 200 {object} DTO.ResponsePayloadSignUp
// @Failure 400 {object} DTO.Response
// @Failure 401 {object} DTO.Response
// @Failure 403 {object} DTO.Response
// @Failure 405 {object} DTO.Response
// @Failure 500 {object} DTO.Response
// @Router /api/v1/auth/signup [post]
func (a Application) HandlerSignUp(c *gin.Context) {
	var response DTO.Response
	var err error
	a.SetLogger(a.l.WithFields(context.LogFieldsFromContext(c)))
	defer a.ProcessEndpointResult(c, &response, &err, "auth_sign_up")

}

// HandlerSignUpCheck
// @Summary check email existing in the system
// @Description use to find out weather email registered or not
// @Param email query string true "email to check"
// @Success 200 {object} DTO.ResponsePayloadSignUpCheck
// @Failure 400 {object} DTO.Response
// @Failure 401 {object} DTO.Response
// @Failure 403 {object} DTO.Response
// @Failure 405 {object} DTO.Response
// @Failure 500 {object} DTO.Response
// @Router /api/v1/auth/signup/check [get]
func (a Application) HandlerSignUpCheck(c *gin.Context) {
	var response DTO.Response
	var err error
	a.SetLogger(a.l.WithFields(context.LogFieldsFromContext(c)))
	defer a.ProcessEndpointResult(c, &response, &err, "auth_sign_up_check")

}

// HandlerSignOut
// @Summary invalidate session
// @Description invalidate user auth session
// @Param X-Access-Token header string true "auth token"
// @Success 200 {object} DTO.ResponsePayloadSignOut
// @Failure 400 {object} DTO.Response
// @Failure 401 {object} DTO.Response
// @Failure 403 {object} DTO.Response
// @Failure 405 {object} DTO.Response
// @Failure 500 {object} DTO.Response
// @Router /api/v1/auth/signup/check [get]
func (a Application) HandlerSignOut(c *gin.Context) {
	var response DTO.Response
	var err error
	a.SetLogger(a.l.WithFields(context.LogFieldsFromContext(c)))
	defer a.ProcessEndpointResult(c, &response, &err, "auth_sign_out")

}
