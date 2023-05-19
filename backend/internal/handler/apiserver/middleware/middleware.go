package middleware

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	ErrorTokenNotFound    = errors.New("user token not found")
	ErrorTokenInvalidType = errors.New("user token is of invalid type")
)

func ErrorHandler(c *gin.Context) {
	c.Next()
	if len(c.Errors) > 0 {
		err := c.Errors[0]
		curErr := err.Unwrap()
		for errors.Unwrap(curErr) != nil {
			curErr = errors.Unwrap(curErr)
		}
		errStr := curErr.Error()
		if errors.Is(curErr, http.ErrNoCookie) {
			errStr = ""
		}
		c.JSON(-1, gin.H{"error": errStr})
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
