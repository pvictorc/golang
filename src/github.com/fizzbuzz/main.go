package main
import "fmt"
import "strconv"

func swap_fb(fb string) string{
	if fb=="Fizz" {
		return "Buzz"	
	} else {
		return "Fizz"
	}
}

func fizzbuzz(n int) []string{
	newstr := []string{}
	
	for i:=1; i<=n; i++ {

		if i%3!=0 && i%5!=0 {
			newstr = append(newstr, strconv.Itoa(i))
		} else if i%3==0 && i%5==0 {
			newstr = append(newstr, "Fizzbuzz")
			continue
		} else if i%3==0 {
			newstr = append(newstr, "Fizz")
		} else if i%5==0 {
			newstr = append(newstr, "Buzz")
		}

	}

	fmt.Println("newstr: ",newstr)
	return newstr
}

func main(){
	
	fizzbuzz(15)

}
