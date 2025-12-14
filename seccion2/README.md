# Go Essentials

Este documento resume los conceptos fundamentales de Go vistos en las primeras secciones del curso.

## Contenidos

* Componentes claves
* Valores y tipos
* Variables y constantes
* Funciones
* Estructuras de control

---

## Componentes claves

### `package main`

El `package main` indica a Go que este paquete es el **punto de entrada** de la aplicación que estamos construyendo. Todo programa ejecutable en Go debe tener:

* Un paquete llamado `main`
* Una función `main()`

Ejemplo:

```go
package main

func main() {
    // punto de entrada de la aplicación
}
```

---

### Módulos en Go

Un **módulo** es un proyecto Go que puede contener **múltiples packages**. Los módulos permiten:

* Gestionar dependencias
* Versionar el código
* Definir el alcance del proyecto

Para inicializar un módulo se utiliza:

```bash
go mod init <path>
```

Ejemplo:

```bash
go mod init github.com/maxirobledo/curso-go
```

Donde:

* `<path>` es la ruta donde el módulo puede ser encontrado (normalmente un repositorio Git)

Este comando crea el archivo `go.mod`, que define:

* El nombre del módulo
* La versión de Go
* Las dependencias del proyecto

---

### Build de un ejecutable

Una vez definido el módulo, se puede crear un **archivo ejecutable** de Go para que el programa pueda ejecutarse en sistemas que **no tengan Go instalado**:

```bash
go build
```

Este comando genera un binario en el directorio actual (por ejemplo, `.exe` en Windows).

---

## Variables y convenciones

### Convención de nombres

Para nombrar variables en Go se utiliza **camelCase**:

```go
var estoEsUnaVariable int
var variableUno string
var nombreApellido string
```

---

### Declaración de variables

Forma estándar:

```go
var tasaRendimiento float64 = 5.5
```

Forma corta (recomendada dentro de funciones):

```go
tasaRendimiento := 5.5
```

> La sintaxis corta (`:=`) infiere automáticamente el tipo de la variable.

---

### Valores por defecto (zero values)

Cuando una variable es declarada pero no inicializada, Go le asigna un **valor por defecto** según su tipo:

* `int` → `0`
* `float64` → `0.0`
* `string` → `""` (string vacío)
* `bool` → `false`

Ejemplo:

```go
var edad int        // 0
var promedio float64 // 0.0
var nombre string   // ""
var activo bool     // false
```

---

Estos conceptos forman la base para comenzar a desarrollar aplicaciones en Go de manera clara, tipada y predecible.
