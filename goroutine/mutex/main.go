package main

import (
	"fmt"
	"sync"
)
var mutex sync.Mutex // 互斥锁，加锁后，代码块变串行
// 问题：假设10w的并发，请求的并不是同一件商品，也会加锁
// 只对单个服务器有效，多个服务器无效
func main()  {
	for i:=0;i<100;i++ {
		mutex.Lock()

		go func() {
			defer mutex.Unlock()

			fmt.Println(i)
		}()

	}
}
