package main

import "fmt"

func mFunc() {
	a := 66666
	fmt.Println("a=", a)

}

//有参无返回值函数的定义，普通参数列表
//定义函数时，在函数名后面定义的参数叫形参
//参数传递，只能由实参传递给形参，不能反过来， 单向传递
func test(a int) {

	fmt.Println("a=", a)
}

func test1(a int, b int) {
	fmt.Printf("a =%d, b = %d \n", a, b)
}

//不定参数类型
func test2(args ...int) {
	fmt.Println("打印参数： ", args)

	fmt.Println("*******************************")

	for i := 0; i < len(args); i++ {
		fmt.Printf("args[%d] = %d \n", i, args[i])
	}

	fmt.Println("*******************************")
	for i, data := range args {
		fmt.Printf("args[%d] = %d \n", i, data)
	}

	fmt.Println("*********传递参数***********")
	test3(args...)
	fmt.Println("*********传递下标为2之前的参数***********")
	test3(args[:2]...)
	fmt.Println("*********传递下标为2之后且包括args[2]本身参数***********")
	test3(args[2:]...)
}

func test3(tmp ...int) {
	for i, data := range tmp {
		fmt.Printf("tmp[%d] = %d \n", i, data)
	}
}
func main() {
	mFunc()

	//调用函数传递的参数叫实参
	test(88888)

	test1(6666, 8888)

	test2(2, 36, 584, 285)
}
