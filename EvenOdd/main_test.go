package main

import "testing"

func GenSliceTest(t *testing.T) {
	n, f := genSlice()
	if n%2 == 0 && f == 1 {
		t.Errorf("Expected Even but got Odd")
	}
}
