package main

import "fmt"

type FuncType func(int, int) int

func Add(a, b int) int {
	return a + b
}

//回调函数，函数有一个参数是函数类型，这个函数就是回调函数
func Calc(a, b int, fTest FuncType) (result int) {
	fmt.Println("Calc")
	result = fTest(a, b)
	return
}

func main() {
	a := Calc(2, 3, Add)
	fmt.Printf("a=%d \n", a)
}
