/**
封装：通过方法实现

继承：通过匿名字段实现

多态：通过接口实现
**/
package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	*Person //只有类型，没有名字，匿名字段，继承Person成员
	id      int
	addr    string
}

func main() {

	var s1 Student = Student{&Person{"Json", 'm', 22}, 12, "sz"}

	//%+v 显示更详细
	fmt.Printf("s1 = %+v \n", s1)

	fmt.Println(s1.Person)

	// var s2 Student

}
