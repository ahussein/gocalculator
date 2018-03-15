// Unitests for the calculator module
package calculator

import (
	"testing"
)

// TestAdd will test the Add function of the calculator
func TestAdd(t *testing.T) {
	expectedResult := float32(15)
	actualResult := Add(10, 5)
	if expectedResult != actualResult {
		t.Error("Add function does not return the expected result. Expected ", expectedResult, " got ",
			actualResult)
	}
}

// TestSubtract will test the Subtract function of the calculator
func TestSubtract(t *testing.T) {
	expectedResult := float32(5)
	actualResult := Subtract(10, 5)
	if expectedResult != actualResult {
		t.Error("Subtract function does not return the expected result. Expected ", expectedResult, " got ",
			actualResult)
	}
}

// TestMultiply will test the Multiply function of the calculator
func TestMultiply(t *testing.T) {
	expectedResult := float32(50)
	actualResult := Multiply(10, 5)
	if expectedResult != actualResult {
		t.Error("Multiply function does not return the expected result. Expected ", expectedResult, " got ",
			actualResult)
	}
}

// TestDivideOK will test the Divide function of the calculator
func TestDivideOK(t *testing.T) {
	expectedResult := float32(2)
	actualResult, err := Divide(10, 5)
	if err != nil {
		t.Error("Divide function raises an error while it was not expected. Error ", err)
	}
	if expectedResult != actualResult {
		t.Error("Divide function does not return the expected result. Expected ", expectedResult, " got ",
			actualResult)
	}
}

// TestDivideByZero will test the Divide function of the calculator when dividing by zero
func TestDivideByZero(t *testing.T) {
	_, err := Divide(10, 0)
	if err == nil {
		t.Error("Divide function should raise an error when dividing by zero")
	}
}
