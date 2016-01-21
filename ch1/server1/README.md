#Server 1
Server1 is a minimal "echo" server.

### Getting Started
1. `go run Server1.go &`  // to run in background
2. Visit `localhost8080` in your browser
3. Send message to Server by `/your/message/here`.


### To Kill Server
1. `echo $!`    // to print the PID of the process
2. `kill <PID>` // e.g: `kill 14425`
