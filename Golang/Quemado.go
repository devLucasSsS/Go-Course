package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"regexp"
	"strings"
	"time"
)

var pl = fmt.Println

var hmArr = [7]string{
	" +---+\n" +
		" 	  |\n" +
		" 	  |\n" +
		" 	  |\n" +
		"    ===\n",
	" +---+\n" +
		" 0	  |\n" +
		" 	  |\n" +
		" 	  |\n" +
		"    ===\n",
	" +---+\n" +
		" 0	  |\n" +
		" |	  |\n" +
		" 	  |\n" +
		"    ===\n",
	" +---+\n" +
		" 0	  |\n" +
		"/|   |\n" +
		" 	  |\n" +
		"    ===\n",
	" +---+\n" +
		" 0	  |\n" +
		"/|\\  |\n" +
		" 	  |\n" +
		"    ===\n",
	" +---+\n" +
		" 0	  |\n" +
		"/|\\ |\n" +
		"/   |\n" +
		"    ===\n",
	" +---+\n" +
		" 0	  |\n" +
		"/|\\ |\n" +
		"/ \\  |\n" +
		"    ===\n",
}
var wordArr = [7]string{
	"HOLA",
	"AUTOS",
	"PYTHON",
	"JAVA",
	"GOLANG",
	"JAZZ",
	"ZODIAC",
}

var randWord string
var guessedLetters string
var correctLetters []string
var wrongGuesses []string

func getRandWord() string {
	seedSecs := time.Now().Unix()
	rand.Seed(seedSecs)
	randWord = wordArr[rand.Intn(7)]
	correctLetters = make([]string, len(randWord))
	return randWord
}

func showGameBoard() {
	fmt.Println(hmArr[len(wrongGuesses)])
	fmt.Print("Palabra Secreta: ")
	for _, v := range correctLetters {
		if v == "" {
			fmt.Print("_")
		} else {
			fmt.Print(v)
		}
	}
	fmt.Print("\nLetras Erroneas: ")
	if len(wrongGuesses) > 0 {
		for _, v := range wrongGuesses {
			fmt.Print(v + " ")
		}
	}
	pl()
}

func getUserLetter() string {
	reader := bufio.NewReader(os.Stdin)
	var guess = ""
	for true {
		fmt.Print("\n Ingresa una letra: ")
		guess, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		guess = strings.ToUpper(guess)
		guess = strings.TrimSpace(guess)
		var IsLetter = regexp.MustCompile(`^[a-zA-Z]+$`).MatchString
		if len(guess) != 1 {
			pl("Por favor ingresa solo una letra")
		} else if !IsLetter(guess) {
			pl("Por favor ingresa solo texto")
		} else if strings.Contains(guessedLetters, guess) {
			pl("Por favor ingresa una palabra que no hayas ingresado")
		} else {
			return guess
		}
	}
	return guess
}

func getAllIndexes(thStr, subStr string) (indices []int) {
	if (len(subStr) == 0) || (len(thStr) == 0) {
		return indices
	}
	offset := 0
	for {
		i := strings.Index(thStr[offset:], subStr)
		if i == -1 {
			return indices
		}
		offset += i
		indices = append(indices, offset)
		offset += len(subStr)
	}
}

func updateCorrectLetters(letter string) {
	indexMatches := getAllIndexes(randWord, letter)
	for _, v := range indexMatches {
		correctLetters[v] = letter
	}
}

func slicesHasEmptya(TheSlice []string) bool {
	for _, v := range TheSlice {
		if len(v) == 0 {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(getRandWord())
	for true {
		showGameBoard()
		guess := getUserLetter()
		if strings.Contains(randWord, guess) {
			updateCorrectLetters(guess)
			if slicesHasEmptya(correctLetters) {
				pl("Mas letras para adivinar")
			} else {
				pl("SI!, la palabra secreta es: ", randWord)
				break
			}
		} else {
			guessedLetters += guess
			wrongGuesses = append(wrongGuesses, guess)
			if len(wrongGuesses) >= 6 {
				pl("Perdiste, la palabra secreta era: ", randWord)
			}
		}
	}
}
