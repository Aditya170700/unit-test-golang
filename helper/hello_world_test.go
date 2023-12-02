package helper

import (
	"fmt"
	"testing"
)

func TestHelloWorldAditya(t *testing.T) {
	result := HelloWorld("Aditya")

	if result != "Hello Aditya" {
		t.Error("Result must be '" + result + "'")
	}

	fmt.Println("DONE")
}

func TestHelloWorldRicki(t *testing.T) {
	result := HelloWorld("Ricki")

	if result != "Hello Ricki" {
		t.Fatal("Result must be '" + result + "'")
	}

	fmt.Println("DONE")
}
