package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"time"
)

const (
	HOST = "localhost"
	PORT = "8080"
	TYPE = "tcp"
)

type CreateSiteRequest struct {
	url string
}

type CreateSiteResponse struct {
	erreur error
}

type GetSiteRequest struct {
	params string
}

func main() {
	listen, err := net.Listen(TYPE, HOST+":"+PORT)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	// close listener
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Fatal(err)
			os.Exit(1)
		}
		go handleRequest(conn)
	}
}

func handleRequest(conn net.Conn) {
	// incoming request
	buffer := make([]byte, 1024)
	qt, err := conn.Read(buffer)
	if err != nil {
		log.Fatal(err)
	}

	// write data to response
	fmt.Printf("Reveived %v\n", string(buffer[0:qt]))
	time := time.Now().Format(time.ANSIC)
	responseStr := fmt.Sprintf("Your message is: %v. Received time: %v", string(buffer[0:qt]), time)
	conn.Write([]byte(responseStr))

	// close conn
	conn.Close()
}
