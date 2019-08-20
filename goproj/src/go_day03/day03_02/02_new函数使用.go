package main

import "fmt"

func main() {
	var a *int
	a = new(int)

	b := new(int)
	*b = 6666
	fmt.Println("b=", b)
	fmt.Println("a = ", a)
	fmt.Printf("b=%v,&b=%v \n", b, *b)

	d, c := 10, 20
	swap(&d, &c)
	fmt.Printf(" *d = %v,*c = %v \n", &d, &c)
}

func swap(d, c *int) {
	*d, *c = *c, *d
	fmt.Printf("swap *d = %v,*c = %v \n", d, c)
	fmt.Printf("swap *d = %v,*c = %v \n", *d, *c)
}
