package main

import "testing"

func TestAdd(t *testing.T) {
	result := Add(2, 2)
	if result != 4 {
		t.Errorf("Ожидалось 4, но получили %d", result)
	}
}
