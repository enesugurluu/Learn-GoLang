package main

import "fmt"

func main() {

	//grade := 3
	// Switch

	// Case typeleri ile switchte kullandıgımız değişkenini type ı aynı olmak zorunda

	switch grade := 3; grade {
	case 5: //if grade == 5 { fmt.Println("Pekiyi") }
		fmt.Println("Pekiyi")

	case 4:
		fmt.Println("İyi")

	case 3:
		fmt.Println("Orta")

	case 2:
		fmt.Println("Geçer")

	case 1:
		fmt.Println("Başarısız")

	default:
		fmt.Printf("%v Geçersiz bir not", grade)
	}

}
