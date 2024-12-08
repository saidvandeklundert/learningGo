package main

import (
	"bufio"
	"io"
	"os"
	"strings"
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
	r, w, _ := os.Pipe()

	// set os.Stdout to our write pipe
	os.Stdout = w

	prompt()

	// close our writer
	_ = w.Close()

	// reset os.Stdout
	os.Stdout = oldOut

	// read output of out prompt function
	out, _ := io.ReadAll(r)

	// perform the test
	if string(out) != "-> " {
		t.Errorf("incorrect prompt: expected '-> ' but got %s", string(out))
	}

}

func Test_intro(t *testing.T) {

	oldOut := os.Stdout

	r, w, _ := os.Pipe()

	os.Stdout = w

	intro()

	_ = w.Close()

	os.Stdout = oldOut

	out, _ := io.ReadAll(r)

	if !strings.Contains(string(out), "Enter a whole number") {
		t.Errorf("Expected 'Enter a whole number' to be present in the 'intro()'")
	}
}

// simulate user input and then test checkNumbers
func Test_checkNumbers(t *testing.T) {
	input := strings.NewReader("7")
	reader := bufio.NewScanner(input)
	res, _ := checkNumbers(reader)
	if !strings.EqualFold(res, "7 is a prime number!") {
		t.Error("incorrect value returend, got:", res)
	}
}

func Test_checkNumbers_table_test(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{name: "empty", input: "", expected: "Please enter a whole number!"},
		{name: "zeroy", input: "0", expected: "0 is not prime, by definition!"},
	}

	for _, test := range tests {
		input := strings.NewReader(test.input)
		reader := bufio.NewScanner(input)
		res, _ := checkNumbers(reader)
		if !strings.EqualFold(res, test.expected) {
			t.Errorf("%s expected %s but got: %s", test.name, test.expected, res)
		}
	}

}
