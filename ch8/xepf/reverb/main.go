package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

var port = flag.String("port", "8080", "port number to be listened")

func main() {

	flag.Parse()

	ln, err := net.Listen("tcp", ":"+*port)

	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := ln.Accept()
		log.Println("new connection...")
		if err != nil {
			log.Fatal(err)
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {

	var wg sync.WaitGroup
	input := bufio.NewScanner(conn)
	doneWriting := make(chan struct{})

	// each scan equilv to 1 echo job
	for input.Scan() {
		wg.Add(1)
		go func() {
			echo(conn, input.Text(), 1*time.Second, &wg)
		}()
	}

	go func() {
		wg.Wait()
		fmt.Println("the wait is done")
		doneWriting <- struct{}{}
	}()

	fmt.Println("waiting all echo worker to finish")
	<-doneWriting
	fmt.Println("okay! they have completed! I can safely closeWrite now!")
	conn.(*net.TCPConn).CloseWrite()
}

// echo echoes "scream" for 3 times, each echo delays for x duration, and finally,
// report to waitgroup that job is completed.
func echo(conn net.Conn, scream string, delay time.Duration, wg *sync.WaitGroup) {
	fmt.Fprintln(conn, "\t", strings.ToUpper(scream))
	time.Sleep(delay)
	fmt.Fprintln(conn, "\t", strings.ToUpper(scream))
	time.Sleep(delay)
	fmt.Fprintln(conn, "\t", strings.ToUpper(scream))
	time.Sleep(delay)
	wg.Done() //test *wg or wg
}
