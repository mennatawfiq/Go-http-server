package main

import (
	"fmt"
	"net"
	"strings"
)

func handler(conn net.Conn, dir string) {

	defer conn.Close()

	buf := make([]byte, 1024)
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading connection: ", err.Error())
		return
	}

	req := string(buf)
	lines := strings.Split(req, "\r\n")
	path := strings.Split(lines[0], " ")[1]
	tokens := strings.Split(path, "/")

	res := route(lines, tokens, path, dir)
	conn.Write([]byte(res))

}
