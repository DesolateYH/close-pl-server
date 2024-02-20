package main

import (
	"encoding/json"
	"github.com/DesolateYH/libary-yh-go/logger"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
)

type getTokenResp struct {
	Data struct {
		Token  string `json:"token"`
		Socket string `json:"socket"`
	} `json:"data"`
}

func getToken(cookie string) (string, error) {
	url := "https://panel.vatzj.com/api/client/servers/661a539a-b5b5-4955-bcf6-737740a6b270/websocket"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		logger.Get().Error("fail to new request", zap.Error(err))
		return "", err
	}
	req.Header.Add("authority", "panel.vatzj.com")
	req.Header.Add("accept", "application/json")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("cookie", cookie)
	req.Header.Add("pragma", "no-cache")
	req.Header.Add("referer", "https://panel.vatzj.com/server/661a539a")
	req.Header.Add("sec-ch-ua", "\"Not A(Brand\";v=\"99\", \"Brave\";v=\"121\", \"Chromium\";v=\"121\"")
	req.Header.Add("sec-ch-ua-mobile", "?0")
	req.Header.Add("sec-ch-ua-platform", "\"Windows\"")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-site", "same-origin")
	req.Header.Add("sec-gpc", "1")
	req.Header.Add("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/121.0.0.0 Safari/537.36")
	req.Header.Add("x-requested-with", "XMLHttpRequest")
	req.Header.Add("x-xsrf-token", "eyJpdiI6Ii9VNU94ZzNwWXFTUGx3Mlh2WWdRK2c9PSIsInZhbHVlIjoibXZGZEJYNnI2c3I3WXpBU2R1eVBTaEd0cTFNN0pyVi9qQWkzZkdyNWRqcGZhMDJCMkJvcisyRWphK0F3MWNEVHg3c1pEUkVGSzNrM0Y2VWo3dlZab21aVVg2Z2NUMTNMVlVYSVUwYURnYjVDQ3dQczI1UVFrdlgyT1F0RWVvWFQiLCJtYWMiOiIwNTdjMzUyNjA1NzE2YTE2Zjc2YmIzYTE2ZTdjYjg1MmUxZjk1MDhiOGEzMjQ0MDNlNmU0Y2E5NDczOTA3YjY0IiwidGFnIjoiIn0=")

	res, err := client.Do(req)
	if err != nil {
		logger.Get().Error("fail to do request", zap.Error(err))
		return "", err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.Get().Error("fail to read body", zap.Error(err))
		return "", err
	}

	var t getTokenResp
	err = json.Unmarshal(body, &t)
	if err != nil {
		logger.Get().Error("fail to unmarshal", zap.Error(err), zap.String("body", string(body)))
		return "", err
	}

	return t.Data.Token, nil
}
