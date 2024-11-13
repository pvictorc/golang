package main
import "fmt"

func main(){
	for i := 0; i <= 20; i += 2 {
		fmt.Printf("%d",i)
		if i<20{
			fmt.Printf(", ")
		}
	}
	fmt.Printf("\n")
}
