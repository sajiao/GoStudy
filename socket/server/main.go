package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func HandleCunc(c net.Conn) {
	defer c.Close()
	addr := c.RemoteAddr().String()
	fmt.Println(addr, " connect success")
	buf := make([]byte, 1024)
	for {
		n, err := c.Read(buf)
		if err != nil {
			fmt.Println(err)
			return
		}
		result := buf[:n]
		fmt.Printf("来自<%s>的数据:%s\n", addr, string(result))
		if "exit " == string(result) {
			fmt.Println(addr, "退出连接")
			return
		}

		c.Write([]byte(fmt.Sprintf("%s 服务端返回数据：%s", time.Now(), string(result))))
	}
}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:9000")
	if err != nil {
		log.Println(err)
		return
	}
	defer listener.Close()
	for {
		connect, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}
		go HandleCunc(connect)
	}
}
