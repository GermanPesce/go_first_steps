package main

import (
	"container/list"

	"fmt"

	"reflect"

	"github.com/google/uuid"
)

func main() {
	fmt.Println("Hello World")
	fmt.Println(uuid.New().String())

	//se comenta asi
	/*
		en varias lineas asi
	*/

	//Variables (GO ES DE TIPADO FUERTE)
	var string1 string
	string1 = "sarasa"
	fmt.Println(string1) //si no se usan las variables pincha
	string1 = "sarasa2"
	fmt.Println(string1)

	var string2 = "se puede declarar asi tambien"
	fmt.Println(string2)

	var int1 = 7
	fmt.Println(int1)
	//para enteros no es raro trabajar con diferentes tipos. por ej; int8, int32, etc

	fmt.Println(string1, int1)

	fmt.Println(reflect.TypeOf(int1))

	var myfloat float64 = 33.5
	fmt.Println(float64(int1) + myfloat)

	var myBool bool = false
	myBool = true
	fmt.Println(myBool)

	varAutoDeclarada := "te cabio"
	fmt.Print(varAutoDeclarada)

	//CONSTANTES
	const myConst = 123 //se pueden declarar sin utilizar

	//CONTROL DE FLUJO (if, else)
	if myBool == true {
		varAutoDeclarada = "cabio"
	} else if myBool == false || string1 == "Hola" {
		varAutoDeclarada = "no cabio"
	} else {
		varAutoDeclarada = "nada"
	}

	//ESTRUCTURAS DE DATOS
	//Array
	var myArray [3]int
	myArray[0] = 1
	fmt.Println(myArray)
	fmt.Println(myArray[2])

	//Map
	myMap := make(map[string]int) //tipo // clave // valor
	myMap["Pepe"] = 37
	myMap["Sarasa"] = 29
	fmt.Println(myMap)
	fmt.Println(myMap["Pepe"])

	myMap2 := map[string]int{"sarasa": 23, "pepe": 29}
	fmt.Println(myMap2)

	//List (trabaja con punteros!!!)
	//a medida que agregamos tipos de datos se van importando librerias
	myList := list.New()
	myList.PushBack(1)
	myList.PushBack(2)
	myList.PushBack(3)
	fmt.Println(myList)
	fmt.Println(myList.Back().Value)

	//BUCLES
	var myArrayBucle [3]int
	myArrayBucle[0] = 1
	myArrayBucle[1] = 2
	myArrayBucle[2] = 3

	for i := 0; i < len(myArrayBucle); i++ {
		myArray = myArrayBucle
	}
	fmt.Println(myArray)
	fmt.Println(myArrayBucle)

	for key, value := range myMap {
		fmt.Println(key, value)
	}

	//llamada a la func
	fmt.Println(myFunction(string1))

	//Structures
	type MyStructure struct {
		name string
		age  int8
	}

	myStructure := MyStructure{"German", 37}
	fmt.Println(myStructure)
}

// FUNCTION
func myFunction(algo string) string {
	return "sarasa" + algo
}
