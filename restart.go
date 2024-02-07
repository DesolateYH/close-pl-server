package main

import (
	"github.com/DesolateYH/libary-yh-go/logger"
	"github.com/gorilla/websocket"
	"time"
)

func restartServer(conn *websocket.Conn) error {
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second * 10)
		logger.Get().Info("send broadcast server_will_restart")
		_, err := sendCommend(conn, Body{
			Event: eventSendCommend,
			Args: []string{
				"Broadcast server_will_restart",
			},
		})
		if err != nil {
			continue
		}
	}

	for {
		time.Sleep(time.Second * 2)
		resp, err := getResp(conn)
		if err != nil {
			return err
		}
		if resp.Event != eventStats {
			continue
		}
		args, err := resp.getStatsEventArgs()
		if err != nil {
			return err
		}
		if args.State != serverStatsRunning {
			continue
		}

		_, err = sendCommend(conn, Body{
			Event: "set state",
			Args:  []string{"restart"},
		})
		if err != nil {
			return err
		}
		logger.Get().Info("restart server success")
		return nil
	}

}
