package main
import "fmt"

func main(){
	// 0 1 1 2 3 5
	var fibo [10]int
	fibo[0] = 0
	fibo[1] = 1

	for i:=2;i<10;i++ {
		
		fibo[i]=fibo[i-2]+fibo[i-1]
		// fmt.Printf("%d ",fibo[i])
	}

	fmt.Println(fibo)
	fmt.Println(fibo[0:4])
}