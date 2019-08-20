package main

import "fmt"

type Person struct {
	name string
	age  int
	sex  byte
}

//带有接受者的函数叫方法
func (tmp Person) PrintlnInfo() {
	fmt.Println("tmp=", tmp)
}

//通过一个函数，给成员赋值
func (p *Person) SetInfo(n string, age int, sex byte) {
	p.name = n
	p.age = age
	p.sex = sex
}

func main() {
	p1 := Person{"Tom", 22, 'm'}
	p1.PrintlnInfo()

	var p2 Person
	(&p2).SetInfo("yoyo", 18, 'f')
	p2.PrintlnInfo()
}
