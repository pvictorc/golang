package main

import "fmt"

func main() {
	number := 8
	numberText := "eight"

	if number == 8 && numberText == "eight" {
		fmt.Println("Success!")
	} else {
		fmt.Println("Fail!")
	}
}