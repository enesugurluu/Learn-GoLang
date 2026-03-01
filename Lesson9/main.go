package main

import "fmt"

func main() {
	//Comparison Operators
	//== -> Equal != -> Not Equal
	// < -> Less than > -> Greater than
	// <= Less than or Equal >= -> Greater than or Equal

	x, y := 1, 2

	fmt.Printf("%T, %v\n", x == y, x == y) // Output: bool, false
	fmt.Printf("%T, %v\n", x != y, x != y) // Output: bool, true
	fmt.Printf("%T, %v\n", x < y, x < y)   // Output: bool, true
	fmt.Printf("%T, %v\n", x <= y, x <= y) // Output: bool, true
	fmt.Printf("%T, %v\n", x > y, x > y)   // Output: bool, false
	fmt.Printf("%T, %v\n", x >= y, x >= y) // Output: bool, false

	//We can compare characters

	a, b := "a", "b" //a:97 b:98 as ASCII
	fmt.Println("-----A and B--------")
	fmt.Printf("%T, %v\n", a == b, a == b) // Output: bool, false
	fmt.Printf("%T, %v\n", a != b, a != b) // Output: bool, true
	fmt.Printf("%T, %v\n", a < b, a < b)   // Output: bool, true
	fmt.Printf("%T, %v\n", a <= b, a <= b) // Output: bool, true
	fmt.Printf("%T, %v\n", a > b, a > b)   // Output: bool, false
	fmt.Printf("%T, %v\n", a >= b, a >= b) // Output: bool, false

	//We cant compare different typed variables

	//Logical Operators
	//We can use them with boolean variables

	// &&			conditional AND			p && q 	is   "if p'then q else false"
	// ||			conditional OR			p || q 	is   "if p then true else q"
	// !			NOT						!p 		is   "not p"

	c, d := 15, 20
	set1 := c == d
	set2 := c < d

	fmt.Printf("%T, %v\n", set1, set1) // Output: bool, false
	fmt.Printf("%T, %v\n", set2, set2) // Output: bool, true

	fmt.Printf("%T, %v\n", set1 && set2, set1 && set2) // Output: bool, false
	fmt.Printf("%T, %v\n", set1 || set2, set1 || set2) // Output: bool, true
}
