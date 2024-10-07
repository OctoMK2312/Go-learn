package multiplicar

import "fmt"

func Multiplicar() {
	var tabla int
	var rango int
	var contador int

	fmt.Println("Está a punto de poder ver las tablas de multiplicar por un rango.")

	// Leer la tabla de multiplicar que desea ver
	fmt.Print("Ingresa la tabla de multiplicar que deseas ver: ")
	_, err := fmt.Scan(&tabla) // Usar Scan para leer un solo valor
	if err != nil {
		fmt.Println("Error al leer la tabla:", err)
		return
	}

	// Leer el valor inicial para el contador
	fmt.Print("Ingresa donde deseas que inicien las tablas de multiplicar: ")
	_, err = fmt.Scan(&contador) // Usar Scan para leer un solo valor
	if err != nil {
		fmt.Println("Error al leer el valor inicial del contador:", err)
		return
	}

	// Leer el rango hasta donde llegarán las tablas
	fmt.Print("Ingresa hasta donde deseas que lleguen las tablas de multiplicar: ")
	_, err = fmt.Scan(&rango) // Usar Scan para leer un solo valor
	if err != nil {
		fmt.Println("Error al leer el rango:", err)
		return
	}

	// Generar la tabla de multiplicar desde el contador hasta el rango
	for contador <= rango {
		resultado := tabla * contador
		fmt.Println(tabla, "*", contador, "=", resultado)
		contador++
	}
}
