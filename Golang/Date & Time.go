package main

import (
	"fmt"
	"time"
)

var pl = fmt.Println

func main() {
	now := time.Now()
	pl(now.Year(), now.Month(), now.Day())
	pl(now.Hour(), now.Minute(), now.Second())
	loc, err := time.LoadLocation("America/Santiago")
	if err != nil {
		pl(err)
	}
	fmt.Printf("Time in Santiago %s\n", now.In(loc))

	Birthday := time.Date(2003, time.April, 14, 02, 30, 0, 0, time.Local)
	pl(Birthday)
	diff := now.Sub(Birthday)
	fmt.Printf("Days Alive: %d days\n", int(diff.Hours()/24))
	fmt.Printf("Hours Alive: %d hours\n", int(diff.Hours()))
}
