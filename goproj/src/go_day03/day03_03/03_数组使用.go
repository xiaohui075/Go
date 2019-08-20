package main

import "fmt"

func main() {
	var a [10]int
	var b [5]int

	fmt.Printf("len(a)=%d,len(b)=%d \n", len(a), len(b))

	// var i int
	// i = 20
	// var c [i]int

	for i := 0; i < len(a); i++ {
		a[i] = a[i] + i
	}

	for j, data := range a {
		fmt.Printf("a[%d] = %d \n", j, data)
	}

	//初始化
	var c [5]int = [5]int{1, 2, 3, 4, 5}
	fmt.Println("c = ", c)

	//部分初始化，没有初始化的元素自动赋值0
	d := [5]int{3, 5, 24}
	fmt.Println("d = ", d)

	//指定摸个元素初始化
	f := [6]int{3: 54, 5: 32}
	fmt.Println("f = ", f)

	//二维数组
	var h [4][3]int
	var k int
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			k++
			h[i][j] = k
			fmt.Printf("h[%d][%d] = %d \n", i, j, h[i][j])
		}
	}
}
