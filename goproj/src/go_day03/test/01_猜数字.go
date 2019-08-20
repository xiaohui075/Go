package main

import (
	"fmt"
	"math/rand"
	"time"
)

func creatNum(randNum *int) {
	var p int
	rand.Seed(time.Now().UnixNano())
	for {
		p = rand.Intn(10000)
		if p > 999 {
			break
		}
	}
	*randNum = p
}

func getNum(randSlice []int, randNum int) {

	randSlice[0] = randNum / 1000
	randSlice[1] = randNum % 1000 / 100
	randSlice[2] = randNum % 100 / 10
	randSlice[3] = randNum % 10

}

func onGame(randSlice []int) {
	var n int

	for {
		for {
			fmt.Println("请输入一个4位数：")
			fmt.Scan(&n)
			if n > 999 && n < 10000 {
				break
			}
			fmt.Println("您输入的数不符合要求，请重新输入！")
		}
		// fmt.Println("n = ", n)
		keySlice := make([]int, 4)
		getNum(keySlice, n)

		j := 0

		for i := 0; i < 4; i++ {
			if randSlice[i] > keySlice[i] {
				fmt.Printf("第%d数小了点 \n", i+1)
			} else if randSlice[i] < keySlice[i] {
				fmt.Printf("第%d数大了点 \n", i+1)
			} else {
				fmt.Printf("第%d数猜对了了\n", i+1)
				j++
			}
		}
		if j == 4 {
			fmt.Printf("全部猜对了 \n")
			break
		}
	}

}

func main() {
	var randNum int
	creatNum(&randNum)

	// fmt.Println("randNum=", randNum)

	randSlice := make([]int, 4)

	getNum(randSlice, randNum)

	// fmt.Println("randSlice = ", randSlice)

	onGame(randSlice)

}
