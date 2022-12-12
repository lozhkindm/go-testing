package main

import (
	"bufio"
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

func TestMain_checkNumbers(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
		quit     bool
	}{
		{name: "empty", input: "", expected: "Please enter a whole number."},
		{name: "zero", input: "0", expected: "0 is not prime, by definition."},
		{name: "one", input: "1", expected: "1 is not prime, by definition."},
		{name: "seven", input: "7", expected: "7 is a prime number."},
		{name: "negative", input: "-2", expected: "Negative numbers are not prime, by definition."},
		{name: "nan", input: "asd", expected: "Please enter a whole number."},
		{name: "decimal", input: "2.5", expected: "Please enter a whole number."},
		{name: "quit", input: "q", expected: "", quit: true},
		{name: "Quit", input: "Q", expected: "", quit: true},
		{name: "symbol", input: "©", expected: "Please enter a whole number."},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			scanner := bufio.NewScanner(strings.NewReader(tt.input))
			res, quit := checkNumbers(scanner)

			if !strings.EqualFold(res, tt.expected) {
				t.Error("wrong result:", res, tt.expected)
			}
			if quit != tt.quit {
				t.Error("wrong quit:", quit, tt.quit)
			}
		})
	}
}
