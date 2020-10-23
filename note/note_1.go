package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	equal()
	quote()
	splicingStr()
	_rune()
	_defer()
}

// = 和 := 的区别
func equal() {
	fmt.Println("test = ：=")
	// = 用于变量的赋值
	// := 变量的声明+赋值
	var name string
	name = "equal1"
	v := "equal"
	v = "equal2"

	log.Println(name, v)
}

// 指针的作用
func quote() {
	fmt.Println("test 指针")
	// 指针用来保存变量的地址。
	x := 5

	// &x 为变量x的地址
	var p *int = &x
	c := &x

	// 使用*c 、*p 可取回地址中的值
	log.Println(x, p, *p, *c)
	fmt.Printf("%+v \n",x)
	//* 运算符，也称为解引用运算符，用于访问地址中的值。
	//＆运算符，也称为地址运算符，用于返回变量的地址
}

// Go 没有异常类型，只有错误类型（Error），通常使用返回值来表示异常状态。
func showErr() {
	_, err := os.Open("test.txt")
	if err != nil {
		log.Fatal(err)
	}
}

// Goroutine 是与其他函数或方法同时运行的函数或方法。 Goroutines 可以被认为是轻量级的线程。
// 与线程相比，创建 Goroutine 的开销很小。 Go应用程序同时运行数千个 Goroutine 是非常常见的做法。

// 如何高效拼接字符串
func splicingStr() {
	fmt.Println("test splicingStr")
	// Go 语言中，字符串是只读的，也就意味着每次修改操作都会创建一个新的字符串。
	//如果需要拼接多次，应使用 strings.Builder，最小化内存拷贝次数。

	var str strings.Builder
	for i := 0; i < 100; i++ {
		str.WriteString("a")
	}
	fmt.Println(str.String())

}

// 什么是rune类型
func _rune() {
	fmt.Println("test _rune")
	// int32 类型的别名 ，跟字节长度有关
	fmt.Println(len("Go语言"))         // utf-8 一个中文占3个字节  8
	fmt.Println(len([]rune("Go语言"))) // 4
}

// defer 的执行顺序
func _defer() ( i int) {
	fmt.Println("test _defer")
	// defer遵循后进先出原则，后声明的defer会先得到执行
	i = 0

	defer func() {
		fmt.Println("defer1")
	}()

	defer func() {
		i += 1
		fmt.Println("defer2",i)
	}()

	fmt.Println(i)
	return i
	// defer 在函数退出前、函数返回之后执行
	// 但在函数退出前还可以修改局部变量的值，前提必须的是有名返回值，即返回值声明了变量 ( i int )
}
