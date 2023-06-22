package main

import "testing"

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
