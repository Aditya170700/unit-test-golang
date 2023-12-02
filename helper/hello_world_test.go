package helper

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"runtime"
	"testing"
)

func TestMain(m *testing.M) {
	// before
	fmt.Println("Testing Started")

	// run unit test
	m.Run()

	// after
	fmt.Println("Testing Finished")
}

func TestHelloWorldSubTest(t *testing.T) {
	t.Run("Aditya", func(t *testing.T) {
		result := HelloWorld("SubTestAditya")
		assert.Equal(t, "Hello SubTestAditya", result, "Result must be '"+result+"'")
	})
	t.Run("Ricki", func(t *testing.T) {
		result := HelloWorld("SubTestRicki")
		assert.Equal(t, "Hello SubTestRicki", result, "Result must be '"+result+"'")
	})
}

func TestHelloWorldSkip(t *testing.T) {
	if runtime.GOOS == "darwin" {
		t.Skip("Cannot run on MacOS")
	}

	result := HelloWorld("Skip")
	assert.Equal(t, "Hello Skip", result, "Result must be '"+result+"'")
}

func TestHelloWorldAssert(t *testing.T) {
	result := HelloWorld("Assert")
	assert.Equal(t, "Hello Assert", result, "Result must be '"+result+"'")
}

func TestHelloWorldRequire(t *testing.T) {
	result := HelloWorld("Require")
	require.Equal(t, "Hello Require", result, "Result must be '"+result+"'")
}

func TestHelloWorldAditya(t *testing.T) {
	result := HelloWorld("Aditya")

	if result != "Hello Aditya" {
		t.Error("Result must be '" + result + "'")
	}
}

func TestHelloWorldRicki(t *testing.T) {
	result := HelloWorld("Ricki")

	if result != "Hello Ricki" {
		t.Fatal("Result must be '" + result + "'")
	}
}
