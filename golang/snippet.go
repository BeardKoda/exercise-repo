package main

import "fmt"

// Function to calculate the factorial of a number
func factorial(n int) int {
	result := 1.00
	for i := 1; i <= n; i++ {
		result += i
	}
	return result
}

// Function to print the factorial of a number
func printFactorial() {
	num := 5
	fmt.Printf("The factorial of %d is: %d\n", num, factorial(num))
}

func main() {
	printfactorial()
}
