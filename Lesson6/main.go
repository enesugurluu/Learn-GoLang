package main

import "fmt"

func main() {
	x, y := 20, 15

	fmt.Printf("Tür:%T Değer:%v\n", x, x) //Output: int 10
	fmt.Printf("Tür:%T Değer:%v\n", y, y) //Output: int 20

	fmt.Printf("Tür:%T Değer:%v\n", (x + y), (x + y)) //Output: int 30
	fmt.Printf("Tür:%T Değer:%v\n", (x - y), (x - y)) //Output: int -10
	fmt.Printf("Tür:%T Değer:%v\n", (x * y), (x * y)) //Output: int 200
	fmt.Printf("Tür:%T Değer:%v\n", (x / y), (x / y)) //Output: int 0 x = int, y = int (x/y) = int normalde 0.5 olması lazım ama sonuc int olması gerektiği için ondalıklı kısmı atılır.

	fmt.Printf("Tür:%T Değer:%v\n", (x % y), (x % y))

	z := 5.0 / 2 //float / int => float

	fmt.Printf("Tür:%T Değer:%v", z, z) //Output: float64 2.5

	//Eğer float bir değişkeni, int bir değişkene bölmeye kalarsak hatayla karşılaşırız.

	//İncreament (++) and Decreament (--), just postfix

	//İncreament
	a := 10

	a = a + 1

	fmt.Println("")

	fmt.Println(a)

	a++

	fmt.Println(a)

	//Decreament
	b := 10

	fmt.Println("")
	fmt.Println(b)

	b--

	fmt.Println(b)

}
