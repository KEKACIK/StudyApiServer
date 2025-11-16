package router

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func (r *Router) AuthMiddleware(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, ErrorResponse{
			Message: "Authorization header required",
		})
		return
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		c.AbortWithStatusJSON(http.StatusUnauthorized, ErrorResponse{
			Message: "Invalid authorization format",
		})
		return
	}

	token := authHeader[7:]
	if token == "" {
		c.AbortWithStatusJSON(http.StatusUnauthorized, ErrorResponse{
			Message: "Token required",
		})
		return
	}

	if token != r.token {
		c.AbortWithStatusJSON(http.StatusUnauthorized, ErrorResponse{
			Message: "Token invalid",
		})
		return
	}
}
