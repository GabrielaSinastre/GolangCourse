// EXERCÍCIO 04
/*
- Crie um tipo. O tipo adjacente deve ser int.
- Crie uma variável para este tipo, com o identificador "x", utilizando a palavra chave var.

Na função main:
Demonstre o valor da variável "x"
Demonstre o tipo da variável "x"
Atribua 42 para a variável "x" utilizando o operador "="
Demonstre o valor da variável "x"
*/

package main

import "fmt"

// NAO PRECISA COLOCAR =
type bolodemurango int

var x bolodemurango

func main() {

	fmt.Println("X: ", x)
	fmt.Printf("Tipo de X: %T \n", x)
	x = 42
	fmt.Println("X: ", x)
	fmt.Printf("Tipo de X: %T", x)

}
