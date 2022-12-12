package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func TestMain_isPrime(t *testing.T) {
	tests := []struct {
		name     string
		num      int
		expected bool
		msg      string
	}{
		{name: "prime", num: 7, expected: true, msg: "7 is a prime number."},
		{name: "not prime", num: 8, expected: false, msg: "8 is not a prime number because it is divisible by 2."},
		{name: "zero", num: 0, expected: false, msg: "0 is not prime, by definition."},
		{name: "one", num: 1, expected: false, msg: "1 is not prime, by definition."},
		{name: "negative", num: -11, expected: false, msg: "Negative numbers are not prime, by definition."},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			res, msg := isPrime(tt.num)
			if res != tt.expected {
				t.Fatal("result and expected are not equal:", res, tt.expected)
			}
			if msg != tt.msg {
				t.Fatal("messages are not equal:", msg, tt.msg)
			}
		})
	}
}

func TestMain_prompt(t *testing.T) {
	oldOut := os.Stdout
	r, w, _ := os.Pipe()

	os.Stdout = w
	prompt()
	_ = w.Close()
	os.Stdout = oldOut

	out, _ := io.ReadAll(r)
	if string(out) != "-> " {
		t.Error("wrong prompt:", string(out))
	}
}

func TestMain_intro(t *testing.T) {
	oldOut := os.Stdout
	r, w, _ := os.Pipe()

	os.Stdout = w
	intro()
	_ = w.Close()
	os.Stdout = oldOut

	out, _ := io.ReadAll(r)
	if !strings.Contains(string(out), "Enter a whole number") {
		t.Error("wrong intro text:", string(out))
	}
}
