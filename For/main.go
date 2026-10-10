package main

import "fmt"

func main() {

	var numero int
	fmt.Print("Digite um número: ")
	fmt.Scanf("%d", &numero)

	for i := 0; i <= numero; i++ {
		fmt.Println(i)
	}

}
