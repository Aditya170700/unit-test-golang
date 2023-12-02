package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Assert")
	assert.Equal(t, "Hello Assert", result, "Result must be '"+result+"'")
	fmt.Println("Executed")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Require")
	require.Equal(t, "Hello Require", result, "Result must be '"+result+"'")
	fmt.Println("Executed")
}

func TestHelloWorldAditya(t *testing.T) {
	result := HelloWorld("Aditya")

	if result != "Hello Aditya" {
		t.Error("Result must be '" + result + "'")
	}

	fmt.Println("Executed")
}

func TestHelloWorldRicki(t *testing.T) {
	result := HelloWorld("Ricki")

	if result != "Hello Ricki" {
		t.Fatal("Result must be '" + result + "'")
	}

	fmt.Println("Executed")
}
