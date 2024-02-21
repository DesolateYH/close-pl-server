package main

import (
	"close-pl-server/internal/connection"
	"close-pl-server/internal/processor"
	"os"
)

func main() {
	cookie := os.Getenv("cookie")
	if cookie == "" {
		panic("cookie is empty")
	}
	conn, err := connection.NewConnection(cookie)
	if err != nil {
		panic(err)
	}
	proc := processor.NewProcessor(conn)

	proc.Run()
}
