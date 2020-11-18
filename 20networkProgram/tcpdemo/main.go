package main

import (
	"bufio"
	"fmt"
	"net"
)

// tcp/server/main.go

// TCP server端

//处理函数
func process(conn net.Conn) {
	defer conn.Close() //处理完后要关闭连接
	//针对当前连接做数据的发送和接受操作
	for {
		//创建一个从当前tcp连接（conn）中进行读操作的对象reader
		reader := bufio.NewReader(conn)
		//声明一个128大小的字节类型数组buf
		var buf [128]byte
		//把数组buf转换为切片buf传入reader对象的Read方法中
		//Read方法会返回一个n值（读取的字节长度）和一个错误信息
		//并把读取到的字节数据存储切片buf中
		n, err := reader.Read(buf[:])
		if err != nil { //错误处理
			fmt.Printf("read from conn failed, err:%v\n", err)
			break
		}
		recv := string(buf[:n]) //把切片buf中的数据转换为字符串conn
		fmt.Printf("接收到的数据；%v\n", recv)
		conn.Write("ok") //向客户端写一个ok信息
	}
}

func main() {
	//1、启动一个tcp监听
	listen, err := net.Listen("tcp", "127.0.0.1:2000")
	if err != nil {
		fmt.Println("listen failed, err:%v\n", err)
		return
	}
	for {
		//2、等待客户端来连接
		conn, err := listen.Accept()
		if err != nil {
			fmt.Printf("accept failed, err:%v\n", err)
			continue
		}
		//3、启动一个单独的goroutine去处理连接
		go process(conn)
	}
}
