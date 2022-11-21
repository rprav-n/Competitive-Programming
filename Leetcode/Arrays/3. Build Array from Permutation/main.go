package main

import "fmt"

func main() {
	nums := []int{0, 2, 1, 5, 3, 4}
	fmt.Println(buildArray(nums))
}

func buildArray(nums []int) []int {
	l := len(nums)
	ans := make([]int, l)

	for i := range nums {
		ans[i] = nums[nums[i]]
	}

	return ans
}
