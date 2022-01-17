package context

import (
	"context"

	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/common"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/metadata"
)

func GetStringFromGin(c *gin.Context, key string) string {
	value, ok := c.Get(key)
	if !ok {
		return ""
	}

	stringValue, ok := value.(string)
	if !ok {
		return ""
	}

	return stringValue
}

func LogFieldsFromContext(c *gin.Context) map[string]interface{} {
	fields, ok := c.Get(KeyLogFields)
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

func NewFromGin(c *gin.Context) *Context {
	requestId := GetStringFromGin(c, KeyRequestID)
	if requestId == "" {
		requestId = common.NewUUIDv4()
	}

	return &Context{
		RequestId: requestId,
	}
}

func NewGrpcContext(c *Context) (ctx context.Context) {
	md := metadata.New(map[string]string{KeyRequestID: c.RequestId})
	ctx = metadata.NewOutgoingContext(context.Background(), md)
	return
}

func LogFieldsFromGrpcContext(c context.Context) map[string]interface{} {
	result := make(map[string]interface{})
	meta, ok := metadata.FromIncomingContext(c)
	if ok {
		IncomingRequestId := meta.Get(KeyRequestID)
		if len(IncomingRequestId) > 0 {
			result[KeyRequestID] = IncomingRequestId[0]
		}
	}

	return result
}

func NewGrpcFromGrpc(c context.Context) (ctx context.Context) {
	requestId := common.NewUUIDv4()
	meta, ok := metadata.FromIncomingContext(c)
	if ok {
		IncomingRequestId := meta.Get(KeyRequestID)
		if len(IncomingRequestId) > 0 {
			requestId = IncomingRequestId[0]
		}
	}
	md := metadata.New(map[string]string{KeyRequestID: requestId})
	ctx = metadata.NewOutgoingContext(context.Background(), md)
	return
}
