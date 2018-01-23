package main

import (
	"fmt"
	"sort"
)

func Solution(A []int) int {

	// Ordenar el array
	sort.Ints(A)

	var numeroImpar int

	for i := 0; i < len(A); i++ {
		fmt.Println(i)
		if i < len(A)-1 {
			if A[i] != A[i+1] && A[i] != A[i-1] {
				numeroImpar = A[i]
				fmt.Println("Numero impar:", A[i])
			}
		}

		if i == len(A)-1 {
			if A[i] != A[i-1] {
				//				fmt.Println("Numeo impar otro:", A[i])
				numeroImpar = A[i]
			}
		}
	}

	return numeroImpar
}

func main() {
	N := []int{3, 2, 3, 2, 4, 7, 4}
	fmt.Printf("El valor impar es: %d", Solution(N))
}

// https://app.codility.com/demo/results/trainingXGS5JZ-DQD/
