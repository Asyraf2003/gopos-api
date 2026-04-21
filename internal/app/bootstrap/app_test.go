package bootstrap

import (
	"context"
	"strings"
	"testing"

	"pos-go/internal/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

func TestNew_RegistersHealthRouteAlways(t *testing.T) {
	cfg := config.Config{
		AppEnv:      "test",
		HTTPPort:    "8080",
		DatabaseURL: testDatabaseURL(t),
		Auth: config.AuthConfig{
			Google: config.GoogleConfig{},
		},
	}

	app, err := New(context.Background(), cfg)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	defer app.DB.Close()

	if !hasRoute(app, "GET", "/api/health") {
		t.Fatal("expected GET /api/health to be registered")
	}

	if hasRoute(app, "GET", "/api/auth/google/start") {
		t.Fatal("did not expect auth route when google config is incomplete")
	}
}

func TestNew_RegistersGoogleAuthRoutesWhenConfigured(t *testing.T) {
	cfg := config.Config{
		AppEnv:      "test",
		HTTPPort:    "8080",
		DatabaseURL: testDatabaseURL(t),
		Auth: config.AuthConfig{
			Google: config.GoogleConfig{
				Issuer:       "https://accounts.google.com",
				ClientID:     "client-id",
				ClientSecret: "client-secret",
				RedirectURL:  "http://127.0.0.1:8081/api/auth/google/callback",
			},
			JWT: config.JWTConfig{
				Issuer: "pos-go",
				Aud:    "pos-go-client",
				Kid:    "local-dev-key",
				Secret: "test-secret-123",
				TTL:    15,
			},
			StateTTL:   10,
			SessionTTL: 720,
		},
	}

	app, err := New(context.Background(), cfg)
	if err != nil {
		t.Fatalf("New() error = %v", err)
	}
	defer app.DB.Close()

	if !hasRoute(app, "GET", "/api/auth/google/start") {
		t.Fatal("expected GET /api/auth/google/start to be registered")
	}

	if !hasRoute(app, "GET", "/api/auth/google/callback") {
		t.Fatal("expected GET /api/auth/google/callback to be registered")
	}
}

func hasRoute(app *App, method, path string) bool {
	for _, route := range app.Echo.Routes() {
		if route.Method == method && route.Path == path {
			return true
		}
	}
	return false
}

func testDatabaseURL(t *testing.T) string {
	t.Helper()

	cfg, err := pgxpool.ParseConfig("postgres://posgo_app:posgo_local_dev_123@127.0.0.1:5432/posgo_app_db?sslmode=disable")
	if err != nil {
		t.Fatalf("ParseConfig() error = %v", err)
	}

	if !strings.Contains(cfg.ConnString(), "posgo_app_db") {
		t.Fatal("expected test database url to target posgo_app_db")
	}

	return cfg.ConnString()
}
