# Go Essentials

---

## Arrays, slice y maps

### Arrays

Un **array** es una estructura de datos que contiene valores pontencialmente diferente pero que describen la misma cosa.

sintaxis

```go
prices := [4]float64{10.99, 9.99, 45.99, 20.0}
```
Esto define que se almacenaran cuatro valores de punto flotante 

Los valores del array tienen una posicion dentro del array llamado **indice**, empezando por **0**. A travez del indice se puede acceder a un valor determinado.

Por ejemplo, quiero acceder al valor del array que se encuentra en la posicion numero 3

```go
fmt.Println("Valor en la posicion numero 3: ", prices[2])
```

La caracteristica de **slice** en arrays permite dividir el array indicando la posicion del indice desde(incluido) .. hasta(excluido)

Por ejemplo, seleccionar los precios desde la posicion 2 hasta la posicion 3 

```go
prices := [4]float64{10.99, 9.99, 45.99, 20.0}

newPrices := prices[1:3]
```

---
