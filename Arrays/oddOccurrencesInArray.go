package main

// Autor: @osmandi
// 55% : https://app.codility.com/demo/results/trainingXNMBTH-UA9/
/*
Número de índices de array: 1 - 1,000,000
Rango de cada elemento: 1 - 1,000,000,000
*/
import (
	"fmt"
	"runtime"
	"sync"
)

func Solution(A []int) int {
	var conteoMap = make(map[int]int)
	var result = make(chan int, len(A))
	var mutex = &sync.Mutex{}
	var numberImpar int
	// Ordenar el array
	//	sort.Ints(A)

	for w := range A {
		go func(A []int, w int, result chan<- int) {
			llave := A[w]
			val := 1
			mutex.Lock()
			// Presencia en el map
			v, ok := conteoMap[llave]
			if ok {
				conteoMap[llave] = v + 1
			} else {
				conteoMap[llave] = val
			}
			mutex.Unlock()
			result <- w
			runtime.Gosched()

		}(A, w, result)
	}

	//	time.Sleep(3000 * time.Millisecond)
	// Para tener un control sobre el tiempo
	for a := 0; a <= len(A)-1; a++ {
		<-result
	}

	mutex.Lock()
	// Recorrer el mapa
	for i := range conteoMap {
		if conteoMap[i] == 1 {
			return i
		}
	}
	mutex.Unlock()

	// Identificar el impar
	return numberImpar

}

func main() {
	array := []int{3, 2, 3, 2, 4, 7, 4}
	fmt.Println("Número impar:", Solution(array))
}
