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
	tokens := strings.Split(path, "/")
	res := ""

	if path == "/" {
		res = "HTTP/1.1 200 OK\r\n\r\n"
	} else if len(tokens) > 1 && tokens[1] == "echo" {
		res = fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(tokens[2]), tokens[2])
	} else if len(tokens) > 1 && strings.ToLower(tokens[1]) == "user-agent" {
		user_agent := strings.Split(lines[3], " ")[1]
		res = fmt.Sprintf("HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s", len(user_agent), user_agent)
	} else {
		res = "HTTP/1.1 404 Not Found\r\n\r\n"
	}

	conn.Write([]byte(res))
}
