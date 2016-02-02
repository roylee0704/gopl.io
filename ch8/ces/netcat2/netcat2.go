package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn := mustDial("tcp", "localhost:8080")
	defer func() {
		conn.Close()
	}()

	// go command is just to `unblock` io.read()
	// flipping them doesn't affect the running state.

	// Mumbling
	// but it affect the closing state. i.e if we go routine stdin,
	// the main program probably still be running?
	go mustCopy(os.Stdout, conn, "listen to network input")
	mustCopy(conn, os.Stdin, "listen to user input")
}
func mustDial(network string, address string) net.Conn {
	conn, err := net.Dial(network, address)

	if err != nil {
		log.Fatal(err)
	}

	return conn
}

func mustCopy(dst io.Writer, src io.Reader, who string) {

	n, err := io.Copy(dst, src)
	fmt.Println(n, who)
	//question, i close server connection, n is returned, does it contains err?
	if err != nil {
		log.Fatal(err)
	}
	if n == 0 {
		fmt.Println("user closed the program with ctrl-c.", err)

		return
	}

}
