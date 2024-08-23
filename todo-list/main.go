package main

import (
	"bufio"
	"fmt"
	"os"
)

// func main() {
// ❤ matrices

// declarar matrices
//inicializo
// var a = [...]int{10, 20, 30, 40, 50} // si no se con cuantos datos se va a inicializar el array, con ...
// var a = [5]int{10, 20, 30, 40, 50}
//modifico
// a[0] = 100
// a[1] = 200

//con len
// for i := 0; i < len(a); i++ {
// 	fmt.Println(a[i])
// }

// //con range, se devuelve indice y valor
// for index, value := range a {
// 	fmt.Printf("Indice: %d, valor: %d\n", index, value)
// }

// si solo queremos el valor, no sacar el indice
// for _, value := range a {
// 	fmt.Printf("valor: %d\n", value)
// }

// ❤  matriz bidimensional:[numero de filas] [numero de columnas] tipo de dato a almacenar
//las llaves representan filas y los elementos representan columnas
// var matriz = [3][3]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
// fmt.Println(matriz)

// ❤ rebanadas o slice, trabajar con partes de un arreglo, se puede agregar y quitar elementos de un mismo tipo, array no se pueden
// rebanadaDias := []string{"Domingo", "lunes", "martes",
// 	"miercoles", "jueves", "viernes", "sabado"}

// //crear rebanada a partir de otra rebanada
// diaRebanada := rebanadaDias[0:5]

// //agregar elementos, recibe un slicen y la cantidad de elemtentos que agregaremos, es dinamica crece su capacidad
// diaRebanada = append(diaRebanada, "viernes", "sabado", "festivo")

// //eliminar el segundo elemento del indice 2: martes
// diaRebanada = append(diaRebanada[:2], diaRebanada[3:]...)

// fmt.Println(rebanadaDias)
// fmt.Println(diaRebanada)

// fmt.Println(len(diaRebanada)) //longitud de la reabanada
// fmt.Println(cap(diaRebanada)) //capacidad de la reabanada, teniendo en cuenta que es una rebanada hija de la rebanada padre

// ❤ make: crea rebanadas, con la longitud 5 y su capacidad 10
// nombres := make([]string, 5, 10)
// nombres[0] = "alex"
// fmt.Println(nombres)

// // ❤ copy: copiar elementos de una rebanada a otra
// rebanada1 := []int{1, 2, 3, 4, 5}
// rebanada2 := make([]int, 5)

// copy(rebanada2, rebanada1) // devuelve cantidad de elementos copiados

// fmt.Println(rebanada1)
// fmt.Println(rebanada2)

//  ❤ MAPAS: almacenar y recuperar claves por clave y valor, [inidcar que tipo de dato es la llave, string] y que tipo de dato sera su valor, string]
// colors := map[string]string{
// 	"rojo":  "#FF0000",
// 	"verde": "#00FF00",
// 	"azul":  "#0000FF",
// }

// fmt.Println(colors["rojo"])
// //agregar elementos
// colors["negro"] = "#000000"
// fmt.Println(colors)

//validar si existe la clave

// fmt.Println(valor, ok) //devuelve vacio y false ya que este color no existe

// if valor, ok := colors["verde"]; ok {
// 	fmt.Println(valor)
// } else {
// 	fmt.Println("no existe la clave")
// }

// delete: eliminar elementos, recibe el mapa y la clave
// delete(colors, "azul")
// fmt.Println(colors)

//iterar por los elementos del mapa
// for clave, valor := range colors {
// 	fmt.Printf("clave: %s, valor: %s\n", clave, valor)
// }
// }

// ❤ ESTRUCTURAS: struct, tipo de dato personalizado, compuestos por diferentes campos, es como una clase,
// type Persona struct {
// 	nombre string
// 	edad   int
// 	correo string
// }

// metodos utilizamos el struct Persona, un metodo pertenece a una estructura colocando el receptor, en cambio las funciones son independientes, en cambio los metodos debemos aceder mediante la instancia, en este caso persona
// func (p *Persona) diHola() { //recepciona el puntero de la struct persona para acceder a los atributos de este
// 	fmt.Println("Hola, mi nombre es", p.nombre)
// }

// func main() {
// p := Persona{"Monica", 30, "moni@gmail.com"}
// p.edad = 31
// fmt.Println(p)

// p2 := Persona{"Juan", 40, "juan@gmail.com"}
// fmt.Println(p2)

