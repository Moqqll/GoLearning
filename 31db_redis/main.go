package main

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8" //最新版本的go-redis库的相关命令都需要传递context.Context参数，
)

//err 不要使用全局变量声明
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

	//检测连接是否成功
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	_, err = rdbConn.Ping(ctx).Result()
	return err
}

//GetsetDemo ...
func GetsetDemo(ctx context.Context) {

	err := rdbConn.Set(ctx, "moqqll", "大帅比", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdbConn.Get(ctx, "moqqll").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("moqqll：", val)

	val2, err := rdbConn.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2：", val2)
	}
}

//ZsetDemo ...
func ZsetDemo(ctx context.Context) {
	zsetKey := "language_rank"
	languages := []*redis.Z{
		{Score: 90.0, Member: "Golang"},
		{Score: 98.0, Member: "Java"},
		{Score: 95.0, Member: "Python"},
		{Score: 97.0, Member: "JavaScript"},
		{Score: 99.0, Member: "C/C+++"},
	}

	//Z-Add
	num, err := rdbConn.ZAdd(ctx, zsetKey, languages...).Result()
	if err != nil {
		fmt.Printf("z-add failed, err:%v\n", err)
		return
	}
	fmt.Printf("z-add %d succ.\n", num)

	//把golang的分数加10
	newScore, err := rdbConn.ZIncrBy(ctx, zsetKey, 10, "Golang").Result()
	if err != nil {
		fmt.Printf("zincy failed, err:%v\n", err)
		return
	}
	fmt.Printf("Golang's score is %f now.\n", newScore)

	//取分数最高的3个
	ret, err := rdbConn.ZRevRangeWithScores(ctx, zsetKey, 0, 2).Result()
	if err != nil {
		fmt.Printf("zrevrange failed, err%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}

	//取95~100分的
	opt := redis.ZRangeBy{
		Min: "95",
		Max: "100",
	}
	ret, err = rdbConn.ZRangeByScoreWithScores(ctx, zsetKey, &opt).Result()
	if err != nil {
		fmt.Printf("zrangebyscore failed, err:%v\n", err)
		return
	}
	for _, z := range ret {
		fmt.Println(z.Member, z.Score)
	}
}

func main() {
	//init redis
	if err := initClient(); err != nil {
		fmt.Printf("init redis failed, err:%v\n", err)
		return
	}
	defer rdbConn.Close()
	ctx := context.Background()

	//

	//ZsetDemo
	ZsetDemo(ctx)

	//GetsetDemo
	// GetsetDemo(ctx)

}
