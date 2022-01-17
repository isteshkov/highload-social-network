package application

import (
	"github.com/gin-gonic/gin"
)

// HandlerHealth
// @Summary Health
// @Description Health
// @Success 200
// @Router /api/health [get]
func (a Application) HandlerHealth(c *gin.Context) {
	a.ProcessEndpointResult(c, nil, nil, "health")
}
