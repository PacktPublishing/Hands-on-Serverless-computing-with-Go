package chatsess

import "testing"

func TestPass(t *testing.T) {
	h := NewPassword("hello")

	if !CheckPassword("hello", h) {
		t.Errorf("hello no match")
	}

	if CheckPassword("goodbye", h) {
		t.Errorf("goodbye matches")
	}
}
