//go:build cmain

package main

import "fmt"

func removeElement(nums []int, index int) []int {
	if index < 0 || index >= len(nums) {
		return nums
	}
	return append(nums[:index], nums[index+1:]...)
}

func main() {
	input := []int{10, 20, 30, 40, 50}
	index := 2
	output := removeElement(input, index)
	fmt.Println(output)
}
