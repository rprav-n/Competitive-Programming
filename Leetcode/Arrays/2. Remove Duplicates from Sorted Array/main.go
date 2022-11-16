package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}

	result := removeDuplicates(nums)

	fmt.Println(result)
}

func removeDuplicates(nums []int) int {
	var k int
	numsMap := make(map[int]struct{})

	for _, n := range nums {
		if _, ok := numsMap[n]; !ok {
			numsMap[n] = struct{}{}
			nums[k] = n
			k++
		}
	}

	return k
}
