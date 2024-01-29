package main

import (
	"github.com/gorilla/websocket"
	"log"
)

func sendCommend(conn *websocket.Conn, body Body) Body {
	err := conn.WriteJSON(body)
	if err != nil {
		log.Fatal(err)
	}
	var response Body
	err = conn.ReadJSON(&response)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(response)
	return response
}

func getResp(conn *websocket.Conn) Body {
	var response Body
	err := conn.ReadJSON(&response)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(response)
	return response
}
