package su

import (
	"testing"
)

func TestMap(t *testing.T) {
	orig := []int{1, 3, 4, 5}
	mapped := Map(orig, func(item int) int { return item + 1 })
	if mapped[0] != 2 || mapped[1] != 4 || mapped[2] != 5 || mapped[3] != 6 {
		t.Error("Expected to be mapped")
	}
}

func TestIncludes(t *testing.T) {
	slice := []int{1, 3, 4, 5}
	if !Includes(slice, 1) {
		t.Error("Expected to found 1")
	}
	if Includes(slice, 2) {
		t.Error("Expected not to found 2")
	}
}

func TestIndexOf(t *testing.T) {
	slice := []int{1, 3, 4, 5}
	if !(IndexOf(slice, 1) == 0) {
		t.Error("Expected to found index")
	}
}
