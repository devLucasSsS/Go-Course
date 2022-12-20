package main

import (
	"fmt"
)

var pl = fmt.Println

func main() {
	sl1 := make([]string, 6)
	sl1[0] = "Sociedad"
	sl1[1] = "Anonima"
	sl1[2] = "De"
	sl1[3] = "Los"
	sl1[4] = "Creadores"
	// sl1[5] = "Vetadores"
	pl("Slice Size: ", len(sl1))
	for i := 0; i < len(sl1); i++ {
		pl(sl1[i])
	}
	for _, x := range sl1 {
		pl(x)
	}
	sl2 := []int{14, 04, 2003}
	pl(sl2)
	sArr := [5]int{1, 2, 3, 4, 5}
	sl3 := sArr[0:2]
	pl(sl3)
	sArr[0] = 10
	pl("sl3 : ", sl3)
	sl3 = append(sl3, 12)
	pl("sl13 :", sl3)
	pl("sArr: ", sArr)
	sl4 := make([]string, 6)
	pl("sl4", sl4)
	pl("sl4[0", sl4[0])
}
