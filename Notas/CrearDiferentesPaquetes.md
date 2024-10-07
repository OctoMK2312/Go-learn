
# Cómo crear funciones en diferentes archivos en Go

## 1._ Crear un módulo en Go

Primero, es necesario inicializar un módulo en Go. Esto es fundamental para gestionar las dependencias y la estructura de tu proyecto. Aquí te muestro cómo inicializar un módulo:

```bash
go mod init nombre-de-tu-modulo
```

Esto creará un archivo `go.mod` en tu directorio de trabajo, que ayudará a Go a gestionar las dependencias del proyecto.

---

## 2._ Crear los archivos con funciones

Ahora, necesitamos crear el archivo principal `main.go`, que contendrá la función principal, y otro archivo con las funciones externas.

### Ejemplo de la función principal en `main.go`

```go
package main

import (
    "fmt"
    "nombre-de-tu-modulo/externo"  // Importa el paquete "externo"
)

func main() {
    fmt.Println("Hola, soy la función principal")
    externo.MensajeExterno()  // Llamada a la función externa
}
```

En este caso, estamos importando el paquete `externo` (el archivo que crearemos a continuación) y usando su función `MensajeExterno`.

---

### Ejemplo de función externa en el archivo `externo.go`

```go
package externo

import "fmt"

// MensajeExterno es una función exportada que puede ser usada desde otro paquete
func MensajeExterno() {
    fmt.Println("Hola, soy un mensaje externo")
}
```

En este archivo, definimos el paquete `externo` y una función `MensajeExterno`. Al comenzar con una letra mayúscula, la función se exporta y puede ser utilizada en otros paquetes o archivos.

---

### 3._ Usar la función externa en la principal

En el archivo `main.go`, ya estamos llamando a `externo.MensajeExterno()`. Como `MensajeExterno` está exportada, se puede utilizar desde el paquete `externo` en el archivo `main.go`.

---

## 4._ Exportar Funciones

En Go, para que una función o variable sea accesible desde otro archivo o paquete, debe comenzar con una **letra mayúscula**. Esta convención es la manera de exportar funciones en Go. Cualquier función que comience con una letra minúscula será privada y solo podrá ser usada dentro del mismo paquete.

En el ejemplo anterior, la función `MensajeExterno` comienza con "M" mayúscula, lo que significa que está exportada y puede ser usada fuera del paquete `externo`.

---

## Estructura del Proyecto

Una vez que tengas todo listo, tu proyecto debería tener una estructura como esta:

```plainText
.
├── go.mod
├── main.go
└── externo.go
```

Si quisieras organizar mejor tu proyecto, podrías colocar `externo.go` en un subdirectorio llamado `externo` para crear un paquete más formal:

```painText
.
├── go.mod
├── main.go
└── externo
    └── externo.go
```

En este caso, la línea de importación en `main.go` cambiaría a:

```go
import "nombre-de-tu-modulo/externo"
```
