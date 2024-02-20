package connection

import (
	"close-pl-server/internal/consts"
	"close-pl-server/internal/model"
	"close-pl-server/internal/myerrors"
	"github.com/DesolateYH/libary-yh-go/logger"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

type Connection interface {
	SendCommend(body model.Body) (model.Body, error)
	GetResp() (model.Body, error)
}

type connection struct {
	socketConn *websocket.Conn
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
	conn := &connection{socketConn: socketConnection}

	authResp, err := conn.SendCommend(model.Body{
		Event: "auth",
		Args: []string{
			token,
		},
	})
	if err != nil {
		return nil, err
	}
	if authResp.Event != consts.EventAuthSuccess {
		logger.Get().Error("fail to auth", zap.Any("resp", authResp))
		return nil, myerrors.FailToAuthError
	}

	return conn, nil
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
