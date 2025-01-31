# Tests y Benchmarks

Este documento explica cómo se han implementado los tests y benchmarks para el proyecto Go TCP Example.

## Estructura de los Tests

Los tests están organizados en los siguientes archivos:

- `server/main_test.go`: Contiene los tests y benchmarks para el servidor TCP.
- `client/main_test.go`: Contiene los tests y benchmarks para el cliente TCP.

## Ejecutar Tests

Para ejecutar los tests, utiliza el comando `go test` en la raíz del proyecto o en los directorios específicos:

```sh
go test ./server -v
go test ./client -v
```

El flag `-v` es para obtener una salida detallada de los tests.

## Ejecutar Benchmarks

Para ejecutar los benchmarks, utiliza el comando `go test` con el flag `-bench`:

```sh
go test -bench=. ./server
go test -bench=. ./client
```

El flag `-bench=.` ejecuta todos los benchmarks en el paquete especificado.

## Script `test.sh`

Para facilitar la ejecución de los tests y benchmarks, se ha creado un script `test.sh`:

```sh
#!/bin/bash

echo "========== 🔍 Ejecutando tests... ============"
go test ./server -v
go test ./client -v

echo "========== 🏎 Ejecutando benchmarks... ======="
echo ""
go test -bench=. ./server
go test -bench=. ./client
```

Para ejecutar el script, primero asegúrate de que tenga permisos de ejecución:

```sh
chmod +x test.sh
```

Luego, ejecuta el script:

```sh
./test.sh
```

## Detalles de Implementación

### Tests del Servidor

En `server/main_test.go`, se ha implementado un test para verificar que el servidor se inicia correctamente y un benchmark para medir el tiempo de conexión al servidor.

### Tests del Cliente

En `client/main_test.go`, se ha implementado un test para verificar que el cliente puede conectarse al servidor de prueba y un benchmark para medir el tiempo de conexión del cliente.

---

Con esta guía, deberías poder ejecutar y entender los tests y benchmarks del proyecto. 🚀
