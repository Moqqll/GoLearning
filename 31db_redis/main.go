package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
)

var rdbConn *redis.Client

//初始化连接
func initClient() (err error) {
	rdbConn = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",  //no psasword
		DB:       0,   //use default DB
		PoolSize: 100, //连接池大小
	})

	// _, err := rdbCon.Ping().Result()
	// if err != nil {
	// 	return err
	// }

	// return nil

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err = rdbConn.Ping(ctx).Result()
	return err

}

//V8Example ...
func V8Example() {
	ctx := context.Background()
	if err := initClient(); err != nil {
		return
	}

	err := rdbConn.Set(ctx, "moqqll", "大帅比", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdbConn.Get(ctx, "moqqll").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("moqqll", val)

	val2, err := rdbConn.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}
}

func main() {
	V8Example()
}
