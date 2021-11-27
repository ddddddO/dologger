package main

import (
	"fmt"
	"io"
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
		fmt.Println("connected")

		// 複数の接続を扱うためgoroutine
		go func() {
			defer conn.Close()

			// 1接続中のdologgerから送られてきたログを1件ずつ処理
			for {
				buf := make([]byte, 2048)
				_, err := conn.Read(buf)
				if err != nil {
					if err == io.EOF {
						fmt.Println("connection closed...")
						return
					}

					fmt.Println("cannot read", err)
				}

				fmt.Println("received log")
				fmt.Println(string(buf))
			}
		}()
	}
}
