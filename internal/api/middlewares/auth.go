package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func AuthMiddleware(context *gin.Context) {

	secretKey := os.Getenv("SECRET_KEY")
	if secretKey == "" {
		context.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{"message": "Secret key not configured"})
		return
	}
	// Logic
	token := context.Request.Header.Get("Authorization")

	if token == "" {
		context.AbortWithStatusJSON(http.StatusUnauthorized,
			gin.H{"message": "token is required"})
		return
	}

	// Validate token
	if secretKey != token {
		context.AbortWithStatusJSON(http.StatusUnauthorized,
			gin.H{"message": "invalid token"})
		return
	}

	context.Set("isAuth", true)
	context.Next()
}
