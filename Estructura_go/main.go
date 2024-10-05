//Aprendiendo la estructura básica de un programa en Go y su ejecución en la terminal

package main //Declaración del paquete principal

import "fmt" //Importación de la librería fmt para la impresión de mensajes en consola

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

	fmt.Println("Hola mundo") //Impresión del mensaje "Hola mundo" en consola

	printVariables() //Llamado a la función printVariables

	controlFlujo() //Llamado a la función controlFlujo

	ciclos() //Llamado a la función ciclos

} //Cierre de la función principal
