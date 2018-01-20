# Soluciones de codility usando en lenguaje Go
- [x] DemoTicket (Test Score= 100%)

***

# Notas importantes del lenguaje de Go
Para romper un slice en un valor y continuar en otro, se usa de la misma manera que cuando adicionamos elementos:
```
E := []int{}
E = append(E[:4], E[8:]...)
```

En Go los ***...*** representan como un bucle for pero continuado de una forma implícita.

***

Los ***<<*** y ***>>*** son operadores usados para crear númeross.

El operador ***n << m*** es equivalente a "n multiplicado por 2, m veces" = n * 2^m.
El operador ***n >> m*** es equivalente a "n dividido por 2, m veces" = n * 2^(-m).
