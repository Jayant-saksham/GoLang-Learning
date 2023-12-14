package main

import (
	"fmt"
	"time"
)

func main() {

	presentTime := time.Now();
	fmt.Println(presentTime);
	

	year, month, date := presentTime.Date();
	fmt.Println(year);
	fmt.Println(month);
	fmt.Println(date);


}