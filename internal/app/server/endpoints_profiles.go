package application

import (
	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/context"
	DTO "github.com/isteshkov/highload-social-network/internal/app/server/dto"

	"github.com/gin-gonic/gin"
)

// HandlerProfile
// @Summary get profile
// @Description returns profile by uid
// @Param profile_uid path string false "profile uid"
// @Param X-Access-Token header string true "auth token"
// @Success 200 {object} DTO.ResponsePayloadProfile
// @Failure 400 {object} DTO.Response
// @Failure 401 {object} DTO.Response
// @Failure 403 {object} DTO.Response
// @Failure 405 {object} DTO.Response
// @Failure 500 {object} DTO.Response
// @Router /api/v1/profiles/{profile_uid} [get]
func (a Application) HandlerProfile(c *gin.Context) {
	var response DTO.Response
	var err error
	a.SetLogger(a.l.WithFields(context.LogFieldsFromContext(c)))
	defer a.ProcessEndpointResult(c, &response, &err, "profile_by_uid")

}

// HandlerSetProfile
// @Summary set profile data
// @Description set profile data
// @Param profile body DTO.RequestSetProfile true "profile model"
// @Param X-Access-Token header string true "auth token"
// @Success 200 {object} DTO.ResponsePayloadProfile
// @Failure 400 {object} DTO.Response
// @Failure 401 {object} DTO.Response
// @Failure 403 {object} DTO.Response
// @Failure 405 {object} DTO.Response
// @Failure 500 {object} DTO.Response
// @Router /api/v1/profiles [post]
func (a Application) HandlerSetProfile(c *gin.Context) {
	var response DTO.Response
	var err error
	a.SetLogger(a.l.WithFields(context.LogFieldsFromContext(c)))
	defer a.ProcessEndpointResult(c, &response, &err, "profile_set")

}

// HandlerProfiles
// @Summary get profiles list
// @Description returns profiles list according to filter and pagination
// @Param friends query bool false "weather my friends or not"
// @Param limit query string false "elements count by page"
// @Param offset query string false "pagination offset"
// @Param X-Access-Token header string true "auth token"
// @Success 200 {object} DTO.ResponsePayloadProfiles
// @Failure 400 {object} DTO.Response
// @Failure 401 {object} DTO.Response
// @Failure 403 {object} DTO.Response
// @Failure 405 {object} DTO.Response
// @Failure 500 {object} DTO.Response
// @Router /api/v1/profiles [get]
func (a Application) HandlerProfiles(c *gin.Context) {
	var response DTO.Response
	var err error
	a.SetLogger(a.l.WithFields(context.LogFieldsFromContext(c)))
	defer a.ProcessEndpointResult(c, &response, &err, "profiles_list")

}

// HandlerBecomeFriends
// @Summary become a friends with profile
// @Description set current user's profile as a friend with passed profile
// @Param X-Access-Token header string true "auth token"
// @Success 200 {object} DTO.ResponsePayloadFriendship
// @Failure 400 {object} DTO.Response
// @Failure 401 {object} DTO.Response
// @Failure 403 {object} DTO.Response
// @Failure 405 {object} DTO.Response
// @Failure 500 {object} DTO.Response
// @Router /api/v1/profiles/friendship/{profile_uid} [post]
func (a Application) HandlerBecomeFriends(c *gin.Context) {
	var response DTO.Response
	var err error
	a.SetLogger(a.l.WithFields(context.LogFieldsFromContext(c)))
	defer a.ProcessEndpointResult(c, &response, &err, "profile_become_friends")

}

// HandlerCeaseFriends
// @Summary cease a friendship with profile
// @Description cease current user's profile friendship with passed profile
// @Param X-Access-Token header string true "auth token"
// @Success 200 {object} DTO.ResponsePayloadFriendship
// @Failure 400 {object} DTO.Response
// @Failure 401 {object} DTO.Response
// @Failure 403 {object} DTO.Response
// @Failure 405 {object} DTO.Response
// @Failure 500 {object} DTO.Response
// @Router /api/v1/profiles/friendship/{profile_uid} [delete]
func (a Application) HandlerCeaseFriends(c *gin.Context) {
	var response DTO.Response
	var err error
	a.SetLogger(a.l.WithFields(context.LogFieldsFromContext(c)))
	defer a.ProcessEndpointResult(c, &response, &err, "profile_cease_friendship")

}
