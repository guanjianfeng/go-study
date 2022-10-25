package main

import (
	"fmt"
	"sync"
)
var wg	sync.WaitGroup

func main()  {
	count :=0
	for i:=0;i<100;i++{
		wg.Add(1)
		go func() {
			count+=1
			wg.Done()
		}()
	}
	wg.Wait() // 防止主进程退出，导致协程退出
	fmt.Println(count)
}