package middlewares

import (
	"io"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/valiant1012/transaction-service/src/auth"
	"github.com/valiant1012/transaction-service/src/config"
	"github.com/valiant1012/transaction-service/src/utility/logger"
)

const (
	BearerSchema = "Bearer"
)

func GinLoggerMiddleware() gin.HandlerFunc {
	ginLogFile, err := os.Create(config.GetGinLogFilePath())
	if err != nil {
		logger.Errorln("could not create gin log file:", err.Error())
	} else {
		gin.DefaultWriter = io.MultiWriter(os.Stdout, ginLogFile)
	}

	ginErrLogFile, err := os.Create(config.GetGinErrLogFilePath())
	if err != nil {
		logger.Errorln("could not create gin error log file:", err.Error())
	} else {
		gin.DefaultErrorWriter = io.MultiWriter(os.Stderr, ginErrLogFile)
	}

	return gin.Logger()
}

func CORSMiddleware() gin.HandlerFunc {
	origins := []string{"*"}
	return cors.New(cors.Config{
		AllowOrigins:     origins,
		AllowMethods:     []string{"OPTIONS", "GET", "POST", "PUT", "PATCH"},
		AllowHeaders:     []string{"Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With"},
		AllowCredentials: true,
		AllowWildcard:    true,
		MaxAge:           12 * time.Hour,
	})
}

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if len(authHeader) < len(BearerSchema)+1 {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "Access denied!",
			})
			return
		}

		tokenString := authHeader[len(BearerSchema)+1:]
		token, claims, err := auth.VerifyJWT(tokenString, config.GetJWTSigningKey())
		if err == nil && token.Valid {
			c.Set("claims", claims)

		} else {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": errors.Wrap(err, "verify token"),
			})
			return
		}
	}
}
