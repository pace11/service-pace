package middlewares

import (
	"net/http"
	"service-pace11/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" || !strings.HasPrefix(authHeader, "Bearer ") {
			utils.HttpResponse(c, nil, http.StatusUnauthorized, "Token", c.Request.Method, nil)
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, claims, err := utils.ParseToken(tokenString)
		if err != nil || !token.Valid {
			utils.HttpResponse(c, nil, http.StatusUnauthorized, "Token", c.Request.Method, nil)
			c.Abort()
			return
		}

		userID, ok := claims["user_id"].(string)
		if !ok {
			utils.HttpResponse(c, nil, http.StatusUnauthorized, "Token", c.Request.Method, nil)
			c.Abort()
			return
		}

		name, ok := claims["name"].(string)
		if !ok {
			utils.HttpResponse(c, nil, http.StatusUnauthorized, "Token", c.Request.Method, nil)
			c.Abort()
			return
		}

		c.Set("user_id", userID)
		c.Set("name", name)
		c.Next()
	}
}
