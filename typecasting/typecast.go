package main
import "fmt"

func main() {
	var x int = 10
	var y float64 = float64(x) // Type casting from int to float64
	fmt.Println("Value of y:", y)

	str := "123"
	num := 123
	// Type casting from string to int is not directly possible in Go.
	// You need to use strconv package for that, but here we are just demonstrating type casting.
	// If you want to concatenate a string and a number, you need to convert the number
	// to a string first.
	// You can't do this: fmt.Println(str + num)
	// Instead, convert num to string:
	fmt.Println(str + fmt.Sprint(num))
}