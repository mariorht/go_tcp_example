package main

import (
	"net"
	"testing"
	"time"
)

// Inicia un servidor TCP en un puerto aleatorio
func startTestServer() (net.Listener, error) {
	listener, err := net.Listen("tcp", "localhost:0") // Puerto aleatorio
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				return // Salir si hay un error
			}
			conn.Close() // No hacemos nada en los tests, solo aceptar conexiones
		}
	}()

	return listener, nil
}

// Test para verificar que el servidor se inicia correctamente
func TestServerStart(t *testing.T) {
	listener, err := startTestServer()
	if err != nil {
		t.Fatalf("No se pudo iniciar el servidor de prueba: %v", err)
	}
	defer listener.Close()

	addr := listener.Addr().String()
	time.Sleep(500 * time.Millisecond) // Esperar un poco para que el servidor esté listo

	conn, err := net.Dial("tcp", addr)
	if err != nil {
		t.Fatalf("No se pudo conectar al servidor: %v", err)
	}
	conn.Close()
}

// Benchmark de conexión al servidor
func BenchmarkServerConnection(b *testing.B) {
	listener, err := startTestServer()
	if err != nil {
		b.Fatalf("No se pudo iniciar el servidor de prueba: %v", err)
	}
	defer listener.Close()

	addr := listener.Addr().String()
	time.Sleep(500 * time.Millisecond) // Esperar un poco para que el servidor esté listo

	b.ResetTimer() // Resetear el timer para medir solo las iteraciones

	for i := 0; i < b.N; i++ {
		conn, err := net.Dial("tcp", addr)
		if err != nil {
			b.Fatalf("Error al conectar al servidor de prueba: %v", err)
		}
		conn.Close()
	}
}
