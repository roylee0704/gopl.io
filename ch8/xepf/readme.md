#Exercise 8.4 & 8.5

Check the exercise from the book: The Go programming language.

## Abstract

The exercise demonstrates the sychronizations between multi go-routines that runs concurrently.

Basically, there are 2 programs in the folder, each represents server (reverb), and another one represents as client (nc), which communicates via tcp socket.

The challenge is, how do you properly handle tcp close when connection is no longer needed between client & server.

### Server 

It is capable of handling request from multiple clients, in other words, one `thread` is created to serve(write/read) one client.

### Client




