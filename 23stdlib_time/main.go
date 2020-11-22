package main

import (
	"fmt"
	"time"
)

func timeDemo() {
	now := time.Now()
	fmt.Printf("current time:%v\n", now)

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%-2d\n", year, month, day, hour, minute, second)
}

func timestampDemo() int64 {
	now := time.Now()
	timestamp1 := now.Unix()     //时间戳
	timestamp2 := now.UnixNano() //纳秒时间戳
	fmt.Printf("current timestamp1:%v\n", timestamp1)
	fmt.Printf("current timestamp2:%v\n", timestamp2)
	return timestamp1
}

func timestampDemo2(timestamp int64) {
	timeObj := time.Unix(timestamp, 0)
	fmt.Println(timeObj)
	year := timeObj.Year()
	month := timeObj.Month()
	day := timeObj.Day()
	hour := timeObj.Hour()
	minute := timeObj.Minute()
	second := timeObj.Second()
	fmt.Printf("%d-%02d-%02d %02d:%02d:%-2d\n", year, month, day, hour, minute, second)
}

func tickDemo() {
	ticker := time.Tick(time.Second)
	for i := range ticker {
		hour := i.Hour()
		minute := i.Minute()
		second := i.Second()
		fmt.Printf("%02d-%02d-%02d\n", hour, minute, second)
	}
}

func formatDemo() {
	now := time.Now()
	fmt.Println(now.Format("2006-01-02 15:04:05.000 Mon Jan"))
	fmt.Println(now.Format("2006-01-02 03:04:05.000 PM Mon Jan"))
	fmt.Println(now.Format("2006/01/02 15:04"))
	fmt.Println(now.Format("15:04 2006/01/02"))
	fmt.Println(now.Format("2006/01/02"))
}

func parseStringTimeDemo() {
	now := time.Now()
	fmt.Println(now)
	//加载时区
	loc, err := time.LoadLocation("Asia/Shanghai")
	if err != nil {
		fmt.Println(err)
		return
	}
	//按照指定时区和指定格式解析字符串时间
	timeObj, err := time.ParseInLocation("2006/01/02 15:04", "2020/11/22 03:10", loc)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(timeObj)
	fmt.Println(timeObj.Sub(now))
}

func main() {
	// timeDemo()
	// tmptimestamp := timestampDemo()
	// timestampDemo2(tmptimestamp)
	// tickDemo()
	// formatDemo()
	parseStringTimeDemo()
}
