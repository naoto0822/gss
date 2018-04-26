package modules

import (
	"testing"
)

func TestNewDecoder(t *testing.T) {
	d := NewDecoder()
	if d == nil {
		t.Error("TestNewDecoder: not expected variable")
	}
}

// TODO: test code !
