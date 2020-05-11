package main

import "log"

type People struct {
    Name string
    Age string
}

//  Go 语言面向对象编程不像 PHP、Java 那样支持隐式的 this 指针，所有的东西都是显式声明的，在 GetXXX 方法中，由于不需要对类的成员变量进行修改，所以不需要传入指针，而 SetXXX 方法需要在函数内部修改成员变量的值，并且作用到该函数作用域以外，所以需要传入指针类型（结构体是值类型，不是引用类型，所以需要显式传入指针）。

// 无所谓指针类型
func (p *People) GetName() string {
    return p.Name
}

// 涉及到修改成员变量的值需要传入指针类型
func (p *People) SetName(name string) {
    p.Name = name
}

func newPeople(name , Age string) *People {
    return &People{
        Name: name,
        Age:  Age,
    }
}

func main() {
    n := newPeople("asda","10")
    n.SetName("ttttt")

    log.Println(n,n.GetName())
}
