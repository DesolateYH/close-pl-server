package main

import "os"

func main() {
	cookie := os.Getenv("cookie")
	if cookie == "" {
		panic("cookie is empty")
	}
	token, err := getToken(cookie)
	if err != nil {
		panic(err)
	}
	conn, err := getConnection(token)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	loopSendMemory(conn)
}
