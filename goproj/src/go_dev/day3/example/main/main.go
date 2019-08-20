package main

import(
	"fmt"
	"go_dev/day3/example/goroute_example"
)

func main()  {
	go goroute_example.Sum(100)
	
	fmt.Println("**********************************")

	  dada(100)

	}

	func dada(n int)  {
		for i := 0; i <=n; i++ {
			fmt.Printf("%d+%d=%d \t  main \n", i ,  n-i, n)
		}
	}
//goto 不得已采用的
//const 申明一个常量
//fallthough switch 一起 用，如果不想break还想继续走就用fallthough
//range  便利数组
//type
//continue
//for
//struct 申明一个结构体
//defer 清理资源，捕获
//select 用在chan管道
//case 
// default
