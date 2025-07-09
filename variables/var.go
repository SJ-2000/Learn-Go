package main
import "fmt"

func main() {
	var myString string = "Hello, Go!"
	var myInt int = 100
	var myFloat float64 = 3.14
	var myBool bool = true
	fmt.Println(myString, myInt, myFloat, myBool)

	//Multiple variable declaration
	var (
		employeeId int = 35106
		firstName, lastName string = "Suryajith", "Sujith"
	)
	fmt.Println("Employee ID:", employeeId, "Name:", firstName, lastName)

	//Short variable declaration
	name:="Suryajith Sujith"
	age,salary,isCoder := 25, 100000.50, true
	fmt.Println("Name:", name, "Age:", age, "Salary:", salary, "Is Coder:", isCoder)

}

