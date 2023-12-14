package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to pizza app");

	reader := bufio.NewReader(os.Stdin);

	fmt.Println("Please rate between 1 and 5");

	input, _ := reader.ReadString('\n');
	fmt.Println(input);

	number, err := strconv.ParseFloat(strings.TrimSpace(input), 64);

	if(err!=nil) {
		fmt.Println(err);
	} else {
		fmt.Println("Adding two numbers %d", number+1);
	}
}