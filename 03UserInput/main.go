package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input program";
	fmt.Println(welcome);

	reader := bufio.NewReader(os.Stdin);
	
	fmt.Println("Enter the rating for the food : ");
	input, _ := reader.ReadString('\n');
	fmt.Println("Thank you for %d rating ", input);
	fmt.Printf("Type of rating was %T", input);

}	