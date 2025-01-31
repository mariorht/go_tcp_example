## 🔹 Paso 1: Instalar Go en Ubuntu
### 1.1. Descargar e instalar Go
Abre una terminal y ejecuta los siguientes comandos para instalar Go:

```bash
sudo apt update
sudo apt install -y golang
```

### 1.2. Verificar la instalación
Para comprobar que Go está instalado correctamente, ejecuta:

```bash
go version
```

Debería mostrar algo como:

```
go version go1.21.4 linux/amd64
```

Si necesitas la última versión de Go, descárgala desde la [página oficial](https://go.dev/dl/) y sigue sus instrucciones.

---

## 🔹 Paso 2: Configurar el entorno de desarrollo en VS Code
### 2.1. Instalar VS Code (si no lo tienes)
Si aún no tienes instalado Visual Studio Code, puedes instalarlo con:

```bash
sudo snap install --classic code
```

O descargarlo desde la [página oficial](https://code.visualstudio.com/Download).

### 2.2. Instalar la extensión de Go
Abre VS Code y busca la extensión **"Go"** de **Go Team at Google** en la pestaña de extensiones (Ctrl + Shift + X). Instálala.


## 🔹 Paso 3: Crear el proyecto en Go
Vamos a estructurar nuestro proyecto de la siguiente manera:

```
go-tcp-example/
│── server/
│   ├── main.go
│── client/
│   ├── main.go
│── go.mod
```

### 3.1. Crear la carpeta del proyecto
```bash
mkdir go-tcp-example
cd go-tcp-example
```

### 3.2. Inicializar un módulo Go
```bash
go mod init example.com/tcp
```

Esto creará un archivo `go.mod`, que es necesario para manejar dependencias.

---

## 🔹 Paso 4: Crear el servidor TCP
Crea la carpeta `server` y el archivo `main.go`:

```bash
mkdir server
nano server/main.go
```

Copia y pega el siguiente código:

```go
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Client connected:", conn.RemoteAddr())

	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		message := scanner.Text()
		fmt.Println("Received:", message)

		// Responder al cliente
		conn.Write([]byte("Message received: " + message + "\n"))
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading:", err)
	}
}

func main() {
	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
	defer listener.Close()
	fmt.Println("Server listening on port 9000...")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn) // Manejar cada cliente en una goroutine
	}
}
```

Guárdalo y sal (`Ctrl + X`, `Y`, `Enter`).

---

## 🔹 Paso 5: Crear el cliente TCP
Crea la carpeta `client` y el archivo `main.go`:

```bash
mkdir client
nano client/main.go
```

Copia y pega este código:

```go
package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:9000")
	if err != nil {
		fmt.Println("Error connecting:", err)
		os.Exit(1)
	}
	defer conn.Close()

	fmt.Println("Connected to server. Type messages:")
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		text := scanner.Text()
		_, err := fmt.Fprintf(conn, text+"\n")
		if err != nil {
			fmt.Println("Error sending message:", err)
			break
		}

		// Recibir respuesta del servidor
		reply, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("Server reply:", reply)
	}
}
```

Guárdalo y sal.

---

## 🔹 Paso 6: Ejecutar el servidor y el cliente
### 6.1. Iniciar el servidor
Abre una terminal y ejecuta:

```bash
cd go-tcp-example/server
go run main.go
```

Verás un mensaje indicando que el servidor está escuchando en el puerto `9000`.

### 6.2. Iniciar el cliente
En otra terminal, ejecuta:

```bash
cd go-tcp-example/client
go run main.go
```

Escribe mensajes en la terminal del cliente y verás cómo el servidor los recibe y responde.

---

Con esto, ya tienes un servidor y un cliente TCP funcional en Go. 🚀
