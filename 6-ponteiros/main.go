package main

import "fmt"

func main() {
	var x *int // declaração de um ponteiro
	y := 10
	x = &y                                  // x e y apontam para o mesmo endereço, ao alterar o valor de y, o valor de x também é alterado
	fmt.Println("Valor de y:", *x)          // 10
	fmt.Println("Valor do ponteiro x:", *x) // 10
	y++                                     // incrementando o valor de y
	fmt.Println("Valor de y:", *x)          // 11
	fmt.Println("Valor do ponteiro x:", *x) // 11
	// &: referenciação
	// *: desreferenciação
}
