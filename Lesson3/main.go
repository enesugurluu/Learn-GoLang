package main

import (
	"fmt"
	"strconv"
)

func main() {
	x := 10
	y := 10.0

	fmt.Printf("%v %T\n", x, x)
	fmt.Printf("%v %T\n", y, y)

	//fmt.Printf(x + y) //Error: invalid operation: x + y (mismatched types int and float64)

	//Type Conversion type(value) => int(y) --> 10.0 => 10
	fmt.Println(x + int(y)) // Output: 20

	fmt.Printf("%v %T\n", y, y) // Output: 10 float64 yani bir değişkeni type conversion yapmış olsak bile orijinal değişkenin türünde herhangi bir değişim olmaz
	// sadece yeni bir değer yaratılır ve o kullanılır type converison sırasında.

	var a int8 = 10
	var b int16 = 10

	//fmt.Println(a + b) //Error: invalid operation: a + b (mismatched types int8 and int16)
	fmt.Println(a + int8(b)) //Output: 20

	fmt.Println("\n--- Float mı Int'e, Int mi Float'a Dönüştürülmeli? ---")

	f := 15.8
	i := 5

	// DURUM 1: Float'ı Int'e Dönüştürmek (Veri Kaybı / Data Loss)
	// Go, float'ı int'e çevirirken yuvarlama (rounding) yapmaz, ondalık kısmı direkt kesip atar (truncation).
	// Bu yüzden 15.8 değeri 15 olur.
	// Sonuç: 15 + 5 = 20 (Gerçek sonuç olan 20.8'den saptık)
	fmt.Printf("Float'i Int'e çevirme sonucu: %v (Veri kaybı yaşandı)\n", int(f)+i)

	// DURUM 2: Int'i Float'a Dönüştürmek (Güvenli Yol / Daha Kesin Sonuç)
	// Int olan 5 değeri float64(5) yapılarak 5.0 gibi işleme sokulur.
	// Sonuç: 15.8 + 5.0 = 20.8
	fmt.Printf("Int'i Float'a çevirme sonucu: %v (Hassasiyet korundu)\n", f+float64(i))

	// int8: Sadece -128 ile 127 arasındaki değerleri alabilir (Çok dar kapasite)
	var kucukKutu int8 = 120

	// int32: -2.1 milyar ile +2.1 milyar arasındaki değerleri alabilir (Çok geniş kapasite)
	var buyukKutu int32 = 1000

	fmt.Println("--- Dar Türü Genişe Çevirmek (GÜVENLİ) ---")
	// kucukKutu'yu int32'ye çeviriyoruz. 120 sayısı int32'nin devasa kapasitesi içinde rahatça kaybolur.
	// Hiçbir veri kaybı olmaz.
	guvenliSonuc := int32(kucukKutu) + buyukKutu
	fmt.Printf("İşlem: %v + %v = %v\n", kucukKutu, buyukKutu, guvenliSonuc) // Çıktı: 1120

	fmt.Println("\n--- Geniş Türü Dara Çevirmek (TEHLİKELİ / TAŞMA) ---")
	// 1000 değerindeki buyukKutu'yu maksimum 127 alabilen int8'e zorla sığdırmaya çalışıyoruz.
	// 1000 sayısı kutuya sığmaz, limit aşılır (overflow) ve Go sayıyı başa sararak anlamsız negatif bir değere çevirir (-24 olur).
	tehlikeliSonuc := kucukKutu + int8(buyukKutu)

	// Beklediğimiz sonuç 1120'ydi ama çıkan sonuca bak:
	fmt.Printf("İşlem: %v + int8(%v) = %v (Veri tamamen bozuldu!)\n", kucukKutu, buyukKutu, tehlikeliSonuc) // Çıktı: 96

	c := 10
	d := "20"

	fmt.Printf("%v %T\n", c, c)
	fmt.Printf("%v %T\n", d, d)

	//fmt.Println(c + int(d)) // We cant convert string type to int type

	num1 := 106
	str1 := string(num1)

	fmt.Printf("%v %T\n", num1, num1) //Output: 106 int
	fmt.Printf("%v %T\n", str1, str1) //Output: j int

	// YANLIŞ YÖNTEM: Sayının Unicode karakter karşılığını verir ("j")
	yanlisCevrim := string(num1)
	fmt.Printf("Yanlış: %v %T\n", yanlisCevrim, yanlisCevrim) // Çıktı: j string

	// DOĞRU YÖNTEM: Sayıyı direkt metne çevirir ("106")
	// Itoa = Integer TO ASCII (Integer'dan String'e) anlamına gelir
	dogruCevrim := strconv.Itoa(num1)
	fmt.Printf("Doğru: %v %T\n", dogruCevrim, dogruCevrim) // Çıktı: 106 string

}
