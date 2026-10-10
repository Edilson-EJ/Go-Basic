package main

import (
	"fmt"
)

func main() {

	var idade int

	println("Digite sua idade: ")
	fmt.Scanf("%d", &idade)

	if idade >= 18 {
		fmt.Println("Você é maior de idade")
	} else {
		fmt.Println("Você é menor de idade")
	}

}
