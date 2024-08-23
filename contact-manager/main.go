package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

// ❤  errores personalizados
/* el error que podemos ver aca es que si se manda a dividendo a 0 saldria el error ya que no se puede dividir por cero,  entonces aca personalizamos el manejo de error*/
// (dividendo, divisor int) lo que recibe, (int, error) lo que retornara
// func divide(dividendo, divisor int) (int, error) {
// 	if divisor == 0 {
// 		// return 0, errors.New("No se puede dividir por 0") //funcion errors.new  del paquete errors, recibe un string y devuelve un error, retorna el entero y el error
// 		return 0, fmt.Errorf("No se puede dividir por 0") //tabn podemos utilizar el fmt.Errorf
// 	}
// 	return dividendo / divisor, nil //nil valor predeterminado de tipo de dato error
// }

// func dividir(dividendo, divisor int) {
// 	defer func() { //funcion anonima que capturara cualquier panico que ocurra dentro de esta funcion, dentro de la variable r, no interrumpe la ejecucion de la aplicacion
// 		if r := recover(); r != nil {
// 			fmt.Println(r)
// 		}
// 	}()
// 	validateZero(divisor) //si no llamo la funcion sale el error sin personalizar runtime error: integer divide by zero
// 	fmt.Println(dividendo / divisor)
// }

// func validateZero(divisor int) {
// 	if divisor == 0 {
// 		panic("No puedes dividir por cero")
// 	}
// }

// ❤  manejo de errores
// func main() {
// str := "123f"                 //colocamos mal el numero para convertir de str a int para que se genere el error
// num, err := strconv.Atoi(str) //atoi recibe string y devuelve int y error
// if err != nil {
// 	fmt.Println("Error", err) //imprime: Error strconv.Atoi: parsing "123f": invalid syntax
// 	return
// }

// fmt.Println("Numero", num)

// retorna si hay error: Error:  No se puede dividir por 0, si no hay errores retorna el resultado 5
// result, error := divide(10, 0) // enviamos dividendo, divisor
// if error != nil {
// 	fmt.Println("Error: ", error)
// 	return
// }

// fmt.Println("resultado: ", result)

// ❤  defer: posponer la ejecucion de una funcion hasta que la funcion que la contiene haya finalizado, se coloca defer antes de la llamada de una funcion, se agrega a una pila de ejecucion
// con defer la primera funcion en llegar es la ultima en ejecutarse, es como en una pila sacaria 1,2,3 en vez de 3,2,1
// defer fmt.Println(3)
// defer fmt.Println(2)
// defer fmt.Println(1)

// vamos a crear un archivo y escribir en ese archivo
// file, err := os.Create("hola.txt") //crear un archivo, recibe un string y devuelve el file y error
// if err != nil {
// 	fmt.Println(err)
// 	return
// }

// defer file.Close() //cierra el flujo antes de que la funcion main se termine de ejecutar, para no dejarlo abierto si sale algun error y no repetir codigo

// _, err = file.Write([]byte("Hola, Monica Alejandra")) //recibe byte[], devuelve int y error, usarmos _ no necesitamos la cantidad de caracteres dentro del archivo, no devuelva el string
// if err != nil {
// 	fmt.Println(err)
// file.Close() //utilizamos defer no repetir codigo
// }
// file.Close() //utilizamos defer no repetir codigo

/* ❤  panic y recover: se utiliza para manejar situaciones excepcionales o errores graves que ocurran en la ejecucion de un programa
panic: se utiliza para generar una interrupcion inmediata, sale los errores en pila,  en la ejecucion del programa, cuando ocurre un error
recover: maneja y captura un panico*, No se debe utilizar para manejo de errores solo para situaciones inesperadas o limpieza y recuperacion antes de finalizar la ejecucion del programa */

// dividir(100, 10)
// dividir(200, 25)
// dividir(34, 0) //ejecuta los primeros dos y estalla en esta linea, leer los errores de abajo hacia arrriba
// dividir(100, 4)

/* ❤ registro de errores: paquete log, se utiliza para registro basico de mensajes, por consola o archivo indica feha, minuto y segundo*/
// log.SetPrefix("main()") //esta funcion se utiliza para agregar un prefijo a los mensajes de los registros del programa, indicarle el error en una funcion, nos dice en que funcion es el error: main()2024/08/16 13:21:25 Puedes verme?
// log.Print("Puedes verme?")                    //log 2024/08/16 13:11:18 Puedes verme?
// log.Fatal("!oye, soy un registro de errores") // detiene la ejecucion de la funcion
// log.Panic("!oye, soy un log")                 // detiene la ejecucion de la funcion

