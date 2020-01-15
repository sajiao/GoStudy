package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	addr := "127.0.0.1:9000"
	conn, err := net.Dial("tcp", addr)
	if err != nil {
		log.Fatal(err.Error())
		return
	}
	defer conn.Close()
	buf := make([]byte, 2048)
	readBuf := make([]byte, 2048)
	for {

		fmt.Println("please input:")
		fmt.Scan(&buf)
		fmt.Println(fmt.Sprintf("%s 客户端数据：%s", time.Now(), string(buf)))
		conn.Write(buf)
		n, err := conn.Read(readBuf)
		if err != nil {
			fmt.Println(err)
			return
		}
		result := readBuf[:n]
		fmt.Println(string(result))
	}
}
