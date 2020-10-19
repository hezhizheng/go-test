package main

import (
	"fmt"
	"time"
)

// select 语句会阻塞当前 Goroutine 并等待多个 Channel 中的一个达到可以收发的状态
// select 能在 Channel 上进行非阻塞的收发操作；
// select 在遇到多个 Channel 同时响应时会随机挑选 case 执行

func main()  {
	c:= make(chan int)
	quit:= make(chan int)

	go enterQ(quit)
	go enterQ(quit)

	go quitC(c)
	go quitC(c)

	go fibonacci(c, quit)

	time.Sleep(2 * time.Second)
}

func enterQ(quit chan int)  {
	quit <- 2
}

func quitC(c chan int)  {
	<-c
}

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x: // 有进要有出
			x, y = y, x+y
			fmt.Println("set c ",x, y)
			//return
		case <-quit: // 有出要先有进
			fmt.Println("quit")
			return
		}
	}
}