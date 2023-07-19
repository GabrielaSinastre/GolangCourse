// EXERCÍCIO 02
/* Use var para declarar três variáveis. Elas devem ter package-level scope. Não atribua valores
a estas variáveis. Utilize os seguintes identificadores e tipos para estas variáveis

1. Identificador "x" deverá ter tipo int
2. Identificador "x" deverá ter tipo string
3. Identificador "x" deverá ter tipo bool
Na função main:
Demonstre cada identificador
O compilador atribuiu valores para essas variáveis. Como esses valores se chamam? */

package main

import "fmt"

// NAO PRECISA COLOCAR =
var x int
var y string
var z bool

func main() {

	fmt.Println("X: ", x)
	fmt.Println("Y: ", y)
	fmt.Println("Z: ", z)

	fmt.Println("X: ", y, "Y: ", x, "Z: ", z)
	fmt.Printf("%v \n %v \n %v", x, y, z)
	fmt.Println("Esses valores resultantes são o 'zero' de cada tipo!")
	fmt.Println("X é 0, Y é '' e Z é false")
}
