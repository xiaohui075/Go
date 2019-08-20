package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func minus(a, b int) int {
	return a - b
}

type funcType func(int, int) int //没有函数名字，没有{}

func main() {

	var result int
	result = add(3, 4)
	fmt.Println("result=", result)

	//声明一个函数类型的变量
	var fTest funcType
	fTest = add
	result = fTest(5, 6)
	fmt.Println("result2=", result)

}
