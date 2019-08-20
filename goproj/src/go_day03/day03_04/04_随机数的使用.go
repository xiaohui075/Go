package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	var a [10]int
	n := len(a)

	fmt.Println("排序前：")

	for i := 0; i < n; i++ {
		a[i] = rand.Intn(100)
		fmt.Printf(" a[%d] = %d \n", i, a[i])
	}

	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if a[j] > a[j+1] {
				a[j], a[j+1] = a[j+1], a[j]
			}
		}
	}

	fmt.Println("\n 排序后：")
	for i := 0; i < n; i++ {
		fmt.Printf("a[%d] = %d \n", i, a[i])
	}

}
