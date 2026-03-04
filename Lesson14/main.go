package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	var x int = 10
	fmt.Println(x)

	var moment time.Time = time.Now() // Now() time paketine ait bir method
	fmt.Println(moment)

	fmt.Println("Lütfen sınav sonucunuzu giriniz. :")
	reader := bufio.NewReader(os.Stdin)
	value, _ := reader.ReadString('\n') // _ --> Blank identifier

	fmt.Println(value)

	bolum, kalan := bolme(104, 5)
	fmt.Println(bolum, kalan)
}

func bolme(bolunen int, bolen int) (bolum, kalan int) {
	bolum = bolunen / bolen
	kalan = bolunen % bolen
	return bolum, kalan
}
