package main

import "fmt"

func test(a int) {
	if a == 1 {
		fmt.Println("a = ", a)
		return
	}

	fmt.Println("a=", a)
	test(a - 1)

}

//递归叠加
func test1() (sum int) {
	for i := 0; i <= 100; i++ {
		sum += i
	}
	return sum
}

func test2(i int) int {
	if i == 1 {
		return 1
	}
	return i + test2(i-1)
}

func test3(i int) int {
	if i == 100 {
		return 100
	}

	return i + test3(i+1)
}
func main() {
	test(5)

	sum := test1()
	sum1 := test2(100)
	sum2 := test3(1)
	fmt.Printf("sum=%d, sum1 = %d, sum2 = %d \n", sum, sum1, sum2)
}
