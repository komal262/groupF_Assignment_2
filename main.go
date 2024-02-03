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

	//Testing Fibonacci Series Function (Function Two)
	fmt.Println("Output of the function created by Navneet Kaur ")
	fmt.Println(Fibonacci(8))
}

// function One
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

// function Two
// Created by Navneet Kaur
// Function generates a Fibonacci series up to n terms
// Input: n (int)
// Output: fibSer ([]int)
func Fibonacci(n int) []int {
	fibSer := make([]int, n)
	fibSer[0], fibSer[1] = 0, 1

	for i := 2; i < n; i++ {
		fibSer[i] = fibSer[i-1] + fibSer[i-2]
	}

	return fibSer
}
