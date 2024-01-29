package main

type Body struct {
	Event string   `json:"event,omitempty"`
	Args  []string `json:"args,omitempty"`
}

const token = "eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiIsImp0aSI6IjZhNjg1ZGEyY2RiYTE1OWU3Y2Q0ZmVkM2ZmYmVmNTI0In0.eyJpc3MiOiJodHRwczovL3BhbmVsLnZhdHpqLmNvbSIsImF1ZCI6WyJodHRwczovL3B0LTUwLnZhdHpqLmNvbTo4MDgyIl0sImp0aSI6IjZhNjg1ZGEyY2RiYTE1OWU3Y2Q0ZmVkM2ZmYmVmNTI0IiwiaWF0IjoxNzA2NTI3MzQ0LCJuYmYiOjE3MDY1MjcwNDQsImV4cCI6MTcwNjUyNzk0NCwic2VydmVyX3V1aWQiOiI2NjFhNTM5YS1iNWI1LTQ5NTUtYmNmNi03Mzc3NDBhNmIyNzAiLCJwZXJtaXNzaW9ucyI6WyIqIl0sInVzZXJfdXVpZCI6IjUwMzMxMzA0LWQ5YTEtNDFkNS1hM2Y5LWYwZTExNTk3ZTAxYSIsInVzZXJfaWQiOjE3MDMsInVuaXF1ZV9pZCI6IkRMZVI2U3VydUU5ZENTcjQifQ.RXW7oCyjWV34G-wFTvPGqq1DBKt2vgya5zchayK3N5U"

func main() {
	conn := getConnection(token)
	defer conn.Close()

	loopSendMemory(conn)

	//c := cron.New()
	//wg := sync.WaitGroup{}
	//wg.Add(1)
	//_, err := c.AddFunc("22 18 * * *", func() {
	//	log.Println("begin stop server job")
	//	processor()
	//	log.Println("success stop server job")
	//})
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//c.Start()
	//
	//log.Println("wait for stop server")
	//wg.Wait()
}

//func processor() {
//	//start
//	// stop
//	if authResp.Event == "auth success" {
//		for {
//			baseResp := readResp(conn)
//			if baseResp.Event == "status" {
//				//offline
//				// running
//				if len(baseResp.Args) > 0 && baseResp.Args[0] == "running" {
//					_ = sendCommend(conn, Body{
//						Event: "set state",
//						Args:  []string{"stop"},
//					})
//				}
//			}
//			time.Sleep(time.Second * 1)
//		}
//	}
//}
