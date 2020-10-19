package main

import "log"

// 注意：无缓冲通道的收发操作必须在不同的两个goroutine间进行，因为通道的数据在没有接收方处理时，数据发送方会持续阻塞，所以通道的接收必定在另外一个 goroutine 中进行。

func main() {
    ch := make(chan int, 5)

    // 写入channel 不能超过预定的容量 否则 all goroutines are asleep - deadlock!
    for i := 0; i <= 4; i++ {
        ch <- i
    }

    // 遍历读取前先close掉 ，否则 all goroutines are asleep - deadlock!
    close(ch)

    for v := range ch {
        log.Println(v, len(ch), ch)
    }

}
