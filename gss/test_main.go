package gss

import (
	"os"
	"testing"
)

// TestMain test setup and teardown
func TestMain(m *testing.M) {
	// TODO: setup
	code := m.Run()
	defer os.Exit(code)
	// TODO: teardown
}
