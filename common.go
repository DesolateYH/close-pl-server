package main

import (
	"github.com/DesolateYH/libary-yh-go/logger"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
)

func sendCommend(conn *websocket.Conn, body Body) (Body, error) {
	err := conn.WriteJSON(body)
	if err != nil {
		logger.Get().Error("fail to write json", zap.Error(err), zap.Any("body", body))
		return Body{}, err
	}
	return getResp(conn)
}

func getResp(conn *websocket.Conn) (Body, error) {
	var response Body
	err := conn.ReadJSON(&response)
	if err != nil {
		logger.Get().Error("fail to read json", zap.Error(err))
		return Body{}, err
	}
	if response.Event == eventDaemonError || response.Event == eventJwtError {
		logger.Get().Warn("daemon error", zap.Any("args", response.Args))
		return Body{}, TokenExpireError
	}

	logger.Get().Info("get resp success", zap.Any("resp", response))
	return response, nil
}
