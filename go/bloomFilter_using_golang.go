package main

import (
	"fmt"
	"log"
	"math"
	"reflect"
	"strconv"
)

var pl = fmt.Println
var pf = fmt.Printf

func errorCheck(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func hash1(word string, arrSize int64) int64 {

	var wordHash int64 = 0

	for i := 0; i < len(word); i++ {

		pl("here", word[i:i+1], reflect.TypeOf(word[i:i+1]))
		intConver, err := strconv.Atoi(word[i : i+1])
		errorCheck(err)
		pl("here")

		wordHash = (wordHash + int64(intConver))
		wordHash %= arrSize

	}

	return wordHash
}

func hash2(word string, arrSize int64) int64 {

	var wordHash int64 = 1

	for i := 0; i < len(word); i++ {

		intConver, err := strconv.Atoi(word[i : i+1])
		errorCheck(err)

		wordHash = (wordHash + int64(intConver)*int64(math.Pow(19, float64(i))))
		wordHash %= arrSize

	}

	return wordHash % arrSize
}

func hash3(word string, arrSize int64) int64 {

	var wordHash int64 = 7

	for i := 0; i < len(word); i++ {

		intConver, err := strconv.Atoi(word[i : i+1])
		errorCheck(err)

		wordHash = wordHash*31 + int64(intConver)
		wordHash %= arrSize
	}

	return wordHash % arrSize
}

func hash4(word string, arrSize int64) int64 {

	var wordHash int64 = 3

	for i := 0; i < len(word); i++ {

		intConver, err := strconv.Atoi(word[i : i+1])
		errorCheck(err)

		wordHash = wordHash*7 + int64(intConver)*int64(math.Pow(7, float64(i)))
		wordHash %= arrSize
	}

	return wordHash
}

func lookup(bitarray [100]bool, arrSize int64, word string) bool {

	var a int64 = hash1(word, arrSize)
	var b int64 = hash2(word, arrSize)
	var c int64 = hash3(word, arrSize)
	var d int64 = hash4(word, arrSize)

	if bitarray[a] && bitarray[b] && bitarray[c] && bitarray[d] {
		return true
	} else {
		return false
	}

}

func insert(bitarray [100]bool, arrSize int64, word string) {

	if lookup(bitarray, arrSize, word) {
		pf("%s is probably already present!\n", word)
	} else {

		var a int64 = hash1(word, arrSize)
		var b int64 = hash2(word, arrSize)
		var c int64 = hash3(word, arrSize)
		var d int64 = hash4(word, arrSize)

		bitarray[a] = true
		bitarray[b] = true
		bitarray[c] = true
		bitarray[d] = true

		pf("%s inserted!\n", word)

	}

}

func main() {

	var bitarray [100]bool
	var arrSize int = 100

	sarray := [33]string{"abound", "abounds", "abundance",
		"abundant", "accessible", "bloom",
		"blossom", "bolster", "bonny",
		"bonus", "bonuses", "coherent",
		"cohesive", "colorful", "comely",
		"comfort", "gems", "generosity",
		"generous", "generously", "genial",
		"bluff", "cheater", "hate",
		"war", "humanity", "racism",
		"hurt", "nuke", "gloomy",
		"facebook", "geeksforgeeks", "twitter",
	}

	for i := 0; i < 33; i++ {

		insert(bitarray, int64(arrSize), sarray[i])
	}
}
