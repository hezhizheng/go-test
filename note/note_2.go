// 一个 go 文件的执行顺序 import –> const –> var –> init() –> main()
// 同一个包内多个 init() 函数的执行顺序不作保证。
package main

import "fmt"

func init()  {
	fmt.Println(`init1`)

}

func init()  {
	fmt.Println(`init2`)
}

func main()  {
	//
}