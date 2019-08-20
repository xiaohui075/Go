package main

import "fmt"

func main() {

	a := 20
	str := "jake"

	f1 := func() {
		fmt.Printf("a = %d , str = %v \n", a, str)
	}

	f1()

	type FuncType func()
	var f2 FuncType
	f2 = f1
	f2()

	//定义匿名函数，同时调用
	func() {
		fmt.Printf("a = %d ,str = %v \n", a, str)
	}() // 后面的（）代表调用此匿名函数

	//带参数的匿名函数
	func(i, j int) {
		sum := i + j
		fmt.Println("sum=", sum)
	}(30, 23)

	//带参数带返回值的匿名函数
	x, y := func(c, d int) (max, min int) {
		if c > d {
			max = c
			min = d
		} else if c < d {
			max = d
			min = c
		} else {
			max, min = c, d
		}
		return max, min
	}(23, 56)

	fmt.Printf("x=%d,y= %d \n", x, y)
}
