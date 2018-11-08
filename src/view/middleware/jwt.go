package middleware

import (
	"github.com/gin-gonic/gin"
)

func JWTMiddleware(c *gin.Context) {
	c.Next()
}
