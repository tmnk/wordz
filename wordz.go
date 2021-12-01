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

func GenWord(length int) string {
	if length == 0 {
		length = random(3, 10)
	}
	word := ""
	for i:= 0; i < length; i++ {
		word += genRandLetter()
	}
	return word
}

func GenString(length int, lenthOfWord int, sep string) string {
	if length == 0 {
		length = random(3, 10)
	}
	result := ""
	for i:= 0; i < length; i++ {
		if len(result) != 0 {
			result += sep +  GenWord(lenthOfWord)
		} else {
			result += GenWord(lenthOfWord)
		}
	}
	return result
}

func init() {
	var SEED = time.Now().Unix()
	rand.Seed(SEED)
}