package main

import (
	"fmt"
	"sync"
	"time"
)

// var x int64
// var wg sync.WaitGroup

// func add() {
// 	for i := 0; i < 5; i++ {
// 		x += i
// 	}
// 	wg.Done()
// }

// func main() {
// 	wg.Add(2)
// 	go add()
// 	go add()
// 	wg.Wait()
// 	fmt.Println(x)
// }
var (
	x      int64
	wg     sync.WaitGroup
	lock   sync.Mutex
	rwlock sync.RWMutex
)

func write() {
	// lock.Lock()
	rwlock.Lock() //加写锁
	x = x + 1
	time.Sleep(10 * time.Millisecond) //假设读操作耗时10ms
	rwlock.Unlock()                   //解写锁
	// lock.Unlock()
	wg.Done()
}

func read() {
	// lock.Lock()
	rwlock.RLock()               //加读锁
	time.Sleep(time.Millisecond) //假设读操作耗时1ms
	rwlock.RUnlock()             //解读锁
	// lock.Unlock()
	wg.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go write()
	}
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go read()
	}
	wg.Wait()
	fmt.Println(x)
	end := time.Now()
	fmt.Println(end.Sub(start))
}
