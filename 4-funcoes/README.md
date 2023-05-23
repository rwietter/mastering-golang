## Funções

A função pode ter zero ou mais argumentos. O tipo vem após o nome da variável.

```go
package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
```

### Funções continuação

Quando dois ou mais parâmetros nomeados consecutivos de uma função compartilham o mesmo tipo, você pode omitir o tipo de todos menos o último. Neste exemplo, vamos encurtar

`x int, y int = x, y int`


## Resultados Múltiplos

Uma função pode retornar qualquer número de resultados.

```go
package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	a, b := swap("hello", "world")
	fmt.Println(a, b)
}
```
