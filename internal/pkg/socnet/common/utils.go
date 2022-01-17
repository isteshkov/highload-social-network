package common

import (
	"github.com/gin-gonic/gin"
	"github.com/isteshkov/highload-social-network/internal/pkg/socnet/models"
	"math/rand"
	"time"

	uuid "github.com/satori/go.uuid"
)

var letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func NewUUIDv4() string {
	return uuid.NewV4().String()
}

func CopyStringInterfaceMap(source map[string]interface{}) map[string]interface{} {
	result := make(map[string]interface{}, len(source))
	for key, value := range source {
		result[key] = value
	}

	return result
}

//nolint:gosec
func RandInt(min int, max int) int {
	rand.Seed(time.Now().UTC().UnixNano())
	return min + rand.Intn(max-min)
}

func RandString(length int) (result string) {
	var bytes []byte
	for i := 0; i < length; i++ {
		bytes = append(bytes, letters[RandInt(0, len(letters))])
	}

	return string(bytes)
}

func KeysFromStringBoolMap(data map[string]bool) []string {
	result := make([]string, 0, len(data))
	for key := range data {
		result = append(result, key)
	}

	return result
}

func KeysFromStringMap(data map[string]string) []string {
	result := make([]string, 0, len(data))
	for key := range data {
		result = append(result, key)
	}

	return result
}

func ValuesFromStringMap(data map[string]string) []string {
	result := make([]string, 0, len(data))
	for _, value := range data {
		result = append(result, value)
	}

	return result
}

func ResponseToGin(c *gin.Context, statusCode int, response interface{}) {
	c.Set(models.ResponseBufferKey, &models.ResponseBuffer{
		Response:   response,
		StatusCode: statusCode,
	})
}

func ResponseBuffer(c *gin.Context) (statusCode int, response interface{}) {
	buffer, ok := c.Get(models.ResponseBufferKey)
	if !ok {
		statusCode = 500
		return
	}

	return buffer.(*models.ResponseBuffer).StatusCode, buffer.(*models.ResponseBuffer).Response
}

func RemoveFromStringsSlice(s []string, elem string) []string {
	for i, v := range s {
		if v == elem {
			return append(s[:i], s[i+1:]...)
		}
	}
	return s
}