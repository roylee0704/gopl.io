// clock2 accepts more than one client connection, simutaneously.
package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {

	listener, err := net.Listen("tcp", "localhost:4747")

	if err != nil {
		log.Fatal(err)
		return
	}

	for {
		conn, err := listener.Accept()

		if err != nil {
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}

func handleConn(c net.Conn) {
	defer c.Close()
	for {
		if _, err := io.WriteString(c, time.Now().Format("15:04:05\n")); err != nil { // client disconnected
			log.Print(err)
			return
		}
		time.Sleep(1 * time.Second)
	}
}
