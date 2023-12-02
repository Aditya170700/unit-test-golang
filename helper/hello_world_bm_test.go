package helper

import "testing"

// run benchmark dari package
// go test -v -run=NoUnitTest -bench=.
// run benchmark dari root
// go test -v -run=NoUnitTest -bench=. ./...

func BenchmarkHelloWorldTable(b *testing.B) {
	payloads := []struct {
		name    string
		request string
	}{
		{
			name:    "Aditya",
			request: "Aditya",
		},
		{
			name:    "Ricki",
			request: "Ricki",
		},
		{
			name:    "Julianto",
			request: "Julianto",
		},
		{
			name:    "Aditya Ricki Julianto",
			request: "Aditya Ricki Julianto",
		},
	}

	for _, payload := range payloads {
		b.Run(payload.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				HelloWorld(payload.request)
			}
		})
	}
}

func BenchmarkHelloWorldSub(b *testing.B) {
	b.Run("Aditya", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Aditya")
		}
	})
	b.Run("Ricki", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			HelloWorld("Ricki")
		}
	})
}

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
