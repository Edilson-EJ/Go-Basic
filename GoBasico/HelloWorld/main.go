package main

import "fmt"

func main() {
	fmt.Println("Inteiros sem sinal")

	var u1 uint8 = 255
	var u2 uint16 = 65535
	var u3 uint32 = 4294967295
	var u4 uint64 = 18446744073709551615

	fmt.Println(u1, u2, u3, u4)

	fmt.Printf("Inteiros com sinal")

	var i1 int8 = 127
	var i2 int16 = 32767
	var i3 int32 = 2147483647
	var i4 int64 = 9223372036854775807

	fmt.Println("", i1)
	fmt.Printf("%d\n", i2)
	fmt.Printf("%d\n", i3)
	fmt.Printf("%d\n", i4)

	var i5 int = 2147483647
	fmt.Printf("%d\n", i5)

	var i6 uint = 9223372036854775807
	fmt.Printf("%d\n", i6)

	fmt.Println("Números de ponto flutuante")

	var f1 float32 = 3.14159
	fmt.Printf("%f\n", f1)

	var f2 float64 = 3.141592653589793
	fmt.Printf("%f\n", f2)

	var f3 complex64 = 1 + 2i
	fmt.Printf("%v\n", f3)

	fmt.Println("Strings")

	var s1 string = "Hello, World!"
	fmt.Printf("%s\n", s1)

	var s2 string = `This is a multi-line
string literal.`
	fmt.Printf("%s\n", s2)

	println("Inferência de tipo")

	NomeCompleto := "John Doe"
	fmt.Printf("%s\n", NomeCompleto)

	const Pi = 3.14159
	fmt.Printf("%f\n", Pi)

}
