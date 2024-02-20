package processor

import "github.com/gorilla/websocket"

type Processor interface {
	MonitorMemory()
}

type processor struct {
	cookie string
	conn   *websocket.Conn
}

func NewProcessor(cookie string, conn *websocket.Conn) Processor {
	return &processor{cookie: cookie, conn: conn}
}
