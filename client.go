package main

import (
	"io"
	"log"
	"net"
	"os"
)

func main()  {
	conn, _ := net.Dial("tcp", "localhost:8888")

	log.Println("connected")

	defer conn.Close()

	io.Copy(os.Stdout,conn)
}