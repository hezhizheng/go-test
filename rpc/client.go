package main

import (
	"log"
	"net/rpc"
)

type Req1 struct {
	Num, Ans int
}

type Res1 struct {
	Num, Ans, Val int
}

func main() {
	client, _ := rpc.DialHTTP("tcp", "localhost:1234")

	result := Res1{}

	args := Req1{
		Num: 4,
		Ans: 3,
	}
	if err := client.Call("Calculation.ExecRpc", args, &result); err != nil {
		log.Fatal("Failed to call Calculation.ExecRpc. ", err)
	}

	log.Printf("%v^%v=%v", result.Num, result.Ans, result.Val)
}
