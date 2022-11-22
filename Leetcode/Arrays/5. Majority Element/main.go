package main

import "fmt"

func main() {
	nums := []int{3, 2, 3}

	fmt.Println(majorityElement(nums))

}

// Boyer–Moore majority vote algorithm - https://en.wikipedia.org/wiki/Boyer–Moore_majority_vote_algorithm
func majorityElement(nums []int) int {
	var count, result int

	for i := 0; i < len(nums); i++ {
		if count == 0 {
			result = nums[i]
			count++
		} else {
			if count == nums[i] {
				count++
			} else {
				count--
			}
		}
	}

	return result
}
