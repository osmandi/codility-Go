package main

// Estado:
/*
	Demo ticket:
	Correctness= 100%
	Performance= 100%

	Test score= 100%

	En codility no se incluye la función main
*/



import (
	"fmt"
	"sort"
)

// Elimina los valores repetidos y el cero
func repetidos(A []int) []int {

	encontrado := make(map[int]bool)
	arrayModificado := []int{}
	for i := range A {
		if encontrado[A[i]] == false && A[i] > 0 {
			encontrado[A[i]] = true
			arrayModificado = append(arrayModificado, A[i])
		}
	}

	return arrayModificado
}

func Solution(A []int) int {

	// Ordenar array de mayor a menor
	sort.Ints(A)

	// Limpiar datos
	arrayModificado := repetidos(A)

	var resultado int = 0
	for i := range arrayModificado {

		// Para valores faltantes
		if arrayModificado[i] > 0 && arrayModificado[i] == resultado+1 {
			resultado += 1
		} else {
			return resultado + 1
		}

		// Para valores secuenciales
		if arrayModificado[i] > 0 && i == len(arrayModificado)-1 {
			return arrayModificado[i] + 1
		}
	}

	// Para valores negativos
	return 1
}

// Función main y definición de casos
func main() {
	A := []int{2, 3, 4}
	B := []int{1, 3, 6, 4, 1, 2}
	C := []int{1, 2, 3, 4, 5, 9}
	D := []int{-1, -3}
	E := []int{}
	for i := 1; i <= 2000000; i++ {
		E = append(E, i)
	}

	E = append(E[:132498], E[132857:]...)

	//fmt.Println(E)

	tests := [][]int{A, B, C, D, E}

	for i := range tests {
		resultado := Solution(tests[i])
		fmt.Println(resultado)
	}

}
