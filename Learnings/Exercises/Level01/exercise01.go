// EXERCÍCIO 01
/* Utilizando o operador curto de declaração, atribua estes valores às variáveis
com os identificadores "x", "y", "z",
1. 42      2. "James Bond"     3. true
Demonstre os valores nestas variáveis:
Uma única declaração print e Múltiplas declarações print
*/

package main

import "fmt"

func main() {
	x := 42
	y := "James Bond"
	z := true

	fmt.Println("X: ", x)
	fmt.Println("Y: ", y)
	fmt.Println("Z: ", z)

	fmt.Println("Nome: ", y, "\nIdade: ", x, "\nHomem: ", z)
}
