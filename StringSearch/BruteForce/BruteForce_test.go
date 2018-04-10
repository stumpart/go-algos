package BruteForce

import (
	"testing"
)

func TestBruteForce(t *testing.T) {
	v := "Me Name Barry Henry"
	res := brute("Barry", v)

	if res != 8 {
		t.Error("Expected 8, got", res)
	}
}

func TestBruteForceNotFound(t *testing.T) {
	v := "Me Name Carl Henry"
	res := brute("Barry", v)

	if res != -1 {
		t.Error("Expected -1. got", res)
	}
}
