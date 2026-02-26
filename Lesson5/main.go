package main

var x = 10
var Y = 20 //Değişken ismini büyük harf ile
// başlatarak diğer paketler tarafından kullanılabilir hale getirdik.

func main() {

	//Variable Naming

	/*
		--- Go'da Değişken İsimlendirme Kuralları (Variable Naming Conventions) ---

		1. CamelCase (MixedCaps) Standardı:
		   - Go'da kelimeleri ayırmak için alt çizgi (snake_case) KULLANILMAZ.
		   - Bunun yerine kelimeler bitişik yazılır ve yeni kelimenin ilk harfi büyük olur (CamelCase).
		   - Doğru : kullaniciAdi, maksimumDeger
		   - Yanlış: kullanici_adi, maksimum_deger

		2. Görünürlük (Visibility - Dışa Aktarma Mantığı):
		   - Go'da bir değişkenin diğer paketlerden erişilebilir (Public) olup olmadığını İLK HARFİ belirler.
		   - Büyük Harfle Başlarsa (Exported): Diğer paketlerden erişilebilir. API veya kütüphane yazarken dışarı açmak istediğiniz değişkenlerde kullanılır.
		     Örn: KullaniciAdi, ToplamTutar, Sayac
		   - Küçük Harfle Başlarsa (Unexported): Private gibidir. Sadece tanımlandığı paketin (package) içinde kullanılabilir, dışarıdan erişilemez.
		     Örn: kullaniciAdi, toplamTutar, sayac

		3. Kısaltmaların (Acronyms) Doğru Yazımı:
		   - Yaygın kısaltmalar (ID, HTTP, URL, JSON, API vb.) kendi içinde tutarlı olarak ya tamamen büyük ya da tamamen küçük yazılır.
		   - Doğru : userID, HTTPClient, jsonParser, requestURL
		   - Yanlış: userId, HttpClient, JsonParser, requestUrl

		4. Kapsam ve İsim Uzunluğu (Scope & Length):
		   - Dar kapsamlarda (kısa fonksiyonlar, döngüler vb.): Kısa ve tek harfli isimler tercih edilir.
		     Örn: index için 'i', hata (error) için 'err', value için 'v'.
		   - Geniş kapsamlarda (paket seviyesindeki global değişkenler): Ne işe yaradığını açıkça belli eden, daha uzun ve açıklayıcı isimler kullanılır.
		     Örn: dbConnectionPool, defaultTimeoutValue
	*/

}
