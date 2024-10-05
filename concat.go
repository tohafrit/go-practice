//go:build cmain

package main

import "fmt"

func concatenate(cl1 []int, cl2 []int) []int {
	return append(cl1, cl2...)
}

func main() {
	s1 := []int{1, 2, 3}
	s2 := []int{4, 5, 6}
	result := concatenate(s1, s2)
	fmt.Println(result)
}
