//The following are some steps to store user input to JSON object
/*
1. Import the required packages
   encoding/json packafe for JSON handling
   fmt to format source code files

2. Collect user input, using fmt package
3. Create a struct that corresponds to the structure of JSON object
4. Assign User input to struct
5. Use json.Marshal() function to sterialize the struct into a JSON encoded byte array
*/

package main

import(
	"encoding/json"
	"fmt"
)

//Define a struct to store user data
//field tag are specified within backticks {`}
type UserData struct{
	Name string `json:"name"`
	Address string `json:"address"`
}

func main(){
	//Get user input
	var name, addresses string
	fmt.Print("Name?")
	fmt.Scanln(&name)

	fmt.Print("Address?")
	fmt.Scanln(&addresses)

	//Create a struct instances and assign user input
	user := UserData{
		Name: name,
		Address: addresses,
	}

	//the json.MarshalIndent() function from the "encoding/json" package is use
	// to serialize a Go struct (user) into a JSON-formatted byte array (jsonData)
	//then printed to the console
	jsonData, err := json.MarshalIndent(user, "", "  ")
	if err != nil{
		fmt.Println("Error", err)
		return
	}

	fmt.Println(string(jsonData))
}