package strategy

import "testing"

func TestStrategy(t *testing.T) {
	amount := 100.50
	
	// Test Credit Card Payment
	creditCard := NewCreditCardPayment("1234567890123456")
	context := NewPaymentContext(creditCard)
	result := context.ExecutePayment(amount)
	expected := "Paid $100.50 using Credit Card ending in 3456"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
	
	// Test PayPal Payment
	paypal := NewPayPalPayment("user@example.com")
	context.SetStrategy(paypal)
	result = context.ExecutePayment(amount)
	expected = "Paid $100.50 using PayPal account user@example.com"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
	
	// Test Bank Transfer Payment
	bankTransfer := NewBankTransferPayment("ACC123456789")
	context.SetStrategy(bankTransfer)
	result = context.ExecutePayment(amount)
	expected = "Paid $100.50 using Bank Transfer from account ACC123456789"
	if result != expected {
		t.Errorf("Expected %s, got %s", expected, result)
	}
}