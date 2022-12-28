package main

import (
	"fmt"
)

var pl = fmt.Println

func Hello() {
	pl("Hello")
}

func Suma(x int, y int) int {
	return x + y
}

func getQuotient(x float64, y float64) (ans float64, err error) {
	if y == 0 {
		return 0, fmt.Errorf("no se puede divir por 0")
	} else {
		return x / y, nil
	}
}

func getTwo(x int) (int, int) {
	return x + 1, x + 2
}

func SumaArray(arr []int) int {
	sum := 0
	for _, num := range arr {
		sum += num
	}
	return sum
}
func main() {
	Hello()
	pl(Suma(1, 1))
	ans, err := getQuotient(5, 0)
	if err == nil {
		fmt.Printf("%f / %f = %f", 5.0, 0.0, ans)
	} else {
		pl(err)
	}
	f1, f2 := getTwo(5)
	fmt.Printf("%d %d\n", f1, f2)
	vArr := []int{1, 2, 3, 4, 5}
	pl("Unknow sum: ", SumaArray(vArr))
}
