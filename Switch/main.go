package main

import (
	"fmt"
)

func main() {

	var mes int

	println("Digite o número do mês: ")
	fmt.Scanf("%d", &mes)

	switch mes {
	case 1:
		fmt.Println("Janeiro")
	case 2:
		fmt.Println("Fevereiro")
	case 3:
		fmt.Println("Março")
	}

}
