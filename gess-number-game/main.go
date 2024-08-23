package main

import (
	"fmt"
	"math/rand"
)

func main() {
	// if else
	// inicializo variable en el if, y realizo la condicion despues del ;
	// if t := time.Now(); t.Hour() < 12 {
	// 	fmt.Println("Mañana")
	// } else if t.Hour() < 17 {
	// 	fmt.Println("Tarde")
	// } else {
	// 	fmt.Println("Noche")
	// }

	// Con el switch el ejercicio anterior
	// switch t := time.Now(); {
	// case t.Hour() < 12:
	// 	fmt.Println("Mañana")
	// case t.Hour() < 17:
	// 	fmt.Println(t.Hour(), "Tarde")
	// default:
	// 	fmt.Println("Noche")
	// }

	// switch para validar en que SO se esta ejecutando GO (runtime.GOOS)
	// switch os := runtime.GOOS; os {
	// case "windows":
	// 	fmt.Println("Go run -> Windows")
	// case "linux":
	// 	fmt.Println("Go run -> linux")
	// //macOs
	// case "darwin":
	// 	fmt.Println("Go run -> macOS")
	// default:
	// 	fmt.Println("Go run -> otro OS")
	// }

	// for
	// break: salir de un bucle
	// continue: salta a la siguiente iteraccion de un bucle, sin ejecutar el codigo que esta despues de esa instrucción, en el ejemplo no imprime el 5 y continua con el 6 y no imprime el print

	// break
	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(i)
	// 	// salir del bucle si i es igual a 5
	// 	if i == 5 {
	// 		break
	// 	}
	// }

	// continue
	// for i := 1; i <= 10; i++ {
	// 	// salir del bucle si i es igual a 5
	// 	if i == 5 {
	// 		continue
	// 	}
	// 	fmt.Println(i)
	// }
	// saludo := hello("monica")
	// fmt.Println(saludo)

	// sum, mult := calc(4, 5)
	// fmt.Println("la suma es: ", sum)
	// fmt.Println("la multiplicación es: ", mult)

	// ❤ Protecto de la seccion: generar un numero aleatorio entre 0 y 100
	// fmt.Println(rand.Intn(100)) // generar un numero aleatorio
	jugar()
}

// esta funcion va a recibir un valor de tipo string y lo que va a retornar es valor string luego del parentesis
// func hello(name string) string {
// 	return "Hola, " + name
// }

// se pude hacer asi
//
//	func calc(a, b int) (int, int) {
//		sum := a + b
//		mult := a * b
//		return sum, mult
//	}
//
// o asi, retorna sum mult de una vez despues del parentesis, y ya se asigna con el = ya no se declaran
// func calc(a, b int) (sum, mult int) {
// 	sum = a + b
// 	mult = a * b
// 	return
// }

// ❤ funcion jugar
func jugar() {
	numAleatorio := rand.Intn(100)
	var numIngresado int //numero que ingrese el usuario
	var intentos int     //intentos que hizo el jugador
	const maxIntentos = 10

	for intentos < maxIntentos {
		intentos++
		fmt.Printf("Ingresa un numero (intentos restantes: %d): ", maxIntentos-intentos+1)
		fmt.Scanln(&numIngresado) //referencia al numero ingresado

		if numIngresado == numAleatorio {
			fmt.Println("!Felicitaciones, adivinaste el numero")
			jugarNuevamente()
			return
		} else if numIngresado < numAleatorio {
			fmt.Println("El numero a adivinar es mayor")
		} else if numIngresado > numAleatorio {
			fmt.Println("El numero a adivinar es menor")
		}
	}

	fmt.Println("Se acabaron los intentos. el numero era: ", numAleatorio)
	jugarNuevamente()
}

func jugarNuevamente() {
	var eleccion string
	fmt.Println("¿Quieres jugar nuevamente? (si/no): ")
	fmt.Scanln(&eleccion)

	switch eleccion {
	case "si":
		jugar()
	case "no":
		fmt.Println("Gracias por jugar")
	default:
		fmt.Println("Eleccion invalida. Intentalo nuevamente.")
		jugarNuevamente()
	}
}
