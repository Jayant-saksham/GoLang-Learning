package main

import "fmt"

// Capital l (L) = It means that it is public 
const Login string = "logintoken";

func main() {
	var username string = "This is a string"
	fmt.Printf("The variable is of type: %T\n", username);

	var isLoggedIn bool = false;
	fmt.Printf("The variable is of type: %T\n", isLoggedIn);

	var integer uint8 = 255;
	fmt.Println(integer);

	var inte int;
	fmt.Println(inte);

	var strings string;
	fmt.Println(strings);

	// Implicit types 
	var website = "flydeck.com";
	fmt.Println(website);

	// NO var keyword - Walrus opeator (:=) - It can be declared "only" inside methods 
	numUser := 300;
	fmt.Println(numUser);
}
