package main

import (
	"go.uber.org/zap"
	"io"
	"log"
	"net/http"
)

type getTokenResp struct {
	Data struct {
		Token  string `json:"token"`
		Socket string `json:"socket"`
	} `json:"data"`
}

func getToken() (string, error) {
	req, err := http.NewRequest("GET", "https://panel.vatzj.com/api/client/servers/661a539a-b5b5-4955-bcf6-737740a6b270/websocket", nil)
	if err != nil {
		log.Println("fail to new request in get token", zap.Error(err))
		return "", err
	}
	req.Header.Set("Authority", "panel.vatzj.com")
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Language", "zh-CN,zh;q=0.9")
	req.Header.Set("Cache-Control", "no-cache")
	req.Header.Set("Cookie", "remember_web_59ba36addc2b2f9401580f014c7f58ea4e30989d=eyJpdiI6Ill6K0hhSFpUSUptNmtmdU9PUlhSVUE9PSIsInZhbHVlIjoiWUZJc0EvS0xjcjJaRDJBVjdxaytDNnRwZVZJRVczeXVwNUlHc3daSmVySGg3VWdMSHZQdGJYUGtHMWJDcGsyL254Vm80VjdmcGwzajhpWWcwT2pka0NVNHMwY3hrWmlUVEVKRmUwaEZyYmNvcTJMcy9mbjhhZ2N6dExJL2dDczVDNTN2UDdhTFVXTzhxREZla1BteVJXNTBzTDZIK0Zxd2loUFVyNTNMOVdZQ25PSjBXWHFNK2NZZGthaWlZVngyM1FteEU2OFFIbmlFRmpiaWtJZDlJMGE4bUN3Qm9vRDZGeWl2Skt2V0tiYz0iLCJtYWMiOiI3MWFkYTliZDRjY2JiMGYxYzQzNWYxNDQxZTQwOTE5MDNjYzU3YjJmZTkwYWI4NTRiOTBmYzQwZjExOTk1MmQ5IiwidGFnIjoiIn0^%^3D; XSRF-TOKEN=eyJpdiI6IjNlMWFXanVqUDU2dUY2aWhTd0JDSXc9PSIsInZhbHVlIjoiQ1cxcC9tQWFDN2Z0QWFRcG85QVQ3OXhCRERoOXpMZ0krYlgzUDFpTHBzWDF1YUh4dit3RmdJcHZMMGttYlREWVdBWXVnK0R6SFlHekFlTm5IRFF0RjMreXdyQkwwc0doQkdadFJNMll2ajM3aWZNd3JMdGRLL05kcmZvR0hhV3kiLCJtYWMiOiIxZmI0NTQ3Yzc0MjdlNTlkOWQ3OTdlMjYyNDE0ZjgwNmRjNTExZDAwMThmZWVhYTRiZWRiMzkwOGJkNDMwMjBjIiwidGFnIjoiIn0^%^3D; pterodactyl_session=eyJpdiI6IkZBWHloOEFCOUpEcXZSZy84a2g0VkE9PSIsInZhbHVlIjoiUlN5c1hWeE1DWTQ0aVAyTlB4MjNybTRlU2w1UU4zYUZRc0x3NDJlRW1OZUIzQldjcGJpMk1qTE1zVW9oYUR2OFZTSW05S0llWGs0bXlXSGRXeDdkeDEzcnkwTm4rc1NHSVpld3FGcVNGRkd2YzJoUFB0S3pYOVplbnZuUGxkN3YiLCJtYWMiOiIyNTgxMjY0MmNlMTVkMGMyMThiZmY0YTQyZjBjZmMwNjM4ZmU2NThkMzZiYTFiZDBhMDdmOTdhNTQzZGVjZjUwIiwidGFnIjoiIn0^%^3D")
	req.Header.Set("Pragma", "no-cache")
	req.Header.Set("Referer", "https://panel.vatzj.com/server/661a539a")
	req.Header.Set("Sec-Ch-Ua", "^^Not")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println("fail to do request in get token", zap.Error(err))
		return "", err
	}
	defer resp.Body.Close()

	all, err := io.ReadAll(resp.Body)
	if err != nil {

		return "", err
	}
}
