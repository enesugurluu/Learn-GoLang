package main //Package Clause
//Burada yazacagimiz kod main paketi icerisinde

import "fmt" //Import Statement

//Go dilinin standart paketlerini veya baskalari tarafindan yazilmis paketleri import ederek kodumuz icerisinde kullanabiliriz.

// Kodumuz
func main() { //main: Fonksiyon, func: Go da fonskiyon oluşturma anahtar kelimesi

	var name string = "Enes" // Var - Unique name of variable - Static Variable Type

	var name2 string
	name2 = "Muhammed"

	var age int
	age = 22
	age = 21

	var isMarried bool
	fmt.Println(isMarried) //output is "false"

	isMarried = true

	var firstName, lastName string = "Enes", "Uğurlu"

	name3 := "Muhammed Enes Uğurlu" //Decleratıon and assıgnment operator ":="

	var test1 = "Test1"

	fmt.Println(age)

	fmt.Println(name)
	fmt.Println(name2)
	fmt.Println(isMarried)
	fmt.Println(firstName, lastName)
	fmt.Println(name3)
	fmt.Println(test1)

	fmt.Printf("%T\n", name)      //Bir değişkenin type ini yazdırmak istediğimizde "Printf" kullanırız ve "%T" özel anahtarını kullanırız.
	fmt.Printf("%T\n", isMarried) // "\n" kullanarak alt satıra geçiyoruz.
	fmt.Printf("%T\n", age)

	/*
		--- BASIC TYPES (BUILT-IN TYPES)
		NUMERIC TYPES

		unsigned -> 0 ve positif değerler için

		uint8               the set of all unsigned 8-bit integers (0 to 255)
		uint16              the set of all unsigned 16-bit integers (0 to 65535)
		uint32              the set of all unsigned 32-bit integers (0 to 4294967295)
		uint64              the set of all unsigned 64-bit integers (0 to 18446744073709551615)

		signed -> negatif, 0 ve positif değerler için

		int8                the set of all signed 8-bit integers (-128 to 127)
		int16               the set of all signed 16-bit integers (-32768 to 32767)
		int32               the set of all signed 32-bit integers (-2147483648 to 2147483647)
		int64               the set of all signed 64-bit integers (-9223372036854775808 to 9223372036854775807)

		float32             the set of all IEEE-754 32-bit floating-point numbers
		float64             the set of all IEEE-754 64-bit floating-point numbers

		complex64           the set of all complex numbers with float32 real and imaginary parts
		complex128          the set of all complex numbers with float64 real and imaginary parts

		byte                alias for uint8
		rune                alias for int32


	*/

	var test2 uint8 = 255
	fmt.Printf("%T\n", test2)

	var test3 uint16 = 65535
	fmt.Printf("%T\n", test3)

	var weight float32 = 77.5
	fmt.Printf("%T\n", weight)
	/*
		STRING TYPES
		A string type represents the set of string values. A string value is a (possibly empty) sequence of bytes.
		The number of bytes is called the length of the string and is never negative.
	*/

	var test4 string = "Test4"
	fmt.Printf("%T\n", test4)

	/*
		BOOLEAN TYPES
		A boolean type represents the set of Boolean truth values denoted by the predeclared constants true and false.
	*/

	var isSigned bool = true
	fmt.Printf("%T\n", isSigned)

	//Variables Declarations
	/*
		var test5 string = "Test5"
		var test6 = true
		test7 := true
	*/

	var ( //Var parantezine alarak değişkenleri tanımlayabiliriz.
		test8  string  = "Test8"
		test9  int     = 9
		test10 float32 = 9.0
		test11 bool    = true
	)

	fmt.Printf("%T\n", test8)
	fmt.Printf("%T\n", test9)
	fmt.Printf("%T\n", test10)
	fmt.Printf("%T\n", test11)

	var name4, age3, isMarried3, weight3, height3 = "Enes", 22, false, 174, 77

	name5, age4, isMarried4, weight4, height4 := "Enes", 22, false, 174, 77

	fmt.Println(name4, age3, isMarried3, weight3, height3)
	fmt.Println(name5, age4, isMarried4, weight4, height4)

	var name6 string
	var age5 int
	var weight6 float32
	var isMarried5 bool

	fmt.Println(name6)      //string --> "" Zero Values
	fmt.Println(age5)       //int --> 0 Zero Value
	fmt.Println(weight6)    //float --> 0 Zero Value
	fmt.Println(isMarried5) //bool --> false Zero Value
}

//Go dili derlenene bir dildir ve kod, build işlem sonrasında executable bir exe dosyasına çevrilir.
