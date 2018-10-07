# Soluciones de codility usando en lenguaje Go
- [x] DemoTicket (Test Score= 100%)
- [x] Iterations
  - [x] BinaryGap
- [ ] Arrays
  - [ ] CyclicRotation  
  - [ ] OddOccurrencesInArray
- [ ] Time Complexility
   - [ ] FrogJmp
   - [ ] PermMissingElem
   - [ ] TapeEquilibrium

***

# Notas importantes del lenguaje de Go

## 1

Para romper un slice en un valor y continuar en otro, se usa de la misma manera que cuando adicionamos elementos:
```
E := []int{}
E = append(E[:4], E[8:]...)
```

En Go los ***...*** representan como un bucle for pero continuado de una forma implícita.

## 2

Los ***<<*** y ***>>*** son operadores usados para crear númeross.

El operador ***n << m*** es equivalente a "n multiplicado por 2, m veces" = n * 2^m.
El operador ***n >> m*** es equivalente a "n dividido por 2, m veces" = n * 2^(-m).

## 3

En Golang para convertir un número decimal a binario:
	strconv.FormatInt(int64(numero), 2)

Para saber el tipo de variable usar reflect.TypeOf(variable)

Para convertir string a int usar strconv.Atoi("string")

Para convertir int a string usar Z,_ :=strconv.Itoa(int)

## 4

%q En Printf para incluir arrays

