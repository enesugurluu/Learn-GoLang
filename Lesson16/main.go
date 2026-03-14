package main

//Kendi paketlerimiz ve go nun kendi eklediği paketler C:/Users/<pc_adi>/go/src içeriisnde bulunur.
//Ben bu src klasörünün içine "merhaba" diye bir kalsör daha oluşturdum şu an.

/*
	Ve içerisinde bu kodu yazdım. Yazmış oldugum fonksiyonların baş harfleri büyük oldugun için diğerleri tarafından erişilenilir.

package merhaba

import "fmt"

	func Hola() {
		fmt.Println("Hola")
	}

	func Hello() {
		fmt.Println("Hello")
	}

	func Merhaba() {
		fmt.Println("Merhaba")
	}
*/

import "merhaba" //Paket isimleri her zaman küçük harfle oluşturulmalıdır.

func main() {
	merhaba.Hola()
	merhaba.Hello()
	merhaba.Merhaba()

}
