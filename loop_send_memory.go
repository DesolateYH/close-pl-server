package main

import (
	"errors"
	"fmt"
	"github.com/DesolateYH/libary-yh-go/logger"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"time"
)

const loopTime = time.Second * 10

// {"event":"stats","args":["{\"cpu_absolute\":305.601,\"disk_bytes\":2675046943,\"memory_bytes\":20518100992,\"memory_limit_bytes\":25804800000,\"network\":{\"rx_bytes\":704862335,\"tx_bytes\":3384315373},\"state\":\"running\",\"uptime\":24150801}"]}
// {"event":"console output","args":["\u003e Broadcast 3"]}
func loopSendMemory(conn *websocket.Conn) {
	lastMemoryUsage := float64(0)
	logger.Get().Info("begin loop send memory")
	for {
		time.Sleep(loopTime)
		resp, err := getResp(conn)
		if err != nil {
			if errors.Is(err, TokenExpireError) {
				token, err := getToken()
				if err != nil {
					return
				}
				_conn, err := getConnection(token)
				if err != nil {
					return
				}
				conn = _conn
			}
			return
		}
		if resp.Event == eventStats && len(resp.Args) > 0 {
			args, err := resp.getStatsEventArgs()
			if err != nil {
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

			logger.Get().Info("begin restart server")
			err = restartServer(conn)
			if err != nil {
				continue
			}
		}
	}
}
