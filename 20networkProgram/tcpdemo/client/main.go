package main

import (
	"fmt"
	"net"
)

// tcp/client/main.go

// TCP client端
func main() {
	//与服务端建立连接
	conn, err := net.Dial("tcp", "127.0.0.1:2000")
	if err != nil {
		fmt.Printf("dial failed, err%v\n", err)
		return
	}
	//利用该连接进行数据的发送和接受

}
