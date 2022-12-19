package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var pl = fmt.Println

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter the first number : ")
	num1, err1 := reader.ReadString('\n')
	if err1 != nil {
		log.Fatal(err1)
	}
	num1 = strings.TrimSpace(num1)
	num1Parse, err3 := strconv.Atoi(num1)
	if err3 != nil {
		log.Fatal(err3)
	}
	fmt.Print("Enter the second number : ")
	num2, err2 := reader.ReadString('\n')
	if err2 != nil {
		log.Fatal(err2)
	}
	num2 = strings.TrimSpace(num2)
	num2Parse, err4 := strconv.Atoi(num2)
	if err4 != nil {
		log.Fatal(err4)
	}
	fmt.Printf("%s + %s = %d\n", num1, num2, (num1Parse + num2Parse))
}
