# Go TCP Example

Este proyecto es un ejemplo de un servidor y un cliente TCP implementados en Go.

## Estructura del Proyecto

```
go-tcp-example/
│── builds/
│   ├── linux/
│   │   ├── server
│   │   ├── client
│   ├── windows/
│   │   ├── server.exe
│   │   ├── client.exe
│   ├── mac/
│   │   ├── server
│   │   ├── client
│── server/
│   ├── main.go
│   ├── main_test.go
│── client/
│   ├── main.go
│   ├── main_test.go
│── go.mod
│── build.sh
│── README.md
│── builds.md
│── tutorial.md
```

## Archivos Principales

- `server/main.go`: Implementación del servidor TCP.
- `client/main.go`: Implementación del cliente TCP.
- `build.sh`: Script para compilar los binarios para Linux, Windows y macOS.
- `go.mod`: Archivo de módulo de Go.
- `builds.md`: [Guía para manejar los builds](builds.md).
- `tutorial.md`: [Tutorial paso a paso para configurar el proyecto](tutorial.md).
- `tests.md`: [Guía para ejecutar tests y benchmarks](tests.md).

## Compilación

Para compilar los binarios para diferentes sistemas operativos, ejecuta el script `build.sh`:

```sh
chmod +x build.sh
./build.sh
```

Los binarios generados se almacenarán en la carpeta `builds/`.

## Ejecución

### Servidor

Para ejecutar el servidor, navega a la carpeta `builds/linux/` y ejecuta:

```sh
./server
```

### Cliente

Para ejecutar el cliente, navega a la carpeta `builds/linux/` y ejecuta:

```sh
./client
```

## Tests y Benchmarks

Para ejecutar los tests y benchmarks, utiliza el script `test.sh`:

```sh
chmod +x test.sh
./test.sh
```

Esto ejecutará los tests y benchmarks para el servidor y el cliente.

## Enlaces Útiles

- [Guía para manejar los builds](builds.md)
- [Tutorial paso a paso para configurar el proyecto](tutorial.md)
- [Guía para ejecutar tests y benchmarks](tests.md)

## Licencia

Este proyecto está licenciado bajo la Licencia MIT. Consulta el archivo [LICENSE](LICENSE) para más detalles.
