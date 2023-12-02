package helper

import "testing"

// run benchmark dari package
// go test -v -run=NoUnitTest -bench=.
// run benchmark dari root
// go test -v -run=NoUnitTest -bench=. ./...

func BenchmarkHelloWorldAditya(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Aditya")
	}
}

func BenchmarkHelloWorldRicki(b *testing.B) {
	for i := 0; i < b.N; i++ {
		HelloWorld("Ricki")
	}
}
