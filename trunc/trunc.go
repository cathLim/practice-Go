package main

import (
	"fmt"
)

func main(){
	//user input floating-point number
	var userFloatNum float64
	fmt.Printf("Floating-point number?")

	//Read user input and store it in _variable (which discard the value)
	_, err :=
	fmt.Scan(&userFloatNum)

	//To check if there was an error while reading the input
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	//Print the floating-point number
	fmt.Printf("The original number: %f\n", userFloatNum)

	//Truncate the floating-point number to an integer
	truncatedNumber := int(userFloatNum)
	fmt.Printf("Truncated to integer: %d\n", truncatedNumber)
}