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
	Type     string       `json:"type"`
	Version  string       `json:"version"`
	Port     string       `json:"port"`
	LogFiles LogFilePaths `json:"log_files"`
}

type LogFilePaths struct {
	GinLogger    string `json:"gin_standard"`
	GinErrLogger string `json:"gin_err_logger"`
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

func GetEnvType() string {
	return env.Type
}

func GetVersion() string {
	return env.Version
}

func GetPort() string {
	return env.Port
}

func GetGinLogFilePath() string {
	return cmp.Or(env.LogFiles.GinLogger, "logs/gin.log")
}

func GetGinErrLogFilePath() string {
	return cmp.Or(env.LogFiles.GinErrLogger, "logs/gin.error.log")
}
