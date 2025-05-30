package strategy

import "fmt"

type PaymentStrategy interface {
	Pay(amount float64) string
}

type CreditCardPayment struct {
	cardNumber string
}

func NewCreditCardPayment(cardNumber string) *CreditCardPayment {
	return &CreditCardPayment{cardNumber: cardNumber}
}

func (cc *CreditCardPayment) Pay(amount float64) string {
	return fmt.Sprintf("Paid $%.2f using Credit Card ending in %s", amount, cc.cardNumber[len(cc.cardNumber)-4:])
}

type PayPalPayment struct {
	email string
}

func NewPayPalPayment(email string) *PayPalPayment {
	return &PayPalPayment{email: email}
}

func (pp *PayPalPayment) Pay(amount float64) string {
	return fmt.Sprintf("Paid $%.2f using PayPal account %s", amount, pp.email)
}

type BankTransferPayment struct {
	accountNumber string
}

func NewBankTransferPayment(accountNumber string) *BankTransferPayment {
	return &BankTransferPayment{accountNumber: accountNumber}
}

func (bt *BankTransferPayment) Pay(amount float64) string {
	return fmt.Sprintf("Paid $%.2f using Bank Transfer from account %s", amount, bt.accountNumber)
}

type PaymentContext struct {
	strategy PaymentStrategy
}

func NewPaymentContext(strategy PaymentStrategy) *PaymentContext {
	return &PaymentContext{strategy: strategy}
}

func (pc *PaymentContext) SetStrategy(strategy PaymentStrategy) {
	pc.strategy = strategy
}

func (pc *PaymentContext) ExecutePayment(amount float64) string {
	return pc.strategy.Pay(amount)
}