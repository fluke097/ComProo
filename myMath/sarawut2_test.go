package main

import "testing"

func TestSarawut2(t *testing.T) {
	result := Sarawut2(7, 3)

	if result != 10 {
		t.Error("Expected 6, got : ", result)
	}
}
