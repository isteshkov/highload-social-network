package application

import (
	"net/http"

	"github.com/isteshkov/highload-social-network/internal/app/server/dto"

	"github.com/gin-gonic/gin"
)

func (a *Application) HandlerNotFound(c *gin.Context) {
	a.sendResponse(c, &dto.Response{}, http.StatusNotFound)
}

func (a *Application) HandlerMethodNotAllowed(c *gin.Context) {
	a.sendResponse(c, &dto.Response{}, http.StatusMethodNotAllowed)
}
