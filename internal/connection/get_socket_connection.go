package connection

import (
	"close-pl-server/internal/consts"
	"close-pl-server/internal/model"
	"close-pl-server/internal/myerrors"
	"github.com/DesolateYH/libary-yh-go/logger"
	"github.com/gorilla/websocket"
	"go.uber.org/zap"
	"net/http"
)

func getSocketConnection(token string) (*websocket.Conn, error) {
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

	socketConn, _, err := dialer.Dial("wss://pt-50.vatzj.com:8082/api/servers/661a539a-b5b5-4955-bcf6-737740a6b270/ws", header)
	if err != nil {
		logger.Get().Error("fail to dial", zap.Error(err))
		return nil, err
	}

	conn := &connection{socketConn: socketConn}
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

	return socketConn, nil
}
