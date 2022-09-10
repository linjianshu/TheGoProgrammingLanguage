package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listen.Accept()
		if err != nil {
			//例如 连接停止
			log.Print(err)
			continue
		}

		//一次处理一个连接 没处理完会导致客户端阻塞  //加上go就不会了
		go handleConn(conn)
	}

}

//对连接进行处理
func handleConn(conn net.Conn) {
	defer conn.Close()
	for {
		_, err := io.WriteString(conn, time.Now().Format("2006-01-02 15:04:05\n"))
		if err != nil {
			//例如 连接断开
			return
		}
		time.Sleep(1 * time.Second)
	}
}
