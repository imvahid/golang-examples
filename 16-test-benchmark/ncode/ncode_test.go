package ncode

import (
	"testing"
	"time"
)

func TestIsValid(t *testing.T) {
	_, err := NewNationalID("3356947567")

	if err != nil {
		t.Error(err)
	}
}

func BenchmarkIsValid(b *testing.B) {
	time.Sleep(1 * time.Second)

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		_, err := NewNationalID("3356947567")

		if err != nil {
			b.Error(err)
		}
	}
}
