# net/http

Go 语言内置的`net/http`包十分优秀，提供了 HTTP 客户端和服务端的实现。

## HTTP 协议

超文本传输协议（hypertext transfer protocol）是互联网上应用最为广泛的一种网络传输协议，所有的 www 文件都必须遵守这个标准。设计 HTTP 最初的目的是为了提供一种发布和接收 HTML 页面的方法。

## HTTP client

### 基本的 HTTP/HTTPS 请求

Get，Head，Post，PostForm 函数发出 HTTP/HTTPS 请求。

```go
resp,err:=http.Get("http://example.com")
...
resp,err:=http.Post("http://example.com/upload","image/jpeg",&buf)
...
resp,err:=http.PostForm("http://example.com/form",url.Values{"key":{"value"},"id:":{"1223"}})
```

程序在使用完 response 后必须关闭回复的主体：

```
resp, err:=http.Get("http://example.com/")
if err!=nil{
//handle error
}
defer resp.Body.Close()
body, err:=ioutil.ReadAll(resp.Body)
//...
```

### GET 请求示例

使用`net/http`包编写一个简单的发送HTTP请求的Client端，代码如下：

```go
resp, err := http.Get("https://www.liwenzhou.com/")
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))
```

将上面的代码保存之后编译成可执行文件，执行之后就能在终端打印`liwenzhou.com`网站首页的内容了，我们的浏览器其实就是一个发送和接收HTTP协议的数据的客户端，我们平时通过浏览器访问网页其实就是从网站的服务器接收HTTP数据，然后浏览器会按照HTML、CSS等规则将网页渲染展示出来。

### 带参数的GET请求示例

使用`net/http`标准库编写一个简单的发送 HTTP 请求的 Client，代码如下：

```go
func myGet() {
	apiURL := "http://127.0.0.1:9000/get"
	//URL param
	data := url.Values{}
	data.Set("name", "moqqll")
	data.Set("age", "18")
	u, err := url.ParseRequestURI(apiURL)
	if err != nil {
		fmt.Println("parse url requestURL failed, err:", err)
		return
	}
	//URL 编码
	u.RawQuery = data.Encode()
	fmt.Println(u.String())
	resp, err := http.Get(u.String())
	if err != nil {
		fmt.Printf("post failed, err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}
```

对应的Server端HandlerFunc如下：

```go
func getHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	data := r.URL.Query()
	fmt.Println(data.Get("name"))
	fmt.Println(data.Get("age"))
	answer := `{"status":"ok"}`
	w.Write([]byte(answer))
}
```

### Post请求示例

上面演示了使用`net/http`包发送`GET`请求的示例，发送`POST`请求的示例代码如下：

```go
func myPost() {
	url := "http://127.0.0.1:9000/post"
	//表单数据
	contentType := "application/x-www-form-urlencoded"
	data := "name=小王子&age=18"
	//json
	// contentType := "application/json"
	// data := `{"name":"小王子","age":18}`
	resp, err := http.Post(url, contentType, strings.NewReader(data))
	if err != nil {
		fmt.Printf("post failed ,err:%v\n", err)
		return
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("get resp body failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))
}
```

对应的Server端HandlerFunc如下：

```go
func postHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	//请求类型是application/x-www-form-urlencoded时解析form数据
	r.ParseForm()
	fmt.Println(r.PostForm) //打印form数据
	fmt.Println(r.PostForm.Get("name"), r.PostForm.Get("age"))

	//请求类型是application/json时从r.Body读取数据
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("read request.Body failed, err:%v\n", err)
		return
	}
	fmt.Println(string(b))
	answer := `{"status":"ok"}`
	w.Write([]byte(answer))
}
```

### 自定义Client

要管理HTTP客户端的头域、重定向策略和其他设置，创建一个Client：

```GO
client := &http.Client{
	CheckRedirect: redirectPolicyFunc,
}
resp, err := client.Get("http://example.com")
// ...
req, err := http.NewRequest("GET", "http://example.com", nil)
// ...
req.Header.Add("If-None-Match", `W/"wyzzy"`)
resp, err := client.Do(req)
// ...
```

### 自定义Transport

要管理代理、TLS配置、keep-alive、压缩和其他设置，创建一个Transport：

```go
tr := &http.Transport{
	TLSClientConfig: &tls.Config{RootCAs:pool},
	DisableCompression: true,
}
client := &http.Client{Transport:tr}
resp, err:=client.Get("https://example.com")
```

Client和Transport类型都可以安全的被多个goroutine同时使用。出于效率考虑，应该一次建立、尽量重用。

## HTTP Server

### 默认的Server

ListenAndServer使用指定的监听地址和处理器启动一个HTTP server，处理器参数通常是nil，这表示采用包变量DefaultServeMux作为处理器。

Handle和HandleFunc函数可以向DefaultServeMux添加处理器。

```go
http.Handle("/foo",fooHandler)
http.HnadleFunc("/bar",func(w http.ResponseWriter, r *http.Request){fmt.Fprintf(w, "Hello, %q",html.EscapeString(r.URL.Path))})
log.Fatal(http.ListenAndServe(":8000",nil))
```

### 默认的Server示例

使用Go语言中的`net/http`包来编写一个简单的接收HTTP请求的Server端示例，`net/http`包是对net包的进一步封装，专门用来处理HTTP协议的数据。具体的代码如下：

```go
func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello,沙河小王子！")
}

func main() {
	http.HandleFunc("/", sayHello)
	err := http.ListenAndServe("127.0.0.1:9000", nil)
	if err != nil {
		fmt.Printf("http server start failed, err:%v\n", err)
	}
}
```

### 自定义Server

```go
s := &http.Server{
	Addr:			":8000",
	Handler:		myHandler,
	ReadTimeout:	10*time.Second,
    WriteTimeout:	10*time.Second,
    MaxHeaderBytes:	1 << 20,
}
log.Fatal(s.ListenAndServe())
```

