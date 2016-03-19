package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

func mustWrite(conn net.Conn) error {
	if n, err := io.Copy(conn, os.Stdin); err != nil {
		return fmt.Errorf("write to conn: %d bytes: %v", n, err) //only err when you send again after close conn.
	}
	return nil
}

func main() {
	c, err := net.Dial("tcp", ":8080")
	conn := c.(*net.TCPConn)

	if err != nil {
		log.Fatal(err)
	}

	readCompleted := make(chan struct{})
	go func() {
		io.Copy(os.Stdout, conn)
		fmt.Println("okay folk, the server has just shut-down the write conn")
		// you will never reach here, unless server closeWrite
		readCompleted <- struct{}{}
	}()

	mustWrite(conn)
	conn.CloseWrite()
	<-readCompleted
	fmt.Println("cool, let's go home, the show has ended")

}
