package main

import (
	"fmt"
	"time"
)

func worker(id int, in <-chan int, out chan<- int) {
	for i := range in {
		fmt.Printf("worker:%d start job:%d\n", id, i)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d stop job:%d\n", id, i)
		out <- i * i
	}
}

func main() {
	// ch1 := make(chan int)
	// ch2 := make(chan int)

	// //开启goroutine将0~100的数发送到ch1中
	// go func() {
	// 	for i := 0; i < 100; i++ {
	// 		ch1 <- i
	// 	}
	// 	close(ch1)
	// }()

	// //开启goroutine取出ch1的值，并求平方后发送到ch2中
	// //利用for无限循环从通道中依次取值
	// go func() {
	// 	for {
	// 		i, ok := <-ch1
	// 		if !ok { //如果值去完了，就退出for循环
	// 			break
	// 		}
	// 		ch2 <- i * i
	// 	}
	// 	close(ch2)
	// }()

	// //利用for range从通道中依次取值
	// for i := range ch2 { //取完值后，自动退出for range循环
	// 	fmt.Println(i)
	// }

	jobs := make(chan int, 100)
	results := make(chan int, 100)

	//开启3个goroutine
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	//5个任务
	for i := 1; i <= 5; i++ {
		jobs <- i
	}
	close(jobs)

	//打印结果
	for r := 0; r < 5; r++ {
		ret := <-results
		fmt.Println(ret)
	}
}
