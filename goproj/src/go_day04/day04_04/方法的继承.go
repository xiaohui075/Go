package main

import "fmt"

type Person struct {
	name string
	sex  byte
	age  int
}

type Student struct {
	Person
	id   int
	addr string
}

func (tmp *Person) PrintlnInfo() {
	fmt.Printf("name=%s,sex=%c,age=%d\n", tmp.name, tmp.sex, tmp.age)
}

func (tmp *Student) PrintlnInfo() {
	fmt.Println("Student: tmp = ", tmp)
	fmt.Printf("Pointer: %p , %v \n", tmp, tmp)
}

func main() {
	s := Student{Person{"Tom", 'm', 22}, 66, "bj"}
	//就近原则 ，先找本作用域的方法，找不到再用继承的方法
	s.PrintlnInfo()
	//显示调用继承的方法
	s.Person.PrintlnInfo()

	sFunc := s.PrintlnInfo //这个就是方法值，调动函数时，无需再传递接受者，隐藏了接受者
	sFunc()
}
