package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const Url = "https://disease.sh/v3/covid-19/all"

func main() {
	response, err := http.Get(Url)

	if err != nil {
		panic(err)
	}
	fmt.Println("Response type : %T\n", response)

	defer response.Body.Close()

	databyte, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	txtData := string(databyte)

	file, err := os.Create("web_content.txt")

	if err != nil {
		panic(err)
	}

	defer file.Close()

	length, err := file.WriteString(txtData)
	if err != nil {
		panic(err)
	}
	fmt.Println(length)
	readData("web_content.txt")
}

func readData(file string) {
	content, err := os.ReadFile(file)
	if err != nil {
		panic(err)
	}
	fmt.Println(content)
	fmt.Println(len(content))
}
