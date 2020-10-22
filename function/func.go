package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

var wg = sync.WaitGroup{}

var ch = make(chan int, 3)

func main() {

	wg.Add(1)
	// only one goroutine
	go task()

	// 等待所有的协程完成
	wg.Wait()

	log.Println("ok2")
}

func task() {

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go Uniqid("id_")
		//time.Sleep(time.Second)
	}
	//wg.Wait()
	// 如果 Uniqid 里面没有使用channel 这里不在Wait 后面做输出，Uniqid 中的输出都不会有

	wg.Done()
	log.Println("ok")

}

func Uniqid(prefix string) string {
	ch <- 1

	now := time.Now()
	sec := now.Unix()
	usec := now.UnixNano() % 0x100000
	s := fmt.Sprintf("%s%08x%05x", prefix, sec, usec)

	time.Sleep(time.Second)

	<-ch
	wg.Done()
	log.Println(s)
	return s
}
