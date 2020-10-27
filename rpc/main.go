package main

import (
	"log"
	"net/http"
	"net/rpc"
)

type Calculation struct {
}

type Req struct {
	Num, Ans int
}

type Res struct {
	Num, Ans, Val int
}

// 计算数值得次方数
func (cal *Calculation) exec(num, ans int) *Res {

	multi := num

	if ans == 0 {
		multi = 1
	}

	for i := 1; i <= ans-1; i++ {
		multi *= num
	}

	return &Res{
		Num: num,
		Ans: ans,
		Val: multi,
	}
}

// func (t *T) MethodName(argType T1, replyType *T2) error
func (cal *Calculation) ExecRpc(req Req, res *Res) error {

	multi := req.Num

	if req.Ans == 0 {
		multi = 1
	}

	for i := 1; i <= req.Ans-1; i++ {
		multi *= req.Num
	}

	res.Val = multi
	res.Num = req.Num
	res.Ans = req.Ans

	return nil
}

func main() {
	// make 关键字的作用是创建切片、哈希表和 Channel 等内置的数据结构，
	//而 new 的作用是为类型申请一片内存空间，并返回指向这片内存的指针。与 var 没有啥区别 ，new返回的是指针, var 返回的是具体值

	cal := new(Calculation) // => var cal Calculation
	var varCal Calculation // => *cal
	log.Println(varCal,cal,*cal)
	val := cal.exec(12, 2)

	log.Printf("%v^%v=%v", val.Num, val.Ans, val.Val)
}

func main3() {
	cal := new(Calculation)

	var res = Res{}
	cal.ExecRpc(Req{
		Num: 2,
		Ans: 3,
	}, &res)

	log.Printf("%v^%v=%v", res.Num, res.Ans, res.Val)
}

func main5() {
	rpc.Register(new(Calculation))
	rpc.HandleHTTP()

	log.Printf("Serving RPC server on port %d", 1234)
	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("Error serving: ", err)
	}
}
