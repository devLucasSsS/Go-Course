package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var pl = fmt.Println

func main() {
	// var st = "What is your age?: "
	fmt.Print("What is your age? : ")
	reader := bufio.NewReader(os.Stdin)
	age, err := reader.ReadString('\n')
	agetrim := strings.TrimSpace(age)
	if err == nil {
		ageParse, err := strconv.Atoi(agetrim)
		if err == nil {
			if ageParse < 5 {
				pl("Too young for school")
			} else if ageParse == 5 {
				pl("Go to kinder")
			} else if (ageParse > 5) && (ageParse <= 17) {
				grade := ageParse - 5
				pl("go to grade %d\n", grade)
			} else {
				pl("go to college")
			}
		}
	}

}
