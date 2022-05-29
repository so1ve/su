package su

import (
	"testing"
)

func TestIf(t *testing.T) {
	if If(true, 1, 2) != 1 {
		t.Error("Expected to be 1")
	}
	if If(false, 1, 2) != 2 {
		t.Error("Expected to be 2")
	}
}
