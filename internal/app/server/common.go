package application

import (
	"github.com/gin-gonic/gin"
	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/context"
)

func logFieldsFromContext(c *gin.Context) map[string]interface{} {
	fields, ok := c.Get(context.KeyLogFields)
	if !ok {
		return map[string]interface{}{
			"default_log_fields": "missing",
		}
	}

	result, ok := fields.(map[string]interface{})
	if !ok || len(result) == 0 {
		return map[string]interface{}{
			"default_log_fields": "missing",
		}
	}

	return result
}
