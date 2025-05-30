package chain

import (
	"testing"
)

func TestChainOfResponsibility(t *testing.T) {
	logger := NewLogger()

	// Test different log levels
	logger.Log(1, "This is an info message")
	logger.Log(2, "This is a warning message")
	logger.Log(3, "This is an error message")
	logger.Log(4, "This is a critical error message")

	// The test passes if no panic occurs and messages are handled
}