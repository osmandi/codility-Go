package main

// Notas
/*
Score = 55 %

Errors:
- n=201 -> got 0 expected 42
- n=100,003 -> Runtime 6sec, time limit 0.1 sec
- n=999,999 -> Runtime 2.88 sec, time limit 0.3
- n=999,999 -> Runtime 6 sec, time limit 0.48 sec
*/

import "fmt"

/*
Autor: @osmandi

Objetivo:
- Identificar el número que no se encuentra repetido en un slice impar

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

	for i := range A {
		for j := range B {
			if B[j] == A[i] {
				encontrado[B[j]] += 1
			}
		}
	}

	for i := range encontrado {
		if encontrado[i] == 1 {

			return i
		}
	}

	fmt.Println(encontrado)

	return 0
}

// Función del ejercicio inmodificable
func Solution(A []int) int {

	arrayLimpio := eliminarRepetidos(A)
	encontrado := encontrarImpar(A, arrayLimpio)

	//	fmt.Printf("El valor impar es %d:", encontrado)

	return encontrado
}

// Función main
func main() {
	N := []int{3, 2, 3, 2, 4, 7, 4}
	fmt.Printf("El número impar es: %d", Solution(N))

}
