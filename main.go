//Aprendiendo la estructura básica de un programa en Go y su ejecución en la terminal

package main //Declaración del paquete principal

import (
	"fmt" //Importación de la librería fmt

	multiplicar "Aprendiendo_Go/Operaciones" //Importación del paquete multiplicar
)

func printVariables() {
	var nombre string = "Miguel" //Declaración de una variable de tipo string
	//De igual manera se pueden declarar variables con go infiriendo el tipo de dato
	edad := 21 //Declaración de una variable de tipo entero

	fmt.Println("Mi nombre es", nombre, "y tengo", edad, "años") //Impresión de los valores de las variables
}

func controlFlujo() {
	//Preguntar edad al usuario y almacenarla en la variable edad
	fmt.Println("Ingresa tu edad:")
	var edad int
	fmt.Scanf("%d", &edad)

	//Usamos el flujo condicional if para determinar si el usuario es mayor de edad
	if edad >= 18 {
		fmt.Println("Eres mayor de edad")
	} else {
		fmt.Println("Eres menor de edad")
	}
}

func ciclos() {
	//Ciclo for
	for i := 0; i < 5; i++ {
		fmt.Println("Iteración", i)
	}

	//Ciclo while
	j := 0
	for j < 5 {
		fmt.Println("Iteración", j)
		j++
	}

	//Ciclo infinito
	k := 0
	for {
		fmt.Println("Iteración", k)
		k++
		if k == 5 {
			break
		}
	}
}

func main() { //Función principal del programa
	//Declaración de variables locales
	var opcion int

	fmt.Println("Bienvenido al programa en Go") //Impresión de mensaje en consola

	fmt.Println("A continuación se mostrarán una serie de opciones, elige una para continuar:") //Impresión de mensaje en consola

	fmt.Println("1.- Variables") //Impresión de mensaje en consola

	fmt.Println("2.- Control de flujo") //Impresión de mensaje en consola

	fmt.Println("3.- Ciclos") //Impresión de mensaje en consola

	fmt.Println("4.- Multiplicar") //Impresión de mensaje en consola

	fmt.Println("Ingresa el número de la opción que deseas ejecutar:") //Impresión de mensaje en consola

	fmt.Scanf("%d", &opcion) //Lectura de la opción ingresada por el usuario

	switch opcion { //Estructura de control switch
	case 1: //Caso 1
		printVariables() //Llamado a la función printVariables
	case 2: //Caso 2
		controlFlujo() //Llamado a la función controlFlujo
	case 3: //Caso 3
		ciclos() //Llamado a la función ciclos
	case 4: //Caso 4
		multiplicar.Multiplicar() //Llamado a la función Multiplicar del paquete multiplicar
	default: //Caso por defecto
		fmt.Println("Opción no válida") //Impresión de mensaje en consola
	}
} //Cierre de la función principal
