package calc

func Add(a int ,b int ,p chan int)  {
	sum:= a+b
	p <- sum
}


func Sub(c int ,d int ,p2 chan int ){
	sub :=c-d
	p2 <- sub
}
