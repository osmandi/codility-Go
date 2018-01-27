package main

import (
	"fmt"
	"sort"
)

/*
v2:
- Mejorar el performance (hacer más rápidos las soluciones)
- Habilitar la opción de 1 a 100,000 índices y un número de 1 a 1,000,000

*/

// Eliminar valores repetidos
func eliminarRepetidos(A []int) []int {

	// Eliminar valores repetidos
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

// Identificar número impar
func encontrarImpar(A, B []int) int {

	// Cantidad encontrados
	encontrado := make(map[int]int)
	encontrado[0] = 0

	// Canal
	canalEncontrado := make(chan map[int]int)

	// Correr A[i] con el array B en cada goroutina
	// para guardar los números repetidos en el mapa
	for i := range A {
		for j := range B {

		}
	}

	for i := range encontrado {
		if encontrado[i] == 1 {
			return i
		}
	}

	//	fmt.Println(encontrado)

	return 0
}

// Función del ejercicio inmodificable
func Solution(A []int) int {

	// Ordenar array
	sort.Ints(A)

	arrayLimpio := eliminarRepetidos(A)
	encontrado := encontrarImpar(A, arrayLimpio)

	return encontrado
}

// Función main
func main() {
	N := []int{3, 2, 3, 2, 4, 7, 4}
	fmt.Printf("El número impar es: %d", Solution(N))

}
