/*

	// 1-) 1 ile 10 arasındaki sayıları if yapısıyla tek - çift olarak yazdırınız.
	package main

	import "fmt"

	func main() {

		for i := 1; i <= 10; i++ {
			if i%2 == 0 {
				fmt.Println(i, "is even")
			} else {
				fmt.Println(i, "is odd")
			}
		}

	}

*/
/*
	// 2-) for yapısını kullanarak Go'da olmayan while döngüsüne örnek veriniz.

	package main

	import "fmt"

	func main() {
		x := 0
		for x < 10 {
			fmt.Println(x)
			x++
		}
	}

*/
/*
// 3-) Switch fallthrough ifadesini açıklayınız.

package main

import "fmt"

func main() {
	switch x := 25; {
	case x < 20:
		fmt.Println("x is less than 20")
		fallthrough

	case x < 50:
		fmt.Println("x is less than 50")
		fallthrough

	case x < 100:
		fmt.Println("x is less than 100")
	}

}
*/

// 5-) 1 ile 50 arasındaki asal sayıları gösteren bir program yazınız.

package main

import "fmt"

func main() {
	for i := 1; i <= 50; i++ {
		isPrime := i != 1
		for j := 2; j <= i/2; j++ {
			if i%j == 0 {
				isPrime = false
				break
			}
			if i == 2 {
				isPrime = true
			}
		}

		if isPrime {
			fmt.Printf("%d is prime\n", i)
		} else {
			fmt.Printf("%d is not prime\n", i)
		}
	}
}
