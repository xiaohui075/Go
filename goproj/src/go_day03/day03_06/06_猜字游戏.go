package main

import (
	"fmt"
	"math/rand"
	"time"
)

func creatNum(p *int) {
	rand.Seed(time.Now().UnixNano())
	var num int
	for {
		num = rand.Intn(10000)
		if num >= 1000 {
			break
		}
	}
	*p = num
}

func getNum(slice []int, n int) {
	slice[0] = n / 1000
	slice[1] = n % 1000 / 100
	slice[2] = n % 100 / 10
	slice[3] = n % 10
}
func main() {
	var randNum int

	creatNum(&randNum)

	// fmt.Println("randNum = ", randNum)

	randSlice := make([]int, 4)
	getNum(randSlice, randNum)
	fmt.Println("randSlice = ", randSlice)

	onGame(randSlice)

}

func onGame(randSlice []int) {
	var num int
	for {
		for {
			fmt.Printf("请输入一个4位数：")
			fmt.Scan(&num)
			if 999 < num && num < 10000 {
				break
			}
			fmt.Println("请输入符合要求的数！")
		}
		// fmt.Println("num = ", num)
		keySlice := make([]int, 4)

		getNum(keySlice, num)

		fmt.Println("keySlice = ", keySlice)

		n := 0
		for i := 0; i < 4; i++ {
			if randSlice[i] > keySlice[i] {
				fmt.Printf("第%d位数小了点 \n", i+1)
			} else if randSlice[i] < keySlice[i] {
				fmt.Printf("第%d位数大了点 \n", i+1)
			} else {
				fmt.Printf("第%d位数猜对了 \n", i+1)
				n++
			}
		}
		if n == 4 {
			fmt.Println("全部猜对!")
			break
		}
	}
}
