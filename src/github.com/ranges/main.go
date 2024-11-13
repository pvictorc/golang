package main
import "fmt"

func reduce(slice []int) int{
	sum := 0
	for _, n := range slice {
		sum+=n
	}
	return sum
}

func main(){

	numbers := []int{30, 11, 40, 23, 12}

	// for _, num := range numbers {
	// 	if num%2 == 0 {
	// 		fmt.Println(num)
	// 	}
	// }

	fmt.Println( reduce(numbers) )
}
