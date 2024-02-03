package main

import (
	"fmt"
)

func main() {
	// Test the power function
	// Created by Dharmik Darji
	base, exponent := 2, 3
	result := power(base, exponent)
	fmt.Println("::DHARMIK DARJI functions's output::")
	fmt.Println(base, "raised to the power of", exponent, "is", result)
	fmt.Println("")
}

// function Two
// Created by dharmik darji
// Calculates and returns the result of raising a base to a given exponent
// INPUT: base - the base value to be raised
//
//	exponent - the exponent value to which the base is raised
//
// OUTPUT: result (integer) - the result of base raised to the exponent
func power(base, exponent int) int {
	result := 1
	for i := 0; i < exponent; i++ {
		result *= base
	}
	return result
}
