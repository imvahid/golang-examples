package main

import "testing"

/*
 * Useful Commands
 * 01. go test .
 * 02. go test ./...
 * 03. go test -bench=.
 * 04. go test -v -bench=.
 * 05. go test -cover .
 * 06. go test -cover ./...
 * 07. go test -count=10 .
 * 08. go test -count=10 ./...
 */

func TestAbs(t *testing.T) {
	o := Abs(-1)

	if o != 1 {
		t.Error("wrong number")
	}
}

func Abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func BenchmarkAbs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Abs(-1)
	}
}
