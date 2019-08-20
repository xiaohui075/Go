package main

import (
	"fmt"
	"go_dev/day2/goroute"
)

func test() (a, b, c int) {
	return 1, 2, 3
}
func main() {
	pipe := make(chan int, 1)
	go goroute.Add(200, 352, pipe)

	sum := <-pipe
	fmt.Println("sum=", sum)

	var c, d, e int
	c, d, e = test()
	fmt.Printf("c = %d, d = %d ,e =%d \n", c, d, e)

	_, d, _ = test()
	fmt.Printf("d = %d \n", d)

	//iota 常量自动生成器，每个一行，自动累加1
	const (
		l       = iota
		p       = iota
		m, k, u = iota, iota, iota //同一行，iota值相同
	)

	const n = iota
	fmt.Printf("l = %d , p = %d , m = %d  , k = %d , u = %d , n = %d\n", l, p, m, k, u, n)

	var ch byte
	ch = 97
	fmt.Printf("%c, %d \n", ch, ch)

	ch = 'a'

	fmt.Printf("%c  ，%d \n", ch, ch)

	fmt.Printf("大写： %c  ， 小写：%d \n", 'A', 'a')

	//大小写相差32  大转小加，小转大减
	fmt.Printf("大写转小写：%c \n", 'A'+32)

	fmt.Printf("小写转大写：%c \n", 'a'-32)

	/*
		字符串
	*/
	str := "hello go "

	fmt.Printf("str[0] = %c , str[1] = %c \n", str[0], str[1])

	/*
		复数
	*/

	var t complex128
	t = 2.1 + 3.14i
	fmt.Println("t = ", t)

	t1 := 2.5 + 4.8i
	fmt.Printf("t1 type is %T \n", t1)

	/*
		格式化输出
	*/

	a := 100
	b := "abc"
	q := 'a'
	w := 3.15
	fmt.Printf("%T , %T, %T ,%T \n", a, b, q, w)

	fmt.Printf("a = %d, b = %s ,q = %c ,w = %f\n", a, b, q, w)

	//%v 读取默认的格式
	fmt.Printf("a = %v , b = %v ,q = %v ,w = %v\n", a, b, q, w)

	//变量的输入
	var a1 int
	fmt.Println("请输入变量a1：")
	// fmt.Scanf("%d", &a1)
	fmt.Scan(&a1)
	fmt.Printf("a1=%d \n", a1)
}
