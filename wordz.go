package wordz

import (
	"math/rand"
	"time"
)

func genRandLetter() string {
	return string(rand.Intn(123 - 97) + 97)
}

func random(min, max int) int {
	return rand.Intn(max-min) + min
}

func GenWord(n ...int) string {
	var length int
	if len(n) == 0 {
		length = random(3, 10)
	} else {
		length = n[0]
	}
	word := ""
	for i:= 0; i < length; i++ {
		word += genRandLetter()
	}
	return word
}

func GenString(n ...int) string {
	var length int
	if len(n) == 0 {
		length = random(3, 10)
	} else {
		length = n[0]
	}
	result := ""
	for i:= 0; i < length; i++ {
		if len(result) != 0 {
			result += " " +  GenWord()
		} else {
			result += GenWord()
		}
	}
	return result
}

func init() {
	var SEED = time.Now().Unix()
	rand.Seed(SEED)
}