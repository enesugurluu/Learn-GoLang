package main

import "fmt"

func main() {
	//(Print - Println) - Printf

	/*
		fmt.Print("Test1")
		fmt.Println("Test2")
		fmt.Printf("Test3")
	*/

	/*
		--- Go'da Print, Println ve Printf Arasındaki Farklar ---

		1. fmt.Print
		   - Verilen değerleri olduğu gibi, yan yana ekrana yazar.
		   - İşlemi bitirdikten sonra otomatik olarak alt satıra geçmez.
		   - Birden fazla string değer yazdırıldığında aralarına otomatik boşluk eklemez.
		   - Örnek:
		     fmt.Print("Merhaba", "Dünya")
		     // Çıktı: MerhabaDünya

		2. fmt.Println (Print Line)
		   - Verilen değerleri yazar ve işlemin en sonuna otomatik olarak bir alt satıra geçme karakteri (\n) ekler.
		   - Birden fazla değer verildiğinde, değerlerin arasına otomatik olarak birer boşluk koyar.
		   - Örnek:
		     fmt.Println("Merhaba", "Dünya")
		     // Çıktı: Merhaba Dünya (ve ardından alt satıra geçer)

		3. fmt.Printf (Print Format)
		   - Çıktıyı belirli bir formata göre biçimlendirerek ekrana yazdırmak için kullanılır.
		   - İçerisinde yer tutucular veya format belirleyiciler (verbs) kullanılır. Örneğin: %s (string), %d (tam sayı), %v (değerin varsayılan hali).
		   - Print gibi, otomatik olarak alt satıra geçmez. Alt satıra geçmek istiyorsanız format metninin sonuna manuel olarak \n eklemeniz gerekir.
		   - Örnek:
		     isim := "Ali"
		     yas := 30
		     fmt.Printf("Benim adım %s ve ben %d yaşındayım.\n", isim, yas)
		     // Çıktı: Benim adım Ali ve ben 30 yaşındayım. (ve ardından alt satıra geçer)
	*/

	name := "Enes"

	fmt.Print("Benim adım", name)
	fmt.Println("Benim adım", name)
	fmt.Printf("Benim adım %v ve değişken türü %T", name, name)

	/*
		--- Go'da Format Belirleyiciler (Verbs) ---

		Go'da Printf, Sprintf vb. fonksiyonlarda değerlerin ekrana veya belleğe
		nasıl yansıtılacağını belirleyen yer tutuculara "verb" denir.

		1. Genel ve Veri Tipi Belirleyiciler (Hata Ayıklama İçin İdeal)
		   %v  : Değerin varsayılan (default) hali. Örn: {Ali 30}
		   %+v : Struct (yapı) yazdırırken alan isimlerini de ekler. Örn: {Ad:Ali Yas:30}
		   %#v : Değerin Go kodundaki tam sözdizimsel karşılığını gösterir. Örn: main.Kullanici{Ad:"Ali", Yas:30}
		   %T  : Verinin tipini (Type) yazdırır. Örn: main.Kullanici, int, float64

		2. Metinler ve Mantıksal Değerler
		   %s  : Düz metin (string) veya byte slice. Örn: Merhaba
		   %q  : Metni çift tırnak içine alarak yazdırır. Örn: "Merhaba"
		   %c  : Sayının karakter (Unicode) karşılığı. Örn: 65 sayısı için A
		   %t  : Mantıksal (boolean) değer. Örn: true, false

		3. Sayılar (Tam Sayı ve Ondalıklı)
		   %d  : 10'luk tabanda (Base 10) tam sayı. Örn: 42
		   %b  : 2'lik tabanda (Binary) tam sayı. Örn: 42 için 101010
		   %x  : 16'lık tabanda (Hexadecimal) sayı (küçük harf). Örn: 2a
		   %X  : 16'lık tabanda (Hexadecimal) sayı (büyük harf). Örn: 2A
		   %f  : Ondalıklı sayı. (Not: %.2f şeklinde kullanılarak virgülden sonraki basamak sayısı kısıtlanabilir. Örn: 3.14)

		4. Bellek Yönetimi (Pointer)
		   %p  : Değişkenin bellekteki adresi. (Başına otomatik 0x eklenir). Örn: 0xc000018030

		Not: Çıktıyı ekrana yazdırmak yerine bir metin (string) değişkeni olarak kaydetmek isterseniz 'fmt.Sprintf' kullanabilirsiniz.
	*/

}
