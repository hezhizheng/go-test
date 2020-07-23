package main

import (
	"log"
)

// 定义二分查找的接口
type BinaryQueryInterface interface {
	main() int
}

// 使用递归的方式实现接口
type Recursive struct {
	BinaryQueryInterface
}

// 定义递归的参数
type RecursiveParam struct {
	Ary    [5]int
	Target int
	Low    int
	Top    interface{}
}

// 实现递归二分的方法
func (r *Recursive) main() int {
	ary := [...]int{1, 2, 4, 5, 6}
	// 如果需要默认值最好是通过结构体去定义
	rp := RecursiveParam{
		Ary:    ary,
		Target: 1,
		Low:    0,
		Top:    nil,
	}
	return recursive(rp)
}

// 递归二分查找算法
func recursive(param RecursiveParam) int {
	top := param.Top
	low := param.Low
	ary := param.Ary
	target := param.Target

	if param.Top == nil {
		top = len(ary)
	}

	// 返回的是一个不大于mid的int值
	mid := (top.(int) + low) / 2

	if ary[mid] == target {
		return mid
	} else if ary[mid] > target {
		param.Top = mid - 1
		return recursive(param)
	} else {
		param.Low = mid + 1
		return recursive(param)
	}

}

func main() {
	// 实现递归接口
	c := BinaryQueryInterface.main(&Recursive{})
	log.Printf("通过递归查找的数组角标为 %d", c)
}