//enviar registro de errores en archivo, paquete os.OpenFile(), recibe un string, una bandera,
// bandera os.O_CREATE: valida la creacion del archivo si no existe si existe lo abre
// bandera os.O_APPEND abre el archivo en modo adjunto, los nuevos datos se agregan al final del archivo
// os.O_WRONLY solo escritura
// 0644 permisos de lectura y de escritura
// file, err := os.OpenFile("info.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
// if err != nil {
// 	log.Fatal(err)
// }

// defer file.Close()

//si no hay errores creando el archivo ni abriendolo..
// log.SetOutput(file) // enviamos archivo para el registro de errores y que lo escriba en file
// log.Print("!oye, soy un log")

// }

/* ❤ proyecto de sección: gestor de contactos, agregar contactos, ver contactos */

type Contact struct {
	Name  string `json:"name"` // `json:"name"`: para que se mire en el json con minusculas
	Email string `json:"email"`
	Phone string `json:"phone"`
}

// agregar/guardar registros en un archivo de tipo json, recibe un slice de tipo Contact, y retornara error si no se ha podido guardar en el archivo

// 1 funcion: crear un archivo json y guardar dato de tipo slice
func saveContactsToFile(contacts []Contact) error {
	// crear el archivo
	file, err := os.Create("contacts.json") // abrimos el flujo
	if err != nil {
		return err
	}

	defer file.Close()

	//crear encode de json para escribir datos, FUNCION json.NewEncoder DEL PAQUETE JSON, de slice a json
	//serializar un dato de tip slice a un dato de tipo json, guardarlo en un archivo de tipo json
	encoder := json.NewEncoder(file) //Recibe objeto tipe file, el archivo donde escribirimos los datos, el encoder se utiliza para convertir estructura de datos en formato json
	err = encoder.Encode(contacts)   // recuperar error, codificar el slice de contactos en formato json y escribirlo en un archivo
	if err != nil {
		return err
	}

	return nil
}

// cargar contactos desde un archivo json, serialzar un dato de tipo json a slice
func loadContacts(contacts *[]Contact) error {
	// abre el archivo creado
	file, err := os.Open("contacts.json")
	if err != nil {
		return err
	}

	defer file.Close()

	// codificar file json a tipo slice, lo contrario
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&contacts)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	var contacts []Contact

	// cargar contactos existentes desde el archivo json
	err := loadContacts(&contacts)
	if err != nil {
		fmt.Println("Error al cargar los contactos: ", err)
	}

	// crear instancia de bufio, // instancia de bufio para entrada de datos: con NewReader creamos un nuevo lector que implementa io.Reader, con os.stdin lector de teclados, se utiliza por limitacion con fmt
	reader := bufio.NewReader(os.Stdin)

	//bucle infinito
	for {
		fmt.Print("==== GESTOR DE CONTACTOS ====\n",
			"1. Agregar un contacto\n",
			"2. Mostrar todos los contactos\n",
			"3. Salir\n",
			"Elige una opcion: ")

		var option int
		_, err = fmt.Scanln(&option) //retorna un entero de cantidad de leidos y error, obviaremos la cantidad
		if err != nil {
			fmt.Println("Error al leer la opcion: ", err)
			return
		}

		// Manejar las opciones del usuario
		switch option {
		case 1:
			// ingresar y crear un contacto
			var c Contact
			fmt.Print("Nomre: ")
			c.Name, _ = reader.ReadString('\n')
			fmt.Print("Email: ")
			c.Email, _ = reader.ReadString('\n')
			fmt.Print("Telefono: ")
			c.Phone, _ = reader.ReadString('\n')

			//agregar un contacto o slice
			contacts = append(contacts, c)

			//guardar en un archivo json
			if err := saveContactsToFile(contacts); err != nil {
				fmt.Println("Error al guardar el contacto: ", err)
			}
		case 2:
			// mostrar todos los contactos
			fmt.Println("================================================================")
			for index, contact := range contacts {
				fmt.Printf("%d. Nombre: %s Email: %s Telefono: %s\n",
					index+1, contact.Name, contact.Email, contact.Phone)
			}
			fmt.Println("================================================================")
		case 3:
			// salir del programa
			return
		default:
			fmt.Println("Opcion invalida")
		}

	}

}
