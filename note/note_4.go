package main

import "fmt"

func main()  {
	// 在同一个 const group 中，如果常量定义与前一行的定义一致，则可以省略类型和值。
	//编译时，会按照前一行的定义自动补全。即等价于
	//const (
	//		a, b = "golang", 100
	//		d, e = "golang", 100
	//		f bool = true
	//		g bool = true
	//	)
	const (
		a, b = "golang", 100
		d, e
		f bool = true
		g
	)

	fmt.Println(d,e,g)
	complement()
}

func complement()  {
	var a int8 = -1
	var b int8 = -128 / a
	fmt.Println(b) // 计算机补码
}
