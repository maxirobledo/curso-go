# Go Essentials

---

## Concurrency

Una de las principales ventajas de **Go** es su modelo de **concurrencia**, que permite ejecutar múltiples tareas de forma **simultánea y coordinada**, de manera simple y eficiente.

> **Nota:** Concurrencia no es lo mismo que paralelismo.
> Go ejecuta código concurrente y puede hacerlo en paralelo si el runtime y el hardware lo permiten.

---

## Goroutines

Una **goroutine** es una función que se ejecuta de manera concurrente con otras funciones.

Se crea anteponiendo la palabra clave `go` a la llamada de una función.

```go
func greet(phrase string) {
	fmt.Println("Hello!", phrase)
}

func slowGreet(phrase string) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello!", phrase)
}

func main() {
	go greet("Nice to meet you!")
	go greet("How are you?")
	go slowGreet("How... are... you...?")
	go greet("I hope you're liking the course")
}
```

En este ejemplo, las funciones se ejecutan como goroutines, pero **el programa puede finalizar antes de que terminen**, ya que la función `main()` no espera su ejecución.

Cuando `main()` finaliza, **el proceso termina y todas las goroutines se detienen**, incluso si aún no completaron su trabajo.

---

## Channels

Un **channel** es un mecanismo de comunicación entre goroutines que permite:

* Enviar y recibir valores
* Sincronizar la ejecución
* Coordinar cuándo una goroutine ha finalizado

Los channels son **tipados**, por ejemplo: `chan bool`, `chan int`, `chan string`.

---

### Sincronización con Channels

En el siguiente ejemplo, usamos un channel para notificar a `main()` cuando una goroutine termina su ejecución.

```go
package main

import (
	"fmt"
	"time"
)

func greet(phrase string, done chan bool) {
	fmt.Println("Hello!", phrase)
	done <- true
}

func slowGreet(phrase string, done chan bool) {
	time.Sleep(3 * time.Second)
	fmt.Println("Hello!", phrase)
	done <- true
}

func main() {
	done := make(chan bool)

	go greet("Nice to meet you!", done)
	go greet("How are you?", done)
	go slowGreet("How... are... you...?", done)
	go greet("I hope you're liking the course", done)

	// Esperamos a que todas las goroutines envíen su señal
	for i := 0; i < 4; i++ {
		<-done
	}
}
```

### Puntos importantes:

* Cada goroutine envía un valor al channel cuando finaliza
* `main()` espera recibir exactamente la cantidad de señales esperadas
* **El channel no se cierra desde las goroutines**
* El control de sincronización queda en `main()`

---

## Conclusión

* Las goroutines permiten ejecutar código concurrente de forma simple
* Los channels permiten comunicar y sincronizar goroutines
* El flujo del programa debe garantizar que `main()` espere la finalización de las goroutines
* El cierre de un channel debe hacerse únicamente desde quien controla todos los envíos
