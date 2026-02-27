package main

import "fmt"

func main() {
	//Typeless Constant

	const x int16 = 5 //Not Typeless Constant
	fmt.Printf("%T %v\n", x, x)

	const y = 5 // Typeless Constant
	fmt.Printf("%T %v\n", y, y)

	//var z int16 = 12
	//const w int8 = 20
	//fmt.Printf("%T %v\n", z+w, z+w) //Error: Invalid operation: z+w (mismatched types int16 and int8)

	const a = 3 //Typeless
	var b int8 = 4
	fmt.Printf("%T %v\n", a+b, a+b) // Output: int8 7
	//Normalde typeless olan a değişkeninin türü default olarak int atanmiş olmasına rağmen a+b işleminde farklı türlerle işlem yapmamıza rağmen hata almamamızın sebebi Go nun otomatik olarak type conversion uygulamsıdır. int8(a) + b

	const c = 3
	const d = 4.0

	fmt.Printf("%T %v\n", c, c)
	fmt.Printf("%T %v\n", d, d)
	//fmt.Printf("%T %v\n", c+d, c+d) Eğer farklı türde olurlarsa bu işlemi yapamayız
	fmt.Printf("%T %v\n", c+d, c+d) //Output: float64 7 Go type conversion yaparak uygun veri tipini atar

}
