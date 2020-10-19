package main

import "time"


/**
errCh := make(chan error, len(tasks))
wg := sync.WaitGroup{}
wg.Add(len(tasks))
for i := range tasks {
    go func() {
        defer wg.Done()
        if err := tasks[i].Run(); err != nil {
            errCh <- err
        }
    }()
}
wg.Wait()

// 当 select 中仅包含两个 case，并且其中一个是 default 时，Go 语言的编译器就会认为这是一次非阻塞的收发操作
select {
case err := <-errCh:
    return err
default:
    return nil
}
 */

func main() {
	ch := make(chan int)
	go func() {
		for range time.Tick(1 * time.Second) {
			ch <- 0
		}
	}()

	for {
		// 多个 <-ch 同时满足可读或者可写条件时会随机选择一个 case 执行其中的代码
		select {
		case <-ch:
			println("case1")
		case <-ch:
			println("case2")
		//default: // channel没有收发信号的时候回直接执行 default
		//	println("default")
		}

	}
}
