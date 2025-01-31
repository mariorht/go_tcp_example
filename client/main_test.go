package main

import (
	"net"
	"testing"
	"time"
)

// Inicia un servidor TCP de prueba en un puerto aleatorio
func startMockServer() (net.Listener, error) {
	listener, err := net.Listen("tcp", "localhost:0") // Puerto aleatorio
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				return // Terminar la goroutine si hay un error
			}
			go func(c net.Conn) {
				defer c.Close()
				for {
					buffer := make([]byte, 4096)
					n, err := c.Read(buffer)
					if err != nil {
						return
					}
					c.Write(buffer[:n])
				}
			}(conn)
		}
	}()

	return listener, nil
}

// Test para verificar que el cliente puede conectarse al servidor de prueba
func TestClientConnection(t *testing.T) {
	listener, err := startMockServer()
	if err != nil {
		t.Fatalf("No se pudo iniciar el servidor de prueba: %v", err)
	}
	defer listener.Close()

	addr := listener.Addr().String() // Obtener la dirección del servidor
	time.Sleep(500 * time.Millisecond) // Esperar un poco para que el servidor esté listo

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		t.Fatalf("El cliente no pudo conectarse: %v", err)
	}
	conn.Close()
}


// Test para verificar la integridad del mensaje enviado para diferentes tamaños de mensajes
func TestMessageIntegrity(t *testing.T) {
	listener, err := startMockServer()
	if err != nil {
		t.Fatalf("No se pudo iniciar el servidor de prueba: %v", err)
	}
	defer listener.Close()

	addr := listener.Addr().String() // Obtener la dirección del servidor
	time.Sleep(500 * time.Millisecond) // Esperar un poco para que el servidor esté listo

	tests := []struct {
		name    string
		message string
	}{
		{"ShortMessage", "Short message"},
		{"MediumMessage", "This is a medium length message for testing purposes."},
		{"LongMessage", "This is a much longer message intended to test the integrity of the message transmission over the TCP connection. It should be long enough to ensure that the message is properly handled by the server and client without any truncation or corruption."},
		{"SuperLongMessage", string(make([]byte, 2*1024*1024))}, // Mensaje de 2MB
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			conn, err := net.Dial("tcp", addr)
			if err != nil {
				t.Fatalf("El cliente no pudo conectarse: %v", err)
			}
			defer conn.Close()

			_, err = conn.Write([]byte(tt.message + "\n"))
			if err != nil {
				t.Fatalf("Error al enviar el mensaje: %v", err)
			}

			var receivedMessage []byte
			buffer := make([]byte, 4096)
			for {
				n, err := conn.Read(buffer)
				if err != nil {
					break
				}
				receivedMessage = append(receivedMessage, buffer[:n]...)
				if len(receivedMessage) >= len(tt.message)+1 {
					break
				}
			}

			if string(receivedMessage) != tt.message+"\n" {
				t.Errorf("El mensaje recibido no coincide. Esperado: %s, Recibido: %s", tt.message, string(receivedMessage))
			}
		})
	}
}

/////////////////////////////////////////////////////////////////////////////////////
// Benchmarks


// Benchmark para medir el tiempo de conexión del cliente
func BenchmarkClientConnection(b *testing.B) {
	listener, err := startMockServer()
	if err != nil {
		b.Fatalf("No se pudo iniciar el servidor de prueba: %v", err)
	}
	defer listener.Close()

	addr := listener.Addr().String() // Obtener la dirección del servidor
	time.Sleep(500 * time.Millisecond) // Esperar a que el servidor se inicie

	b.ResetTimer() // Resetear el timer para medir solo las iteraciones

	for i := 0; i < b.N; i++ {
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			b.Fatalf("Error al conectar al servidor de prueba: %v", err)
		}
		conn.Close()
	}
}