package main

import (
	"fmt"
	"sort"
)

func repetidos(A []int) []int {

	// Eliminar elementos repetidos
	encontrado := make(map[int]bool)
	arrayModificado := []int{}
	for i := range A {
		if encontrado[A[i]] == false {
			encontrado[A[i]] = true
			arrayModificado = append(arrayModificado, A[i])
		}
	}

	return arrayModificado
}

func Solution(A []int) (int, []int) {

	// Ordenar array de mayo a menor
	sort.Ints(A)

	arrayModificado := repetidos(A)

	// Determinar el resultado
	// 	Si son negativos: 1
	// 	Retornar el menor valor faltante
	// 	En una secuencia retornar el sigueinte último valor

	var resultado int = 0
	for i := range arrayModificado {

		// Para valores faltantes
		if arrayModificado[i] > 0 && arrayModificado[i] == resultado+1 {
			resultado += 1
		} else {
			return resultado + 1, arrayModificado
		}

		// Para valores secuenciales
		if arrayModificado[i] > 0 && i == len(arrayModificado)-1 {
			return arrayModificado[i] + 1, arrayModificado
		}
	}

	// Para valores negativos
	return 1, arrayModificado
}

func main() {
	A := []int{1, 2, 3}
	B := []int{1, 3, 6, 4, 1, 2}
	C := []int{1, 2, 3, 4, 5, 9}
	D := []int{-1, -3}
	E := []int{}
	for i := 0; i <= 100000; i++ {
		E = append(E, i)
		fmt.Println(E[i])
	}

	tests := [][]int{A, B, C, D, E}

	for i := range tests {
		resultado, array := Solution(tests[i])
		fmt.Println(resultado, array)
	}

}
