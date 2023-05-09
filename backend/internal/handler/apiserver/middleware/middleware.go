package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
)

var (
	ErrorTokenNotFound    = errors.New("user token not found")
	ErrorTokenInvalidType = errors.New("user token is of invalid type")
)

func ErrorHandler(c *gin.Context) {
	c.Next()
	rc := -1
	if len(c.Errors) > 0 {
		err := c.Errors[0]
		e := err.Unwrap()
		c.JSON(rc, gin.H{"error": e.Error()})
	}
}

func GetUserToken(c *gin.Context) (string, error) {
	rowToken, ok := c.Get("userToken")
	if !ok {
		return "", fmt.Errorf("%w", ErrorTokenNotFound)
	}
	token, ok := rowToken.(string)
	if !ok {
		return "", fmt.Errorf("%w", ErrorTokenInvalidType)
	}
	return token, nil
}
