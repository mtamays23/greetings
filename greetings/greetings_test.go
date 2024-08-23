// pruebas para el archivo greetings, test, probar la funcion hello
// ejecutar en el archivo go test o go test -v

package greetings

import (
	"regexp"
	"testing"
)

func TestHelloName(t *testing.T) {
	// probar la funcion hello que recibe un nombre
	name := "Ismael"
	// expresion regular que busca una coincidencia exacta con el nombre, delimitado por los caracteres de la palabra completa con \
	want := regexp.MustCompile(`\b` + name + `\b`)
	// capturo lo que retorna hello
	msg, err := Hello("Ismael")
	// si name es igual o no, analiza el mensaje si coincide con el nombre o no, si coincide devuelve bool si no devuelve error
	if !want.MatchString(msg) || err != nil {
		t.Fatalf(`Hello("Ismael) = %q, %v, quiere coincidencia para %#q, nil`, msg, err, want)
	}
}

// valido prueba cuando va vacio
func TestHelloEmpty(t *testing.T) {
	msg, err := Hello("") //retorna message vacio y error
	if msg != "" || err == nil {
		t.Fatalf(`Hello("") = %q, %v, quiere "", error`, msg, err)
	}
}
