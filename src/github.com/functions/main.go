package main

import (
	"fmt"
	"strings"
)

func add(a, b int) int{
	return (a+b)
}

func sayLoudly(phrase string) string{
	return strings.ToUpper(phrase)
}

func main(){

	a:=1
	b:=2

	fmt.Println("",add(a,b))

	fmt.Println(sayLoudly("Olá Mundo!"))
}