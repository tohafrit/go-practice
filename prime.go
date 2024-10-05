//go:build cmain

package main

import (
	"fmt"
	"math"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	sqrtN := int(math.Sqrt(float64(n)))
	for i := 2; i <= sqrtN; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)
	if isPrime(n) {
		fmt.Printf("%d prime", n)
	} else {
		fmt.Printf("%d not prime", n)
	}
}
