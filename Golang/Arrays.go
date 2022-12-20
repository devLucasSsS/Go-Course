package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	var arr1 [5]int
	arr1[0] = 1
	arr2 := [5]int{1, 2, 3, 4, 5}
	pl("Index 0: ", arr2[0])
	pl("Array length: ", len(arr2))
	for i := 0; i < len(arr2); i++ {
		pl(arr2[i])
	}
	for i, v := range arr2 {
		fmt.Printf("%d : %d", i, v)
	}
	arr3 := [2][2]int{
		{1, 2},
		{3, 4},
	}
	for i := 0; i < 2; i++ {
		for j := 0; j < 2; j++ {
			pl(arr3[i][j])
		}
	}

	aStr1 := "abcde"
	rArr := []rune(aStr1)
	for _, v := range rArr {
		fmt.Printf("Rune array: %d\n", v)
	}
	byteArr := []byte{'a', 'b', 'c'}
	bStr := string(byteArr[:])
	pl("I'm a string: ", bStr)
}
