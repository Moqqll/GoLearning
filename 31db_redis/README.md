# go-redis

Redis是一个开源的内存数据库，Redis提供了多种不同类型的数据结构，很多业务场景下的问题都可以很自然地映射到这些数据结构上。除此之外，通过复制、持久化和客户端分片等特性，我们可以很方便地将Redis扩展成一个能够包含数百GB数据、每秒处理上百万次请求的系统。

## Redis支持的数据结构

Redis支持诸如字符串（strings）、哈希（hashes）、列表（lists）、集合（sets）、带范围查询的排序集合（sorted sets）、位图（bitmaps）、hyperloglogs、带半径查询和流的地理空间索引等数据结构（geospatial indexes）。

## Redis应用场景

- 缓存系统，减轻主数据库（MySQL）的压力。
- 计数场景，比如微博、抖音中的关注数和粉丝数。
- 热门排行榜，需要排序的场景特别适合使用ZSET。
- 利用LIST可以实现队列的功能。

## 准备Redis环境

* docker
* windows安装

## 安装go-redis

区别于另一个比较常用的Go语言redis client库：[redigo](https://github.com/gomodule/redigo)，我们这里采用https://github.com/go-redis/redis连接Redis数据库并进行操作，因为`go-redis`支持连接哨兵及集群模式的Redis。

使用以下命令下载并安装:

```go
go get -u github.com/go-redis/redis
```

## 普通连接

```go
func initClient() (err error) {
	rdbCon = redis.NewClient(&redis.Options{
		Addr:     "1270.0.1:6379",
		Password: "", //no psasword
		DB:       0,  //use default DB
	})

	_, err := rdbCon.Ping().Result()
	if err != nil {
		return err
	}

	return nil
}
```

## V8新版本相关

### 基础连接

最新版本的`go-redis`库的相关命令都需要传递`context.Context`参数，例如：

```go
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
```

### 连接redis哨兵模式









### 连接redis集群





## 基本使用

### set/get示例

```go
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
```

### zset示例

```

```

