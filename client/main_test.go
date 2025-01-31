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
			conn.Write([]byte("Test message\n"))
			conn.Close()
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
