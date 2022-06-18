package su_slice

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

func TestIndexOf(t *testing.T) {
	slice := []int{1, 3, 4, 5}
	if !(IndexOf(slice, 1) == 0) {
		t.Error("Expected to found 1")
	}
	if !(IndexOf(slice, 6) == -1) {
		t.Error("Expected not to found 6")
	}
}

func TestIndexOfReflect(t *testing.T) {
	slice := [][]int{{2, 3, 4}, {1, 2, 3}}
	if !(IndexOfReflect(slice, []int{2, 3, 4}) == 0) {
		t.Error("Expected to found []int{2, 3, 4}")
	}
	if !(IndexOfReflect(slice, []int{2, 3, 1}) == -1) {
		t.Error("Expected not to found []int{2, 3, 1}")
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

func TestIncludesReflect(t *testing.T) {
	slice := [][]int{{2, 3, 4}, {1, 2, 3}}
	if !IncludesReflect(slice, []int{2, 3, 4}) {
		t.Error("Expected to found []int{2, 3, 4}")
	}
	if IncludesReflect(slice, []int{2, 3, 1}) {
		t.Error("Expected not to found []int{2, 3, 1}")
	}
}

func TestEvery(t *testing.T) {
	slice := []int{1, 3, 4, 5}
	if !Every(slice, func(item int, _ int) bool { return item > 0 }) {
		t.Error("Expected to be every")
	}
}

func TestSome(t *testing.T) {
	slice := []int{1, 3, 4, 5}
	if !Some(slice, func(item int, _ int) bool { return item > 3 }) {
		t.Error("Expected to be some")
	}
}
