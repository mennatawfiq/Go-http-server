package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	defer l.Close()

	conn, er := l.Accept()
	if er != nil {
		fmt.Println("Error accepting connection: ", err.Error())
		os.Exit(1)
	}

	buf := make([]byte, 1024)
	_, err = conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading connection: ", err.Error())
		os.Exit(1)
	}

	req := string(buf)
	lines := strings.Split(req, "\r\n")
	path := strings.Split(lines[0], " ")[1]
	res := ""
	if path == "/" {
		res = "HTTP/1.1 200 OK\r\n\r\n"
	} else {
		res = "HTTP/1.1 404 Not Found\r\n\r\n"
	}

	conn.Write([]byte(res))

}
