package connection

import (
	"close-pl-server/internal/consts"
	"close-pl-server/internal/model"
	"close-pl-server/internal/myerrors"
	"github.com/DesolateYH/libary-yh-go/logger"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"time"
)

type Connection interface {
	SendCommend(body model.Body) (model.Body, error)
	GetResp() (model.Body, error)

	RestartServer() error

	Refresh() error
}

type connection struct {
	socketConn *websocket.Conn
	cookie     string
}

func NewConnection(cookie string) (Connection, error) {
	token, err := getToken(cookie)
	if err != nil {
		return nil, err
	}
	socketConnection, err := getSocketConnection(token)
	if err != nil {
		return nil, err
	}

	return &connection{
		socketConn: socketConnection,
		cookie:     cookie,
	}, nil
}

func (c *connection) RestartServer() error {
	for i := 0; i < 3; i++ {
		time.Sleep(time.Second * 10)
		logger.Get().Info("send broadcast server_will_restart")
		_, err := c.SendCommend(model.Body{
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
		resp, err := c.GetResp()
		if err != nil {
			return err
		}
		if resp.Event != consts.EventStats {
			continue
		}
		args, err := resp.GetStatsEventArgs()
		if err != nil {
			return err
		}
		if args.State != consts.ServerStatsRunning {
			continue
		}

		_, err = c.SendCommend(model.Body{
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

func (c *connection) Refresh() error {
	token, err := getToken(c.cookie)
	if err != nil {
		return err
	}
	socketConnection, err := getSocketConnection(token)
	if err != nil {
		return err
	}
	c.socketConn = socketConnection
	return nil
}

func (c *connection) SendCommend(body model.Body) (model.Body, error) {
	err := c.socketConn.WriteJSON(body)
	if err != nil {
		logger.Get().Error("fail to write json", zap.Error(err), zap.Any("body", body))
		return model.Body{}, err
	}
	return c.GetResp()
}

func (c *connection) GetResp() (model.Body, error) {
	var response model.Body
	err := c.socketConn.ReadJSON(&response)
	if err != nil {
		logger.Get().Error("fail to read json", zap.Error(err))
		return model.Body{}, err
	}
	if response.Event == consts.EventDaemonError || response.Event == consts.EventJwtError {
		logger.Get().Warn("daemon error", zap.Any("args", response.Args))
		return model.Body{}, myerrors.TokenExpireError
	}

	logger.Get().Info("get resp success", zap.Any("resp", response))
	return response, nil
}
