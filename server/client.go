package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main() {
	//连接服务器
	conn, _ := net.Dial("tcp", "localhost:8888")
	log.Println("connected")
	//
	io.Copy(os.Stdout, conn)
	//最后执行关闭连接
	defer conn.Close()
}
