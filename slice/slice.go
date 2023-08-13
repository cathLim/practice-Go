package main

import(
	"fmt"
	"sort"
)

func main(){
	var num1, num2, num3, num4 int

	fmt.Print("Ënter the first integer")
	fmt.Scan(&num1)

	fmt.Print("Ënter the second integer")
	fmt.Scan(&num2)

	fmt.Print("Ënter the third integer")
	fmt.Scan(&num3)

	fmt.Print("Ënter the forth integer")
	fmt.Scan(&num4)

	//Create a slice and append the four distinct integers
	numbers := []int{num1,num2,num3,num4}

	//sort the slice in ascending order
	sort.Ints(numbers)

	fmt.Println("The sorted slice:", numbers)


}