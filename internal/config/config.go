// audit:allow-oversize reason=config-entrypoint
package config

import (
	"errors"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv      string
	HTTPPort    string
	DatabaseURL string
	Auth        AuthConfig
}

type AuthConfig struct {
	Google     GoogleConfig
	JWT        JWTConfig
	Debug      DebugConfig
	StateTTL   time.Duration
	SessionTTL time.Duration
}

type DebugConfig struct {
	Enabled bool
}

type GoogleConfig struct {
	Issuer       string
	ClientID     string
	ClientSecret string
	RedirectURL  string
}

func (c GoogleConfig) IsConfigured() bool {
	return c.ClientID != "" && c.ClientSecret != "" && c.RedirectURL != ""
}

type JWTConfig struct {
	Issuer string
	Aud    string
	Kid    string
	Secret string
	TTL    time.Duration
}

func Load() (Config, error) {
	_ = godotenv.Load()

	jwtTTLMinutes, err := getEnvInt("AUTH_JWT_TTL_MINUTES", 15)
	if err != nil {
		return Config{}, err
	}

	authStateTTLMinutes, err := getEnvInt("AUTH_STATE_TTL_MINUTES", 10)
	if err != nil {
		return Config{}, err
	}

	authSessionTTLHours, err := getEnvInt("AUTH_SESSION_TTL_HOURS", 720)
	if err != nil {
		return Config{}, err
	}

	authDebugEnabled, err := getEnvBool("AUTH_DEBUG_ENABLED", false)
	if err != nil {
		return Config{}, err
	}

	cfg := Config{
		AppEnv:      getEnv("APP_ENV", "local"),
		HTTPPort:    getEnv("HTTP_PORT", "8080"),
		DatabaseURL: os.Getenv("DATABASE_URL"),
		Auth: AuthConfig{
			Google: GoogleConfig{
				Issuer:       getEnv("AUTH_GOOGLE_ISSUER", "https://accounts.google.com"),
				ClientID:     os.Getenv("AUTH_GOOGLE_CLIENT_ID"),
				ClientSecret: os.Getenv("AUTH_GOOGLE_CLIENT_SECRET"),
				RedirectURL:  os.Getenv("AUTH_GOOGLE_REDIRECT_URL"),
			},
			JWT: JWTConfig{
				Issuer: getEnv("AUTH_JWT_ISSUER", "pos-go"),
				Aud:    getEnv("AUTH_JWT_AUDIENCE", "pos-go-client"),
				Kid:    getEnv("AUTH_JWT_KID", "local-dev-key"),
				Secret: os.Getenv("AUTH_JWT_SECRET"),
				TTL:    time.Duration(jwtTTLMinutes) * time.Minute,
			},
			Debug: DebugConfig{
				Enabled: authDebugEnabled,
			},
			StateTTL:   time.Duration(authStateTTLMinutes) * time.Minute,
			SessionTTL: time.Duration(authSessionTTLHours) * time.Hour,
		},
	}

	if cfg.DatabaseURL == "" {
		return Config{}, errors.New("DATABASE_URL is required")
	}

	return cfg, nil
}

func (c Config) HTTPAddr() string {
	return ":" + c.HTTPPort
}

func getEnv(key string, fallback string) string {
	value := os.Getenv(key)
	if value == "" {
		return fallback
	}

	return value
}

func getEnvInt(key string, fallback int) (int, error) {
	value := os.Getenv(key)
	if value == "" {
		return fallback, nil
	}

	n, err := strconv.Atoi(value)
	if err != nil {
		return 0, errors.New(key + " must be a valid integer")
	}

	return n, nil
}

func getEnvBool(key string, fallback bool) (bool, error) {
	value := strings.TrimSpace(os.Getenv(key))
	if value == "" {
		return fallback, nil
	}

	switch strings.ToLower(value) {
	case "1", "true", "yes", "on":
		return true, nil
	case "0", "false", "no", "off":
		return false, nil
	default:
		return false, errors.New(key + " must be a valid boolean")
	}
}
