package main

import "fmt"

func main() {
	nums := []int{1, 3, 2, 1}
	fmt.Println(getConcatenation(nums))
}

func getConcatenation(nums []int) []int {
	nums = append(nums, nums...)

	return nums
}
