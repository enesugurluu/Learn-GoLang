# Go Çalışma Mantığı ve Build Alma

## Go Nasıl Çalışır?
Go, derlenen statik tipli bir dildir. Yazdığınız kod doğrudan makine koduna derlenir, bu yüzden sanal makineye ihtiyaç duymaz ve çok hızlı çalışır.

## Build Alma İşlemi
Go projelerini derlemek ve çalıştırılabilir hale getirmek için `go build` komutu kullanılır.

* **go run main.go:** Kodu derlemeden (geçici olarak) doğrudan çalıştırır. Geliştirme aşamasında kullanılır.
* **go build main.go:** İşletim sisteminize uygun çalıştırılabilir bir dosya oluşturur. Bu dosyayı sunucuya atıp doğrudan çalıştırabilirsiniz.