package main

import "testing"

func TestReverse(t *testing.T) {
	slice := []rune{'a', 'b', 'c', 'd', 'e'}
	reversed := reverse(slice)
	expected := []rune{'e', 'd', 'c', 'b', 'a'}
	if string(reversed) != string(expected) {
		t.Errorf("expected %s, got '%s'", string(expected), string(reversed))
	}
}
