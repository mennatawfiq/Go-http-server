package main

import (
	"fmt"
	"net"
	"os"
)

func main() {

	l, err := net.Listen("tcp", "0.0.0.0:9999")
	if err != nil {
		fmt.Println("Failed to bind to port 9999")
		os.Exit(1)
	}

	dir := ""
	for i, arg := range os.Args {
		if arg == "--directory" && i+1 < len(os.Args) {
			dir = os.Args[i+1]
		}
	}

	defer l.Close()

	for {

		conn, er := l.Accept()
		if er != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			continue
		}

		go handler(conn, dir)

	}

}
