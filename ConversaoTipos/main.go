package main

import (
	"fmt"
	"strconv"
)

func main() {
	var texto string
	fmt.Println("Digite um texto")
	fmt.Scanln(&texto)
	fmt.Println("O texto digitado foi: ", texto)

	numero, err := strconv.Atoi(texto)
	if err != nil {
		fmt.Println("Erro ao converter para inteiro")
	} else {
		fmt.Println("O número digitado foi: ", numero)
	}

	numero2, _ := strconv.ParseInt(texto, 10, 64)
	fmt.Println("O número digitado foi: ", numero2)

}
