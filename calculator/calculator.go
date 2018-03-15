package calculator

import (
	"errors"
)

// Add adds tow integers and return a float result	
func Add(param1 int32, param2 int32) float32{
	return float32(param1 + param2)
}

// Subtract substracts two integers and return a float result	
func Subtract(param1 int32, param2 int32) float32{
	return float32(param1 - param2)
}

// Multiply multiplies two integers and return a float result
func Multiply(param1 int32, param2 int32) float32{
	return float32(param1 * param2)
}


// Divide divides two integers and return a float result
func Divide(param1 int32, param2 int32) (float32, error){
	if param2 == 0 {
		return float32(0), errors.New("Division by zero is not allowed")
	}
	return (float32(param1) / float32(param2)), nil
}