package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	numMap := make(map[int]int)

	for i, num := range nums {
		complement := target - num

		if index, found := numMap[complement]; found {
			return []int{index, i}
		}

		numMap[num] = i
	}

	return nil
}

func main() {
	numbers := []int{14, 18, 3, 8, 71}
	target := 79

	result := twoSum(numbers, target)
	fmt.Println("Indices: ", result)
}
