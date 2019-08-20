package main

import "fmt"

import "time"

func main() {

	str := "abcde"

	//通过for打印每个字符

	//普通写法
	for i := 0; i < len(str); i++ {
		fmt.Printf("str[%d]=%c\n", i, str[i])
	}

	fmt.Println("range写法 \n ")
	//range的写法   迭代  打印每个元素，默认返回2个值：一个是元素的位置，一个是元素本身
	for i, data := range str {
		fmt.Printf(" str[%d]=%c\n", i, data)
	}

	fmt.Println("****************************************")
	for i := range str { //第2个返回值，默认丢弃，返回元素的位置（下标）
		fmt.Printf("str[%d]=%c \n", i, str[i])
	}

	fmt.Println("****************************************")

	for i, _ := range str {
		fmt.Printf("str[%d]=%c \n", i, str[i])
	}

	j := 0
	for {
		j++
		time.Sleep(time.Second)

		if j == 5 {
			break
			// continue
		}

		fmt.Println("i=", j)
	}

	goto End

	fmt.Println("2222222222222222")

End:
	fmt.Println("33333333333333")
}
