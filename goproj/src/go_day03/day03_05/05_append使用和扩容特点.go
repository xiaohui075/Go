package main

import "fmt"

import "math/rand"

import "time"

func main() {

	s := make([]int, 10)

	InitData(s)

	fmt.Printf("排序前：cap(s) = %d,len(s)=%d,s = %d\n", cap(s), len(s), s)

	bubbleSort(s)

	fmt.Printf("\n 排序后：cap(s) = %d,len(s)=%d,s = %d\n", cap(s), len(s), s)
	// oldcap := cap(s)

	// for i := 0; i < 20; i++ {
	// 	s = append(s, i)
	// 	if newcap := cap(s); oldcap < newcap {
	// 		fmt.Printf("cap : %d ====>%d \n", oldcap, newcap)
	// 		oldcap = newcap
	// 	}
	// }
}

func InitData(s []int) {

	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 10; i++ {
		s[i] = rand.Intn(100)
	}
}

func bubbleSort(s []int) {
	n := len(s)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-1-i; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}
