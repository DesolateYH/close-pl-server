package main

import (
	"encoding/json"
	"fmt"
	"github.com/DesolateYH/libary-yh-go/logger"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"time"
)

const loopTime = time.Second * 10

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
	lastMemoryUsage := float64(0)
	logger.Get().Info("begin loop send memory")
	for {
		time.Sleep(loopTime)
		resp, err := getResp(conn)
		if err != nil {
			continue
		}
		if resp.Event == eventStatus && len(resp.Args) > 0 {
			var args statusEventArgs
			err := json.Unmarshal([]byte(resp.Args[0]), &args)
			if err != nil {
				logger.Get().Error("fail to unmarshal statusEventArgs", zap.Error(err), zap.Any("resp", resp))
				time.Sleep(time.Second * 10)
				continue
			}

			memoryUsage := float64(args.MemoryBytes) / float64(args.MemoryLimitBytes) * 100
			if memoryUsage < 95 {
				logger.Get().Info("memory usage is less than 95%", zap.Float64("memoryUsage", memoryUsage))
				continue
			}
			if memoryUsage < lastMemoryUsage {
				continue
			}

			lastMemoryUsage = memoryUsage
			_, err = sendCommend(conn, Body{
				Event: eventSendCommend,
				Args: []string{
					fmt.Sprintf("Broadcast current_memory_useage_%.2f%%", memoryUsage),
				},
			})
			if err != nil {
				continue
			}
		}
	}
}
