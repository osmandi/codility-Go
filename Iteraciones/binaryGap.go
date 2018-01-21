// Notas
/*
- Retornar la longitud máxima de un número binario (cantidad de 0 entre 1) Ejemplo:
	9 -> 1001 -> Longitud= 2
	529 -> 1000010001 -> Longitud= 4 y 3
	15 -> 1111 -> Longitud= 0

- Para convertir un entero a binario, dividir entre dos y hasta que dé uno se tomará el último cociente y los residuos de forma ascendente

- En Golang para convertir un número decimal a binario:
	strconv.FormatInt(int64(numero), 2)

- Para saber el tipo de variable usar reflect.TypeOf(variable)

- Para convertir string a int usar strconv.Atoi("string")

- Para convertir int a string usar Z,_ :=strconv.Itoa(int)

- Con strings.Split("string", "1") Puedo obtener array de subcadenas fraccinadas en cada 1 sin incluirse. Si empieza por 1 incluir los primeros ceros, si termina con 1 incluir los últimos 0.

La función main no está incluida en la evaluación

*/

package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Convertir decimal a Binario
func decimalToBinary(N int) string {
	S := strconv.FormatInt(int64(N), 2)
	return S
}

// Calcula el espacio binario y retorna los ceros en string
func extractZeros(S string) []string {

	// Verificar si solo tiene 1
	if !strings.Contains(S, "0") {
		return nil
	}

	// Crear una subcadena en array de ceros dentro del array
	arrayFail := strings.Split(S, "1")

	// Eliminar espacios vacíos en el array
	var array []string
	for _, i := range arrayFail {
		if i != "" {
			array = append(array, i)
		}
	}

	// Si el binario empieza por cero
	//  eliminar el primer elemento
	if strings.Index(S, "0") == 0 {
		array = append(array[1:])
		fmt.Println("Empieza por cero", S)
	}

	// Si el binaro termina en cero
	//  Eliminar el último elemento
	lastIndexString := len(S) - 1
	lastIndexArray := len(array) - 1
	if strings.LastIndex(S, "0") == lastIndexString {
		array = append(array[:lastIndexArray])
		fmt.Println("Termina en cero", S)
	}

	//fmt.Println("array", array[0])

	return array
}

// Retorna el tamaño binario máximo si aplica
func spaceBinary(M []string) int {

	// Array que guardara la cantidad de ceros por elemento
	var numberGrap []int
	for i := range M {
		numberGrap = append(numberGrap, strings.Count(M[i], "0"))
	}

	// Número máximo de array
	var numberMax int
	for i, v := range numberGrap {
		if v > numberMax {
			numberMax = numberGrap[i]
		}
	}

	return numberMax
}

// Retorna el resultado
func Solution(N int) int {

	// Convertir núnero decimal a binario
	toBinary := decimalToBinary(N)

	// Extraer los ceros del array
	zerosArray := extractZeros(toBinary)

	if zerosArray == nil {
		fmt.Println("Solo hay número 1, no tiene longitud", toBinary)
		return 0
	}

	grap := spaceBinary(zerosArray)

	//	fmt.Printf("Longitud binario máximo es: %d\n", grap)

	return grap
}

func main() {
	N := 15 // 2
	fmt.Println(Solution(N))
	fmt.Println(Solution(9))

}
