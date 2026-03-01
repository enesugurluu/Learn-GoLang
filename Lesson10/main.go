package main

import "fmt"

func main() {
	//if <boolean expression> { code }
	x := 27
	if x%2 == 0 {
		fmt.Printf("%v is even\n", x)
	}

	if x%2 == 1 {
		fmt.Printf("%v is odd\n", x)
	}

	if !false {
		fmt.Println("Mesaj yazdırılacak")
	}

	y := 51

	if y > 50 {
		fmt.Printf("%v is greater than 50\n", y)
	} else {
		fmt.Printf("%v is less than 50\n", y)
	}

	fmt.Println("Please enter a number")

	var entered_value int
	fmt.Scanf("%d", &entered_value)

	if entered_value < 0 {
		fmt.Printf("%v is negative\n", entered_value)
	} else if entered_value > 0 {
		fmt.Printf("%v is positive\n", entered_value)
	} else {
		fmt.Printf("%v is zero\n", entered_value)
	}

	//if <boolean expression> { code } else if <boolean expression> {code} else {else}

	a := 5
	if a > 5 {
		fmt.Printf("%v is greater than 5\n", a)

	} else {
		fmt.Printf("%v is less than 5\n", a)
	}

	fmt.Println(a) //Bu durumda hala a ya erişebiliyoruz

	if b := 10; b > 20 { //Eğer if statementi içerisinde iki tane ifade kullanacaksak ilkinden sonra ";" kullanmalıyız
		fmt.Printf("%v is greater than 20\n", b)
	} else {
		fmt.Printf("%v is less than 20\n", b)
	}

	//fmt.Println(b) //Error : undefined: b

	//Nested If
	if x := -5; x > 0 {
		fmt.Printf("%v is positive\n", x)
	} else {
		if x == 0 {
			fmt.Printf("%v is zero\n", x)
		} else {
			fmt.Printf("%v is negative\n", x)
		}
	}
}
