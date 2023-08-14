//https://leetcode.com/problems/two-sum/
package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	var indices = make(map[int]int)

	for i, value := range nums{
		x := target - value;

		found, ok := indices[x]
		if ok {
			result := []int {found, i}
			return result
		}
		indices[value] = i
	}
	return nil

}