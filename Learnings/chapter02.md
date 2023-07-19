# CAPÍTULO O1 - Variáveis, Valores e Tipos

## Notação importante:
  pacote.identificador
  *Isso não muda

  nil == nada

  func Println (a ...interface{}) (n int, err error)
    func Println (a ...interface{}) => significa que pode passar ilimitados desse
      tipo, posso colocar quantos elementos eu quiser, de qualquer tipo que eu quiser
    (n int, err error) => o que a função retorna (nesse caso um int e um error)
  
## Variáveis não podem ser declaradas e não utilizadas
  Se o valor pode ser descartado, pode colocar _,
    ex:
      numerodebytes, errors := fmt.Println("Hello world!", "Oi galera", 100)
        return será: Hello world! Oi galera! 100 28 <nil>

      _, errors := fmt.Println("Hello world!", "Oi galera", 100)
        return será: Hello world! Oi galera! 100
        28 não vai aparecer pois é descartado

  ### Tipos de Variáveis
    x := 16 -> int
    y := "strings" -> texto
    z := true -> boolean

## Operador curto de declaração
  := parece uma 'marmota (gopher)' ou o punisher
  usado para declarar variáveis (tipos)
  *tipagem automatica  
    := DECLARAR variáveis
    !==
    = ATRIBUIR valores para variáveis
  *não pode usar := fora do escopo (fora docode block)

## Palavra-chave var
  fora do escopo precisa usar var já que não pode usar :=

## Tipos
  ### int, string, bool
  base de tudo
  são estáticos
  pode ser deduzido, ex: var x = 10.5 -> compilador entende que é um float
  se ao criar uma var fora do escopo e não atribuir valor, só consigo atribuir valor dentro de uma função depois
  tipos compostos: slice, array, struct, map

## Valor zero
  exemplo da caixa postal
  valor presente na variável antes de ela ser inicializada
    os zeros em cada tipo:
      ints: 0
      float: 0.0
      bool:s false
      string: ""
      pointers, functions, interfaces, slices, channels, maps: nil

## Package fmt
  Print -> imprime o resultado/mensagem
  Sprint -> vai retornar uma string com o valor, sem imprimir a mensagem
  Fprint -> Fileprint, escrever em arquivo (não necessariamente arquivo, mas interface writer por exemplo)

## Criar meu próprio tipo
  ### Type is life
  type hotdog int
  var b hotdog = 101 

  func main() {
    fmt.Printf("%T", b)
  }

  O que foi feito: 
    foi criado um tipo hotdog que possui como tipo subjacente (tipo que vem por trás como base) int
    criou uma var do tipo hotdog
  
  b não pode receber nenhum valor que não seja hotdog, mesmo que seu tipo adjacente seja int (não tem equivalência)

  ### Conversions
  para pegar o b (hotdog) e passar para um int, teria que fazer o seguinte:
   
  type hotdog int
  var b hotdog = 101

  func main() {
    x = int(b)
    fmt.Printf("%T", b)
  }