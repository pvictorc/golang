
package main

import "fmt"

func main() {
	var language string = "Go"
	const released int = 2009
	isAProgrammingLanguage:=true

	fmt.Println("language, relased, isAProgrammingLanguage: ",
	language,released,isAProgrammingLanguage)

	language = "C"
	if language == "Java"{
		fmt.Println("language is Java")
	} else if language == "Go"{
		fmt.Println("language is Go")
	} else {
		fmt.Println("don't know what language is")
	}
}

