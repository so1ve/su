package su_string

import "testing"

func TestSubstring(t *testing.T) {
	s := "abcdefg"
	if Substring(s, 1, 3) != "bc" {
		t.Error("Expected to be bc, got", Substring(s, 1, 3))
	}
	s = "嗨害嗨"
	if Substring(s, 1) != "害嗨" {
		t.Error("Expected to be 害嗨, got", Substring(s, 1))
	}
}

func TestCharCodeAt(t *testing.T) {
	s := "abcdefg"
	if CharCodeAt(s, 3) != 100 {
		t.Error("Expected to be bc, got", CharCodeAt(s, 3))
	}
}
