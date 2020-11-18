# Network Programming

现在我们每天都在使用互联网，我们前面已经学习了如何编写 Go 语言程序，但是和才能让我们的程序通过网络互相通信呢？

本章我们就一起来学习下 Go 语言中的网络变成。关于网络编程的内容是一个很庞大的领域，本文只是简单的延时了如何使用 net 包进行 TCP 和 UDP 通信。如需了解更详细的网络编程请自行检索和阅读专业资料。

## 互联网协议的介绍

互联网的核心是一系列的协议，总称为“互联网协议”（Internet protocol suite），正是这样一些协议规定了计算机组合连接与组网。我们理解了这些协议，也就理解了互联网（Internet）的原理。由于这些协议太过庞大和复杂，没有办法在这里一概而全，只能介绍一下我们日常开发中接触较多的几个协议。

### 互联网分层模型

互联网的逻辑实现被分为好几层。每一层都有自己的功能，就像建筑物一样，每一层都靠下一层支持。用户接触到的只是最上面的那一层，根本不会感觉到下面的几层。要理解互联网就需要自下而上理解每一层的实现的功能。

![osi七层模型](https://www.liwenzhou.com/images/Go/socket/osi.png)

## Socket 编程

Socket 是 BSD UNIX 的进程通信机制，通常也称作“套接字”，用户描述 IP 地址和端口，是一个通信链的句柄。Socket 可以理解为 TCP/IP 网络的 API，它定义了许多函数或例程，程序员可以用它们来开发 TCP/IP 网络上的应用程序。电脑上运行的应用程序通常都是通过“套接字”向网络发出请求或者应答网络请求。

###　 Socket 图解

`Socket`是应用层与 TCP/IP 协议族通信的中间软件抽象层。在设计模式中，`Socket`其实就是一个门面模式，它把复杂的 TCP/IP 协议族隐藏在`Scoket`后面，对用户来说只需要调用`Scoket`规定的相关函数，让`Scoket`去组织符合指定的协议的数据，然后进行通信。

![socket图解](https://www.liwenzhou.com/images/Go/socket/socket.png)

## Go 语言实现 TCP 通信

### TCP 协议

TCP/IP(transmission control protocol/Internet protocol)即传输控制协议/网间协议，是一种面向连接（连接导向）的、可靠的、基于字节（byte=8bit）流的传输层（transport layer）通信协议，因为它是面向连接的协议，所以它的数据会像水流一样传输，因此会存在黏包问题。

### TCP 服务端

一个 TCP 服务端可以同时连接很多个客户端，例如世界各地的用户使用自己电脑上的浏览器访问淘宝网。

Go 语言中创建多个`goroutine`实现并发非常方便和高效，所以我们每建立一次链接就创建一个`goroutine`去处理。

TCP 服务端程序的处理流程是：

1、监听端口

2、接收客户端请求并建立链接

3、创建 goroutine 处理连接

我们使用 Go 语言的 net 包实现 TCP 服务端的代码如下：

```go
// tcp/server/main.go

// TCP server端

//处理函数
func process(conn net.Conn){

}
```
