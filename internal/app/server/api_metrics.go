package application

import (
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func (a *Application) buildMetricApi() API {
	router := gin.New()
	router.GET("/metrics", wrapPromHandler)
	return router
}

func wrapPromHandler(c *gin.Context) {
	promhttp.Handler().ServeHTTP(c.Writer, c.Request)
}
