package xmlp

import (
	"bytes"
	"testing"
)

func TestNewDecoder(t *testing.T) {
	data := `<xml></xml>`
	b := []byte(data)
	r := bytes.NewReader(b)
	decoder := NewDecoder(r)
	if decoder == nil {
		t.Error("TestNewDecoder failed factory Decoder")
	}
}

// TODO: test code!