/* ❤ Punteros y metodos: se pueden definir metodos en estructuras para realizar operaciones a las variables de la estructura,
   los metodos se definen como funciones que tiene un receptor, que es un puntero o una variable de la estructura, son receptores de metodos de GO ya que permiten
   que los metodos modifiquen directamente la variable original

   un puntero almacena la direccion en memoria de otra variable, se utilizan para referenciar y acceder a la variable original a travez de su direccion de la memoria */
// var x int = 10
// fmt.Println(x) //10
// var p *int = &x // puntero referencia de la memoria de la variable x
// fmt.Println(p) //0xc000110068
// editar(&x)
// fmt.Println(x) //20

// para acceder al metodo diHola debemos acceder mediante la instancia del puntero, en este caso Persona
// p := Persona{"Aleja", 32, "Ale@gmail.com"}
// p.diHola()
// }

// func editar(x *int) {
// 	*x = 20 //modifique la variable original, la copia de la memoria
// }

// ❤ PROYECTO gestor de tareas, gestionar una lista de tareas pendientes
type Tarea struct {
	nombre      string
	descripcion string
	completado  bool
}

type listaTareas struct {
	tareas []Tarea
}

// metodo para agregar tareas, recibimos dato de tipo tarea
func (l *listaTareas) agregarTarea(t Tarea) {
	l.tareas = append(l.tareas, t)
}

// metodo para marcar como completado una tarea
func (l *listaTareas) marcarCompletado(index int) {
	l.tareas[index].completado = true
}

// metodo para editar tarea
func (l *listaTareas) editarTarea(index int, t Tarea) {
	l.tareas[index] = t
}

// metodo para eliminar tarea
func (l *listaTareas) eliminarTarea(index int) {
	l.tareas = append(l.tareas[:index], l.tareas[index+1:]...)
}
func main() {
	//desarrollo interface con la que interactuara el usuario, y utilizaremos los metodos y las estructuras que hemos creado
	// instancia de la lista de tareas
	lista := listaTareas{}

	// instancia de bufio para entrada de datos: con NewReader creamos un nuevo lector que implementa io.Reader, con os.stdin lector de teclados, se utiliza por limitacion con fmt
	leer := bufio.NewReader(os.Stdin)

	// menu de navegacion, bucle for infinito
	for {
		var opcion int
		fmt.Println("Seleccione una opción:\n",
			"1. Agregar Tarea\n",
			"2. Marcar tarea como completada\n",
			"3. Editar tarea\n",
			"4. Eliminar tarea\n",
			"5. Salir",
		)

		// leemos la opcion que ha ingresado el usuario
		fmt.Println("Ingrese la opción")
		fmt.Scanln(&opcion)

		switch opcion {
		case 1:
			var t Tarea
			fmt.Print("Ingrese el nombre de la tarea: ")
			// lee el string con la variable bufio creada arriba
			t.nombre, _ = leer.ReadString('\n')
			fmt.Print("Ingrese la descripcion de la tarea: ")
			t.descripcion, _ = leer.ReadString('\n')
			lista.agregarTarea(t)
			fmt.Println("Tarea agregada correctamente")
		case 2:
			var index int
			fmt.Println("Ingrese el indice de la tarea que desea marcar como completado: ")
			fmt.Scanln(&index)
			lista.marcarCompletado(index)
			fmt.Println("Tarea marcada como completada correctamente")
		case 3:
			var index int
			var t Tarea
			fmt.Println("Ingrese el indice de la tarea que desea actualizar: ")
			fmt.Scanln(&index)
			fmt.Print("Ingrese el nombre de la tarea: ")
			t.nombre, _ = leer.ReadString('\n')
			fmt.Print("Ingrese la descripcion de la tarea: ")
			t.descripcion, _ = leer.ReadString('\n')
			lista.editarTarea(index, t)
			fmt.Println("Tarea actualizada correctamente")
		case 4:
			var index int
			fmt.Println("Ingrese el indice de la tarea que desea eliminar: ")
			fmt.Scanln(&index)
			lista.eliminarTarea(index)
			fmt.Println("Tarea eliminada correctamente")
		case 5:
			fmt.Println("Saliendo del programa...")
			return
		default:
			fmt.Println("Opcion invalida")
		}
		//listar todas las tareas
		fmt.Println("Lista de tareas:")
		fmt.Println("================================================")

		// recuperamos indice y tareas, recorremos la instancia que creamos lista, el campo tareas que esta dentro de la listaTareas
		for i, t := range lista.tareas {
			// imprimimos indice, nombre de la tarea, descripcion, completado(booleano)
			fmt.Printf("%d. %s - %s - Completado: %t\n", i, t.nombre, t.descripcion, t.completado)
		}
		fmt.Println("================================================")
	}
}
