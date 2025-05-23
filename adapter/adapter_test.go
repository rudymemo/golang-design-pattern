package adapter

import (
	"testing"
)

// TestOldPrinter_Print 测试 OldPrinter 的 Print 方法
func TestOldPrinter_Print(t *testing.T) {
	// Arrange
	oldPrinter := &OldPrinter{}
	input := "hello"

	output := oldPrinter.Print(input)

	expected := "old printer: hello"

	if output != expected {
		t.Errorf("Expected %s, but got %s", expected, output)
	}
}

// TestPrinterAdapter_Print 测试 PrinterAdapter 的 Print 方法
func TestPrinterAdapter_Print(t *testing.T) {
	// Arrange
	oldPrinter := &OldPrinter{}
	adapter := &PrinterAdapter{
		oldPrinter: oldPrinter,
	}
	input := "world"

	output := adapter.Print(input)

	expected := "old printer: world"

	if output != expected {
		t.Errorf("Expected %s, but got %s", expected, output)
	}
}
