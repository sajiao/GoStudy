package main

import (
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

var count = 0      // 俩大爷已经遇见了多少次
var total = 100000 // 总共需要遇见多少次

var z0 = " 吃了没，您吶?"
var z3 = " 嗨！吃饱了溜溜弯儿。"
var z5 = " 回头去给老太太请安！"
var l1 = " 刚吃。"
var l2 = " 您这，嘛去？"
var l4 = " 有空家里坐坐啊。"

var liWriteLock sync.Mutex
var zhangWriteLock sync.Mutex

type RequestResponse struct {
	Serial  uint32 //序号
	Payload string //内容
}

// 序列化 RequestResponse，并发送
// 序列化后的结构如下：
// 	长度	4 字节
// 	Serial 4 字节
// 	PayLoad 变长
func writeTo(r *RequestResponse, conn *net.TCPConn, lock *sync.Mutex) {
	lock.Lock()
	defer lock.Unlock()
	payloadBytes := []byte(r.Payload)
	serialBytes := make([]byte, 4)
	binary.BigEndian.PutUint32(serialBytes, r.Serial)
	length := uint32(len(payloadBytes) + len(serialBytes))
	lengthByte := make([]byte, 4)
	binary.BigEndian.PutUint32(lengthByte, length)

	conn.Write(lengthByte)
	conn.Write(serialBytes)
	conn.Write(payloadBytes)
	//fmt.Println("发送:" + r.Payload)
}

func readFrom(conn *net.TCPConn) (*RequestResponse, error) {
	ret := &RequestResponse{}
	buf := make([]byte, 4)
	if _, err := io.ReadFull(conn, buf); err != nil {
		return nil, fmt.Errorf("读长度故障：s%", err.Error())
	}
	length := binary.BigEndian.Uint32(buf)
	if _, err := io.ReadFull(conn, buf); err != nil {
		return nil, fmt.Errorf("读Serial故障:s%", err.Error())
	}
	ret.Serial = binary.BigEndian.Uint32(buf)
	payloadBytes := make([]byte, length-4)
	if _, err := io.ReadFull(conn, payloadBytes); err != nil {
		return nil, fmt.Errorf("读payload故障:s%", err.Error())
	}
	ret.Payload = string(payloadBytes)
	return ret, nil
}

func zhangDaYeListen(conn *net.TCPConn) {
	for count < total {
		r, err := readFrom(conn)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		if r.Payload == l2 {
			go writeTo(&RequestResponse{r.Serial, z3}, conn, &zhangWriteLock)
		} else if r.Payload == l4 { // 如果收到：有空家里坐坐啊。
			go writeTo(&RequestResponse{r.Serial, z5}, conn, &zhangWriteLock) // 回复：回头去给老太太请安！
		} else if r.Payload == l1 { // 如果收到：刚吃。
			// 不用回复
		} else {
			fmt.Println(" 张大爷听不懂：" + r.Payload)
			break
		}
	}
}

func zhangDayeSay(conn *net.TCPConn) {
	nextSerial := uint32(0)
	for i := 0; i < total; i++ {
		writeTo(&RequestResponse{nextSerial, z0}, conn, &zhangWriteLock)
		nextSerial++
	}
}

// 李大爷的耳朵，实现是和张大爷类似的
func liDaYeListen(conn *net.TCPConn, wg *sync.WaitGroup) {
	defer wg.Done()
	for count < total {
		r, err := readFrom(conn)
		if err != nil {
			fmt.Println(err.Error())
			break
		}
		// fmt.Println(" 李大爷收到：" + r.Payload)
		if r.Payload == z0 { // 如果收到：吃了没，您吶?
			writeTo(&RequestResponse{r.Serial, l1}, conn, &liWriteLock) // 回复：刚吃。
		} else if r.Payload == z3 {
			// do nothing
		} else if r.Payload == z5 {
			//fmt.Println(" 俩人说完走了 ")
			count++
		} else {
			fmt.Println(" 李大爷听不懂：" + r.Payload)
			break
		}
	}
}

// 李大爷的嘴
func liDaYeSay(conn *net.TCPConn) {
	nextSerial := uint32(0)
	for i := 0; i < total; i++ {
		writeTo(&RequestResponse{nextSerial, l2}, conn, &liWriteLock)
		nextSerial++
		writeTo(&RequestResponse{nextSerial, l4}, conn, &liWriteLock)
		nextSerial++
	}
}

func startServer() {
	tcpAddr, _ := net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	tcpListener, _ := net.ListenTCP("tcp", tcpAddr)
	defer tcpListener.Close()
	fmt.Println("张大爷在胡同门口等着...")
	for {
		conn, err := tcpListener.AcceptTCP()
		if err != nil {
			fmt.Println(err)
			break
		}
		fmt.Println("碰见一个黎大爷：" + conn.RemoteAddr().String())
		go zhangDaYeListen(conn)
		go zhangDayeSay(conn)
	}
}

func startClient() {
	var tcpAddr *net.TCPAddr
	tcpAddr, _ = net.ResolveTCPAddr("tcp", "127.0.0.1:9999")
	conn, _ := net.DialTCP("tcp", nil, tcpAddr)

	defer conn.Close()
	var wg sync.WaitGroup
	wg.Add(1)
	go liDaYeListen(conn, &wg)
	go liDaYeSay(conn)
	wg.Wait()

}

func main() {
	go startServer()
	time.Sleep(time.Second)
	t1 := time.Now()
	startClient()
	elapsed := time.Since(t1)
	fmt.Println(" 耗时: ", elapsed)
}
