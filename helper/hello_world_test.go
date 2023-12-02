package helper

import "testing"

func TestHelloWorld(t *testing.T) {
	result := HelloWorld("Aditya")

	if result != "Hello Aditya" {
		panic("Result is not Hello Aditya")
	}
}
