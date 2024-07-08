package main

import (
	"flag"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/valiant1012/transaction-service/src/config"
	"github.com/valiant1012/transaction-service/src/constants"
	"github.com/valiant1012/transaction-service/src/models/postgres"
	"github.com/valiant1012/transaction-service/src/server/middlewares"
	"github.com/valiant1012/transaction-service/src/server/router"
	"github.com/valiant1012/transaction-service/src/utility/logger"
)

func main() {
	// Initialize and parse command line flags
	configLocation := flag.String("config", "cloud/env.json", "Config file location")
	flag.Parse()

	// Initialize config
	config.Init(*configLocation)

	// Set release mode if production
	if config.GetEnvType() == constants.EnvProduction {
		gin.SetMode(gin.ReleaseMode)
	}

	// Init Logger
	err := logger.Init(config.GetServerLogFilePath())
	if err != nil {
		panic(errors.Wrap(err, "init logger"))
	}

	// Init Services
	err = initServices()
	if err != nil {
		panic(errors.Wrap(err, "init services"))
	}

	// Initialize gin engine & attach middlewares
	engine := gin.New()
	engine.SetTrustedProxies(nil)
	engine.Use(gin.Recovery())
	engine.Use(middlewares.GinLoggerMiddleware())
	engine.Use(middlewares.CORSMiddleware())

	// Add Routes
	router.AddRoutes(engine)

	// Run gin engine
	if err = engine.Run(config.GetPort()); err != nil {
		panic(err)
	}
}

func initServices() error {
	err := postgres.Connect()
	if err != nil {
		return errors.Wrap(err, "connect postgres")
	}

	return nil
}
