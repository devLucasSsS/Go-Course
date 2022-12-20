package main

import (
	"bufio"
	f "fmt"
	"log"
	"os"
)

var pl = f.Println

// commentary
/* block  comments  */

func main() {
	pl("hello")
	pl("What is your name?")
	reader := bufio.NewReader(os.Stdin)
	name, err := reader.ReadString('\n')
	if err == nil {
		pl("Hello", name)
	} else {
		log.Fatal(err)
	}

}
