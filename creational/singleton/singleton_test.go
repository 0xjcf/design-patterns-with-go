package singleton

import "testing"

func TestGetInstance(t *testing.T) {
	counter1 := GetInstance()

	if counter1 == nil {
		t.Error("Expected counter1 to not be nil after calling GetInstance()")
	}

	expectedCounter := counter1

	counter2 := GetInstance()

	if counter2 != expectedCounter {
		t.Error("Expected counter2 to be the same instance as counter1")
	}
}

func TestGetCountValue(t *testing.T) {
	counter1 := GetInstance()
	currentCounter := counter1.AddOne()
	if currentCounter != 1 {
		t.Error("Expected currentCounter to be 1 after calling AddOne() on counter1")
	}

	counter2 := GetInstance()
	currentCounter = counter2.AddOne()
	if currentCounter != 2 {
		t.Error("Expected currentCounter to be 2 after calling AddOne() on counter1 & counter2")
	}

	counter3 := GetInstance()
	currentCounter = counter3.SubtractOne()
	if currentCounter != 1 {
		t.Error("Expected currentCounter to be 1 after calling AddOne() on counter1 & counter2, & calling SubtractOne() on counter3 ")
	}
}
