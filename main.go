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

	// Test the calculateCube function
	// Created by Jashanpreet kaur
	inputNumber := 6
	fmt.Println("::JASHANPREET KAUR functions's output::")
	fmt.Printf("The cube of %d is: %d\n", inputNumber, calculateCube(inputNumber))
	fmt.Println("")

	// Test the toUpperString function
	// Created by Sidham Kour
	inputString := "hello, world!"
	fmt.Println("::SIDHAM KOUR functions's output::")
	fmt.Printf("The uppercase string of \"%s\" is: \"%s\"\n", inputString, toUpperString(inputString))
	fmt.Println("")
	// Repeat the word function
	// Created by Komal Negah
	fmt.Println("::Komal Negah functions's output::")
	repeat("\nKomal", 10)

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
	if num&1 == 0 {
		return true
	}
	return false
}

// function Four
// created by Jashanpreet kaur
// purpose: calculateCube takes a number as an argument and gives output in number's cube.
// INPUT: give any number
// OUTPUT: returns cube of given number.
func calculateCube(number int) int {
	return number * number * number
}

// Function five
// Created by Sidham Kour
// Convert all characters in the given string to uppercase
// INPUT: input normal string
// OUTPUT: gives string with all characters in uppercase
func toUpperString(input string) string {
	// Convert the input string to a slice of runes for manipulation
	runes := []rune(input)

	// Iterate through each rune in the slice
	for i := 0; i < len(runes); i++ {
		// Check if the current rune is a lowercase letter
		if runes[i] >= 'a' && runes[i] <= 'z' {
			// Convert the lowercase letter to uppercase
			runes[i] = runes[i] - 'a' + 'A'
		}
	}

	// Convert the slice of runes back to a string and return
	return string(runes)
}

// Function six
// Created by Komal Negah
// Program to repeat the word
// OUTPUT: repeat the given word

func repeat(text string, digit int) {
	for i := 0; i < digit; i++ {
		fmt.Println(text)
	}
}
