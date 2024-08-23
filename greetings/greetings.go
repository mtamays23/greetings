package greetings //nombre del pquete done trabajare

import (
	"errors"
	"fmt"
	"math/rand"
)

// hello devuelve un saludos para la persona especificada
// func Hello(name string) (string, error) {

// 	//manejo de errores
// 	if name == "" {
// 		return "", errors.New("nombre vacio") //retorna message vacio y error
// 	}
// 	message := fmt.Sprintf(randomFormat(), name) //formatea informacion y devover string y guardarlo en la variable mesage
// 	return message, nil                          // retorno mensaje y el default del error
// }

// //llamaremos esta funcion desde otro modulo hello el cual se creo en la raiz y se inicializo modulo con go mod init hello(nombre del proyecto)

// // 💒 saludos aleatorios
// func randomFormat() string {
// 	formats := []string{
// 		"¡Hola, %v! ¡Bienvenido!",
// 		"¡Que bueno verte, %v!",
// 		"¡Saludo, %v! ¡Encantado de conocerte!",
// 	}
// 	return formats[rand.Intn(len(formats))] // devuelve un mensaje aleatorio del slice format
// }

// 💒 devolver saludos para varias personas
// hello devuelve un saludos para la persona especificada
// llamaremos esta funcion desde otro modulo hello el cual se creo en la raiz y se inicializo modulo con go mod init hello(nombre del proyecto)
func Hello(name string) (string, error) {

	//manejo de errores
	if name == "" {
		return "", errors.New("nombre vacio") //retorna message vacio y error
	}
	message := fmt.Sprintf(randomFormat(), name) //formatea informacion y devover string y guardarlo en la variable mesage
	return message, nil                          // retorno mensaje y el default del error
}

// devover saludos para varias personas recibe sile string y retorna un  map: dicc con clave y con valor string y un error
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string) // dicc con clave string y valor string

	// si solo queremos el valor, no sacar el indice
	for _, name := range names {
		message, err := Hello(name)
		// manejo el error
		if err != nil {
			return nil, err // nil en el mensaje y error
		}

		messages[name] = message
	}
	return messages, nil // nil si no hay error
}

// 💒 saludos aleatorios
func randomFormat() string {
	formats := []string{
		"¡Hola, %v! ¡Bienvenido!",
		"¡Que bueno verte, %v!",
		"¡Saludo, %v! ¡Encantado de conocerte!",
	}
	return formats[rand.Intn(len(formats))] // devuelve un mensaje aleatorio del slice format
}
