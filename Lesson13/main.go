package main

import "fmt"

func main() {
	//Functions
	//Fonskiyonlar belirli bir işlemi yapmak için kullanılan kod blocklarıdır.

	var x, y, sum int

	x = 10
	y = 5
	sum = x + y

	fmt.Printf("%d ile %d nin toplamı %d\n", x, y, sum)

	x = 12
	y = 7
	sum = x + y

	fmt.Printf("%d ile %d nin toplamı %d\n", x, y, sum)

	//Fonskiyonları programımızı daha moduler ve daha okunabilir hale getirmek için kullanırız.

	fmt.Println(sumFunc(10, 5)) // Output : 15
	hello()
	sumFunc2(21, 45)
	sumFunc2(13, 2)
	sumFunc2(63, 72)

	hello2("Enes", 22) //Argümanlar fonskiyon çalıştırılırken kullanılanlardır.

	//func function name(params) returnType { code }

	//Fonskiyon isimlendirme kuralları

	//1-)İlk karakter harf olmalı
	//2-)Camel case kullanılır, myBestFunction, mySumFunction
	//3-)Eğer paket dışında da kullanacaksak fonskiynonun ilk harfi büyük olmalı

}

func sumFunc(x, y int) int {
	return x + y
}

func sumFunc2(x, y int) {
	fmt.Println(x, y)
}

func hello() {
	fmt.Println("hello")
}

func hello2(name string, age int) { //Parametre fonskiyon yazarken kullanılır.
	fmt.Printf("Name: %s, Age: %d \n", name, age)
}

func result(grade int) string { //We can use the parametres only in function scope
	if grade >= 50 {
		return "Geçtiniz"
		fmt.Println("Hello") //Return line nin altında yine scope içinde bir kod yazmak anlamsızdır çünkü hiçbir zaman çalışmayacaktır.
	}
	return "Kaldınız"
}
