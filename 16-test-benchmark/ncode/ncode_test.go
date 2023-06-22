package ncode

import (
	"testing"
	"time"
)

/*
 * Useful Commands
 * 01. go test
 * 02. go test .
 * 03. go test -v
 * 04. go test -v .
 * 05. go test -v ./...
 * 06. go test -bench=.
 * 07. go test -v -bench=.
 * 08. go test -cover
 * 09. go test -cover .
 * 10. go test -cover ./...
 * 11. go test -count=10
 * 12. go test -count=10 .
 * 13. go test -count=10 ./...
 * 14. go test -trace trace.out OR go test -trace=trace.out
 * 15. go tool trace trace.out
 */

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
