package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn := mustDial("tcp", "localhost:8080")
	defer conn.Close()

	// go command is just to `unblock` io.read()
	// flipping them doesn't matter.
	go mustCopy(conn, os.Stdin)
	mustCopy(os.Stdout, conn)
}
func mustDial(network string, address string) net.Conn {
	conn, err := net.Dial(network, address)

	if err != nil {
		log.Fatal(err)
	}

	return conn
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}

}
