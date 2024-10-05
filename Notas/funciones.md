# Funciones en go

Las funciones en **Go** es muy parecido a otros lenguajes ya que son fragmentos de codigo reutilizables, a continuación se muestra un ejemplo sencillo del uso de funciones en **Go**:

```go
package main
import "fmt"
    
func saludar("Nombre"){
    fmt.Printl("Hola ", nombre); 
    }

func main(){
    fmt.Println("Bienvenido");

    saludar("Juan");
    }
```

En este cado la salida del programa sera `Bienvenido. Hola Juan`
