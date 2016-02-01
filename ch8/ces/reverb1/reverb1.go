package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8080")

	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Print(err)
			continue //skip to the next connection
		}

		go handleConn(conn)

	}
}

// handleConn is one single gopher,
// and go command enables to duplicate gopher. (meditate and imagine)
func handleConn(c net.Conn) {
	defer c.Close()

	//take the conn (aka phone)
	input := bufio.NewScanner(c)
	for input.Scan() {
		go echo(c, input.Text(), 1*time.Second)
	}
	// NOTE: ignoring potential errors from input.Err()
}

func echo(c net.Conn, shout string, delay time.Duration) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	fmt.Fprintln(os.Stdout, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	fmt.Fprintln(os.Stdout, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
	fmt.Fprintln(os.Stdout, "\t", strings.ToUpper(shout))
}
