package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

var pl = fmt.Println

func main() {
	reader := bufio.NewReader(os.Stdin)
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	randNum := rand.Intn(50) + 1
	for true {
		fmt.Print("Guess a number between 0 and 50 : ")
		pl("Random number is: ", randNum)
		num, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		numTrim := strings.TrimSpace(num)
		numParse, err1 := strconv.Atoi(numTrim)
		if err1 != nil {
			log.Fatal(err1)
		}
		if numParse > randNum {
			pl("Guess Something Lower")
		} else if numParse < randNum {
			pl("Guess Something Higher")
		} else {
			pl("You Guessed it")
			break
		}
	}

}
