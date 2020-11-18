package main

/*
使用接口的方式实现一个既可以往终端打印日志，也可以向文件中写入日志的简易日志库
*/

//Logger ...
type Logger interface {
	Info(string)
}

//FileLogger ...
type FileLogger struct {
}

//ConsoleLogger ...
type ConsoleLogger struct {
}

//Info ...
// func (f *FileLogger) Info(msg string) {
// 	// var tmpf *os.File
// 	// var tmperr error

// }

func main() {}
