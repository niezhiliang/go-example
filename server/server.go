package main

import (
	"io"
	"log"
	"net"
	"time"
)

func main() {
	//监听
	listener, _ := net.Listen("tcp", "localhost:8888")
	for {
		log.Println("server启动成功，等待连接中...")
		//阻塞等待用户连接
		conn, _ := listener.Accept()
		log.Println("新连接加入：", conn)
		/* go */ handle(conn)
	}
}

func handle(conn net.Conn) {
	defer conn.Close()

	for {
		io.WriteString(conn, "hello \t")
		time.Sleep(1 * time.Second)
	}
}
