Aquí tienes una estructura recomendada para manejar los builds de tu proyecto y un script para compilar y organizar los binarios.

---

## 📂 **Estructura de la carpeta de builds**
Organizaremos la carpeta `builds/` con los binarios generados:

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
│── client/
│   ├── main.go
│── go.mod
│── build.sh
```

---

## 📝 **Script `build.sh` para compilar los binarios**
Crea un archivo `build.sh` en la raíz del proyecto:

```bash
nano build.sh
```

Copia y pega el siguiente código:

```bash
#!/bin/bash

# Directorio de builds
BUILD_DIR="builds"
LINUX_DIR="$BUILD_DIR/linux"
WINDOWS_DIR="$BUILD_DIR/windows"
MAC_DIR="$BUILD_DIR/mac"

# Crear las carpetas de builds
mkdir -p "$LINUX_DIR" "$WINDOWS_DIR" "$MAC_DIR"

# Compilar para Linux
echo "🔧 Building for Linux..."
GOOS=linux GOARCH=amd64 go build -o "$LINUX_DIR/server" ./server/main.go
GOOS=linux GOARCH=amd64 go build -o "$LINUX_DIR/client" ./client/main.go

# Compilar para Windows
echo "🔧 Building for Windows..."
GOOS=windows GOARCH=amd64 go build -o "$WINDOWS_DIR/server.exe" ./server/main.go
GOOS=windows GOARCH=amd64 go build -o "$WINDOWS_DIR/client.exe" ./client/main.go

# Compilar para macOS
echo "🔧 Building for macOS..."
GOOS=darwin GOARCH=amd64 go build -o "$MAC_DIR/server" ./server/main.go
GOOS=darwin GOARCH=amd64 go build -o "$MAC_DIR/client" ./client/main.go

echo "✅ Build completed! Check the 'builds/' directory."
```

---

## 🚀 **Ejecutar el script**
1. Dale permisos de ejecución:

   ```bash
   chmod +x build.sh
   ```

2. Ejecútalo para generar los binarios:

   ```bash
   ./build.sh
   ```

Después de ejecutar el script, los binarios estarán organizados dentro de la carpeta `builds/` listos para su distribución. 🚀