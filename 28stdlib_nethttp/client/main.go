package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

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

func main() {
	myGet()
	myPost()
}
