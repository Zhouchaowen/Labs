package main

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v9"
)

/*
1.设置key-value到redis
2.通过key获取value
*/
func main() {
	ctx := context.Background()

	rdb := redis.NewClient(&redis.Options{
		Addr:     "10.2.0.104:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		panic(err)
	}

	err = rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		panic(err)
	} else {
		fmt.Println("key2", val2)
	}

	// Output: key value
	// key2 does not exist
}
