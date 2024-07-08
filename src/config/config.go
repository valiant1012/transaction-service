package config

import (
	"cmp"

	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var (
	env Env
)

type Env struct {
	Type          string       `mapstructure:"type"`
	Version       string       `mapstructure:"version"`
	Port          string       `mapstructure:"port"`
	JWTSigningKey string       `mapstructure:"jwt_signing_key"`
	LogFiles      LogFilePaths `mapstructure:"log_files"`
	Database      DBConfig     `mapstructure:"database"`
}

type LogFilePaths struct {
	GinLogger    string `mapstructure:"gin_standard"`
	GinErrLogger string `mapstructure:"gin_err_logger"`
	ServerLogger string `mapstructure:"server_logger"`
}

type DBConfig struct {
	Postgres Postgres `mapstructure:"postgres"`
}

type Postgres struct {
	DBName   string `mapstructure:"db_name"`
	Host     string `mapstructure:"host"`
	Port     string `mapstructure:"port"`
	Username string `mapstructure:"username"`
	Password string `mapstructure:"password"`
	SSLMode  string `mapstructure:"sslmode"`
}

func Init(configLocation string) {
	viper.SetConfigFile(configLocation)
	viper.SetConfigType("json")
	err := viper.ReadInConfig()
	if err != nil {
		panic(errors.Wrap(err, "could not read config file"))
	}
	err = viper.Unmarshal(&env)
	if err != nil {
		panic(errors.Wrap(err, "could not unmarshal env file"))
	}
}

func GetVar(v string) any {
	return viper.Get(v)
}

func GetEnvType() string {
	return env.Type
}

func GetVersion() string {
	return env.Version
}

func GetPort() string {
	return env.Port
}

func GetJWTSigningKey() string {
	return env.JWTSigningKey
}

func GetGinLogFilePath() string {
	return cmp.Or(env.LogFiles.GinLogger, "logs/gin.log")
}

func GetGinErrLogFilePath() string {
	return cmp.Or(env.LogFiles.GinErrLogger, "logs/gin.error.log")
}

func GetServerLogFilePath() string {
	return cmp.Or(env.LogFiles.ServerLogger, "logs/server.log")
}

func GetPostgresConfig() Postgres {
	return env.Database.Postgres
}
