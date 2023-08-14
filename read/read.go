package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	// Read content of textFile
	content, err := ioutil.ReadFile("names.txt")
	if err != nil {
		fmt.Println("Error", err)
		return
	}

	// Split the content into lines
	lines := strings.Split(string(content), "\n")

	/*
	strings.Fields function splits the current line into fields based on whitespace char
	(spaces/tabs)
	*/
	for _, line := range lines {
		parts := strings.Fields(line)
		if len(parts) >= 2 {
			firstName := parts[0]
			lastName := parts[1]
			fmt.Printf("The First Name is %s, and The Last Nameis %s\n", firstName, lastName)
		}
	}
}
