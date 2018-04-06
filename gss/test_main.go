package gss

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// TODO: setup
	code := m.Run()
	defer os.Exit(code)
	// TODO: teardown
}
