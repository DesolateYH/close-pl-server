package main

func main() {
	token, err := getToken()
	if err != nil {
		panic(err)
	}
	conn, err := getConnection(token)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	loopSendMemory(conn)

	//c := cron.New()
	//wg := sync.WaitGroup{}
	//wg.Add(1)
	//_, err := c.AddFunc("22 18 * * *", func() {
	//	logger.Get().Error("begin stop server job")
	//	processor()
	//	logger.Get().Error("success stop server job")
	//})
	//if err != nil {
	//	log.Fatalln(err)
	//}
	//c.Start()
	//
	//logger.Get().Error("wait for stop server")
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
