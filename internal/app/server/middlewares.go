package application

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"net/http"
	"runtime/debug"
	"time"

	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/common"
	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/context"
	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/errors"
	responseRenderer "github.com/isteshkov/highload-social-network/internal/pkg/socnet/json"
	"github.com/isteshkov/highload-social-network/internal/app/server/dto"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var (
	producer = errors.NewProducer("general error")
)

func (a *Application) MiddlewareResponse(c *gin.Context) {
	logger := a.l.WithFields(logFieldsFromContext(c))

	done := make(chan bool)
	go func() {
		defer func() {
			panicInfo := recover()
			if panicInfo != nil {
				a.l.WithField("stack", string(debug.Stack())).Fatal(fmt.Sprintf("panic recovered: %+v", panicInfo))
				common.ResponseToGin(c, http.StatusInternalServerError, nil)
				c.Abort()
			}
			done <- panicInfo != nil
		}()
		c.Next()
	}()

	ticker := time.NewTicker(time.Second * time.Duration(a.cfg.TimeOutSecond))
	defer ticker.Stop()
	for {
		select {
		case <-ticker.C:
			logger.Error(producer.New("timeout exceeded"))
			common.ResponseToGin(c, http.StatusBadGateway, nil)
			c.Abort()
			return
		case <-done:
			statusCode, response := common.ResponseBuffer(c)
			if statusCode == 0 {
				statusCode = http.StatusInternalServerError
			}
			if response != nil {
				c.Render(statusCode, responseRenderer.EmptySliceRender{Data: response})
			} else {
				c.AbortWithStatus(statusCode)
				c.Abort()
			}
			return
		}
	}
}

func (a *Application) MiddlewareAccess(c *gin.Context) {
	requestId := c.Request.Header.Get(dto.HeaderRequestId)
	if _, err := uuid.Parse(requestId); err != nil {
		requestId = common.NewUUIDv4()
	}

	path := c.Request.URL.Path
	generalLoggerFields := map[string]interface{}{
		"path":       path,
		"remote":     c.Request.RemoteAddr,
		"request_id": requestId,
	}

	logger := a.l.WithFields(generalLoggerFields)

	c.Set(context.KeyLogFields, generalLoggerFields)
	c.Set(context.KeyRequestID, requestId)

	tsBeforeProcess := time.Now().UTC()

	requestData, err := ioutil.ReadAll(c.Request.Body)
	if err == nil {
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(requestData))
	}

	var requestBody interface{}
	err = json.Unmarshal(requestData, &requestBody)
	if err != nil {
		requestBody = struct {
			ErrorMsg string
		}{fmt.Sprintf("error during unmarshal request: %s string request :%s", err.Error(), string(requestData))}
		err = nil
	}

	c.Next()

	//regular requests from load balancer
	if path == "/" {
		return
	}

	statusCode, responseBody := common.ResponseBuffer(c)
	if responseBody == nil {
		responseBody = struct {
			UnMarshalingError string
		}{UnMarshalingError: "missing response struct"}
		err = nil
	}
	latency := math.Floor(time.Now().UTC().Sub(tsBeforeProcess).Seconds()*1000) / 1000
	a.metrics.NotifyRequestDone(path, latency)
	logger.WithFields(
		map[string]interface{}{
			"status_code":   statusCode,
			"latency":       latency,
			"method":        c.Request.Method,
			"request_body":  requestBody,
			"response_body": responseBody,
			"headers":       c.Request.Header,
			"queries":       c.Request.URL.Query(),
			"requester_uid": context.GetStringFromGin(c, context.KeyRequestUserUid),
		}).
		Info("ACCESS")

	return
}

func (a *Application) MiddlewareAuth(c *gin.Context) {
	logger := a.l.WithFields(logFieldsFromContext(c))

	token := c.Request.Header.Get(dto.HeaderAccessToken)
	if token == "" {
		common.ResponseToGin(c, http.StatusUnauthorized, nil)
		c.Abort()
		return
	}

	userUid, err := a.services.Auth.ValidateAccessToken(token, context.NewFromGin(c))
	if err != nil {
		common.ResponseToGin(c, http.StatusUnauthorized, nil)
		c.Abort()
		logger.Error(ErrorProducerInfraService.Wrap(err))
		return
	}

	c.Set(context.KeyRequestUserUid, userUid)
	c.Next()

	return
}

func middlewareCORS() gin.HandlerFunc {
	conf := cors.DefaultConfig()
	conf.AllowAllOrigins = true
	conf.AllowHeaders = append(conf.AllowHeaders, dto.HeaderAccessToken)
	return cors.New(conf)
}
