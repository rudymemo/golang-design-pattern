package state

import (
	"testing"
)

func TestStatePattern(t *testing.T) {
	vm := NewVendingMachine()

	// Test initial state
	if !vm.HasProduct() {
		t.Error("Vending machine should have product initially")
	}

	// Test inserting coin
	vm.InsertCoin()

	// Test selecting product
	vm.SelectProduct()

	// Test dispensing product
	vm.DispenseProduct()

	// After dispensing, should not have product
	if vm.HasProduct() {
		t.Error("Vending machine should not have product after dispensing")
	}

	// Test trying to select product without coin
	vm.SelectProduct()
}