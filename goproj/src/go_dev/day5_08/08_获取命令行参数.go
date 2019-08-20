package main

import io "fmt"
import "os"

func main() {

	list := os.Args

	n := len(list)

	io.Println("n=", n)

}
