// EXERCÍCIO 03
/* Utilizando a solução do exercício anterior:
Em package-level scope, atribua os seguintes valores às variáveis:
1. para "x" atribua 42
2. para "y" atribua "James Bond"
3. para "z" atribua true
Na função main:
Use fmt.Sprintf para atribuir todos esses valores a uma única variável.
Faça essa atribuição de tipo string a uma variável de nome "s" utilizando
o operador curto de declaração.
Demonstre a variável "s"*/

package main

import "fmt"

// NAO PRECISA COLOCAR =
var x = 42
var y = "James Bond"
var z = true

func main() {

	fmt.Println("X: ", x)
	fmt.Println("Y: ", y)
	fmt.Println("Z: ", z)

	s := fmt.Sprintf("%v \n %v \n %v", x, y, z)
	fmt.Printf("Tipo de s: %T\n", s)
	fmt.Println("Valor de s: ", s)
}
