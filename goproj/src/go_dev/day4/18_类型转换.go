package main

import "fmt"

func main() {

	var flag bool
	flag = true
	fmt.Printf("flag= %t\n", flag)

	ch := 'a'
	var a int
	a = int(ch)
	fmt.Printf("a = %d \n", a)

	var b int
	fmt.Println("请输入b 的值：")
	fmt.Scan(&b)
	if b > 0 {
		fmt.Println("b 大于 0 ")
	} else if b < 0 {
		fmt.Println("b小于0")
	} else {
		fmt.Println("b 等于0")
	}

	switch b { //switch 后面写的是变量本身 默认包含break
	case -8:
		fmt.Printf("输入的是负数  b = %d\n", b)
		fallthrough //不跳出switch语句，后面的无条件执行
	case 0:
		fmt.Printf("输入的是0 b = %d\n", b)
		fallthrough
	case 1:
		fmt.Printf("按下的是1 b = %d \n", b)
		fallthrough
	case 8:
		fmt.Printf("按下的是8 b = %d\n", b)
		fallthrough
	default:
		fmt.Printf("按下的是其他按钮 b = %v\n", b)
	}

	score := 90
	switch {
	case score >= 90:
		fmt.Println("成绩优秀")
	case score > 80:
		fmt.Println("成绩良好")
	case score > 70:
		fmt.Println("成绩合格")
	default:
		fmt.Println("成绩不合格")
	}
}
