package main

import "testing"

func TestMain_isPrime(t *testing.T) {
	res, msg := isPrime(0)
	if res {
		t.Error("expected false, but got true", 0)
	}
	if msg != "0 is not prime, by definition." {
		t.Error("wrong message", msg)
	}

	res, msg = isPrime(7)
	if !res {
		t.Error("expected true, but got false", 7)
	}
	if msg != "7 is a prime number." {
		t.Error("wrong message", msg)
	}
}
