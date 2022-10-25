package main

import (
	"fmt"
	"time"

	goredislib "github.com/go-redis/redis/v8"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v8"
	"sync"
)

func main() {
	// Create a pool with go-redis (or redigo) which is the pool redisync will
	// use while communicating with Redis. This can also be any pool that
	// implements the `redis.Pool` interface.
	//这里的变量哪些可以放到global中， redis的配置是否应该在nacos中
	client := goredislib.NewClient(&goredislib.Options{
		Addr: "192.168.0.104:6379",
	})
	pool := goredis.NewPool(client) // or, pool := redigo.NewPool(...)

	// Create an instance of redisync to be used to obtain a mutual exclusion
	// lock.
	rs := redsync.New(pool)

	// Obtain a new mutex by using the same name for all instances wanting the
	// same lock.
	gNum := 2
	mutexname := "421"

	var wg sync.WaitGroup
	wg.Add(gNum)
	for i :=0 ;i<gNum;i++ {
		go func() {
			defer wg.Done()
			/**
			多节点redis实现分布式锁算法RedLock:有效防止单点故障
			假设有5台独立的redis服务器
			1.获取时间戳问题
			2.client尝试按照顺序使用相同的key,value获取所有redis服务的锁，在获取的过程中的获取时间比锁过期时间短很多，这是为了不要过长时间等待已经关闭的redis服务，并且试着获取下一个redis实例。
				比如：TTL为5,设置获取锁最多1秒，所以如果1秒内无法获取锁，就放弃获取这个锁，从而尝试获取下个锁
			3.client获取所有能获取的锁后的时间减去第一步的时间，这个时间差要小于TTL的时间且至少有3个redis实例成功获取锁，才算真正的获取锁成功
			4.如果获取锁成功，则锁的真正有效时间减去第三步的时间差的时间，比如TTL是5秒，获取所有锁用了2秒，则真正的锁的有效时间为3秒（其实应该减去时钟漂移）
			5.如果客户端由于某种原因获取锁失败，便会解锁所有锁的实例；因为有可能获取了 小于3个redis锁实例，必须释放掉，否则影响其他client获取
			**/
			mutex := rs.NewMutex(mutexname)

			fmt.Println("开始获取锁")
			if err := mutex.Lock(); err != nil {
				panic(err)
			}

			fmt.Println("获取锁成功")

			time.Sleep(time.Second*8) // 业务逻辑

			fmt.Println("开始释放锁")
			if ok, err := mutex.Unlock(); !ok || err != nil {
				panic("unlock failed")
			}
			fmt.Println("释放锁成功")
		}()
	}
	wg.Wait()
}