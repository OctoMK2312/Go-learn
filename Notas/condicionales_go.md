
# Condicionales en Go

En Go, los condicionales son similares a los de otros lenguajes de programación, aunque con algunas diferencias en la sintaxis. Vamos a revisar los tipos de condicionales más comunes en Go:

---

## Índice de Condicionales en Go

1. [If](#1-if)
2. [If-Else](#2-if-else)
3. [If-Else If-Else](#3-if-else-if-else)
4. [If con declaración corta](#4-if-con-declaración-corta)
5. [Switch](#5-switch)
    - [Con valor](#5-switch)
    - [Sin valor (Switch condicional)](#6-switch-sin-valor-switch-condicional)
    - [Múltiples valores en un solo caso](#7-switch-con-múltiples-valores-en-un-solo-caso)
    - [Uso de `fallthrough`](#8-uso-de-fallthrough-en-switch)
6. [Condicional Ternario en Go (alternativa)](#9-condicional-ternario-en-go-alternativa)

## 1. If

La estructura básica del condicional `if` en Go no necesita paréntesis alrededor de la condición, pero el bloque de código asociado debe estar entre llaves `{}`.

```go
package main

import "fmt"

func main() {
    x := 10

    if x > 5 {
        fmt.Println("x es mayor que 5")
    }
}
```

## 2. If-Else

El bloque `else` se usa para manejar el caso contrario cuando la condición no se cumple.

```go
package main

import "fmt"

func main() {
    x := 3

    if x > 5 {
        fmt.Println("x es mayor que 5")
    } else {
        fmt.Println("x es menor o igual a 5")
    }
}
```

## 3. If-Else If-Else

Puedes usar `else if` para agregar más de una condición en la misma estructura.

```go
package main

import "fmt"

func main() {
    x := 8

    if x > 10 {
        fmt.Println("x es mayor que 10")
    } else if x > 5 {
        fmt.Println("x es mayor que 5 pero menor o igual a 10")
    } else {
        fmt.Println("x es menor o igual a 5")
    }
}
```

## 4. If con declaración corta

Go permite declarar una variable dentro de la estructura del `if`. Esta variable solo estará disponible dentro del bloque `if` y sus subsecuentes `else if` o `else`.

```go
package main

import "fmt"

func main() {
    if x := 5; x > 3 {
        fmt.Println("x es mayor que 3")
    }
    // No puedes usar 'x' aquí, ya que solo existe dentro del bloque if
}
```

## 5. Switch

El `switch` en Go es una alternativa más limpia al uso de múltiples `if-else if`. No necesitas usar `break` para salir del `switch`, ya que Go lo hace automáticamente. El `switch` evalúa el valor de una variable y ejecuta el caso que coincida.

```go
package main

import "fmt"

func main() {
    x := 2

    switch x {
    case 1:
        fmt.Println("x es 1")
    case 2:
        fmt.Println("x es 2")
    case 3:
        fmt.Println("x es 3")
    default:
        fmt.Println("x no es ni 1, ni 2, ni 3")
    }
}
```

## 6. Switch sin valor (Switch condicional)

Puedes usar `switch` sin un valor explícito. Esto permite usar condicionales en los casos, lo que es muy útil para evaluar rangos de valores.

```go
package main

import "fmt"

func main() {
    x := 8

    switch {
    case x < 5:
        fmt.Println("x es menor que 5")
    case x > 5 && x < 10:
        fmt.Println("x es mayor que 5 pero menor que 10")
    default:
        fmt.Println("x es mayor o igual a 10")
    }
}
```

## 7. Switch con múltiples valores en un solo caso

Un solo caso en un `switch` puede manejar múltiples valores.

```go
package main

import "fmt"

func main() {
    x := 4

    switch x {
    case 1, 2, 3:
        fmt.Println("x es 1, 2 o 3")
    case 4, 5:
        fmt.Println("x es 4 o 5")
    default:
        fmt.Println("x es otro valor")
    }
}
```

## 8. Uso de `fallthrough` en Switch

Por defecto, Go no cae automáticamente al siguiente caso como en otros lenguajes como C. Sin embargo, puedes usar la palabra clave `fallthrough` si quieres que continúe evaluando el siguiente caso.

```go
package main

import "fmt"

func main() {
    x := 2

    switch x {
    case 1:
        fmt.Println("x es 1")
    case 2:
        fmt.Println("x es 2")
        fallthrough
    case 3:
        fmt.Println("x es 3")
    default:
        fmt.Println("Valor no identificado")
    }
}
```

## 9. Condicional Ternario en Go (alternativa)

Go no tiene un operador ternario como en otros lenguajes (ej. `condicion ? valor1 : valor2`), pero puedes lograr algo similar con una función simple o usando `if-else` en una línea:

```go
package main

import "fmt"

func main() {
    x := 10
    resultado := "Es mayor" 
    if x < 5 {
        resultado = "Es menor"
    }
    fmt.Println(resultado)
}
```

Esto es equivalente a un operador ternario en otros lenguajes.
