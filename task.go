package main

import (
	"fmt"
	"math/rand"
	"strings"
)

func main() {
	printSegitiga(4)
	printSegitiga(5)

	fmt.Println(genPass("abcd", "low"))
	fmt.Println(genPass("abcd", "med"))
	fmt.Println(genPass("abcd", "strong"))

	fmt.Println("Movie Number (duration 17) : ")
	fmt.Print(chooseFilmByDuration(17))
	fmt.Println("Movie Number  (duration 10) : ")
	fmt.Print(chooseFilmByDuration(10))
	fmt.Println("Movie Number  (duration 1) : ")
	fmt.Println(chooseFilmByDuration(1))

}

func printSegitiga(value int) {
	fmt.Println("Print segitiga = ", value)
	for i := value; i > 0; i-- {

		for k := i; k <= value-1; k++ {
			fmt.Print(" ")
		}
		for j := 1; j <= 2*i-1; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func genPass(pass string, level string) string {

	charspecial := "!@#$%^&*?"
	passRune := []rune(pass)
	var passNumber string
	var passSpecial string

	if level == "low" || level == "med" || level == "strong" {
		for i := 0; i < 6; i++ {
			passNumber += fmt.Sprint(rand.Intn(9))
		}
	}

	if level == "med" || level == "strong" {
		passSpecial = string(charspecial[rand.Intn(len(charspecial))])

	}

	if level == "strong" {
		for i := 0; i <= rand.Intn(len(pass)); i++ {
			index := rand.Intn(len(pass))
			passRune[index] = []rune(strings.ToUpper(string(passRune[index])))[0]
		}

	}

	return string(passRune) + passSpecial + passNumber

}

func chooseFilmByDuration(hour int) string {
	var listMovieDuration = []int{1, 7, 3, 4, 8, 9}
	var result string = ""
	for index, value := range listMovieDuration {
		for i := index + 1; i < len(listMovieDuration); i++ {
			if value+listMovieDuration[i] == hour {
				result += fmt.Sprint(value, " dan ", listMovieDuration[i], "\n")
			}
		}
	}

	if result == "" {
		return "No movie"
	}
	return result
}

func runHelloWorld() {
	fmt.Println("Hello World")
}
