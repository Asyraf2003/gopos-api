package http

import (
	"os"
	"strings"
	"testing"
)

func TestCapabilityHandler_DoesNotImportPostgresAdapter(t *testing.T) {
	files := []string{
		"capability_handler.go",
		"capability_handler_read.go",
		"capability_handler_write.go",
		"capability_handler_response.go",
	}

	for _, file := range files {
		source, err := os.ReadFile(file)
		if err != nil {
			t.Fatalf("ReadFile(%s) error = %v", file, err)
		}
		if strings.Contains(string(source), "internal/platform/postgres") {
			t.Fatalf("%s imports PostgreSQL adapter", file)
		}
	}
}
