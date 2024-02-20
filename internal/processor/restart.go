package processor

import (
	"close-pl-server/internal/consts"
	"close-pl-server/internal/model"
	"github.com/DesolateYH/libary-yh-go/logger"
	"github.com/gorilla/websocket"
	"time"
)

func restartServer(conn *websocket.Conn) error {
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second * 10)
		logger.Get().Info("send broadcast server_will_restart")
		_, err := connection.SendCommend(conn, model.Body{
			Event: consts.EventSendCommend,
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
		resp, err := connection.GetResp(conn)
		if err != nil {
			return err
		}
		if resp.Event != consts.EventStats {
			continue
		}
		args, err := resp.getStatsEventArgs()
		if err != nil {
			return err
		}
		if args.State != consts.ServerStatsRunning {
			continue
		}

		_, err = connection.SendCommend(conn, model.Body{
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
