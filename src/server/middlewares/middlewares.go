package middlewares

import (
	"io"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/valiant1012/transaction-service/src/config"
)

func GinLoggerMiddleware() gin.HandlerFunc {
	ginLogFile, err := os.Create(config.GetGinLogFilePath())
	if err != nil {
		// todo
	}
	ginErrLogFile, err := os.Create(config.GetGinErrLogFilePath())
	if err != nil {
		// todo
	}

	gin.DefaultWriter = io.MultiWriter(os.Stdout, ginLogFile)
	gin.DefaultErrorWriter = io.MultiWriter(os.Stderr, ginErrLogFile)
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
