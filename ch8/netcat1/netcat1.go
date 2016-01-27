package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	mustCopy(os.Stdout, conn)

}

// mustCopy is another example of utility (i.e: io package) that orchestrates between reader & writer
// meaning: a wrapper to existing implementations
func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
