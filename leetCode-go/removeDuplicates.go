//https://leetcode.com/problems/contains-duplicate/
package main

import(
	"fmt"
)

func containsDuplicate(nums []int) bool {
    //using map method 
	//map keys are integers
	//the values are booleans 
	// to indicate whether the number has been encountered.
	hashset := make(map[int]bool)

	for _, n := range nums {
		if hashset[n] {
			return true
		}
		hashset[n] = true
	}
	return false
}

func main(){
	nums := []int{1,2,3,4,2}
	fmt.Println(containsDuplicate(nums))
}