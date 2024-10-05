# Manejo de variables en GO

En go hay varias maneras de poder declarar variables por ejemplo:

```go
package main

import "fmt"

func main() {
    var nombre string = "Juan"
    edad := 25 // Go infiere el tipo
    fmt.Println("Nombre:", nombre, "Edad:", edad)
}
```

- ```var``` se usa para declarar variables explisitamente

- ```:=``` es una forma corta que permite a go inferir el tipo de dato de la variable
