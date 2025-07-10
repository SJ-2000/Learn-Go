package main
import "fmt"

func main() {
	const pi = 3.14
	const name = "Suryajith Sujith"
	//pi = 3.14159 // This will cause an error because pi is a constant
	fmt.Println("Value of pi:", pi)
	const(
		//grouped constants
		red = "RED"
		green = "GREEN"
		blue = "BLUE"
	)
	const taxRate = 0.18
	price := 100.0
	fmt.Println("FinalPrice", price + (price * taxRate))
}