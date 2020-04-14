package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	wg1 sync.WaitGroup
)

func task2(poolChan chan int) {
	defer wg1.Done()

	// 执行一次就往 管道 中写入一次
	poolChan <- 1

	fmt.Println("this is  channel + WaitGroup task 2 ")
	time.Sleep(time.Second * 2)

	// 从管道中取出
	<-poolChan
}

func main() {

	pool := make(chan int, 5)

	for i := 1; i <= 10; i++ {
		wg1.Add(1)
		go task2(pool)
	}

	//close(pool)
	wg1.Wait()

	// task done
	fmt.Println(" channel + select task done")
	fmt.Println(" channel len", len(pool))

}

// 煎鱼事例

func main3() {
	userCount := 10
	ch := make(chan bool, 2)
	for i := 0; i < userCount; i++ {
		wg1.Add(1)
		go Read(ch, i)
	}

	wg1.Wait()
}

func Read(ch chan bool, i int) {
	defer wg1.Done()

	ch <- true
	fmt.Printf("go func: %d, time: %d\n", i, time.Now().Unix())
	time.Sleep(time.Second)
	<-ch
}
