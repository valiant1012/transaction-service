package postgres

import (
	"context"
	"fmt"
	"time"

	"github.com/pkg/errors"
	"github.com/valiant1012/transaction-service/src/config"
	"github.com/valiant1012/transaction-service/src/utility/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() error {
	var err error
	postgresConfig := config.GetPostgresConfig()

	// postgres://user:password@host:port/db_name?sslmode=disable
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
		postgresConfig.Username,
		postgresConfig.Password,
		postgresConfig.Host,
		postgresConfig.Port,
		postgresConfig.DBName,
		postgresConfig.SSLMode,
	)

	// Build connection
	db, err = gorm.Open(postgres.Open(connStr))
	if err != nil {
		return errors.Wrap(err, "open connection")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err = AutoMigrate(ctx)
	if err != nil {
		logger.Errorln("Auto Migrate Error:", err.Error())
		// can panic if needed
	}

	return nil
}

func AutoMigrate(ctx context.Context) error {
	var err error
	err = MigrateTransactions(ctx)
	if err != nil {
		return errors.Wrap(err, "auto-migrate")
	}
	return nil
}

func DB() *gorm.DB {
	return db
}
