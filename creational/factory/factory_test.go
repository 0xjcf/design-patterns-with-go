package factory

import (
	"strings"
	"testing"
)

func TestGetPaymentMethodCash(t *testing.T) {
	paymentMethod, err := GetPaymentMethod(Cash)
	if err != nil {
		t.Fatal("Payment method is not accepted!")
	}
	msg := paymentMethod.Pay(1.00)
	if !strings.Contains(msg, "paid using cash") {
		t.Errorf("The cash payment message is incorrect. Actual -> %s\n", msg)
	}
	t.Log("LOG:", msg)
}

func TestGetPaymentMethodDebit(t *testing.T) {
	paymentMethod, err := GetPaymentMethod(DebitCard)
	if err != nil {
		t.Fatal("Payment method is not accepted")
	}
	msg := paymentMethod.Pay(10.00)
	if !strings.Contains(msg, "paid using a debit card") {
		t.Errorf("The debit card payment message is incorrect. Actual -> %s\n", msg)
	}
	t.Log("LOG:", msg)
}

func TestGetPaymentMethodUnknown(t *testing.T) {
	_, err := GetPaymentMethod(3)
	if err == nil {
		t.Error("Payment method is unknown")
	}
	t.Log("LOG:", err)
}
