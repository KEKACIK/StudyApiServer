package router

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	ErrAuthHeaderRequired = errors.New("invalid authorization: header required")
	ErrAuthInvalidFormat  = errors.New("invalid authorization: invalid format")
	ErrAuthTokenRequired  = errors.New("invalid authorization: token required")
	ErrAuthInvalidToken   = errors.New("invalid authorization: token invalid")
)

func (r *Router) AuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, ErrorResponse{
			Message: ErrAuthHeaderRequired.Error(),
		})
		return
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		c.AbortWithStatusJSON(http.StatusUnauthorized, ErrorResponse{
			Message: ErrAuthInvalidFormat.Error(),
		})
		return
	}

	token := authHeader[7:]
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, ErrorResponse{
			Message: ErrAuthTokenRequired.Error(),
		})
		return
	}

	if token != r.token {
		c.AbortWithStatusJSON(http.StatusUnauthorized, ErrorResponse{
			Message: ErrAuthInvalidToken.Error(),
		})
		return
	}
}
