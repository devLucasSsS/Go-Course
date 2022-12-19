package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

var pl = fmt.Println

func main() {
	pl("5 + 4 = ", 5+4)
	pl("5 - 4 = ", 5-4)
	pl("5 * 4 = ", 5*4)
	pl("5 / 4 = ", 5/4)
	pl("5 % 4 = ", 5%4)
	mInt := 1
	// mInt += 1
	mInt++
	pl("Float Precision = ", 0.1111+0.1111)
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	randNum := rand.Intn(50) + 1
	pl("Random :", randNum)
	pl("ABS(-10) = ", math.Abs(-10))
	pl("Pow(4,2) = ", math.Pow(4, 2))
	pl("Sqrt(16) = ", math.Sqrt(16))
	pl("Cbrt(8) = ", math.Cbrt(8))
	pl("Ceil(4,2) = ", math.Ceil(4.4))
	pl("Floor(4,4) = ", math.Floor(4.4))
	pl("Round(4,4) = ", math.Round(4.4))
	pl("Log2(8) = ", math.Log2(8))
	pl("Log10(100) = ", math.Log10(100))
	pl("Log(7.389) = ", math.Log(7.389))
	pl("Max(5,4) = ", math.Max(5, 4))
	pl("Min(5,4) = ", math.Min(5, 4))

}
