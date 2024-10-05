//go:build cname

package main

import "fmt"

func reverse(nums []int) []int {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}

func main() {
	input := []int{1, 2, 3, 4, 5}
	output := reverse(input)
	fmt.Println(output)
}
