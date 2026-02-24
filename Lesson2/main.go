package main

import "fmt"

var packVar = "Package Scope" //Bu kısımda oluşturdugumuz değişkene dosya içerisinde istediğimiz yerden ulaşabilriz.
/*
packVar2 := "Package Scope2" --> Outside a function, every statement begins with a keyword (var, func, and so on)
and so the := construct is not available

syntax error: non-declaration statement outside function body
*/

/*
Peki neden tüm değişkenleri paket düşeyinde tanımlayıp kullanmıyoruz?
Çünkü bu değişkenler programımızın çalıştıgı her süre boyunca hafızada yer kaplarlar.
*/

// var funcVar = "Func (Package) Scope" // Eğer aynı isimde bir değişkeni bu sefer paket düzeyinde tanımlarsak.
// Scope içerisinde yine scope içinde tanımlanmış halini kullanır.

func main() {
	var funcVar = "Func Scope"
	//Bu main fonksiyonu içerisinde oluşturmuş oldugumuz değişkene sadece main fonkisyonu içerisinde ulaşabiliriz.

	if true {
		var blockVar = "Block Scope"
		//İf blogu içerisinde oluşturdugumuz bu değişkeni sadece bu blok içerisinde kullanabiliriz.
		//BU if blogunun dışarısında kullanamaya kalarsak "undefined: blockVar" ile karşılaşırız.
		fmt.Println(blockVar)
	}

	fmt.Println(funcVar)
	fmt.Println(packVar)
	// fmt.Println(blockVar) //Error undefined: blockVar

	var name = "Enes"
	// name := "Muhammed" // Error : No new variables on the left side of ':='

	name, surname := "Enes", "Uğurlu" //Eski olanı (name) sadece günceller. Yeni olanı (surname) ise sıfırdan oluşturur.
	fmt.Print(name, surname)

	myFunc()
}

func myFunc() {
	fmt.Println(packVar) // Output: Package Scope
	//fmt.Println(funcVar) // Error: undefined: funcVar

}
