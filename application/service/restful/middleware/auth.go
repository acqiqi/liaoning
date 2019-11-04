package middleware

import (
	"github.com/gin-gonic/gin"
)

//验证中间件
func ValidateAuthMiddleware(notAuthRouterArr []string) gin.HandlerFunc {
	return func(c *gin.Context) {

	}
}
