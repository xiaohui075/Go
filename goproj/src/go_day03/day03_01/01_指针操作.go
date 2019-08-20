package main

import "fmt"

func main() {
	var a int
	//每个变量有两层含义：变量的内存，变量的地址
	a = 10
	fmt.Printf("a=%d \n", a)

	fmt.Printf("a = %v \n", &a)

	var p *int
	p = &a
	fmt.Printf("&a = %v,p = %v \n", &a, p)

	*p = 666 //*p操作的不是p的内存，而是p所指向的地址的内存

	fmt.Printf("*p = %v,a = %v ，p=%v\n", *p, a, p)

	var b *int
	b = nil
	fmt.Println("b= ", b)
}
