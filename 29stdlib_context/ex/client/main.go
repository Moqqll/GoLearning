package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type respData struct {
	resp *http.Response
	err  error
}

func doCall(ctx context.Context) {
	//新建一个client
	transport := http.Transport{
		DisableKeepAlives: true,
	}
	client := http.Client{
		Transport: &transport,
	}

	//封装request
	respChan := make(chan *respData, 1)
	req, err := http.NewRequest("GET", "http://127.0.0.1:9000", nil)
	if err != nil {
		fmt.Printf("new request failed, err:%v\n", err)
		return
	}
	req = req.WithContext(ctx) //使用带超时的ctx创建一个新的client request

	//开始进行异步request
	var wg sync.WaitGroup
	wg.Add(1)
	defer wg.Wait()
	go func() {
		resp, err := client.Do(req)
		fmt.Printf("client.do resp:%v, err:%v\n", resp, err)
		rd := &respData{
			resp: resp,
			err:  err,
		}
		respChan <- rd
		wg.Done()
	}()

	select {
	case <-ctx.Done():
		fmt.Println("call api timeout")
	case result := <-respChan:
		fmt.Println("call server api success")
		if result.err != nil {
			fmt.Println("call server api data failed, err:", err)
			return
		}
		defer result.resp.Body.Close()
		data, _ := ioutil.ReadAll(result.resp.Body)
		fmt.Printf("resp data:%v\n", string(data))
	}
}

func main() {
	//定义一个100ms的超时
	ctx, cancel := context.WithTimeout(context.Background(), time.Millisecond*100)
	defer cancel()
	doCall(ctx)
}
