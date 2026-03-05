package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

func main() {

	result, err := evenNumber(11)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

	result2, err2 := squareRoot(-5)
	if err2 != nil {
		fmt.Println(err2)
	} else {
		fmt.Printf("%v sayısının karakökü %v dir", result2*result2, result2)
	}

	file, err := os.Open("Lesson15/test.txt")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(file)
	}
}

func evenNumber(num int) (string, error) {
	if num%2 != 0 {
		return "", errors.New("HATA: Çift sayı girmediniz")
	}
	return "Çift sayı girdiniz", nil //nil başlangıç değeri olmayan ifadelere verilen isimdir
}

func squareRoot(num float64) (float64, error) {
	if num < 0 {
		return 0, errors.New("HATA: negative number")
	}
	return math.Sqrt(num), nil
}
