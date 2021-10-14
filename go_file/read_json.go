package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

func main() {
	parseUnstructuredData()
}

//Working with structured Data
func parseStructuredData() {
	jsonFile, err := os.Open("users.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Succ open")
	defer jsonFile.Close()

	//read opened jsonFile as a byte array
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var users Users
	json.Unmarshal(byteValue, &users)

	//print
	for i := 0; i < len(users.Users); i++ {
		fmt.Println("User Type: " + users.Users[i].Type)
		fmt.Println("User Age: " + strconv.Itoa(users.Users[i].Age))
		fmt.Println("User Name: " + users.Users[i].Name)
		fmt.Println("Facebook Url: " + users.Users[i].Social.Facebook)
	}
}

//Working with Unstructured Data

func parseUnstructuredData() {
	jsonFile, err := os.Open("users.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Succ open")
	defer jsonFile.Close()

	//read opened jsonFile as a byte array
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var result map[string]interface{}
	json.Unmarshal(byteValue, &result)

	fmt.Println(result["users"])
}

// Users struct which contains
// an array of users
type Users struct {
	Users []User `json:"users"`
}

// User struct which contains a name
// a type and a list of social links
type User struct {
	Name   string `json:"name"`
	Type   string `json:"type"`
	Age    int    `json:"Age"`
	Social Social `json:"social"`
}

// Social struct which contains a
// list of links
type Social struct {
	Facebook string `json:"facebook"`
	Twitter  string `json:"twitter"`
}

//ref: https://tutorialedge.net/golang/reading-writing-files-in-go/
