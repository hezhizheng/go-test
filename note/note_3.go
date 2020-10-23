//对于无缓冲的 channel，发送方将阻塞该信道，直到接收方从该信道接收到数据为止，而接收方也将阻塞该信道，直到发送方将数据发送到该信道中为止。

//对于有缓存的 channel，发送方在没有空插槽（缓冲区使用完）的情况下阻塞，而接收方在信道为空的情况下阻塞。

package main

var ch1 = make(chan int)
var ch2 = make(chan int , 2)



func main()  {


	// 没有缓冲的channel，接收方需要在goroutine中收到信号，否则 发送方阻塞直到接收方接收到数据。
	go func() {
		<-ch1
		<-ch1
	}()

	ch1 <- 1
	ch1 <- 1

	 ch2 <- 1
	 ch2 <- 1

	<-ch2
	<-ch2

	 //ch2 <- 1

//	select {
//	default:
//
//}



}
