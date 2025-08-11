package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func route(lines, tokens []string, path, dir string) string {

	res := ""

	if path == "/" {
		res = "HTTP/1.1 200 OK\r\n\r\n"
	} else if len(tokens) > 1 && tokens[1] == "echo" {
		res = fmt.Sprintf(
			"HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s",
			len(tokens[2]),
			tokens[2],
		)
	} else if len(tokens) > 1 && strings.ToLower(tokens[1]) == "user-agent" {
		user_agent := ""
		for _, line := range lines {
			if strings.HasPrefix(strings.ToLower(line), "user-agent:") {
				user_agent = strings.TrimSpace(strings.TrimPrefix(line, "User-Agent:"))
				break
			}
		}
		res = fmt.Sprintf(
			"HTTP/1.1 200 OK\r\nContent-Type: text/plain\r\nContent-Length: %d\r\n\r\n%s",
			len(user_agent),
			user_agent,
		)
	} else if len(tokens) > 1 && tokens[1] == "files" {
		// GET /files/foo HTTP/1.1
		if len(tokens) < 3 {
			res = "HTTP/1.1 404 Not Found\r\n\r\n"
		}
		file_name := tokens[2]
		full_path := filepath.Join(dir, file_name)
		file, er := os.ReadFile(full_path)
		if er != nil {
			res = "HTTP/1.1 404 Not Found\r\n\r\n"
		}
		header := fmt.Sprintf(
			"HTTP/1.1 200 OK\r\nContent-Type: application/octet-stream\r\nContent-Length: %d\r\n\r\n",
			len(file),
		)
		res = header + string(file)
	} else {
		res = "HTTP/1.1 404 Not Found\r\n\r\n"
	}

	return res
}
