package main

import (
	"os"
	"testing"
)

func Test_isPrime(t *testing.T) {
	result, msg := isPrime(0)
	if result {
		t.Errorf("with %d as test parameter, got true, but expected false", 0)
	}

	if msg != "0 is not prime, by definition!" {
		t.Error("wrong message returned:", msg)
	}
}

func Test_isPrimeTableTests(t *testing.T) {
	primeTests := []struct {
		name     string
		testNum  int
		expected bool
		msg      string
	}{
		{"prime", 7, true, "7 is a prime number!"},
		{"not prime", 8, false, "8 is not a prime number because it is divisible by 2"},
	}

	for _, tc := range primeTests {
		result, msg := isPrime(tc.testNum)
		if result != tc.expected {
			t.Errorf("%s: expected %t got %t", tc.name, tc.expected, result)
		}

		if msg != tc.msg {
			t.Errorf("%s: expected %s got %s", tc.name, tc.msg, msg)
		}
	}
}



func Test_prompt(t *testing.T) {

	// save a copy of os.Stdout
	oldOut := os.Stdout

	// create a read and write pipe
	r,w,_ := os.Pipe()

	// set os.Stdout to our write pipe
	os.Stdout = w
	
}