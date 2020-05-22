package factory

import (
	"errors"
	"fmt"
)

const (
	Cash = iota
	DebitCard
)

type PaymentMethod interface {
	Pay(amount float32) string
}

func GetPaymentMethod(m int) (PaymentMethod, error) {
	switch m {
	case Cash:
		return &CashPayment{}, nil
	case DebitCard:
		return &DebitPayment{}, nil
	default:
		return nil, errors.New(fmt.Sprintf("Payment method %d not recognized.\n", m))
	}
}

type CashPayment struct{}

func (c *CashPayment) Pay(amount float32) string {
	return fmt.Sprintf("$%0.2f paid using cash.\n", amount)
}

type DebitPayment struct{}

func (d *DebitPayment) Pay(amount float32) string {
	return fmt.Sprintf("$%0.2f paid using a debit card.\n", amount)
}
