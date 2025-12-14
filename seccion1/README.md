# Introducción a Go

Go es un lenguaje de **programación open‑source** desarrollado por **Google**, diseñado para ser simple, eficiente y fácil de mantener.

---

## Características principales

* **Open Source**: el código fuente de Go es abierto y mantenido por la comunidad junto con Google.
* **Lenguaje compilado**: el código Go se compila a binarios nativos.
* **Tipado fuerte**: Go es un lenguaje **tipado**, lo que significa que respeta estrictamente el tipo de los valores y variables. Esto ayuda a detectar errores en tiempo de compilación.

Ejemplo:

```go
var edad int = 30
// edad = "treinta"  // ❌ Error de compilación
```

---

## Ejecutar un programa en Go desde la terminal

Para ejecutar un programa Go directamente desde la terminal sin generar un binario, se utiliza el comando:

```bash
go run app.go
```

Este comando:

* Compila el código de forma temporal
* Ejecuta el programa
* No deja archivos binarios generados

Es ideal para **pruebas rápidas y aprendizaje**.

---

Más adelante, cuando el proyecto crezca, se utilizará `go build` para generar ejecutables reutilizables.
