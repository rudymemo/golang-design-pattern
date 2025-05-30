package decorator

import (
	"math"
	"testing"
)

func TestDecorator(t *testing.T) {
	// Simple coffee
	coffee := &SimpleCoffee{}
	if coffee.Cost() != 2.0 {
		t.Errorf("Expected cost 2.0, got %.2f", coffee.Cost())
	}
	if coffee.Description() != "Simple Coffee" {
		t.Errorf("Expected 'Simple Coffee', got %s", coffee.Description())
	}
	
	// Coffee with milk
	coffeeWithMilk := NewMilkDecorator(coffee)
	if coffeeWithMilk.Cost() != 2.5 {
		t.Errorf("Expected cost 2.5, got %.2f", coffeeWithMilk.Cost())
	}
	if coffeeWithMilk.Description() != "Simple Coffee, Milk" {
		t.Errorf("Expected 'Simple Coffee, Milk', got %s", coffeeWithMilk.Description())
	}
	
	// Coffee with milk and sugar
	coffeeWithMilkAndSugar := NewSugarDecorator(coffeeWithMilk)
	if coffeeWithMilkAndSugar.Cost() != 2.7 {
		t.Errorf("Expected cost 2.7, got %.2f", coffeeWithMilkAndSugar.Cost())
	}
	if coffeeWithMilkAndSugar.Description() != "Simple Coffee, Milk, Sugar" {
		t.Errorf("Expected 'Simple Coffee, Milk, Sugar', got %s", coffeeWithMilkAndSugar.Description())
	}
	
	// Coffee with all decorators
	fullCoffee := NewWhipDecorator(coffeeWithMilkAndSugar)
	expectedCost := 3.4
	if math.Abs(fullCoffee.Cost()-expectedCost) > 0.001 {
		t.Errorf("Expected cost %.2f, got %.2f", expectedCost, fullCoffee.Cost())
	}
	if fullCoffee.Description() != "Simple Coffee, Milk, Sugar, Whip" {
		t.Errorf("Expected 'Simple Coffee, Milk, Sugar, Whip', got %s", fullCoffee.Description())
	}
}