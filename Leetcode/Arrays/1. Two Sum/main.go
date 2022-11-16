package main

import "fmt"

func main() {
	nums := []int{3, 2, 4}
	target := 6

	result := twoSum(nums, target)

	fmt.Println(result)
}

func twoSum(nums []int, target int) []int {
	var result []int

	numsMap := make(map[int]int)

	for i, n := range nums {
		if j, ok := numsMap[target-n]; ok {
			result = append(result, []int{j, i}...)
			break
		} else {
			numsMap[n] = i
		}
	}

	return result
}
