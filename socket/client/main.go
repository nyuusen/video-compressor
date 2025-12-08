package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	fmt.Println("exec socket client")

	// ストリームソケットを作成し、サーバへの接続を確立
	addr := "localhost:8080"
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal("dial error:", err)
	}
	defer conn.Close()
	fmt.Println("connect success: " + addr)

	// データ送信
	msg := "hello,world!"
	_, err = conn.Write([]byte(msg))
	if err != nil {
		log.Fatal("write error:", err)
	}
}
