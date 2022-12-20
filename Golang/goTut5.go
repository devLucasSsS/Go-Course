package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	// for x := 1; x <= 5; x++ {
	// 	pl(x)
	// }
	// aNums := []int{1, 2, 3}
	// for _, num := range aNums {
	// 	pl(num)
	// }
	xVal := 1
	for true {
		if xVal == 5 {
			break
		}
		pl(xVal)
		xVal++
	}
}
