package config

import (
	"log"
	"strconv"
	"strings"

	"github.com/base-org/pessimism/internal/api/server"
	"github.com/base-org/pessimism/internal/api/service"

	"github.com/base-org/pessimism/internal/logging"
	"github.com/joho/godotenv"

	"os"
)

type FilePath string

type Env string

const (
	Development Env = "development"
	Production  Env = "production"
	Local       Env = "local"
)

// Config ... Application level configuration defined by `FilePath` value
// TODO - Consider renaming to "environment config"
type Config struct {
	Environment Env

	SvcConfig    *service.Config
	ServerConfig *server.Config

	LoggerConfig *logging.Config
}

// NewConfig ... Initializer
func NewConfig(fileName FilePath) *Config {
	if err := godotenv.Load(string(fileName)); err != nil {
		log.Fatalf("config file not found for file: %s", fileName)
	}

	config := &Config{

		Environment: Env(getEnvStr("ENV")),

		SvcConfig: &service.Config{
			L1RpcEndpoint:  getEnvStr("L1_RPC_ENDPOINT"),
			L2RpcEndpoint:  getEnvStr("L2_RPC_ENDPOINT"),
			L1PollInterval: getEnvInt("L1_POLL_INTERVAL"),
			L2PollInterval: getEnvInt("L2_POLL_INTERVAL"),
		},

		LoggerConfig: &logging.Config{
			UseCustom:         getEnvBool("LOGGER_USE_CUSTOM"),
			Level:             getEnvInt("LOGGER_LEVEL"),
			DisableCaller:     getEnvBool("LOGGER_DISABLE_CALLER"),
			DisableStacktrace: getEnvBool("LOGGER_DISABLE_STACKTRACE"),
			Encoding:          getEnvStr("LOGGER_ENCODING"),
			OutputPaths:       getEnvSlice("LOGGER_OUTPUT_PATHS"),
			ErrorOutputPaths:  getEnvSlice("LOGGER_ERROR_OUTPUT_PATHS"),
		},

		ServerConfig: &server.Config{
			Host:            getEnvStr("SERVER_HOST"),
			Port:            getEnvInt("SERVER_PORT"),
			ListenLimit:     getEnvInt("SERVER_LISTEN_LIMIT"),
			KeepAlive:       getEnvInt("SERVER_KEEP_ALIVE_TIME"),
			ReadTimeout:     getEnvInt("SERVER_READ_TIMEOUT"),
			WriteTimeout:    getEnvInt("SERVER_WRITE_TIMEOUT"),
			ShutdownTimeout: getEnvInt("SERVER_SHUTDOWN_TIME"),
		},
	}

	return config
}

// IsProduction ... Returns true if the env is production
func (cfg *Config) IsProduction() bool {
	return cfg.Environment == Production
}

// IsDevelopment ... Returns true if the env is development
func (cfg *Config) IsDevelopment() bool {
	return cfg.Environment == Development
}

// IsLocal ... Returns true if the env is local
func (cfg *Config) IsLocal() bool {
	return cfg.Environment == Local
}

// getEnvStr ... Reads env var from process environment, panics if not found
func getEnvStr(key string) string {
	envVar, ok := os.LookupEnv(key)

	// Not found
	if !ok {
		log.Fatalf("could not find env var given key: %s", key)
	}

	return envVar
}

// getEnvBool ... Reads env vars and converts to booleans
func getEnvBool(key string) bool {
	if val := getEnvStr(key); val == "1" {
		return true
	}
	return false
}

// getEnvSlice ... Reads env vars and converts to string slice
func getEnvSlice(key string) []string {
	return strings.Split(getEnvStr(key), ",")
}

// getEnvInt ... Reads env vars and converts to int
func getEnvInt(key string) int {
	val := getEnvStr(key)
	intRep, err := strconv.Atoi(val)
	if err != nil {
		log.Fatalf("env val is not int; got: %s=%s; err: %s", key, val, err.Error())
	}
	return intRep
}
