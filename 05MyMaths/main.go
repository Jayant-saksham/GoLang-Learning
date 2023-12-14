package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	// "math/rand"
)

func main() {
	fmt.Println("Welcome to Maths commands in GoLang");

	// fmt.Println(rand.Intn(100));

	myRandomNumber, _ := rand.Int(rand.Reader, big.NewInt(5));
	fmt.Println(myRandomNumber);
}