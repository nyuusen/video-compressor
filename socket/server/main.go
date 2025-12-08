package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("exec socket server")

	// リスナーを作成
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	defer listener.Close()
	fmt.Println("server start")

	// 接続待ち
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}
		go func(c net.Conn) {
			defer c.Close()
			// データ受信
			buf := make([]byte, 1024)
			// ユーザー空間バッファにコピー
			n, err := c.Read(buf)
			if err != nil {
				log.Fatal("read error:", err)
			}
			fmt.Printf("receive data is %s (size: %d)", string(buf), n)
		}(conn)
	}
}
