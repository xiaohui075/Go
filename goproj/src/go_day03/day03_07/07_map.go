package main

import "fmt"

func main() {

	var m1 map[int]string
	fmt.Println("m1=", m1)

	m2 := make(map[int]string, 2)
	m2[1] = "mike"
	m2[2] = "go"
	m2[3] = "take"
	fmt.Println("m2=", m2)
	fmt.Printf("len(m2)= %d\n", len(m2))

	m3 := map[int]string{1: "dkks", 2: "ksdi", 3: "eidj"}
	fmt.Printf("m3=%v,len(m3)=%d \n", m3, len(m3))

	for key, value := range m3 {
		fmt.Printf("%d ====> %s \n", key, value)
	}

	/*
		如何判断一个key值是否存在
		第一个返回值为key所对应的value，第二个返回值为key是否存在的条件，存在ok为true
	*/
	value, ok := m3[1]
	if ok == true {
		fmt.Println("m3[1]=", value)
	} else {
		fmt.Println("key不存在")
	}

	//删除某个key值
	delete(m3, 1)
	fmt.Println("m3=", m3)
	test1(m3)
	fmt.Println("m3=", m3)
}

func test1(m map[int]string) {
	m[1] = "iii"
	m[2] = "ddd"
}
