// Notas
/*
Score 44%

n=201, 2001 => Error
n=100003,999999 (6sec -> 0.1sec), (6sec -> 0.3sec or 0.5 sec)

*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Eliminar números repetidos
func eliminatePaired(A []int) []int {

	// Mapa que ayudará a elimina repetidos
	pareado := make(map[int]bool)

	// Array que almacenará los valores
	var numero []int

	// Eliminar valores repetidos
	for i := range A {
		if pareado[A[i]] == false {
			pareado[A[i]] = true
			numero = append(numero, A[i])
		}
	}

	return numero
}

// Convertir Array int a string
func convertToString(A []int) []string {

	var B []string

	for i := range A {
		B = append(B, strconv.Itoa(A[i]))
	}

	return B
}

// Retornar el valor que solo aparece una sola vez
func countString(A, B []string) int {

	// Une el array A
	arrayAJoin := strings.Join(A, "")

	// Donde se guardará el valor impar
	var numberImpar int

	// Detectar cuál valor está solo una vez
	for i := range B {
		if strings.Count(arrayAJoin, B[i]) == 1 {
			numberImpar, _ := strconv.Atoi(B[i])
			return numberImpar
		}
	}

	return numberImpar
}

func Solution(A []int) int {

	// Eliminar elementos repetidos
	arrayNoPaired := eliminatePaired(A)

	// Convertir arrays int a string
	arrayString := convertToString(A)
	arrayNoPairedString := convertToString(arrayNoPaired)

	// Extraer número impar
	number := countString(arrayString, arrayNoPairedString)

	return number
}

func main() {
	A := []int{9, 3, 9, 3, 9, 7, 9}
	fmt.Println(Solution(A))
}
