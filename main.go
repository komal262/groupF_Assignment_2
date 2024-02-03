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

	// Test the factorial function
	// Created by Kashish Maggu
	factorialNum := 5
	fmt.Println("::KASHISH MAGGU functions's output::")
	fmt.Println("Factorial of", factorialNum, "is", factorial(factorialNum))
	fmt.Println("")

	// Test the isEven function
	// Created by Jitender kaushik
	num := 7
	// If it's even, print that it's even, otherwise print that it's odd.
	fmt.Println("::JITENDER KAUSHIK functions's output::")
	if isEven(num) {
		fmt.Println(num, "is even.")
	} else {
		fmt.Println(num, "is odd.")
	}
	fmt.Println("")

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

// function Three
// created by kashish maggu
// purpose: returns factorial of given number
func factorial(n int) int {
    if n == 0 || n == 1 {
        return 1
    }
    return n * factorial(n-1)
}


// function three
// Created by jitender kaushik
// isEven function checks if a number is even or not.
// INPUT: give any number
// OUTPUT: provides given number is odd or even.
func isEven(num int) bool {
	return num%2 == 0 // If num is divisible by 2 with no remainder, it's even.