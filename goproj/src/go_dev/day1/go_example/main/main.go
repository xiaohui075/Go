package main

import(
	"go_dev/day1/go_example/calc"
	"fmt"
)

func main(){
	pipe := make(chan  int , 1)
	pipe1 :=make(chan int ,1)
	go calc.Add(285,632,pipe)
	go calc.Sub(9987,8752,pipe1)
	sum:=  <- pipe
	sub :=  <- pipe1

	for i:=0;i<100;i++{
		fmt.Println("Hello World!",sum,"\t",sub,"\t",i)
	}

	c := 0303
	fmt.Printf("c type is %T \n",c)
	// fmt.Println("Hello World!",sum,"\t",sub)
}
