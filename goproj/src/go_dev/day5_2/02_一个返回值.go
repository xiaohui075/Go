package main

import "fmt"

//单个返回值
func myFunc01() int {
	a := 666
	return a
}

//多个返回值
func myFunc02() (a, b, c int) {
	a, b, c = 12, 13, 16
	return a, b, c
}
func main() {
	
	fmt.Println("a = ", myFunc01())

	a, b, c := myFunc02()
	fmt.Printf("a = %d ,b = %d ,c = %d \n", a, b, c)















}











