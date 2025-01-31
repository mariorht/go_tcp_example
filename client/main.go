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
