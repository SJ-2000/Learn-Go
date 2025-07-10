package main

import "fmt"

func main() {
	var name string
	var age int

	fmt.Print("Enter name: ")
	fmt.Scanln(&name)
	//The & is used to pass the address of the variable (kind of like pointers in C).

	fmt.Print("Enter age: ")
	fmt.Scanln(&age)

	fmt.Println("Hello", name, "you are", age, "years old.")
}