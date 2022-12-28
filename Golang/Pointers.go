package main

import (
	"fmt"
)

var pl = fmt.Println

func changeVal2(myPtr *int) {
	*myPtr = 12
}

func dblArrVals(arr *[4]int) {
	for x := 0; x < 4; x++ {
		arr[x] *= 2
	}
}

func getAvg(num ...float64) float64 {
	var sum float64 = 0.0
	var numsize float64 = float64(len(num))
	for _, val := range num {
		sum += val
	}
	return (sum / float64(numsize))
}

func main() {
	f4 := 10
	pl("f4: ", f4)
	pl("f4 Address: ", &f4)
	var f4Ptr *int = &f4
	pl("f4 Address: ", f4Ptr)
	pl("f4 Value: ", *f4Ptr)
	*f4Ptr = 11
	pl("f4 Value: ", *f4Ptr)
	pl("f4 before Function: ", *f4Ptr)
	changeVal2(&f4)
	pl("f4 after Function: ", *f4Ptr)

	pArr := [4]int{1, 2, 3, 4}
	dblArrVals((*[4]int)(&pArr))
	pl(pArr)
	iSlice := []float64{11, 13, 17}
	fmt.Printf("Average : %.3f\n", getAvg(iSlice...))
}
