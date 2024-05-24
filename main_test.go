package main

import "testing"

func TestAdd(t *testing.T) {
	result := add(3, 4)
	expected := 7
	if result != expected {
		t.Errorf("Addition failed: got %d, expected %d", result, expected)
	}
}
