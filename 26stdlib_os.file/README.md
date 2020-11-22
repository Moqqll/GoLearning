# os.file

本文主要介绍 Go 语言中文件读写相关的操作。

计算机中的文件是存储在外部截至（通常是磁盘）上的数据集合，文件分为文本文件和二进制文件。

## 打开和关闭文件

`os.Open()`函数能够打开一个文件，返回一个`*File`和一个`err`。对得到的文件实例调用`close()`方法能够关闭文件。

```go
func openfile() {
	//只读方式打开当前目录下的main.go文件
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	//关闭文件
	err = file.Close()
	if err != nil {
		fmt.Println("close file failed, err:", err)
	}
	fmt.Println("close file succ.")
}
```

为了防止文件忘记关闭，我们通常使用 defer 语句注册文件关闭语句。

## 读取文件

### file.Read()

#### 基本使用

Read 方法的定义如下：

```go
func (f *Fiel) Read(b []byte) (n int, err error)
```

它接收一个字节切片，返回读取的字节数和可能的具体错误，读到文件末尾时会返回`0`和`io.EOF`。 举个例子：

```go
func readfile() {
	//只读方式打开当前目录下的main.go文件
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	//使用read方法读取数据
	var tmp = make([]byte, 1000)
	n, err := file.Read(tmp)
	if err == io.EOF {
		fmt.Println("文件读完了。")
	}
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Printf("读取了%d字节的数据\n", n)
	fmt.Println(string(tmp[:n]))
}
```

#### 循环读取

使用 for 循环读取文件中的所有数据：

```go
func cycleread() {
	//只读方式打开当前目录下的main.go文件
	file, err := os.Open("./main.go")
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	//循环读取文件
	//tmp无法预先知道文件的大小
	var content []byte
	var tmp = make([]byte, 128) //tmp不一定够大
	for {
		n, err := file.Read(tmp)
		if err == io.EOF {
			fmt.Println("文件读完了。")
			break
		}
		if err != nil {
			fmt.Println("read file failed, err:", err)
			return
		}
		content = append(content, tmp[:n]...)
	}
	fmt.Println(string(content))
}
```

### bufio 读取文件

bufio 是在 file 的基础上封装了一层 API，支持更多功能。

```go
func bioread() {
	file, err := os.Open("./xx.txt")
	if err != nil {
		fmt.Println("open file failed,err:", err)
		return
	}
	defer file.Close()
	//bufio逐行读取示例
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			if len(line) != 0 {
				fmt.Println(line)
			}
			fmt.Println("文件读完了")
			break
		}
		if err != nil {
			fmt.Println("read file failed,err:", err)
			return
		}
		fmt.Println(line)
	}
}
```

## ioutil 读取整个文件

`io/ioutil`包的`ReadFile`方法能够读取完整的文件，只需要将文件名作为参数传入。

```go
func iotil() {
	//ioutil.ReadFile读取整个文件
	content, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Println(string(content))
}
```

## 文件写入操作

`os.OpenFile()`函数能够以指定模式打开文件，从而实现文件写入相关功能。

```go
func OpenFile(name string, flag int, perm FileMode) (*File, error){}
```

其中：

`name:`要打开的文件名；`flag:`打开文件的模式，模式有以下几种：

| 模式        | 含义     |
| ----------- | -------- |
| os.O_WRONLY | 只写     |
| os.O_CREATE | 创建文件 |
| os.O_RDONLY | 只读     |
| os.O_RDWR   | 读写     |
| os.O_TRUNC  | 清空     |
| os.O_APPEND | 追加     |

`perm`：文件权限，一个八进制数。r（读）04，w（写）02，x（执行）01。

### Write 和 WriteString

```go
func writefile() {
	file, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	str := "moqqll!\n"
	file.Write([]byte(str))
	file.WriteString(str)
}
```

### ioutil.NewWriter

```go
func iotilwrite() {
	file, err := os.OpenFile("./xx.txt", os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Println("open file failed, err:", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for i := 0; i < 10; i++ {
		writer.WriteString("moqqll !\n") //将数据线写入缓存
	}
	writer.Flush() //将缓存中的内容写入文件
}
```

### ioutil.WriteFile

```go
func iotilwritefile() {
	str := "沙河娜扎"
	err := ioutil.WriteFile("./xx.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("Write file failed, err:", err)
		return
	}
}
```

## 实现 copyFile

```go
func copyFile(dstName, srcName string) (written int64, err error) {
	defer fmt.Println("copy done!")
	//以只读方式打开源文件
	src, err := os.Open(srcName)
	if err != nil {
		fmt.Println("open src file failed,err:", err)
		return
	}
	defer src.Close()
	//以写|创建方式打开目标文件
	dst, err := os.OpenFile(dstName, os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("open dst file failed, err:", err)
		return
	}
	defer dst.Close()
	//以字节流拷贝
	return io.Copy(dst, src)
}
```

## 实现一个 cat 命令

```

```
