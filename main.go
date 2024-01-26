package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"os"
	"time"
)

type Body struct {
	Event string   `json:"event,omitempty"`
	Args  []string `json:"args,omitempty"`
}

func CheckError(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func main() {
	dialer := &websocket.Dialer{}
	header := http.Header{}

	header.Set("User-Agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:123.0) Gecko/20100101 Firefox/123.0")
	header.Set("Accept", "*/*")
	header.Set("Accept-Language", "zh-CN,zh;q=0.8,zh-TW;q=0.7,zh-HK;q=0.5,en-US;q=0.3,en;q=0.2")
	header.Set("Accept-Encoding", "gzip, deflate, br")
	header.Set("Origin", "https://panel.vatzj.com")
	//header.Set("Sec-WebSocket-Extensions", "permessage-deflate")
	//header.Set("Sec-WebSocket-Key", "/SDUSD46UTvtM6LPdrlOSA==")
	header.Set("DNT", "1")
	header.Set("Sec-GPC", "1")
	header.Set("Sec-Fetch-Dest", "empty")
	header.Set("Sec-Fetch-Mode", "websocket")
	header.Set("Sec-Fetch-Site", "same-site")
	header.Set("Pragma", "no-cache")
	header.Set("Cache-Control", "no-cache")

	conn, resp, err := dialer.Dial("wss://pt-50.vatzj.com:8082/api/servers/661a539a-b5b5-4955-bcf6-737740a6b270/ws", header)
	CheckError(err)
	for key, values := range resp.Header {
		fmt.Printf("%s:%v\n", key, values)
	}
	defer conn.Close()

	authReq := Body{
		Event: "auth",
		Args: []string{
			"eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiIsImp0aSI6IjZhNjg1ZGEyY2RiYTE1OWU3Y2Q0ZmVkM2ZmYmVmNTI0In0.eyJpc3MiOiJodHRwczovL3BhbmVsLnZhdHpqLmNvbSIsImF1ZCI6WyJodHRwczovL3B0LTUwLnZhdHpqLmNvbTo4MDgyIl0sImp0aSI6IjZhNjg1ZGEyY2RiYTE1OWU3Y2Q0ZmVkM2ZmYmVmNTI0IiwiaWF0IjoxNzA2MjU0MTMyLCJuYmYiOjE3MDYyNTM4MzIsImV4cCI6MTcwNjI1NDczMiwic2VydmVyX3V1aWQiOiI2NjFhNTM5YS1iNWI1LTQ5NTUtYmNmNi03Mzc3NDBhNmIyNzAiLCJwZXJtaXNzaW9ucyI6WyIqIl0sInVzZXJfdXVpZCI6IjUwMzMxMzA0LWQ5YTEtNDFkNS1hM2Y5LWYwZTExNTk3ZTAxYSIsInVzZXJfaWQiOjE3MDMsInVuaXF1ZV9pZCI6IkliWkdWRGxiMkVrVnVvZmIifQ.on3wVSAh5BqJVF7ZjRHL4EQBU2CcmCQYqkUS30kIQKI",
		},
	}
	authResp := request(conn, authReq)
	//start
	// stop
	if authResp.Event == "auth success" {
		for {
			baseResp := readResp(conn)
			if baseResp.Event == "status" {
				//offline
				// running
				if len(baseResp.Args) > 0 && baseResp.Args[0] == "running" {
					_ = request(conn, Body{
						Event: "set state",
						Args:  []string{"stop"},
					})
				}
			}
			time.Sleep(time.Second * 1)
		}
	}
}

func readResp(conn *websocket.Conn) Body {
	var response Body
	err := conn.ReadJSON(&response)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(response)
	return response
}

func request(conn *websocket.Conn, body Body) Body {
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
