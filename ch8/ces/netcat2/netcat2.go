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

	go mustCopy(os.Stdout, conn, "listen to network input: if network input closed, eof will be sent")
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

	// if EOF were read, jump out of blocking call: Copy()
	n, err := io.Copy(dst, src)
	fmt.Println(n, who)

	if n == 0 {
		fmt.Println("user closed the program with ctrl-c.", err)
	}

	if err != nil {
		log.Fatal(err)
	}
}
