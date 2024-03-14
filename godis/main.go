package main

import (
	"log"
	"net"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func main() {
	// Initialize TCP Server
	log.Println("Starting TCP Server on localhost:6379")
	l, err := net.Listen("tcp", "localhost:6379")
	if err != nil {
		log.Fatal("Error starting TCP Server: ", err)
	}

	// Accept incoming connection
	_, err = l.Accept()
	if err != nil {
		log.Fatal("Error accepting incoming connection: ", err)
	}

}
