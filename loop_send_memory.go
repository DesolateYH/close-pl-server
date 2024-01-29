package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"log"
	"time"
)

type statusEventArgs struct {
	CpuAbsolute      float64 `json:"cpu_absolute"`
	DiskBytes        int64   `json:"disk_bytes"`
	MemoryBytes      int64   `json:"memory_bytes"`
	MemoryLimitBytes int64   `json:"memory_limit_bytes"`
	Network          struct {
		RxBytes int   `json:"rx_bytes"`
		TxBytes int64 `json:"tx_bytes"`
	} `json:"network"`
	State  string `json:"state"`
	Uptime int    `json:"uptime"`
}

// {"event":"stats","args":["{\"cpu_absolute\":305.601,\"disk_bytes\":2675046943,\"memory_bytes\":20518100992,\"memory_limit_bytes\":25804800000,\"network\":{\"rx_bytes\":704862335,\"tx_bytes\":3384315373},\"state\":\"running\",\"uptime\":24150801}"]}
// {"event":"console output","args":["\u003e Broadcast 3"]}
func loopSendMemory(conn *websocket.Conn) {
	for {
		resp := getResp(conn)
		if resp.Event == eventStatus && len(resp.Args) > 0 {
			var args statusEventArgs
			err := json.Unmarshal([]byte(resp.Args[0]), &args)
			if err != nil {
				log.Println("fail to unmarshal statusEventArgs", zap.Error(err), zap.Any("resp", resp))
				time.Sleep(time.Second * 10)
				continue
			}

			sendCommend(conn, Body{
				Event: eventSendCommend,
				Args: []string{
					fmt.Sprintf("Broadcast {当前内存使用率: %.2f%%}", float64(args.MemoryBytes)/float64(args.MemoryLimitBytes)*100),
				},
			})
			time.Sleep(time.Second * 10)
		}
	}
}
