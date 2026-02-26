package main

import (
	"fmt"
	"math"
)

func main() {
	//Constant
	const pi = 3.1415
	fmt.Println(pi)

	radius := 3.0

	areaOfCircle := pi * (math.Pow(radius, 2))
	fmt.Println(areaOfCircle)

	const x int = 2
	const y float32 = 3.4
	const z string = "Enes"
	const t bool = false

	fmt.Printf("%T %v\n", x, x)
	fmt.Printf("%T %v\n", y, y)
	fmt.Printf("%T %v\n", z, z)
	fmt.Printf("%T %v\n", t, t)

	//Constantlara değer ataması sadece 1 kere yapılır

	//x = 20 //Error: cannot assign to x (neither addressable nor a map index expression)
	fmt.Println(x)

	//const --> compile time
	//var --> run time

	//const test = math.Pow(10, 2) //Error: math.Pow(10, 2) (value of type float64) is not constant
	//fmt.Println(test)

	//Bu hatanın sebebi, Go dilinde sabitlerin (const) derleme zamanında (compile time) kesin olarak bilinmek zorunda olmasıdır. math.Pow(10, 2) bir fonksiyon çağrısıdır.
	//Go'da paket fonksiyonları derleme aşamasında değil, program çalıştırıldığında, yani çalışma zamanında (run time) çalıştırılır.
	//Derleyici (compiler), kodu derlerken bu fonksiyonun sonucunu çalıştırıp hesaplayamadığı için, dönen değeri bir sabite atamanıza izin vermez.

	//const test2 //Error: Missing value in the const declaration
	//test2 = 10

	//Const olarak bir değişken tanımlamak istediğimizde direkt değerinin belirtmek durumundayız.

	//test3 := 5

	//const test4 = test3 //Error: test3 (variable of type int) is not constant
	//fmt.Println(test4)

	const ab = 1
	var yz = 2

	fmt.Printf("%T %v\n", ab+yz, ab+yz) //Output: int 3

	const ac, ad, ae = 1, 2, 3
	fmt.Printf("%T %v\n", ac, ac)
	fmt.Printf("%T %v\n", ad, ad)
	fmt.Printf("%T %v\n", ae, ae)
}
