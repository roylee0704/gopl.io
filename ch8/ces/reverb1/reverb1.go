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
	scanner := bufio.NewScanner(c)
	for scanner.Scan() {
		go echo(c, scanner.Text(), 3*time.Second)
	}

	// NOTE: read error
	if err := scanner.Err(); err != nil {
		fmt.Printf("Invalid read: %s\n", err)
	}
}

func echo(c net.Conn, shout string, delay time.Duration) {
	_, err := fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	fmt.Fprintln(os.Stdout, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	_, err = fmt.Fprintln(c, "\t", shout)
	fmt.Fprintln(os.Stdout, "\t", shout)
	time.Sleep(delay)
	_, err = fmt.Fprintln(c, "\t", strings.ToLower(shout))
	fmt.Fprintln(os.Stdout, "\t", strings.ToUpper(shout))

	// NOTE: write error
	if err != nil {
		fmt.Printf("Invalid write: %s\n", err)
	}
}
