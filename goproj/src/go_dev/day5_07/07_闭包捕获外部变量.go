package main

import "fmt"

func main() {

	a := 30
	str := "json"

	func() {
		a = 56
		str = "mike"
		fmt.Printf("a = %d ,str = %v \n", a, str)

	}()

	fmt.Printf("a = %d  ,str = %v \n", a, str)

	fmt.Println("*********************")

	/*
		返回值为一个匿名函数，返回一个函数类型，通过f来调用返回的匿名 函数，f来调用闭包函数
		它不关心这些捕获了的变量和常量是否已经超出了作用域
		所以只有闭包还在使用它，这些变量就还会存在
	*/
	f := test()
	fmt.Println(f())
	// defer fmt.Println(f())
	// defer fmt.Println(f())
	// defer fmt.Println(f())
	// defer fmt.Println(f())
	// defer fmt.Println(f())

	fmt.Println("*******************************")

	a1 := 100
	b1 := 200

	defer func(x, y int) {
		fmt.Printf("a1=%d,b1=%d \n ", x, y)
	}(a1, b1)

	a1 = 300
	b1 = 400
	fmt.Printf("外部：a1 = %d ,b1 = %d \n", a1, b1)
}

func test() func() int {
	var x int
	return func() int {
		x++
		return x * x
	}
}
