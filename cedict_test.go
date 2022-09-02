package cedict

import "testing"

func TestBytes(t *testing.T) {
	if len(Bytes) < 1000 {
		t.Error("too few bytes")
	}
}
