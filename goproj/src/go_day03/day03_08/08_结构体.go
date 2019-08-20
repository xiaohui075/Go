package main

import "fmt"

func main() {
	var s Student = Student{1, "mike", 'm', 25, "gz"}
	s1 := Student{1, "mike", 'm', 25, "gz"}
	s2 := Student{name: "jams", age: 12, address: "sz"}
	fmt.Printf("s = %v,s1=%v,s2  = %v", s, s1, s2)
}

//结构体
type Student struct {
	id      int
	name    string
	sex     byte
	age     int
	address string
}
