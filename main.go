package main

import (
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/valiant1012/transaction-service/src/config"
	"github.com/valiant1012/transaction-service/src/server/middlewares"
	"github.com/valiant1012/transaction-service/src/server/routers"
)

func main() {
	// Initialize and parse command line flags
	configLocation := flag.String("config", "cloud/env.json", "Config file location")
	flag.Parse()

	// Initialize config
	config.Init(*configLocation)

	// Initialize gin engine & attach middlewares
	engine := gin.New()
	engine.SetTrustedProxies(nil)
	engine.Use(gin.Recovery())
	engine.Use(middlewares.GinLoggerMiddleware())
	engine.Use(middlewares.CORSMiddleware())

	// Add Routes
	routers.AddRoutes(engine)

	// Run gin engine
	if err := engine.Run(config.GetPort()); err != nil {
		panic(err)
	}
}
