package main
import "fmt"

func main() {
	x := 10
	y := 20
	fmt.Println("x + y =", x + y) // Addition
	fmt.Println("x - y =", x - y) // Subtraction
	fmt.Println("x * y =", x * y) // Multiplication
	fmt.Println("x / y =", x / y) // Division
	fmt.Println("x % y =", x % y) // Modulus
	fmt.Println("x & y =", x & y) // Bitwise AND
	fmt.Println("x | y =", x | y) // Bitwise OR
	fmt.Println("x ^ y =", x ^ y) // Bitwise XOR
	fmt.Println("x << 2 =", x << 2) // Left Shift
	fmt.Println("x >> 2 =", x >> 2) // Right Shift
	fmt.Println("x == y:", x == y) // Equality
	fmt.Println("x != y:", x != y) // Inequality
	fmt.Println("x < y:", x < y)   // Less than
	fmt.Println("x <= y:", x <= y) // Less than or equal to
	fmt.Println("x > y:", x > y)   // Greater than
	fmt.Println("x >= y:", x >= y) // Greater than or equal to
	fmt.Println("x && y:", (x > 5) && (y < 30)) // Logical AND
	fmt.Println("x || y:", (x < 5) || (y > 15)) // Logical OR
	fmt.Println("!x:", !(x > 5)) // Logical NOT
	fmt.Println("x++ =", x + 1) // Increment (not a built-in operator
	fmt.Println("x-- =", x - 1) // Decrement (not a built-in operator)
	fmt.Println("x += y =", x + y) // Addition assignment
	fmt.Println("x -= y =", x - y) // Subtraction assignment
	fmt.Println("x *= y =", x * y) // Multiplication assignment
	fmt.Println("x /= y =", x / y) // Division assignment
	fmt.Println("x %= y =", x % y) // Modulus assignment
	fmt.Println("x &= y =", x & y) // Bitwise AND assignment
	fmt.Println("x |= y =", x | y) // Bitwise OR assignment
	fmt.Println("x ^= y =", x ^ y) // Bitwise XOR assignment
}