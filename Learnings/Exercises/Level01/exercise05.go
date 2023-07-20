// EXERCÍCIO 04
/* Utilizando o exercício anterior

Em package-level scope, utilizando a palavra chave var, crie uma variável com o identificador "y".
O tipo desta variável deve ser o tipo subjacente do tipo que você criou no exercício anterior.

Na função main:
(Já está feito)
Demonstre o valor da variável "x"
Demonstre o tipo da variável "x"
Atribua 42 para a variável "x" utilizando o operador "="
Demonstre o valor da variável "x"
(Fazer)
Utilize conversão para transformar o tipo do valor da variável "x" em seu tipo subjacente e, utilizando
o operador "=", atribua o valor de "x" a "y"
Demonstre o valor de "y"
Demonstre o tipo de "y"
*/

package main

import "fmt"

// NAO PRECISA COLOCAR =
type bolodemurango int

var x bolodemurango
var y int

func main() {

	fmt.Println("X: ", x)
	fmt.Printf("Tipo de X: %T \n", x)
	x = 42
	fmt.Println("X: ", x)
	fmt.Printf("Tipo de X: %T\n", x)
	y = int(x)
	fmt.Println("Y: ", y)
	fmt.Printf("Tipo de Y: %T\n", y)
}
