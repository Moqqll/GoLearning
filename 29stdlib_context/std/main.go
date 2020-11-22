// package main

// import (
// 	"context"
// 	"fmt"
// 	"sync"
// 	"time"
// )

// var wg sync.WaitGroup
// var exit bool

// //全局变量方式存在的问题
// //1、使用全局变量在跨包调用时不容易统一
// //2、如果woker中再启动goroutine，就不太好控制

// //管道方式存在的问题
// //1、使用全局变量在跨包调用时不容易实现规范和统一，需要维护一个公用的channel

// func worker() {
// 	for {
// 		fmt.Println("worker")
// 		time.Sleep(time.Second)
// 		//如何接收外部命令实现退出
// 		// if exit {
// 		// 	break
// 		// }
// 	}
// 	// wg.Done()
// }

// func worker1(exitChan chan struct{}) {
// LOOP:
// 	for {
// 		fmt.Println("worker1")
// 		time.Sleep(time.Second)
// 		select {
// 		case <-exitChan:
// 			break LOOP
// 		default:
// 		}
// 	}
// 	wg.Done()
// }

// func ctxWorker(ctx context.Context) {
// 	go ctxWorker2(ctx)
// LOOP:
// 	for {
// 		fmt.Println("ctxWorker")
// 		time.Sleep(time.Second)
// 		select {
// 		case <-ctx.Done():
// 			break LOOP
// 		default:
// 		}
// 	}
// 	wg.Done()
// }

// func ctxWorker2(ctx context.Context) {
// LOOP:
// 	for {
// 		fmt.Println("ctxWorker2")
// 		time.Sleep(time.Second)
// 		select {
// 		case <-ctx.Done():
// 			break LOOP
// 		default:
// 		}
// 	}
// 	wg.Done()
// }

// func main() {
// 	// var exitChan = make(chan struct{})
// 	ctx, cancel := context.WithCancel(context.Background())
// 	wg.Add(1)
// 	// go worker()
// 	// go worker1(exitChan)
// 	go ctxWorker(ctx)
// 	time.Sleep(time.Second * 3)
// 	//如何优雅的结束子goroutine
// 	// exit = true //修改全局变量实现子goroutine的退出
// 	// exitChan <- struct{}{} //给goroutine发送退出信号
// 	// close(exitChan)
// 	cancel() //同志子goroutine结束
// 	wg.Wait()
// 	fmt.Println("over")
// }
// package main

// import (
// 	"context"
// 	"fmt"
// )

// func gen(ctx context.Context) <-chan int {
// 	dst := make(chan int)
// 	n := 1
// 	go func() {
// 		for {
// 			select {
// 			case <-ctx.Done():
// 				return //return结束该goroutine，防止泄漏
// 			case dst <- n:
// 				n++
// 			}
// 		}
// 	}()
// 	return dst
// }

// func main() {
// 	ctx, cancel := context.WithCancel(context.Background())
// 	defer cancel() //当我们取完需要的整数后调用cancel

// 	for n := range gen(ctx) {
// 		fmt.Println(n)
// 		if n == 5 {
// 			break
// 		}
// 	}
// }

// package main

// import (
// 	"context"
// 	"fmt"
// 	"sync"
// 	"time"
// )

// //context WithTimeout

// var wg sync.WaitGroup

// func worker(ctx context.Context) {
// LOOP:
// 	for {
// 		fmt.Println("db connecting...")
// 		time.Sleep(time.Millisecond * 10) //假设正常连接数据库需要10ms
// 		select {
// 		case <-ctx.Done(): //50ms后自动调用
// 			break LOOP
// 		default:
// 		}
// 	}
// 	fmt.Println("worker done!")
// 	wg.Done()
// }

// func main() {
// 	//设置一个50ms的超时
// 	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
// 	wg.Add(1)
// 	go worker(ctx)
// 	time.Sleep(time.Second * 5)
// 	cancel() //通知子goroutine结束
// 	wg.Wait()
// 	fmt.Println("over")
// }

package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

//context.WithValue

//TraceCode ...
type TraceCode string

var wg sync.WaitGroup

func worker(ctx context.Context) {
	key := TraceCode("TRACE_CODE")
	tracecode, ok := ctx.Value(key).(string)
	if !ok {
		fmt.Println("invalid trace code")
	}
LOOP:
	for {
		fmt.Printf("worker, trace code:%s\n", tracecode)
		time.Sleep(time.Millisecond * 10) // 假设正常连接数据库耗时10毫秒
		select {
		case <-ctx.Done(): // 50毫秒后自动调用
			break LOOP
		default:
		}
	}
	fmt.Println("worker done!")
	wg.Done()
}

func main() {
	//设置一个50ms的超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*50)
	ctx = context.WithValue(ctx, TraceCode("TRACE_CODE"), "123456789")
	wg.Add(1)
	go worker(ctx)
	time.Sleep(time.Second * 5)
	cancel()
	wg.Wait()
	fmt.Println("over")
}
