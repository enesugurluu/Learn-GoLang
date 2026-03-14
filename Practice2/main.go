// 1 -) Iki rakam arasında toplama, çıkarma ve çarpma
// işleminin yapıldığı bir fonkiyon yazınız.
//
//package main
//
//import "fmt"
//
//func main() {
//	number1 := 12
//	number2 := 14
//	sum, subt, multiply := calc(number1, number2)
//
//	fmt.Println("Toplam", sum)
//	fmt.Println("Çıkartma", subt)
//	fmt.Println("Çarpma", multiply)
//}
//
//func calc(number1, number2 int) (int, int, int) {
//	sum := number1 + number2
//	subtraction := number1 - number2
//	multiply := number1 * number2
//	return sum, subtraction, multiply
//}

// 2 -) Kullanıcı tarafından girilen nota göre geçtiniz
// veya kaldınız geri dönüşünü yazdırınız.

//package main
//
//import "fmt"
//
//func main() {
//	var entered_grade float64
//
//	fmt.Print("Enter grade: ")
//	fmt.Scan(&entered_grade)
//
//	if entered_grade >= 50 {
//		fmt.Println("Geçtiniz")
//	} else {
//		fmt.Println("Kaldınız")
//	}
//}

// 3 -) 1 ile yüz arasındaki bir sayıyı tahmin etme uygulaması
// yazınız. Toplam tahmin hakkınız 10 olsun.

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	secret := rand.Intn(100)

	var userGuess int

	fmt.Println("Enter the your guess: ")
	fmt.Scan(&userGuess)

	for i := 1; i <= 10; i++ {
		if userGuess > secret {
			fmt.Println("You guess is too big")
			fmt.Printf("Enter the your guess again, remaining %v: ", 10-i)
			fmt.Scan(&userGuess)
		} else if userGuess < secret {
			fmt.Println("You guess is too small")
			fmt.Printf("Enter the your guess again, remaining %v: ", 10-i)
			fmt.Scan(&userGuess)
		} else {
			fmt.Println("Congratulations")
			break
		}
	}
}
