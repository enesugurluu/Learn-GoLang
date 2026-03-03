package main

import "fmt"

func main() {
	for x := 0; x <= 10; x++ {
		fmt.Println(x)
	}

	//for <init statement>; <condition>; <post statement> { code }

	//We don't have to declare in for loop for usign in for loop
	y := 0
	for ; y <= 10; y++ {
		fmt.Println(y)
	}

	//for { //İnfinite Loop
	//	fmt.Println("Benim adım enes")
	//}

	//for i := 0; true; i++ { //İnfinite Loop because of the condition always true
	//	fmt.Println(i)
	//}

	//z := 10
	//for z >= 0 {
	//	fmt.Println(z)
	//	z--
	//}

	//for i := 0; i <= 10; i++ {
	//	if i%3 == 0 {
	//		continue //Return the head of the for loop
	//	}
	//	fmt.Println(i)
	//}

	for i := 0; i <= 10; i++ {
		if i == 3 {
			break //Return the end of the for loop
		}
		fmt.Println(i)
	}
}
