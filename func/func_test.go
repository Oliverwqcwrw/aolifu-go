package main

import "testing"

func TestSum(t *testing.T) {
	a := 10
	b := 20
	result := sum(a, b)
	if result == 30 {
		t.Log("success")
	} else {
		t.Fatalf("failed result=%v", result)
	}
}
