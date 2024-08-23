package main

import (
	"fmt"
	"log"

	"github.com/mtamays23/greetings" // el modulo se llamaria asi pero tengo que remplazar para que lo reconozca con go mod edit -replace github.com/mtamays23/greetings=../greetings
)

// func main() {
// 	//manejo de errores
// 	log.SetPrefix("greetings: ") //prefijo del modulo greetings
// 	log.SetFlags(0)              //Se utiliza para establecer la bandera de formato en 0, Configura el registro para no incluir ninguna marca de tiempo ni metadatos en los mensajes como la fecha y la hora

// 	message, err := greetings.Hello("Ismael") //recibo mensaje y error
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	fmt.Println(message)
// }

// devover saludos para varias personas recibe sile string y retorna un  map: dicc con clave y con valor string y un error
func main() {
	//manejo de errores
	log.SetPrefix("greetings: ") //prefijo del modulo greetings
	log.SetFlags(0)              //Se utiliza para establecer la bandera de formato en 0, Configura el registro para no incluir ninguna marca de tiempo ni metadatos en los mensajes como la fecha y la hora

	names := []string{"Ismael", "Eduardo", "Mery", "Monica"}
	messages, err := greetings.Hellos(names) //recibo mensaje y error

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(messages)
}
