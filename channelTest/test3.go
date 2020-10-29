package main

import (
	"log"
	"time"
)

var ch1 = make(chan interface{},2)

func task1(i interface{}) {

	ch1 <- i
	log.Println("task1",i)
	time.Sleep(time.Second)
	<-ch1

}

func main() {

	//go func() {
	//	<-ch1
	//}()
	//ch1 <- 12
	//log.Println(ch1)
	//for i := 1; i <= 5; i++ {
	//	ch1 <- 12
	//}

	for i := 1; i <= 100; i++ {
		go task1(i)
		go task1(i)
	}

	for  {
		//select {
		//case <-ch1:
		//default:
		//
		//
		//
		//}
		//if len(ch1) == 0 {
		//	log.Println("done")
		//	break
		//}
	}

}
