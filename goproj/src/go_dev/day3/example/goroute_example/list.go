package goroute_example

import(
	"fmt"
)
func Sum(n int)  {
	for i := 0; i <=n; i++ {
		fmt.Printf("%d+%d=%d\n", i ,  n-i, n)
	}
}