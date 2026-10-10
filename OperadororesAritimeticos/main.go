package main

import "fmt"

func main() {

	var numero1 int
	var numero2 int
	fmt.Println("Digite o primero número")
	fmt.Scanln(&numero1)
	fmt.Println("Digite o segundo número")
	fmt.Scanln(&numero2)
	fmt.Printf("%d + %d = %d \n", numero1, numero2, numero1+numero2)
	fmt.Printf("%d - %d = %d \n", numero1, numero2, numero1-numero2)
	fmt.Printf("%d * %d = %d \n", numero1, numero2, numero1*numero2)
	fmt.Printf("%d / %d = %d \n", numero1, numero2, numero1/numero2)
	fmt.Printf("%d %% %d = %d \n", numero1, numero2, numero1%numero2)

}
