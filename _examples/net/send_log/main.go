package main

import (
	"fmt"
	"net"
	"time"

	// "github.com/ddddddO/dologger"
)

const (
	interval = 2 * time.Second

	protocol   = "tcp"
	targetHost = "localhost"
	targetPort = 8080
)

func main() {
	fmt.Println("start sample app")

	conn, err := net.Dial(
		protocol,
		fmt.Sprintf("%s:%d", targetHost, targetPort),
	)
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	// log := dologger.New(conn)
	// log.WithJSON()

	ticker := time.NewTicker(interval)
	for t := range ticker.C {
		fmt.Println("send log")

		// receive_logプロセスへTCPでログを送る
		// log.Debug("for debug").
		// 	Str("name", "ddddd").
		// 	Int("second", t.Second()).
		// 	Out()
	}
}
