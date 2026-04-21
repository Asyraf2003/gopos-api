package config

import (
	"os"
	"testing"
)

func TestLoad_DebugDisabledByDefault(t *testing.T) {
	t.Setenv("DATABASE_URL", "postgres://example")
	t.Setenv("AUTH_DEBUG_ENABLED", "")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	if cfg.Auth.Debug.Enabled {
		t.Fatal("Auth.Debug.Enabled = true, want false")
	}
}

func TestLoad_DebugEnabledFromEnv(t *testing.T) {
	t.Setenv("DATABASE_URL", "postgres://example")
	t.Setenv("AUTH_DEBUG_ENABLED", "true")

	cfg, err := Load()
	if err != nil {
		t.Fatalf("Load() error = %v", err)
	}

	if !cfg.Auth.Debug.Enabled {
		t.Fatal("Auth.Debug.Enabled = false, want true")
	}
}

func TestLoad_InvalidDebugBool(t *testing.T) {
	t.Setenv("DATABASE_URL", "postgres://example")
	t.Setenv("AUTH_DEBUG_ENABLED", "not-bool")

	_, err := Load()
	if err == nil {
		t.Fatal("Load() error = nil, want error")
	}
}

func TestMain(m *testing.M) {
	_ = os.Unsetenv("AUTH_DEBUG_ENABLED")
	_ = os.Unsetenv("DATABASE_URL")
	os.Exit(m.Run())
}
