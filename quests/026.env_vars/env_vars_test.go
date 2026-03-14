package env_vars

import (
	"os"
	"testing"
)

func setEnv(t *testing.T, key, val string) {
	t.Helper()
	prev, ok := os.LookupEnv(key)
	if err := os.Setenv(key, val); err != nil {
		t.Fatalf("cannot set env %s: %v", key, err)
	}
	t.Cleanup(func() {
		if ok {
			os.Setenv(key, prev)
		} else {
			os.Unsetenv(key)
		}
	})
}

func unsetEnv(t *testing.T, key string) {
	t.Helper()
	prev, ok := os.LookupEnv(key)
	if ok {
		if err := os.Unsetenv(key); err != nil {
			t.Fatalf("cannot unset env %s: %v", key, err)
		}
		t.Cleanup(func() {
			os.Setenv(key, prev)
		})
	}
}

func TestLoadConfig(t *testing.T) {
	t.Run("success_all_fields", func(t *testing.T) {
		setEnv(t, "DB_HOST", "db.example.com")
		setEnv(t, "DB_PORT", "9090")
		setEnv(t, "DEBUG_MODE", "true")

		cfg, err := LoadConfig()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if cfg.DBHost != "db.example.com" {
			t.Errorf("expected DBHost %q, got %q", "db.example.com", cfg.DBHost)
		}
		if cfg.DBPort != 9090 {
			t.Errorf("expected DBPort %d, got %d", 9090, cfg.DBPort)
		}
		if !cfg.DebugMode {
			t.Errorf("expected DebugMode true, got false")
		}
	})

	t.Run("success_defaults", func(t *testing.T) {
		setEnv(t, "DB_HOST", "localhost")
		unsetEnv(t, "DB_PORT")
		unsetEnv(t, "DEBUG_MODE")

		cfg, err := LoadConfig()
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if cfg.DBHost != "localhost" {
			t.Errorf("expected DBHost %q, got %q", "localhost", cfg.DBHost)
		}
		if cfg.DBPort != 5432 {
			t.Errorf("expected default DBPort 5432, got %d", cfg.DBPort)
		}
		if cfg.DebugMode {
			t.Errorf("expected default DebugMode false, got true")
		}
	})

	t.Run("fail_missing_host", func(t *testing.T) {
		unsetEnv(t, "DB_HOST")
		_, err := LoadConfig()
		if err == nil {
			t.Error("expected error for missing DB_HOST, got nil")
		}
	})

	t.Run("fail_invalid_port", func(t *testing.T) {
		setEnv(t, "DB_HOST", "localhost")
		setEnv(t, "DB_PORT", "invalid-port")

		_, err := LoadConfig()
		if err == nil {
			t.Error("expected error for invalid DB_PORT, got nil")
		}
	})

	t.Run("fail_invalid_bool", func(t *testing.T) {
		setEnv(t, "DB_HOST", "localhost")
		setEnv(t, "DEBUG_MODE", "not-a-bool")

		_, err := LoadConfig()
		if err == nil {
			t.Error("expected error for invalid DEBUG_MODE, got nil")
		}
	})
}

func TestMain(m *testing.M) {
	code := m.Run()
	if code == 0 {
		colorReset := "\033[0m"
		colorGreen := "\033[32m"
		println(colorGreen, "Success! Completed the env_vars Quest 🎉", colorReset)
	}
	os.Exit(code)
}
