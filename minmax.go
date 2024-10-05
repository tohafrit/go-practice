//go:build cmain

package main

import (
	"errors"
	"fmt"
)

func findMinMax(nums []int) (min, max int, err error) {
	if len(nums) == 0 {
		return 0, 0, errors.New("slice is empty")
	}
	min, max = nums[0], nums[0]
	for _, num := range nums[1:] {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
	}

	return min, max, nil
}

func main() {
	numbers := []int{3, 5, 1, 8, -2, 7}
	min, max, err := findMinMax(numbers)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Printf("Slice: %v\nMinimum: %d, Maximum: %d\n", numbers, min, max)
	}
}
