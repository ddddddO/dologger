package main

import (
	"fmt"
	"net"
)

const (
	protocol = "tcp"
	port     = 8080
)

func main() {
	fmt.Println("start Aggregation log server")

	ln, err := net.Listen(protocol, fmt.Sprintf(":%d", port))
	if err != nil {
		fmt.Println("cannot listen", err)
	}

	// 接続を待ち受け続ける
	for {
		// 1接続分
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("cannot accept", err)
		}
		fmt.Println("received")

		// 1接続中のdologgerから送られてきたログを1件ずつ処理
		for {
			buf := make([]byte, 2048)
			if _, err := conn.Read(buf); err != nil {
				fmt.Println("cannot read", err)
			}

			fmt.Println("received log")
			fmt.Println(string(buf))
		}
	}
}
