package main

import (
	"context"
	"log"
	"time"
)

// 定义 Context 接口，用于主动停止正在运行的 goroutine
type ContextInterFace interface {
	main()
}

// 使用 channel + select 的方式实现停止运行中的 goroutine
type ChanSelect struct {
	ContextInterFace
}

// 使用 context 的方式实现停止运行中的 goroutine
type ContextEntity struct {
	ContextInterFace
}

// 使用 context 的方式实现停止运行中的 多个 goroutine
type Multiple struct {
	ContextInterFace
}

func (c *ChanSelect) main() {

	// make 一个用于停止 goroutine 的 channel
	stop := make(chan bool)

	go func() {
		for {
			select {
			case <-stop: // channel 只要能取出就停止
				log.Println("监控退出，停止了...")
				return
			default:
				log.Println("goroutine监控中...")
				time.Sleep(2 * time.Second)
			}
		}
	}()

	time.Sleep(10 * time.Second)
	log.Println("可以了，通知监控停止")

	// 发送停止 goroutine 的信号
	stop <- true
	// 为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)

	log.Println("来自 ChanSelect")
}

func (c *ContextEntity) main() {

	ctx, cancel := context.WithCancel(context.Background())
	go func(ctx context.Context) {
		for {
			select {
			case <-ctx.Done():
				log.Println("监控退出，停止了...")
				return
			default:
				log.Println("goroutine监控中...")
				time.Sleep(2 * time.Second)
			}
		}
	}(ctx)

	time.Sleep(10 * time.Second)
	log.Println("可以了，通知监控停止")

	// 发送停止 goroutine 的信号
	cancel()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)

	log.Println("来自 ContextEntity")
}

func (c *Multiple) main() {
	ctx, cancel := context.WithCancel(context.Background())

	// 所有goroutine 都传入 Context 进行跟踪
	go watch(ctx, "【监控1】")
	go watch(ctx, "【监控2】")
	go watch(ctx, "【监控3】")

	time.Sleep(10 * time.Second)
	log.Println("可以了，通知监控停止")

	// 所有衍生的context收到cancel信号，所监控的goroutine都会停止
	cancel()
	//为了检测监控过是否停止，如果没有监控输出，就表示停止了
	time.Sleep(5 * time.Second)

	log.Println("来自 Multiple")
}

func watch(ctx context.Context, name string) {
	for {
		select {
		case <-ctx.Done():
			log.Println(name, "监控退出，停止了...")
			return
		default:
			log.Println(name, "goroutine监控中...")
			time.Sleep(2 * time.Second)
		}
	}
}

func main() {
	//ContextInterFace.main(&ChanSelect{})
	//ContextInterFace.main(&ContextEntity{})
	ContextInterFace.main(&Multiple{})
}
