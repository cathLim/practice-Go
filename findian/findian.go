package main

import (
	"fmt"
	"strings"
)

func main(){
	//user input string
	var userString string
	fmt.Println("Input string?")

	//Read user input and store it in _ variable
	_, err :=
	fmt.Scan(&userString)

	//To check if there was an error while reading the input
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	//To check if the string contains the character 'i', 'a' and 'n'
	if(strings.Contains(userString, "i") || strings.Contains(userString, "a") || strings.Contains(userString, "n")){
		fmt.Println("Found!")
	}else{
		fmt.Println("Not Found!")
	}
}