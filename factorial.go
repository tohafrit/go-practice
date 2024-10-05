//go:build cmain

package main

import "fmt"

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}

func main() {
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)
	if n < 0 {
		fmt.Println("Invalid number")
		return
	}
	fmt.Printf("Factorial %d is %d\n", n, factorial(n))
}
