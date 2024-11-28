package main

import "fmt"

func main() {
	var x int = 45 // Initialize x to 45
	var a int = 40 // Set a starting value less than x for the loop to run

	// First loop: Modify condition so it runs
	for i := 10; i < 10; i++ { // Use a new variable `i` to avoid shadowing `a`
		fmt.Printf("Value is %d\n", i)
	}

	// Second loop: Ensure `a < x` at the start
	for a < x {
		a++                                // Increment a
		fmt.Printf("You are lady %d\n", a) // Corrected formatting
	}
}
