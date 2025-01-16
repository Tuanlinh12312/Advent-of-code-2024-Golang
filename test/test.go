package main

import "fmt"

func main() {
	// Define a recursive lambda function using a variable
	var factorial func(int) int
	factorial = func(n int) int {
		if n <= 1 { // Base case
			return 1
		}
		return n * factorial(n-1) // Recursive call
	}

	// Call the lambda function
	fmt.Println(factorial(5)) // Output: 120
}