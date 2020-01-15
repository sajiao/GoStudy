package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hello world")
}

const metersToYards float64 = 1.09361

func main() {

	//go run(4)
	//http.HandleFunc("/", sayHello) //注册URI路径与相应的处理函数
	//log.Println("【默认项目】服务启动成功 监听端口 80")
	//er := http.ListenAndServe("0.0.0.0:80", nil)
	//if er != nil {
	//	log.Fatal("ListenAndServe: ", er)
	//}
	var meters float64
	fmt.Print("Enter meters swam: ")
	fmt.Scan(&meters)
	yards := meters * metersToYards
	fmt.Println(meters, " meters is ", yards, " yards.")
}

func logRecover() {
	if err := recover(); err != nil {
		tmsg := fmt.Sprintf("err msg:%s  stack:%s", err, debug.Stack())
		log.Println(tmsg)
	}
}
func run(num int) {
	defer logRecover()
	if num%4 == 0 {
		panic("请求出错")
	}
	go myPrint(num)
}

func myPrint(num int) {
	defer logRecover()
	if num%4 == 0 {
		panic("请求又出错了")
	}
	fmt.Printf("%d\n", num)
}
