package main

import (
	"net"
	"os"
)

const (
	HOST = "localhost"
	PORT = "8080"
	TYPE = "tcp"
)

type GetSiteRequest struct {
	params string
}

type GetSiteResponse struct {
	url string
	erreur error
}

type UpdateSiteRequest struct {
	url string
	visited bool
	date string
}

type UpdateSiteReponse struct {
	erreur error
}

type CreateFileRequest struct {
	url string
	infos string
}

type CreateFileResponse struct {
	erreur error
}

type GetFileRequest struct {
	params string
}

type GetFileResponse struct {
	files []string
	erreur error
}

type UpdateFileRequest struct {
	url string
	infos string
}

type UpdateFileResponse struct {
	erreur error
}

func main() {
	tcpServer, err := net.ResolveTCPAddr(TYPE, HOST+":"+PORT)

	if err != nil {
		println("ResolveTCPAddr failed:", err.Error())
		os.Exit(1)
	}

	conn, err := net.DialTCP(TYPE, nil, tcpServer)
	if err != nil {
		println("Dial failed:", err.Error())
		os.Exit(1)
	}

	_, err = conn.Write([]byte("This is a message"))
	if err != nil {
		println("Write data failed:", err.Error())
		os.Exit(1)
	}

	// buffer to get data
	received := make([]byte, 1024)
	_, err = conn.Read(received)
	if err != nil {
		println("Read data failed:", err.Error())
		os.Exit(1)
	}

	println("Received message:", string(received))

	conn.Close()
}
