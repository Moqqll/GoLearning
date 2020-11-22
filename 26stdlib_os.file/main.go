package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

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

func iotil() {
	//ioutil.ReadFile读取整个文件
	content, err := ioutil.ReadFile("./main.go")
	if err != nil {
		fmt.Println("read file failed, err:", err)
		return
	}
	fmt.Println(string(content))
}

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

func iotilwritefile() {
	str := "沙河娜扎"
	err := ioutil.WriteFile("./xx.txt", []byte(str), 0666)
	if err != nil {
		fmt.Println("Write file failed, err:", err)
		return
	}
}

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

//使用文件操作相关知识，模拟实现linux平台cat命令
func cat() {

}

func main() {
	// openfile()
	// readfile()
	// cycleread()
	// bioread()
	// iotil()
	// writefile()
	// iotilwrite()
	// iotilwritefile()
	// copyFile("dst.txt", "src.txt")
}
