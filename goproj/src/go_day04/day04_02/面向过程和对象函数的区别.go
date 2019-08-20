package main

import "fmt"

func Add01(a, b int) int {
	return a + b
}

//面向对象，方法：给某个类型绑定一个函数
type long int

//tmp叫接受者，是传递的一个参数
func (tmp long) Add02(result long) long {
	return tmp + result
}

func main() {
	var sum1 int
	sum1 = Add01(2, 4)
	fmt.Println("sum1=", sum1)

	var a long = 5
	//调用方法格式：变量名.函数（所需参数）
	sum2 := a.Add02(9)
	fmt.Println("sum2=", sum2)

}
