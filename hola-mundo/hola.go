package main

// agregar un paquete de tercero
// "rsc.io/quote"

// ❤ declarar variables fuera de la funcion
// var firstName, lastName, age = "Alejandra", "Tamayo", 30

// ❤ declaracion de constantes: constantes cuyo valor no cambia en la ejecucion del programa, palabra reservada const, se debe declarar a nivel de paquete por recomendacion
// const Pi float32 = 3.14 o asi
// const Pi = 3.14

// const (
// 	x = 100
// 	y = 0b1010 //binario
// 	z = 0o12   // octal
// 	w = 0xFF   // hexadecimal
// )

// ❤ valor iota: iota inicia desde cero y los demas valores se incrementan en 1, vierves seria el 6
// const (
// 	Domingo = iota + 1
// 	Lunes
// 	Martes
// 	Miercoles
// 	Jueves
// 	Viernes
// )

func main() {
	// fmt.Println("Hola Mundo")
	// fmt.Println(quote.Hello())

	// ❤ declaracion de variables de esta forma
	// var firstName, lastName string
	// var age int

	// firstName = "Alex"
	// lastName = "Perez"
	// age = 30

	// o de esta otra en conjunto, utilizando la palabra reservada var
	// var firstName, lastName, age = "Alejandra", "Tamayo", 30

	// o utilizamos el :=: se utiliza para declarar una variable nueva, solo dentro de las funciones, var sirver para dentro y fuera de las funciones
	// firstName, lastName, age := "Alejandra", "Tamayo", 30
	// para modificar si es con =
	// age = 30

	// fmt.Println(firstName, lastName, age)
	// fmt.Println(Pi)
	// fmt.Println(x, y, z, w)
	// fmt.Println(Viernes)

	// ❤ tipos de datos basicos
	/* almacenar numeros enteros: int8, int16, int32 bits... uint es lo mismo pero es dependiendo del SO
	numeros mas grandes con decimal es float: float32 y float64*
	con math podemos saber elvalor minimo o el valor maximo en cada variable */
	// fmt.Println(math.MinInt64, math.MaxInt64) //VALOR MAXIMO QUE PUEDE ALMACENAR

	// Booleanos
	// var valueBool bool = false

	// escapar de comillas, generar saltos de lineas, generar espacio
	// fullName := "Aleja Tamayo \t(alias \"mtamays\")\n"
	// fmt.Println(fullName)

	// //numeros positivos hasta 255, codigo tipo asci
	// var numAsci byte = 'a' //valor asci de la letra a
	// fmt.Println(numAsci)

	// num := "hola"
	// fmt.Println(num[0]) //imprime valor decimal de la letra h

	//codigo unicode, caracteres unicode
	// var uni rune = '❤'
	// fmt.Println(uni) //imprime valor unicode

	// ❤ valores predeterminados
	// var (
	// 	defaultInt    int     //0
	// 	defaultUint   uint    //0
	// 	defaultFloat  float32 //0
	// 	defaultBool   bool    //false
	// 	defaultString string  //''
	// )

	// fmt.Println(defaultInt, defaultUint, defaultFloat, defaultBool, defaultString)

	// ❤ conversiones de tipos
	// var num16 int16 = 50
	// var mun32 int32 = 100

	// fmt.Println(num16 + mun32) es invalida porque son de distintos tipos, necesitamos hacer una conversion
	// fmt.Println(int32(num16) + mun32)

	//que pasa si quiero convertir string a int o viceversa, atoi convierte string a int
	// num1 := "100"
	// convert, _ := strconv.Atoi(num1) //devuelve el int y posible error, si no queremos recuperar el error ponemos _
	// fmt.Println(convert)

	// // de entero a cadena, Itoa convierte int a string
	// n := 42
	// num1 = strconv.Itoa(n)
	// fmt.Println(num1)

	// ❤ El PAQUETE FMT
	// https://pkg.go.dev/fmt#pkg-functions, Explicacion funcion del paquete

	// fmt.Print("hola") //no coloca salto de linea

	// name := "Moni"
	// age := 30

	// fmt.Printf("Hola, me llamo %s y tengo %d años. \n", name, age) //formato de salida

	// greeting := fmt.Sprintf("Hola, me llamo %s y tengo %d años.", name, age) //devolver informacion formateada como string
	// fmt.Println(greeting)

	// para que el usuario ingrese datos
	// var name string
	// var lastname string
	// var age int

	// fmt.Print("Ingrese su nombre: ")
	/* SCAN podemos escanear datos, recibe la referencia de la memoria de una variable donde se almacenara lo que se ingresa por teclado, para ingresar la referencia o
	para enviar la referencia utilizamos &*/
	// fmt.Scanln(&name, &lastname) //solo capta una cadena
	// fmt.Print("Ingrese su edad: ")
	// fmt.Scanln(&age)
	// fmt.Printf("Hola, me llamo %s y tengo %d años. \n", name, age) //formato de salida

	// // tipo de dato de variable
	// fmt.Printf("el tipo de name es: %T\n", name)
	// fmt.Printf("el tipo de name es: %T\n", age)

	// ❤ Operadores aritmeticos y paquetes
	// suma +
	// resta -
	// division /
	// modulo %
	// ++ incremento
	// -- decremento

	// a := 10
	// b := 3
	// b++
	// a += 5

	// fmt.Println(a)
	// fmt.Println(math.Pi)

	/* ❤  Ejercicio: Crear un programa que solicite al usuario que ingrese los lados de un triángulo rectángulo y luego calcule e imprima el área y el perímetro del triángulo.*/
	// Solicitar al usuario que ingrese la longitud de los dos lados del triángulo rectángulo.

	// var lado1, lado2 float64

	// fmt.Print("Ingrese longitud lado1 del triangulo: ")
	// fmt.Scanln(&lado1)
	// fmt.Print("Ingrese longitud lado2 del triangulo: ")
	// fmt.Scanln(&lado2)

	// // Calcular la hipotenusa del triángulo usando el teorema de Pitágoras.
	// hip := math.Sqrt(math.Pow(lado1, 2) + math.Pow(lado2, 2))
	// fmt.Println("Hipotenusa: ", hip)

	// // Calcular el área del triángulo usando la fórmula base x altura / 2.
	// area := (lado1 * lado2) / 2
	// fmt.Printf("Area: %.2f \n", area) //Imprimir el área del triángulo con dos decimales de precisión.

	// //Calcular el perímetro del triángulo sumando los lados.
	// // Imprimir el perímetro del triángulo con dos decimales de precisión.
	// peri := lado1 + lado2 + hip
	// fmt.Printf("Perimetro: %.2f \n", peri)

	/* ❤  Ooperadores booleanos:
	Igualdad (==)

	Desigualdad (!=)

	Mayor que (>)

	Menor que (<)

	Mayor o igual que (>=)

	Menor o igual que (<=)

	// Comparación de booleanos
	fmt.Println(true == true)           // true
	fmt.Println(false != true)          // true
	fmt.Println(true && false == false) // true
	fmt.Println(true || false == true)  // true

	Operadores lógicos
	Operador AND lógico (&&): evalúa dos expresiones booleanas y devuelve verdadero (true) si ambas expresiones son verdaderas, y devuelve falso (false) si alguna de las expresiones es falsa
	Operador OR lógico (||): evalúa dos expresiones booleanas y devuelve verdadero (true) si al menos una de las expresiones es verdadera, y devuelve falso (false) si ambas expresiones son falsas.
	Operador NOT lógico (!): niega el valor booleano de una expresión, es decir, si una expresión es verdadera, la niega y devuelve falso, y si una expresión es falsa, la niega y devuelve verdadero.

	Expresiones: El orden en que se resuelven las expresiones en un programa depende de la prioridad de los operadores y los paréntesis utilizados para agrupar las operaciones
	Los paréntesis se evalúan primero. Las expresiones dentro de los paréntesis se resuelven antes que cualquier otra operación.
	Luego se resuelven las operaciones aritméticas, como la multiplicación, la división, la suma y la resta. La multiplicación y la división se resuelven antes que la suma y la resta.
	Finalmente, se resuelven las operaciones de comparación y los operadores lógicos.
	Es importante tener en cuenta que los operadores con la misma prioridad se resuelven de izquierda a derecha. Por ejemplo, en la expresión 2 + 3 * 4, la multiplicación se resuelve primero debido a su mayor prioridad, y el resultado es 14. Si queremos que la suma se resuelva primero, debemos utilizar paréntesis para indicar la prioridad, como en (2 + 3) * 4, lo que da como resultado 20.
	*/
}
