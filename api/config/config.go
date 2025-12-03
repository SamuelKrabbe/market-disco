package config

import (
	"bufio"
	"errors"
	"os"
	"strings"
	"time"
)

// App config struct
type Config struct {
	Server   ServerConfig
	Postgres PostgresConfig
	Logger   Logger
}

// Server config struct
type ServerConfig struct {
	AppVersion        string
	Port              string
	PprofPort         string
	Mode              string
	JwtSecretKey      string
	CookieName        string
	ReadTimeout       time.Duration
	WriteTimeout      time.Duration
	SSL               bool
	CtxDefaultTimeout time.Duration
	CSRF              bool
	Debug             bool
}

// Logger config
type Logger struct {
	Development       bool
	DisableCaller     bool
	DisableStacktrace bool
	Encoding          string
	Level             string
}

// Postgresql config
type PostgresConfig struct {
	PostgresqlHost     string
	PostgresqlPort     string
	PostgresqlUser     string
	PostgresqlPassword string
	PostgresqlDbname   string
	PostgresqlSSLMode  bool
	PostgresqlDriver   string
}

// Load config file from given path
func LoadConfig(path string) error {
	f, err := os.Open(path)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}

		parts := strings.SplitN(line, "=", 2)
		if len(parts) != 2 {
			return errors.New("invalid .env line: " + line)
		}

		key := strings.TrimSpace(parts[0])
		value := strings.TrimSpace(parts[1])

		err := os.Setenv(key, value)
		if err != nil {
			return err
		}
	}
	return scanner.Err()
}

// Parse config file
func ParseConfig() (*Config, error) {
	readTimeout, err := time.ParseDuration(os.Getenv("SERVER_READ_TIMEOUT"))
	if err != nil {
		return nil, err
	}

	writeTimeout, err := time.ParseDuration(os.Getenv("SERVER_WRITE_TIMEOUT"))
	if err != nil {
		return nil, err
	}

	ctxTimeout, err := time.ParseDuration(os.Getenv("SERVER_CTX_TIMEOUT"))
	if err != nil {
		return nil, err
	}

	ssl := os.Getenv("SERVER_SSL") == "true"
	csrf := os.Getenv("SERVER_CSRF") == "true"
	debug := os.Getenv("SERVER_DEBUG") == "true"

	postgresSSL := os.Getenv("POSTGRES_SSL") == "true"

	cfg := &Config{
		Server: ServerConfig{
			AppVersion:        os.Getenv("SERVER_APP_VERSION"),
			Port:              os.Getenv("SERVER_PORT"),
			PprofPort:         os.Getenv("SERVER_PPROF_PORT"),
			Mode:              os.Getenv("SERVER_MODE"),
			JwtSecretKey:      os.Getenv("SERVER_JWT_SECRET"),
			CookieName:        os.Getenv("SERVER_COOKIE_NAME"),
			ReadTimeout:       readTimeout,
			WriteTimeout:      writeTimeout,
			SSL:               ssl,
			CtxDefaultTimeout: ctxTimeout,
			CSRF:              csrf,
			Debug:             debug,
		},
		Postgres: PostgresConfig{
			PostgresqlHost:     os.Getenv("POSTGRES_HOST"),
			PostgresqlPort:     os.Getenv("POSTGRES_PORT"),
			PostgresqlUser:     os.Getenv("POSTGRES_USER"),
			PostgresqlPassword: os.Getenv("POSTGRES_PASSWORD"),
			PostgresqlDbname:   os.Getenv("POSTGRES_DB"),
			PostgresqlSSLMode:  postgresSSL,
			PostgresqlDriver:   os.Getenv("POSTGRES_DRIVER"),
		},
	}

	return cfg, nil
}
