package main

import (
	"close-pl-server/internal"
	"os"
)

func main() {
	cookie := os.Getenv("cookie")
	if cookie == "" {
		panic("cookie is empty")
	}
	token, err := internal.getToken(cookie)
	if err != nil {
		panic(err)
	}
	conn, err := internal.getConnection(token)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	internal.loopSendMemory(conn)
}
