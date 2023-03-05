package main

import "testing"

func Test_isPrime(t *testing.T) {
	// result, msg := isPrime(0)
	// if result {
	// 	t.Errorf("with %d as test parameter, got true, but expected false", 0)
	// }

	// if msg != "0 is not prime!" {
	// 	t.Error("wrong message returned", msg)
	// }

	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is a prime!"},
		{"not prime", 8, false, "8 is not prime!"},
	}

	for _, e := range primeTests {
		result, msg := isPrime(e.testNum)
		if e.expected && !result {
			t.Errorf("with %d as test parameter, got false, but expected true", e.testNum)
		}

		if !e.expected && result {
			t.Errorf("with %d as test parameter, got true, but expected false", e.testNum)
		}

		if e.msg != msg {
			t.Errorf("%s: Expected %s but got %s", e.name, e.msg, msg)
		}
	}

}
