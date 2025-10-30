package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Option 1: return all planes
func returnAllPlanes() {
	// Set up http get
	resp, err := http.Get("http://localhost:8000/planes")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	// Print the response
	fmt.Println(string(body))
}

// Option 2: return plane by ID
func returnPlaneByID() {
	// Get ID number from user
	var id int
	fmt.Print("Enter an ID number: ")
	fmt.Scanln(&id)
	fmt.Println()

	// Set up the http get
	resp, err := http.Get(fmt.Sprintf("http://localhost:8000/planes/%d", id))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	// Print the response
	fmt.Println(string(body))
}

// Option 3: add a plane to the dict
func addPlane() {
	// Get plane information from user
	var model string
	var manu string
	var range_mi float64
	var fixed_gear bool
	fmt.Print("Enter the plane model: ")
	fmt.Scanln(&model)
	fmt.Print("Enter the plane manufacturer: ")
	fmt.Scanln(&manu)
	fmt.Print("Enter the plane's range (in miles, using #'s): ")
	fmt.Scanln(&range_mi)
	fmt.Print("Is the plane fixed gear? ('true' or 'false'): ")
	fmt.Scanln(&fixed_gear)

	// Format the request
	jsonData := []byte(fmt.Sprintf(
		`{"model":"%s","manufacturer":"%s","range_mi":%f,"fixed_gear":%t}`,
		model, manu, range_mi, fixed_gear,
	))

	// Set up the http post
	resp, err := http.Post("http://localhost:8000/planes", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	// Print the response
	fmt.Println(string(body))
}

// Option 4: delete a plane by ID
func deletePlaneByID() {
	// Get ID number from user
	var id int
	fmt.Print("Enter an ID number: ")
	fmt.Scanln(&id)
	fmt.Println()

	// Set up the http request
	client := &http.Client{}
	url := fmt.Sprintf("http://localhost:8000/planes/del/%d", id)
	req, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		panic(err)
	}

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)
	// Print the response
	fmt.Println(string(body))
}

// Display the options to user
func printOptions() int {
	var option int
	fmt.Println("Select an option below:")
	fmt.Println("1. Return all planes")
	fmt.Println("2. Return plane by ID")
	fmt.Println("3. Add a plane to the dict")
	fmt.Println("4. Delete a plane by ID")
	fmt.Println("5. Quit")
	fmt.Print("Enter your desired option (1-5): ")
	fmt.Scanln(&option)
	fmt.Scanln("Option %d chosen...", option)
	fmt.Println()

	return option
}

// Main function for redirecting logic for a particular option
func main() {
	input := printOptions()

	switch input {
	case 1:
		returnAllPlanes()
	case 2:
		returnPlaneByID()
	case 3:
		addPlane()
	case 4:
		deletePlaneByID()
	case 5:
		fmt.Println("See ya!")
	}

	if input < 5 {
		main()
	}
}
