package main

import "testing"

func TestEvenOrOddString(t *testing.T) {
	res, e := evenOrOddString(2), "even\nodd\neven"
	if res != e {
		t.Errorf("Expected even odd even, got %v", res)
	}
}
